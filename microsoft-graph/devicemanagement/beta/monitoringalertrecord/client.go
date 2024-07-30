package monitoringalertrecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringAlertRecordClient struct {
	Client *msgraph.Client
}

func NewMonitoringAlertRecordClientWithBaseURI(sdkApi sdkEnv.Api) (*MonitoringAlertRecordClient, error) {
	client, err := msgraph.NewClient(sdkApi, "monitoringalertrecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitoringAlertRecordClient: %+v", err)
	}

	return &MonitoringAlertRecordClient{
		Client: client,
	}, nil
}
