package devicecompliancepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyClient: %+v", err)
	}

	return &DeviceCompliancePolicyClient{
		Client: client,
	}, nil
}
