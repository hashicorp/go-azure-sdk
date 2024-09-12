package manageddevicedevicecompliancepolicystate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceCompliancePolicyStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceDeviceCompliancePolicyStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicedevicecompliancepolicystate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceDeviceCompliancePolicyStateClient: %+v", err)
	}

	return &ManagedDeviceDeviceCompliancePolicyStateClient{
		Client: client,
	}, nil
}
