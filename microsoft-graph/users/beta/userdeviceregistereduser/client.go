package userdeviceregistereduser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceRegisteredUserClient struct {
	Client *msgraph.Client
}

func NewUserDeviceRegisteredUserClientWithBaseURI(api sdkEnv.Api) (*UserDeviceRegisteredUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdeviceregistereduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceRegisteredUserClient: %+v", err)
	}

	return &UserDeviceRegisteredUserClient{
		Client: client,
	}, nil
}
