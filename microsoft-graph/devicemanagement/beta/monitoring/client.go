package monitoring

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringClient struct {
	Client *msgraph.Client
}

func NewMonitoringClientWithBaseURI(sdkApi sdkEnv.Api) (*MonitoringClient, error) {
	client, err := msgraph.NewClient(sdkApi, "monitoring", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitoringClient: %+v", err)
	}

	return &MonitoringClient{
		Client: client,
	}, nil
}
