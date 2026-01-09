package devicecompliancepolicyassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicyassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyAssignmentClient: %+v", err)
	}

	return &DeviceCompliancePolicyAssignmentClient{
		Client: client,
	}, nil
}
