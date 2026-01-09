package comanageddevicesecuritybaselinestatesettingstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceSecurityBaselineStateSettingStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceSecurityBaselineStateSettingStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicesecuritybaselinestatesettingstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceSecurityBaselineStateSettingStateClient: %+v", err)
	}

	return &ComanagedDeviceSecurityBaselineStateSettingStateClient{
		Client: client,
	}, nil
}
