package devicecompliancescriptrunsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRunSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceComplianceScriptRunSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceComplianceScriptRunSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancescriptrunsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceComplianceScriptRunSummaryClient: %+v", err)
	}

	return &DeviceComplianceScriptRunSummaryClient{
		Client: client,
	}, nil
}
