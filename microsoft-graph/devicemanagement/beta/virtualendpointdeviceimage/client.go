package virtualendpointdeviceimage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointDeviceImageClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointDeviceImageClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointDeviceImageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointdeviceimage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointDeviceImageClient: %+v", err)
	}

	return &VirtualEndpointDeviceImageClient{
		Client: client,
	}, nil
}
