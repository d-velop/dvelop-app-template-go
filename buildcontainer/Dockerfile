FROM debian:stretch

# unzip, zip, python, build-essential (make, gcc and more) ...
RUN apt-get update && \
    apt-get -y install --no-install-recommends ca-certificates unzip zip python curl openssh-client git-core build-essential jq && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOLANG_VERSION 1.12.4
ENV GOLANG_CHECKSUM d7d1f1f88ddfe55840712dc1747f37a790cbcaa448f6c9cf51bbe10aa65442f5
RUN curl -fsSL https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz -o golang.tar.gz && \
    echo "${GOLANG_CHECKSUM} golang.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf golang.tar.gz && \
    rm golang.tar.gz && \
    echo "prebuild standard packages for windows to speed up cross compilation" && \
    GOOS=windows GOARCH=amd64 go install -v std

# go-junit-report https://github.com/jstemmer/go-junit-report
RUN go get -u github.com/jstemmer/go-junit-report

# terraform
ENV TERRAFORM_VERSION 0.11.13
ENV TERRAFORM_CHECKSUM 5925cd4d81e7d8f42a0054df2aafd66e2ab7408dbed2bd748f0022cfe592f8d2
RUN curl -fsSL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip -o terraform.zip  && \
    echo "${TERRAFORM_CHECKSUM} terraform.zip" | sha256sum -c - && \
    unzip terraform.zip -d /usr/local/bin && chmod +x /usr/local/bin/terraform ; rm terraform.zip

# terraform terraform provider plugin (for remote state)
ENV TERRAFORM_TERRAFORM_PLUGIN_VERSION 1.0.2
ENV TERRAFORM_TERRAFORM_PLUGIN_CHECKSUM c4c1e826c85ef45fb18ca75e06284d13b0f72bce61a652e74eb016cf7779eafd
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-terraform/${TERRAFORM_TERRAFORM_PLUGIN_VERSION}/terraform-provider-terraform_${TERRAFORM_TERRAFORM_PLUGIN_VERSION}_linux_amd64.zip -o terraform_terraform_plugin.zip  && \
    echo "${TERRAFORM_TERRAFORM_PLUGIN_CHECKSUM} terraform_terraform_plugin.zip" | sha256sum -c - && \
    unzip terraform_terraform_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_terraform_plugin.zip

# terraform template provider plugin
ENV TERRAFORM_TEMPLATE_PLUGIN_VERSION 2.1.1
ENV TERRAFORM_TEMPLATE_PLUGIN_CHECKSUM de3c45a4fa1f756aa4db3350c021d1c0f9b23640cff77e0ba4df4eeb8eae957f
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-template/${TERRAFORM_TEMPLATE_PLUGIN_VERSION}/terraform-provider-template_${TERRAFORM_TEMPLATE_PLUGIN_VERSION}_linux_amd64.zip -o terraform_template_plugin.zip  && \
    echo "${TERRAFORM_TEMPLATE_PLUGIN_CHECKSUM} terraform_template_plugin.zip" | sha256sum -c - && \
    unzip terraform_template_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_template_plugin.zip

# terraform null provider plugin
ENV TERRAFORM_NULL_PLUGIN_VERSION 2.1.1
ENV TERRAFORM_NULL_PLUGIN_CHECKSUM 5b29f6a87f4dae4dc295a43f499182221333dbe474fcf724b8ea98fd38cfb3e6
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-null/${TERRAFORM_NULL_PLUGIN_VERSION}/terraform-provider-null_${TERRAFORM_NULL_PLUGIN_VERSION}_linux_amd64.zip -o terraform_null_plugin.zip  && \
    echo "${TERRAFORM_NULL_PLUGIN_CHECKSUM} terraform_null_plugin.zip" | sha256sum -c - && \
    unzip terraform_null_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_null_plugin.zip

# terraform aws provider plugin
ENV TERRAFORM_AWS_PLUGIN_VERSION 2.7.0
ENV TERRAFORM_AWS_PLUGIN_CHECKSUM b2a046c82f38c1c9b89503be7d4b83049b1af445d11ea08a1b2b3196d510d02f
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-aws/${TERRAFORM_AWS_PLUGIN_VERSION}/terraform-provider-aws_${TERRAFORM_AWS_PLUGIN_VERSION}_linux_amd64.zip -o terraform_aws_plugin.zip  && \
    echo "${TERRAFORM_AWS_PLUGIN_CHECKSUM} terraform_aws_plugin.zip" | sha256sum -c - && \
    unzip terraform_aws_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_aws_plugin.zip

# aws cli
RUN curl -sSL https://s3.amazonaws.com/aws-cli/awscli-bundle.zip -o awscli-bundle.zip && \
    unzip awscli-bundle.zip && \
    ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws && \
    rm awscli-bundle.zip && rm -rf ./awscli-bundle

WORKDIR /build
ENTRYPOINT [ "/usr/bin/make" ]
