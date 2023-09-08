package userdevicecommand

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceCommandClient struct {
	Client *msgraph.Client
}

func NewUserDeviceCommandClientWithBaseURI(api sdkEnv.Api) (*UserDeviceCommandClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdevicecommand", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceCommandClient: %+v", err)
	}

	return &UserDeviceCommandClient{
		Client: client,
	}, nil
}
