// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecr

import (
	"errors"
	"fmt"
	"github.com/kindlyops/amazon-ssm-credential-helper/ssm-login/api"
	log "github.com/cihub/seelog"
	"github.com/docker/docker-credential-helpers/credentials"
)

var notImplemented = errors.New("not implemented")

type SSMHelper struct {
	ClientFactory api.ClientFactory
}

// ensure SSMHelper adheres to the credentials.Helper interface
var _ credentials.Helper = (*SSMHelper)(nil)

func (SSMHelper) Add(creds *credentials.Credentials) error {
	// This does not seem to get called
	return notImplemented
}

func (SSMHelper) Delete(serverURL string) error {
	// This does not seem to get called
	return notImplemented
}

func (self SSMHelper) Get(serverURL string) (string, string, error) {
	defer log.Flush()

	client := self.ClientFactory.NewClientFromRegion(registry.Region)
	auth, err := client.GetCredentials(serverURL)
	if err != nil {
		log.Errorf("Error retrieving credentials: %v", err)
		return "", "", credentials.NewErrCredentialsNotFound()
	}
	return auth.Username, auth.Password, nil
}

func (self SSMHelper) List() (map[string]string, error) {
	log.Debug("Listing credentials")
	client := self.ClientFactory.NewClientWithDefaults()

	auths, err := client.ListCredentials()
	if err != nil {
		log.Errorf("Error listing credentials: %v", err)
		return nil, fmt.Errorf("Could not list credentials: %v:", err)
	}

	result := map[string]string{}

	for _, auth := range auths {
	        serverURL := auth.ProxyEndpoint
	        result[serverURL] = auth.Username
	}
	return result, nil
}


______

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type UserCredentials struct {
	username string
	password string
}
func getCredentials(name, credType string) string {
	svc := ssm.New(session.New())
	prams := &ssm.GetParameterInput{
		Name:   aws.String(name+credType),
		WithDecryption: aws.Bool(true),
	}
	resp, err := svc.GetParameter(prams)
	if err != nil {
		panic(err)
	}

	return *resp.Parameter.Value
}

func main() {
	// Create the service's client with the session.
	var url string = "gofactory-docker-godev.jfrog.io"
	var u string = getCredentials(url,"-usr")
	var p string = getCredentials(url,"-pwd")
	creds := UserCredentials{u, p}
	fmt.Printf("Found creds for:\n")
	fmt.Println(creds.username)
}

