package monitoringalertrule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringAlertRuleClient struct {
	Client *msgraph.Client
}

func NewMonitoringAlertRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*MonitoringAlertRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "monitoringalertrule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitoringAlertRuleClient: %+v", err)
	}

	return &MonitoringAlertRuleClient{
		Client: client,
	}, nil
}
