package devicecompliancepolicyuserstatusoverview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyUserStatusOverviewClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyUserStatusOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyUserStatusOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicyuserstatusoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyUserStatusOverviewClient: %+v", err)
	}

	return &DeviceCompliancePolicyUserStatusOverviewClient{
		Client: client,
	}, nil
}
