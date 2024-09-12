package comanageddevicecompliancepolicystate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceCompliancePolicyStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceCompliancePolicyStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceCompliancePolicyStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicecompliancepolicystate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceCompliancePolicyStateClient: %+v", err)
	}

	return &ComanagedDeviceCompliancePolicyStateClient{
		Client: client,
	}, nil
}
