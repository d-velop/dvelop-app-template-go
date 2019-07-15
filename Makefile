APP_NAME=shop-middleware
DOMAIN_SUFFIX=.service.d-velop.cloud
BUILD_VERSION=rev.$(shell git rev-parse --short HEAD).date.$(shell date '+%d-%m-%Y-%H.%M.%S')

all: build

generate:
	go generate -tags release ./...

clean:
	rm -rf ./dist
	rm -rf ./terraform/.terraform

init:
	mkdir -p ./dist/test-reports

test: init
	go test ./... -v | tee ./dist/test-reports/testresult.txt && go-junit-report -set-exit-code < ./dist/test-reports/testresult.txt > ./dist/test-reports/testresult.xml

build: clean build-app build-lambda

build-app: generate test
	GOOS=windows GOARCH=amd64 go build -tags release -o ./dist/$(APP_NAME)app.exe  ./cmd/app/

build-lambda: generate test
	GOOS=linux GOARCH=amd64 go build -tags release -o ./dist/lambda ./cmd/lambda/ &&\
	cd ./dist && touch -t $(date +%Y)01010000 lambda && zip -X -j lambda.zip lambda && # for reproducible zip file cf. https://content.pivotal.io/blog/barriers-to-deterministic-reproducible-zip-files\
	cd ..

tf-bucket:
	$(eval BUCKET_NAME=$(APP_NAME)-terraform)
	@aws s3api get-bucket-location --bucket $(BUCKET_NAME) > /dev/null 2>&1; \
	if  [ "$$?" -ne "0" ]; \
	then \
		echo Create terraform state bucket \"$(BUCKET_NAME)\"...; \
		aws s3api create-bucket --bucket $(BUCKET_NAME) --acl private --region eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1 &&\
		aws s3api put-bucket-versioning --bucket $(BUCKET_NAME) --versioning-configuration Status=Enabled &&\
		aws s3api put-public-access-block --bucket $(BUCKET_NAME) --public-access-block-configuration BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true; \
	fi

tf-init: tf-bucket
	cd ./terraform && \
	terraform init -input=false -plugin-dir=/usr/local/lib/custom-terraform-plugins

plan: tf-init build-lambda asset_hash
	$(eval PLAN=$(shell mktemp))
	cd ./terraform && \
	terraform plan -input=false \
	-var "signature_secret=$(SIGNATURE_SECRET)" \
	-var "build_version=$(BUILD_VERSION)" \
	-var "appname=$(APP_NAME)" \
	-var "domainsuffix=$(DOMAIN_SUFFIX)" \
	-var "asset_hash=$(ASSET_HASH)" \
	-out=$(PLAN)

apply: plan
	cd ./terraform && \
	terraform apply -input=false -auto-approve=true $(PLAN)

deploy-assets: asset_hash apply
	# best practice for immutable content: cache 1 year (vgl https://jakearchibald.com/2016/caching-best-practices/)
	aws s3 sync ./web s3://assets.$(APP_NAME)$(DOMAIN_SUFFIX)/$(ASSET_HASH) --exclude "*.html" --cache-control max-age=31536000

asset_hash:
	$(eval ASSET_HASH=$(shell find web -type f ! -path "*.html" -exec md5sum {} \; | sort -k 2 | md5sum | tr -d " -"))

deploy: apply deploy-assets

show: tf-init
	cd ./terraform && \
	terraform show

rename:
	if [ -z $${NAME} ]; then echo "NAME is not set. Usage: rename NAME=NEW_APP_NAME"; exit 1; fi
	@echo Rename App to $(NAME) ...
	find . -name "docker-build.*" -or -name "Makefile" -or -name "*.tf" -or -name "*.go" | while read f; do		\
		echo "Processing file '$$f'";															\
		sed -i 's/$(APP_NAME)/$(NAME)/g' $$f;														\
	done

destroy: tf-init
	echo "destroy is disabled. Uncomment in Makefile to enable destroy."
	#cd ./terraform && \
	#terraform destroy -var "signature_secret=$SIGNATURE_SECRET" -var "build_version=$build_version" -var "appname=$(APP_NAME)" -var "domainsuffix=$(DOMAIN_SUFFIX)" -input=false -force

dns:
	cd ./terraform && terraform output -json | jq "{Domain: .domain.value, Nameserver: .nameserver.value}" > ../dist/dns-entry.json

