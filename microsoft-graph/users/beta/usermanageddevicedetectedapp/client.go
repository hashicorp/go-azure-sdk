package usermanageddevicedetectedapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceDetectedAppClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceDetectedAppClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceDetectedAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicedetectedapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceDetectedAppClient: %+v", err)
	}

	return &UserManagedDeviceDetectedAppClient{
		Client: client,
	}, nil
}
