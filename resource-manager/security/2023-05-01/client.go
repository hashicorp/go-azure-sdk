package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-05-01/servervulnerabilityassessmentssettings"
)

type Client struct {
	ServerVulnerabilityAssessmentsSettings *servervulnerabilityassessmentssettings.ServerVulnerabilityAssessmentsSettingsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	serverVulnerabilityAssessmentsSettingsClient := servervulnerabilityassessmentssettings.NewServerVulnerabilityAssessmentsSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&serverVulnerabilityAssessmentsSettingsClient.Client)

	return Client{
		ServerVulnerabilityAssessmentsSettings: &serverVulnerabilityAssessmentsSettingsClient,
	}
}
