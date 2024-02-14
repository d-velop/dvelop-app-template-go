FROM node:14.20.1-buster

# unzip, zip, python, build-essential (make, gcc and more) ...
RUN apt-get update && \
    apt-get -y install --no-install-recommends ca-certificates unzip zip python3-venv curl openssh-client git-core build-essential jq && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOLANG_VERSION 1.21.7
ENV GOLANG_CHECKSUM 13b76a9b2a26823e53062fa841b07087d48ae2ef2936445dc34c4ae03293702c
RUN curl -fsSL https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz -o golang.tar.gz && \
    echo "${GOLANG_CHECKSUM} golang.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf golang.tar.gz && \
    rm golang.tar.gz && \
    echo "prebuild standard packages for $platform to speed up cross compilation" && \
    GOOS=windows GOARCH=amd64 go install -v std

# go-junit-report https://github.com/jstemmer/go-junit-report
# RUN go get -u github.com/jstemmer/go-junit-report
RUN go install github.com/jstemmer/go-junit-report@latest

# terraform
ENV TERRAFORM_VERSION 1.7.2
ENV TERRAFORM_CHECKSUM e3965584b2347edd294ca97068af573654716abbb2ce5230d309692dab659267
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
ENV TERRAFORM_TEMPLATE_PLUGIN_VERSION 2.2.0
ENV TERRAFORM_TEMPLATE_PLUGIN_CHECKSUM 8a154388f3708e3df5a69122a23bdfaf760a523788a5081976b3d5616f7d30ae
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-template/${TERRAFORM_TEMPLATE_PLUGIN_VERSION}/terraform-provider-template_${TERRAFORM_TEMPLATE_PLUGIN_VERSION}_linux_amd64.zip -o terraform_template_plugin.zip  && \
    echo "${TERRAFORM_TEMPLATE_PLUGIN_CHECKSUM} terraform_template_plugin.zip" | sha256sum -c - && \
    unzip terraform_template_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_template_plugin.zip

# terraform null provider plugin
ENV TERRAFORM_NULL_PLUGIN_VERSION 3.2.2
ENV TERRAFORM_NULL_PLUGIN_CHECKSUM 3248aae6a2198f3ec8394218d05bd5e42be59f43a3a7c0b71c66ec0df08b69e7
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-null/${TERRAFORM_NULL_PLUGIN_VERSION}/terraform-provider-null_${TERRAFORM_NULL_PLUGIN_VERSION}_linux_amd64.zip -o terraform_null_plugin.zip  && \
    echo "${TERRAFORM_NULL_PLUGIN_CHECKSUM} terraform_null_plugin.zip" | sha256sum -c - && \
    unzip terraform_null_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_null_plugin.zip

# terraform aws provider plugin
ENV TERRAFORM_AWS_PLUGIN_VERSION 5.34.0
ENV TERRAFORM_AWS_PLUGIN_CHECKSUM 3b0c9bdbbf846beef2c9573fc27898ceb71b69cf9d2f4b1dd2d0c2b539eab114
RUN curl -fsSL https://releases.hashicorp.com/terraform-provider-aws/${TERRAFORM_AWS_PLUGIN_VERSION}/terraform-provider-aws_${TERRAFORM_AWS_PLUGIN_VERSION}_linux_amd64.zip -o terraform_aws_plugin.zip  && \
    echo "${TERRAFORM_AWS_PLUGIN_CHECKSUM} terraform_aws_plugin.zip" | sha256sum -c - && \
    unzip terraform_aws_plugin.zip -d /usr/local/lib/custom-terraform-plugins ; rm terraform_aws_plugin.zip

#RUN update-alternatives --install /usr/bin/python python /usr/bin/python3.7 2 && \
#    update-alternatives --install /usr/bin/python python /usr/bin/python2.7 1

# aws cli
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install -i /usr/local/aws -b /usr/local/bin && \
    rm awscliv2.zip && rm -rf ./aws

# GUI BuildTools (NodeJS, Yarn, Gulp)
#RUN yarn global add gulp-cli

WORKDIR /build
ENTRYPOINT [ "/usr/bin/make" ]
