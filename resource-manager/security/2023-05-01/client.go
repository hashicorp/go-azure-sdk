package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-05-01/servervulnerabilityassessmentssettings"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ServerVulnerabilityAssessmentsSettings *servervulnerabilityassessmentssettings.ServerVulnerabilityAssessmentsSettingsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	serverVulnerabilityAssessmentsSettingsClient, err := servervulnerabilityassessmentssettings.NewServerVulnerabilityAssessmentsSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerVulnerabilityAssessmentsSettings client: %+v", err)
	}
	configureFunc(serverVulnerabilityAssessmentsSettingsClient.Client)

	return &Client{
		ServerVulnerabilityAssessmentsSettings: serverVulnerabilityAssessmentsSettingsClient,
	}, nil
}
