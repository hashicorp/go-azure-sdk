package virtualendpointexternalpartnersetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointExternalPartnerSettingClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointExternalPartnerSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointExternalPartnerSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointexternalpartnersetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointExternalPartnerSettingClient: %+v", err)
	}

	return &VirtualEndpointExternalPartnerSettingClient{
		Client: client,
	}, nil
}
