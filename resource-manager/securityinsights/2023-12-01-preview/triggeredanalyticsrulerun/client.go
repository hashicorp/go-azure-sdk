package triggeredanalyticsrulerun

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredAnalyticsRuleRunClient struct {
	Client *resourcemanager.Client
}

func NewTriggeredAnalyticsRuleRunClientWithBaseURI(sdkApi sdkEnv.Api) (*TriggeredAnalyticsRuleRunClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "triggeredanalyticsrulerun", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggeredAnalyticsRuleRunClient: %+v", err)
	}

	return &TriggeredAnalyticsRuleRunClient{
		Client: client,
	}, nil
}
