package memanageddevicesecuritybaselinestatesettingstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceSecurityBaselineStateSettingStateClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceSecurityBaselineStateSettingStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicesecuritybaselinestatesettingstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceSecurityBaselineStateSettingStateClient: %+v", err)
	}

	return &MeManagedDeviceSecurityBaselineStateSettingStateClient{
		Client: client,
	}, nil
}
