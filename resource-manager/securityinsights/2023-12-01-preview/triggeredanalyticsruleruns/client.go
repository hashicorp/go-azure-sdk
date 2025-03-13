package triggeredanalyticsruleruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredAnalyticsRuleRunsClient struct {
	Client *resourcemanager.Client
}

func NewTriggeredAnalyticsRuleRunsClientWithBaseURI(sdkApi sdkEnv.Api) (*TriggeredAnalyticsRuleRunsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "triggeredanalyticsruleruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggeredAnalyticsRuleRunsClient: %+v", err)
	}

	return &TriggeredAnalyticsRuleRunsClient{
		Client: client,
	}, nil
}
