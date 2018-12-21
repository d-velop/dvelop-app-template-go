# d.velop app template for Go

This template contains everything you need to write an app for d.velop cloud.

To demonstrate all the aspects of app development a hypothetical but not trivial use case
of *an employee applying for vacation* is implemented.

## Getting Started

Just clone this repo and follow the [build instructions](#build) to get the sample app up and running.
After this adjust the code to fit the purpose of your own business problem/app.

### Prerequisites

A linux docker container is used for the build and deployment process of the app.
So besides docker the only thing you need on your local development system is a git client
and an editor or IDE for golang.

### Build

Execute the build with

```
docker-build build
```

This will build a self contained web application `dist/<appname>app.exe` which can be used to run and test your app 
as a local process on your dev system and a deployment package for aws lambda `dist/lambda` which 
should be used for the production deployment of your app in d.velop cloud.

## Run and test your app locally

Just start `dist/<appname>app.exe` to run and test your app on a local dev environment.
Please keep in mind, that some functions like authentication
which require the presence of additional apps (e.g. IdentityProviderApp), 
won't work because these apps are not available on your local system.

## Rename the app

You should change the name of the app so that it reflects the business problem you would like
to solve.

Each appname in d.velop cloud must be unique. To facilitate this every provider/company chooses
a unique provider prefix which serves as a namespace for the apps of this provider.
The prefix can be selected during the registration process in d.velop cloud.
If you choose a provider prefix which corresponds to your company name or an abbreviation of the company name
it's very likely that it is available when you later register your app in d.velop cloud.

For example if your company is named *Super Duper Software Limited* and the domain of your app 
is *employees applying for vacation* your app should be named
something like `superduperltd-vacationprocess`App. Note that the `App` suffix isn't used in the configuration files. 

Apps belonging to the core d.velop cloud platform don't have a provider prefix. 

For now the following places have to be adjusted manually as soon as the name of the app changes:

1.  `Makefile` change the `APP_NAME` variable. Furthermore change the `DOMAIN_SUFFIX` to a domain you own like `yourcompany.com`
2.  `/terraform/backend.tf` change the `bucket` names (2 occurrences)
3.  `docker-build.bat` and `dockerbuild.sh` change the `APPNAME` variable      
4.  `/domain/plugins/conf/config.go` change the `AppName` const
5.  `go.mod` change the module name. Unfortunately this requires to change the import path in a lot of go files from
    `github.com/d-velop/dvelop-app-template-go` to something like `github.com/<yourcompany>/<appname>`.
    
The 'Replace' function of your IDE should help.

**Please finish at least step 1 and step 2 before you [deploy](#deployment) your app because the names of a lot of
AWS resources are derived from the `APP_NAME` and `DOMAIN_SUFFIX`. Changing them afterwards requires a
redeployment of the AWS resources which takes some time**

## Deployment

**Please read [Rename the app](#rename-the-app) before you proceed with the deployment.**

You need an AWS Account to deploy your app. At the time of writing some of the AWS services are
free to use for a limited amount of time and workload. 
Check the [Free Tier](https://aws.amazon.com/free/) offering from AWS for the current conditions. 

Manually create the S3 bucket for the terraform state configured in `terraform/backend.tf` and an IAM user with
the appropriate rights to create the AWS resources defined by your terraform configuration. 
You could start with a user who has the `arn:aws:iam::aws:policy/AdministratorAccess` policy to start quickly, 
but you **should definitely restrict the rights of that IAM user to a minimum as soon as you go into production**.

Configure the AWS credentials of the created IAM user by using one of the methods described in
[Configuring the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).
For example set the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables.

**Windows**

```
SET AWS_ACCESS_KEY_ID=<YOUR-ACCESS-KEY-ID>
SET AWS_SECRET_ACCESS_KEY=<YOUR-SECRET-ACCESS-KEY>
```

**Linux**

```
export AWS_ACCESS_KEY_ID=<YOUR-ACCESS-KEY-ID>
export AWS_SECRET_ACCESS_KEY=<YOUR-SECRET-ACCESS-KEY>
```

Deploy the lambda function and all other AWS resources like AWS API Gateway.

```
docker-build deploy
```

The build container uses [Terraform](https://www.terraform.io/) to manage the AWS resources and to deploy
your lambda function. This tool implements a desired state mechanism which means the first execution will take some time
to provision all the required AWS resources. Consecutive executions will only deploy the difference between the desired state
(e.g. the new version of your lambda function) and the state which is already deployed (other AWS resources which won't change
between deployments) and will be much quicker.

### Test your endpoint

The endpoint URLs are logged at the end of the deployment. Just invoke them in a browser to test your app.  
 
```
Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

endpoint = [
    https://xxxxxxxxxx.execute-api.eu-central-1.amazonaws.com/prod/vacationprocess/,
    https://xxxxxxxxxx.execute-api.eu-central-1.amazonaws.com/dev/vacationprocess/
]

```

To watch the current deployment state you can invoke

```
docker-build show 
```

at any time without changing your deployment.

### Deployment of a new app version

Just follow the [deployment](#deployment) steps. A new deployment package for the lambda function will be build automatically.

### Additional AWS resources

The terraform deployment configuration contains 2 additonal modules which are disabled by default.
Just uncomment the corresponding lines in `/terraform/main.tf` to use them but **ensure that the DNS resolution
for your hosted zone works before you use these modules**. Read the comments in the terraform file.

#### asset_cdn
This module uses *aws cloudfront* as a CDN for your static assets. Furthermore it allows you to define
a custom domain for your assets instead of the s3 URL. Your deployment should work perfectly without this module.

#### api_custom_domain 
This module allows you to define a custom domain for your app endpoints. A custom domain name is required
as soon as you register your app in the d.velop cloud center because the base path of your app must
begin with the name of your app. So instead of the default endpoints

```
    https://xxxxxxxxxx.execute-api.eu-central-1.amazonaws.com/prod/vacationprocess/
    https://xxxxxxxxxx.execute-api.eu-central-1.amazonaws.com/dev/vacationprocess/
```
which base paths begin with `/prod` or `/dev` you need endpoints like

```
    https://vacationprocess.xyzdomain./vactionprocess
    https://def.vacationprocess.xyzdomain./vactionprocess
```
which are provided by this module.

## Projectstructure

The presented structure is by no means mandatory for d.velop cloud apps and is highly opinionated.
Feel free to change the structure if it doesn't fit your needs.
On the other hand it takes a significant amount of time to invent a logical and useful structure
for apps and we are pretty sure this structure is at least a good starting point.
So we would recommend that you try to use it and get comfortable with it so you don't
waste your time and start immediately to implement a solution for your business problem.

### Go Directories

#### `/cmd`

Contains the main applications for this project. That is the self contained webapplication `/cmd/app` 
which can be run on your local machine and the lambda function `/cmd/lambda` for AWS.

Don't put a lot of code in the application directory. Put that code in the `/domain` directory.

It's common to have a small `main` function which basically wires up the dependencies and apart from this
completely relies on the code from the `/domain` directory.

#### `/domain`

Contains the vast majority of the code for this app.

The structure follows the principles of 
[Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) or
[Hexagonal Architecture](https://alistair.cockburn.us/coming-soon/) 
and separates the core of the domain from external frameworks, the DB and the UI.

The directory root contains the heart of the domain and has no dependencies to external 'things' like http or databases.

##### `/domain/<UseCase>`

Each use case of the domain has its own subdirectory that is named after the use case. So you should be able
to understand the business domain of an app you've never seen before by opening the domain folder and looking
at the directory names.

The use cases don't have any dependencies to external 'things' either.

##### `/domain/mock`

Contains test mocks which are relevant to more than one use case.

##### `/domain/plugins`

Contains the dependencies to external 'things' like a database or the invocation channel e.g. http.
The idea is to treat these external 'things' as plugins to the domain in order to keep 
the domain simple, understandable and separately testable. Last but not least you are able to change external
dependencies like the database later on without rewriting the whole app because the relevant code
is not scattered over the whole codebase.

Again you might want to read
[Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
or [Hexagonal Architecture](https://alistair.cockburn.us/coming-soon/)

### `/buildcontainer`

Contains the `Dockerfile` for the buildcontainer. It is kept in a separate directory to keep
the buildcontext small so that the image can be build as fast as possible.

### `/terraform`

Contains the terraform files

### `/web`

Contains the webfrontend. 

The frontend tooling is kept to a bare minimum to keep the whole project as simple as possible.
Furthermore there are hundreds of possible combinations of frameworks and build tools which
can be used for the frontend. So each developer has his own preferences about the tooling.

Use your favorite tools for the frontend and change the `web` folder accordingly. 
Don't forget to change the `deploy-assets` task in the `makefile` and the go:generate commands in `/domain/plugins/gui/`

It's likely that we'll provide web projects using different tools and frameworks 
which can be used to replace the `web` folder in the future. 

## Go Modules

This project uses [Go Modules](https://golang.org/doc/go1.11#modules). That means you need at least Go 1.11 if you want to compile
this project outside the build container. This means also that your project **must not be located in GOPATH/src**
(cf.[Preliminary module support](https://golang.org/cmd/go/#hdr-Preliminary_module_support))and the **depedencies
must not be checked into source control**.

### IDE Support for Go Modules

In some IDEs, like JetBrains GoLand, Go Modules support must be activated explicitly in order to get IntelliSense.
* Settings > Go > Go Module (vgo) - Enable Go Modules (vgo) integration

## Build mechanism
A linux docker container is used to build and deploy the software. This has the advantage, that the build
doesn't rely on specific tools or tool versions that need to be installed on the local development machine or
build server.

During the build the whole application directory is mmounted in the docker container. The build targets are
implemented in the `Makefile`.

Two wrappers (`docker-build.bat` and `docker-build.sh`) are provided so you don't have to remember the
rather long docker command.
Furthermore these wrappers provide a little utility function to passthrough all environment variables listed
in the `environment` file from the docker host (that is your development machine or buildserver)
to the build container.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

Thanks to the following projects for inspiration

* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
* [How Do You Structure Your Go Apps](https://github.com/katzien/go-structure-examples)
* [GoDDD](https://github.com/marcusolsson/goddd)
* [Starting an Open Source Project](https://opensource.guide/starting-a-project/)
* [README template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)
* [CONTRIBUTING template](https://github.com/nayafia/contributing-template/blob/master/CONTRIBUTING-template.md)

