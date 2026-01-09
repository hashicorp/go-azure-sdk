package deviceshellscriptrunsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptRunSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptRunSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptRunSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscriptrunsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptRunSummaryClient: %+v", err)
	}

	return &DeviceShellScriptRunSummaryClient{
		Client: client,
	}, nil
}
