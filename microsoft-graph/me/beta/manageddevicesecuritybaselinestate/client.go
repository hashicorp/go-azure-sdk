package manageddevicesecuritybaselinestate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceSecurityBaselineStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceSecurityBaselineStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicesecuritybaselinestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceSecurityBaselineStateClient: %+v", err)
	}

	return &ManagedDeviceSecurityBaselineStateClient{
		Client: client,
	}, nil
}
