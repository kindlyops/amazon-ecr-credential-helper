// Copyright 2017 Kindly Ops LLC.
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

package api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	log "github.com/cihub/seelog"
	"github.com/kindlyops/amazon-ssm-credential-helper/ssm-login/cache"
)

type Client interface {
	GetCredentials(serverURL string) (*Auth, error)
	ListCredentials() ([]*Auth, error)
}

type Auth struct {
	Username string
	Password string
}

type defaultClient struct {
	ecrClient       ssmiface.SSMAPI
	credentialCache cache.CredentialsCache
}

// GetCredentials returns username, password, and proxyEndpoint
func (dc defaultClient) GetCredentials(serverURL string) (*Auth, error) {
	sess := session.Must(session.NewSession())
	svc := ssm.New(sess)
	pramsUser := &ssm.GetParameterInput{
		Name:           aws.String(serverURL + "-usr"),
		WithDecryption: aws.Bool(true),
	}
	pramsPass := &ssm.GetParameterInput{
		Name:           aws.String(serverURL + "-pwd"),
		WithDecryption: aws.Bool(true),
	}
	respUser, errUser := svc.GetParameter(pramsUser)
	respPass, errPass := svc.GetParameter(pramsPass)
	if errUser != nil || errPass != nil {
		return nil, nil
	}

	log.Debugf("Retrieving credentials for (%s)", serverURL)

	return &Auth{
		Username: *respUser.Parameter.Value,
		Password: *respPass.Parameter.Value,
	}, nil
}

func (dc defaultClient) ListCredentials() ([]*Auth, error) {
	// TODO: what does this need?
	return nil, nil
}
