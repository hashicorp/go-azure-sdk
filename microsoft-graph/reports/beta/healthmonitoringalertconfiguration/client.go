package healthmonitoringalertconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringAlertConfigurationClient struct {
	Client *msgraph.Client
}

func NewHealthMonitoringAlertConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*HealthMonitoringAlertConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "healthmonitoringalertconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HealthMonitoringAlertConfigurationClient: %+v", err)
	}

	return &HealthMonitoringAlertConfigurationClient{
		Client: client,
	}, nil
}
