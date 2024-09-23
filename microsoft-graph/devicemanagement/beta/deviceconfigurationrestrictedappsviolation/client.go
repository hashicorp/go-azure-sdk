package deviceconfigurationrestrictedappsviolation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationRestrictedAppsViolationClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationRestrictedAppsViolationClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationRestrictedAppsViolationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationrestrictedappsviolation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationRestrictedAppsViolationClient: %+v", err)
	}

	return &DeviceConfigurationRestrictedAppsViolationClient{
		Client: client,
	}, nil
}
