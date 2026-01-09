package deviceconfigurationuserstatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationUserStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationUserStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationuserstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationUserStateSummaryClient: %+v", err)
	}

	return &DeviceConfigurationUserStateSummaryClient{
		Client: client,
	}, nil
}
