package userdeviceregisteredowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceRegisteredOwnerClient struct {
	Client *msgraph.Client
}

func NewUserDeviceRegisteredOwnerClientWithBaseURI(api sdkEnv.Api) (*UserDeviceRegisteredOwnerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdeviceregisteredowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceRegisteredOwnerClient: %+v", err)
	}

	return &UserDeviceRegisteredOwnerClient{
		Client: client,
	}, nil
}
