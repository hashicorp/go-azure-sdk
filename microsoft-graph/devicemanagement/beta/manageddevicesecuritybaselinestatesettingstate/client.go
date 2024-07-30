package manageddevicesecuritybaselinestatesettingstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceSecurityBaselineStateSettingStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceSecurityBaselineStateSettingStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicesecuritybaselinestatesettingstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceSecurityBaselineStateSettingStateClient: %+v", err)
	}

	return &ManagedDeviceSecurityBaselineStateSettingStateClient{
		Client: client,
	}, nil
}
