package virtualendpointusersetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointUserSettingClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointUserSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointUserSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointusersetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointUserSettingClient: %+v", err)
	}

	return &VirtualEndpointUserSettingClient{
		Client: client,
	}, nil
}
