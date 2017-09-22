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

package main

import (
	log "github.com/cihub/seelog"
	"github.com/docker/docker-credential-helpers/credentials"
	ssm "github.com/kindlyops/amazon-ssm-credential-helper/ssm-login"
	"github.com/kindlyops/amazon-ssm-credential-helper/ssm-login/api"
	"github.com/kindlyops/amazon-ssm-credential-helper/ssm-login/config"
)

func main() {
	defer log.Flush()
	config.SetupLogger()
	credentials.Serve(ssm.SSMHelper{ClientFactory: api.DefaultClientFactory{}})
}
