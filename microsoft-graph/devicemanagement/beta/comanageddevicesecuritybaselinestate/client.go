package comanageddevicesecuritybaselinestate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceSecurityBaselineStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceSecurityBaselineStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicesecuritybaselinestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceSecurityBaselineStateClient: %+v", err)
	}

	return &ComanagedDeviceSecurityBaselineStateClient{
		Client: client,
	}, nil
}
