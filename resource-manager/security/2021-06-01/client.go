package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2021-06-01/assessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2021-06-01/assessmentsmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2021-06-01/settings"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Assessments         *assessments.AssessmentsClient
	AssessmentsMetadata *assessmentsmetadata.AssessmentsMetadataClient
	Settings            *settings.SettingsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	assessmentsClient, err := assessments.NewAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Assessments client: %+v", err)
	}
	configureFunc(assessmentsClient.Client)

	assessmentsMetadataClient, err := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssessmentsMetadata client: %+v", err)
	}
	configureFunc(assessmentsMetadataClient.Client)

	settingsClient, err := settings.NewSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %+v", err)
	}
	configureFunc(settingsClient.Client)

	return &Client{
		Assessments:         assessmentsClient,
		AssessmentsMetadata: assessmentsMetadataClient,
		Settings:            settingsClient,
	}, nil
}
