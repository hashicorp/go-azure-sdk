package devicemanagementresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &DeviceManagementResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
