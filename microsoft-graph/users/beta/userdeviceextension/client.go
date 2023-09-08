package userdeviceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserDeviceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserDeviceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdeviceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceExtensionClient: %+v", err)
	}

	return &UserDeviceExtensionClient{
		Client: client,
	}, nil
}
