package deviceconfigurationdevicesettingstatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceSettingStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationDeviceSettingStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationDeviceSettingStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationdevicesettingstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationDeviceSettingStateSummaryClient: %+v", err)
	}

	return &DeviceConfigurationDeviceSettingStateSummaryClient{
		Client: client,
	}, nil
}
