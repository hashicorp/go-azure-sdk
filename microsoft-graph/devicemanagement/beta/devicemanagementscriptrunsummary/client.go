package devicemanagementscriptrunsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptRunSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptRunSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptRunSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptrunsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptRunSummaryClient: %+v", err)
	}

	return &DeviceManagementScriptRunSummaryClient{
		Client: client,
	}, nil
}
