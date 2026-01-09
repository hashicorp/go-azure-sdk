package comanageddevicedevicecompliancepolicystate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceDeviceCompliancePolicyStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceDeviceCompliancePolicyStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicedevicecompliancepolicystate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceDeviceCompliancePolicyStateClient: %+v", err)
	}

	return &ComanagedDeviceDeviceCompliancePolicyStateClient{
		Client: client,
	}, nil
}
