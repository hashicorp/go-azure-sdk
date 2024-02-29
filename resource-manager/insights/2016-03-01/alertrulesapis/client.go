package alertrulesapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRulesAPIsClient struct {
	Client *resourcemanager.Client
}

func NewAlertRulesAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*AlertRulesAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "alertrulesapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AlertRulesAPIsClient: %+v", err)
	}

	return &AlertRulesAPIsClient{
		Client: client,
	}, nil
}
