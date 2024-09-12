package virtualendpointorganizationsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointOrganizationSettingClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointOrganizationSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointOrganizationSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointorganizationsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointOrganizationSettingClient: %+v", err)
	}

	return &VirtualEndpointOrganizationSettingClient{
		Client: client,
	}, nil
}
