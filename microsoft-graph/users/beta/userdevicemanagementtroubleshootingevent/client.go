package userdevicemanagementtroubleshootingevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceManagementTroubleshootingEventClient struct {
	Client *msgraph.Client
}

func NewUserDeviceManagementTroubleshootingEventClientWithBaseURI(api sdkEnv.Api) (*UserDeviceManagementTroubleshootingEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdevicemanagementtroubleshootingevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceManagementTroubleshootingEventClient: %+v", err)
	}

	return &UserDeviceManagementTroubleshootingEventClient{
		Client: client,
	}, nil
}
