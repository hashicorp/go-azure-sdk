package alertruleincidents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRuleIncidentsClient struct {
	Client *resourcemanager.Client
}

func NewAlertRuleIncidentsClientWithBaseURI(sdkApi sdkEnv.Api) (*AlertRuleIncidentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "alertruleincidents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AlertRuleIncidentsClient: %+v", err)
	}

	return &AlertRuleIncidentsClient{
		Client: client,
	}, nil
}
