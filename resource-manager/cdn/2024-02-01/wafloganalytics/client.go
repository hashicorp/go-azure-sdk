package wafloganalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WafLogAnalyticsClient struct {
	Client *resourcemanager.Client
}

func NewWafLogAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*WafLogAnalyticsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "wafloganalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WafLogAnalyticsClient: %+v", err)
	}

	return &WafLogAnalyticsClient{
		Client: client,
	}, nil
}
