package devicelocalcredential

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLocalCredentialClient struct {
	Client *msgraph.Client
}

func NewDeviceLocalCredentialClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceLocalCredentialClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicelocalcredential", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceLocalCredentialClient: %+v", err)
	}

	return &DeviceLocalCredentialClient{
		Client: client,
	}, nil
}
