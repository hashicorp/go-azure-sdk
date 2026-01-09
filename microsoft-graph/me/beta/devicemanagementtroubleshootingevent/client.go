package devicemanagementtroubleshootingevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTroubleshootingEventClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementTroubleshootingEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementtroubleshootingevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementTroubleshootingEventClient: %+v", err)
	}

	return &DeviceManagementTroubleshootingEventClient{
		Client: client,
	}, nil
}
