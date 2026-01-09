package drivebundlecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveBundleContentClient struct {
	Client *msgraph.Client
}

func NewDriveBundleContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveBundleContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivebundlecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveBundleContentClient: %+v", err)
	}

	return &DriveBundleContentClient{
		Client: client,
	}, nil
}
