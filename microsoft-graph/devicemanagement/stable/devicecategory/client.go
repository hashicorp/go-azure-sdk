package devicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewDeviceCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCategoryClient: %+v", err)
	}

	return &DeviceCategoryClient{
		Client: client,
	}, nil
}
