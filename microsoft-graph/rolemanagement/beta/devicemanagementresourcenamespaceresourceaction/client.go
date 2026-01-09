package devicemanagementresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementResourceNamespaceResourceActionClient: %+v", err)
	}

	return &DeviceManagementResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
