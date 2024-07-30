package deviceconfigurationassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationAssignmentClient: %+v", err)
	}

	return &DeviceConfigurationAssignmentClient{
		Client: client,
	}, nil
}
