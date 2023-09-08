package memanageddevicesecuritybaselinestate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceSecurityBaselineStateClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceSecurityBaselineStateClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceSecurityBaselineStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicesecuritybaselinestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceSecurityBaselineStateClient: %+v", err)
	}

	return &MeManagedDeviceSecurityBaselineStateClient{
		Client: client,
	}, nil
}
