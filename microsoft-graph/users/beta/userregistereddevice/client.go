package userregistereddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegisteredDeviceClient struct {
	Client *msgraph.Client
}

func NewUserRegisteredDeviceClientWithBaseURI(api sdkEnv.Api) (*UserRegisteredDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userregistereddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserRegisteredDeviceClient: %+v", err)
	}

	return &UserRegisteredDeviceClient{
		Client: client,
	}, nil
}
