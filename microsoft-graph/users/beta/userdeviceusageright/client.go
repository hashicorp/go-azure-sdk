package userdeviceusageright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceUsageRightClient struct {
	Client *msgraph.Client
}

func NewUserDeviceUsageRightClientWithBaseURI(api sdkEnv.Api) (*UserDeviceUsageRightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdeviceusageright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceUsageRightClient: %+v", err)
	}

	return &UserDeviceUsageRightClient{
		Client: client,
	}, nil
}
