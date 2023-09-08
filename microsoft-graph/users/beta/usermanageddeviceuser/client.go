package usermanageddeviceuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceUserClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceUserClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddeviceuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceUserClient: %+v", err)
	}

	return &UserManagedDeviceUserClient{
		Client: client,
	}, nil
}
