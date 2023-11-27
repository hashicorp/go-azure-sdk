package v2019_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2019-06-01/smartdetectoralertrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	SmartDetectorAlertRules *smartdetectoralertrules.SmartDetectorAlertRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	smartDetectorAlertRulesClient, err := smartdetectoralertrules.NewSmartDetectorAlertRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SmartDetectorAlertRules client: %+v", err)
	}
	configureFunc(smartDetectorAlertRulesClient.Client)

	return &Client{
		SmartDetectorAlertRules: smartDetectorAlertRulesClient,
	}, nil
}
