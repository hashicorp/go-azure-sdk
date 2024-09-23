package deviceconfigurationgroupassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationGroupAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationGroupAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationGroupAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationgroupassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationGroupAssignmentClient: %+v", err)
	}

	return &DeviceConfigurationGroupAssignmentClient{
		Client: client,
	}, nil
}
