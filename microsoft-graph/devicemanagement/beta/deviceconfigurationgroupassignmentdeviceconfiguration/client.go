package deviceconfigurationgroupassignmentdeviceconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationGroupAssignmentDeviceConfigurationClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationGroupAssignmentDeviceConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationGroupAssignmentDeviceConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationgroupassignmentdeviceconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationGroupAssignmentDeviceConfigurationClient: %+v", err)
	}

	return &DeviceConfigurationGroupAssignmentDeviceConfigurationClient{
		Client: client,
	}, nil
}
