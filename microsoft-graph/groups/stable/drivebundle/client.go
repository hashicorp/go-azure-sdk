package drivebundle

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveBundleClient struct {
	Client *msgraph.Client
}

func NewDriveBundleClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveBundleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivebundle", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveBundleClient: %+v", err)
	}

	return &DriveBundleClient{
		Client: client,
	}, nil
}
