package devicehealthscriptrunsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceHealthScriptRunSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceHealthScriptRunSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicehealthscriptrunsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceHealthScriptRunSummaryClient: %+v", err)
	}

	return &DeviceHealthScriptRunSummaryClient{
		Client: client,
	}, nil
}
