package devicemanagementresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementResourceNamespaceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementResourceNamespaceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementResourceNamespaceClient: %+v", err)
	}

	return &DeviceManagementResourceNamespaceClient{
		Client: client,
	}, nil
}
