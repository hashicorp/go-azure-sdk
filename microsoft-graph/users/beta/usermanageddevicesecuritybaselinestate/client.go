package usermanageddevicesecuritybaselinestate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceSecurityBaselineStateClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceSecurityBaselineStateClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceSecurityBaselineStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicesecuritybaselinestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceSecurityBaselineStateClient: %+v", err)
	}

	return &UserManagedDeviceSecurityBaselineStateClient{
		Client: client,
	}, nil
}
