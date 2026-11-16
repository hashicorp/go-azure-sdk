package v2026_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/authenticationsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/discoveryrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/entities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/healthmodels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/relationships"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/signaldefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AuthenticationSettings *authenticationsettings.AuthenticationSettingsClient
	DiscoveryRules         *discoveryrules.DiscoveryRulesClient
	Entities               *entities.EntitiesClient
	HealthModels           *healthmodels.HealthModelsClient
	Relationships          *relationships.RelationshipsClient
	SignalDefinitions      *signaldefinitions.SignalDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	authenticationSettingsClient, err := authenticationsettings.NewAuthenticationSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationSettings client: %+v", err)
	}
	configureFunc(authenticationSettingsClient.Client)

	discoveryRulesClient, err := discoveryrules.NewDiscoveryRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiscoveryRules client: %+v", err)
	}
	configureFunc(discoveryRulesClient.Client)

	entitiesClient, err := entities.NewEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Entities client: %+v", err)
	}
	configureFunc(entitiesClient.Client)

	healthModelsClient, err := healthmodels.NewHealthModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HealthModels client: %+v", err)
	}
	configureFunc(healthModelsClient.Client)

	relationshipsClient, err := relationships.NewRelationshipsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Relationships client: %+v", err)
	}
	configureFunc(relationshipsClient.Client)

	signalDefinitionsClient, err := signaldefinitions.NewSignalDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SignalDefinitions client: %+v", err)
	}
	configureFunc(signalDefinitionsClient.Client)

	return &Client{
		AuthenticationSettings: authenticationSettingsClient,
		DiscoveryRules:         discoveryRulesClient,
		Entities:               entitiesClient,
		HealthModels:           healthModelsClient,
		Relationships:          relationshipsClient,
		SignalDefinitions:      signalDefinitionsClient,
	}, nil
}
