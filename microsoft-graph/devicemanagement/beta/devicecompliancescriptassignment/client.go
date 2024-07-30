package devicecompliancescriptassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceComplianceScriptAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceComplianceScriptAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancescriptassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceComplianceScriptAssignmentClient: %+v", err)
	}

	return &DeviceComplianceScriptAssignmentClient{
		Client: client,
	}, nil
}
