package healthmonitoring

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringClient struct {
	Client *msgraph.Client
}

func NewHealthMonitoringClientWithBaseURI(sdkApi sdkEnv.Api) (*HealthMonitoringClient, error) {
	client, err := msgraph.NewClient(sdkApi, "healthmonitoring", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HealthMonitoringClient: %+v", err)
	}

	return &HealthMonitoringClient{
		Client: client,
	}, nil
}
