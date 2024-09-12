package intentdevicestatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentDeviceStateSummaryClient struct {
	Client *msgraph.Client
}

func NewIntentDeviceStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentDeviceStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentdevicestatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentDeviceStateSummaryClient: %+v", err)
	}

	return &IntentDeviceStateSummaryClient{
		Client: client,
	}, nil
}
