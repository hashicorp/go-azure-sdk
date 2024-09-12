package manageddevicecompliancepolicystate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCompliancePolicyStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceCompliancePolicyStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceCompliancePolicyStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicecompliancepolicystate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceCompliancePolicyStateClient: %+v", err)
	}

	return &ManagedDeviceCompliancePolicyStateClient{
		Client: client,
	}, nil
}
