package metricsobjectfirewall

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricsObjectFirewallClient struct {
	Client *resourcemanager.Client
}

func NewMetricsObjectFirewallClientWithBaseURI(sdkApi sdkEnv.Api) (*MetricsObjectFirewallClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "metricsobjectfirewall", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MetricsObjectFirewallClient: %+v", err)
	}

	return &MetricsObjectFirewallClient{
		Client: client,
	}, nil
}
