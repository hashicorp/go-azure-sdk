package manageddevicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceCategoryClient: %+v", err)
	}

	return &ManagedDeviceCategoryClient{
		Client: client,
	}, nil
}
