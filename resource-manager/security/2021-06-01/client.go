package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2021-06-01/assessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2021-06-01/assessmentsmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2021-06-01/settings"
)

type Client struct {
	Assessments         *assessments.AssessmentsClient
	AssessmentsMetadata *assessmentsmetadata.AssessmentsMetadataClient
	Settings            *settings.SettingsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	assessmentsClient := assessments.NewAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&assessmentsClient.Client)

	assessmentsMetadataClient := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&assessmentsMetadataClient.Client)

	settingsClient := settings.NewSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&settingsClient.Client)

	return Client{
		Assessments:         &assessmentsClient,
		AssessmentsMetadata: &assessmentsMetadataClient,
		Settings:            &settingsClient,
	}
}
