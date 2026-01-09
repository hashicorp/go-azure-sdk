package deviceconfigurationconflictsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationConflictSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationConflictSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationConflictSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationconflictsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationConflictSummaryClient: %+v", err)
	}

	return &DeviceConfigurationConflictSummaryClient{
		Client: client,
	}, nil
}
