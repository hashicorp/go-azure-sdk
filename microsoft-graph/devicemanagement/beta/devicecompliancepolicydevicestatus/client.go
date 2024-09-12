package devicecompliancepolicydevicestatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyDeviceStatusClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyDeviceStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyDeviceStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicydevicestatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyDeviceStatusClient: %+v", err)
	}

	return &DeviceCompliancePolicyDeviceStatusClient{
		Client: client,
	}, nil
}
