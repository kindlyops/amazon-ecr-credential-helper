# Amazon SSM Docker Credential Helper

[![Build Status](https://travis-ci.org/kindlyops/amazon-ssm-credential-helper.svg?branch=master)](https://travis-ci.org/kindlyops/amazon-ssm-credential-helper)

The Amazon SSM Docker Credential Helper is a
[credential helper](https://github.com/docker/docker-credential-helpers)
for the Docker daemon that makes it easier to use
[AWS EC2 Systems Manager Parameter Store](http://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html) to share
credentials for 3rd party private Docker registries such as JFrog Artifactory or Quay.io.

## Prerequisites

You must have at least Docker 1.11 installed on your system.

You also must have AWS credentials available in one of the standard locations:

* The `~/.aws/credentials` file
* The `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables
* An [IAM role for Amazon EC2](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html)
* If you are working with an assumed role please set the environment variable: `AWS_SDK_LOAD_CONFIG=true` also.

The Amazon SSM Docker Credential Helper uses the same credentials as the AWS
CLI and the AWS SDKs. For more information about configuring AWS credentials,
see
[Configuration and Credential Files](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-config-files)
in the *AWS Command Line Interface User Guide*.

The credentials must have a policy applied that
[allows access to Amazon SSM](http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-access.html).

## Installing

To build and install the Amazon SSM Docker Credential Helper, we suggest golang 
1.6+ and `git` and `make` installed on your system.

You can install this via `go get` with:

```
go get -u github.com/kindlyops/amazon-ssm-credential-helper/ssm-login/cli/docker-credential-ssm-login
```


If you already have Docker environment, just clone this repository anywhere
and run `make docker`. This command builds the binary by Go inside the Docker 
container and output it to local directory. 

With `TARGET_GOOS` environment variable, you can also cross compile the binary.

Place the `docker-credential-ssm-login` binary on your `PATH` and set the contents 
of your `~/.docker/config.json` file to be:

```json
{
	"credsStore": "ssm-login"
}
```

This configures the Docker daemon to use the credential helper for all Amazon
ECR registries.

With Docker 1.13.0 or greater, you can configure Docker to use different
credential helpers for different registries. To use this credential helper for
a specific ECR registry, create a `credHelpers` section with the URI of your
ECR registry:

```json
{
	"credHelpers": {
		"your-private-registry.jfrog.io": "ssm-login"
	}
}
```

This is useful if you use `docker` to operate on registries that use different 
authentication credentials.

## Usage

`docker pull your-private-registry.jfrog.io/my-repository:my-tag`

`docker push your-private-registry.jfrog.io/my-repository:my-tag`

There is no need to use `docker login` or `docker logout`.

### Usage with Docker Compose

Note that the current version of Docker Compose (1.14.0 at the time of writing)
does not support `credHelpers`. See
[docker/compose#4948](https://github.com/docker/compose/issues/4948) for more
information. Until this is supported, you need to specify
`"credsStore": "ecr-login"` as described above. Alternatively, you can manually
pull your images with `docker pull <your image>`, before running
`docker-compose up`.

## Troubleshooting

Logs from the Amazon SSM Docker Credential Helper are stored in `~/.ssm/log`.

For more information about Amazon SSM, see the
[Amazon Systems Manager Parameter Store](http://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html)
 documentation.

## License

The Amazon SSM Docker Credential Helper is licensed under the Apache 2.0
License.
