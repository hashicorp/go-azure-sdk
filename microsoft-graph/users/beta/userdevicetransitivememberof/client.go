package userdevicetransitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceTransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewUserDeviceTransitiveMemberOfClientWithBaseURI(api sdkEnv.Api) (*UserDeviceTransitiveMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdevicetransitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceTransitiveMemberOfClient: %+v", err)
	}

	return &UserDeviceTransitiveMemberOfClient{
		Client: client,
	}, nil
}
