package drivebundlecontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveBundleContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveBundleContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveBundleContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivebundlecontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveBundleContentStreamClient: %+v", err)
	}

	return &DriveBundleContentStreamClient{
		Client: client,
	}, nil
}
