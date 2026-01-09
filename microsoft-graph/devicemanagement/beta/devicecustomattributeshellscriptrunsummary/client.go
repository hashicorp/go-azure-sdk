package devicecustomattributeshellscriptrunsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptRunSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptRunSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptRunSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptrunsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptRunSummaryClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptRunSummaryClient{
		Client: client,
	}, nil
}
