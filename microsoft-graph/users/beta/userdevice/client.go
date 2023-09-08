package userdevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceClient struct {
	Client *msgraph.Client
}

func NewUserDeviceClientWithBaseURI(api sdkEnv.Api) (*UserDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceClient: %+v", err)
	}

	return &UserDeviceClient{
		Client: client,
	}, nil
}
