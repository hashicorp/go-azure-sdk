package v2018_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/geographichierarchies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/heatmaps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/profiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/realusermetrics"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Endpoints             *endpoints.EndpointsClient
	GeographicHierarchies *geographichierarchies.GeographicHierarchiesClient
	HeatMaps              *heatmaps.HeatMapsClient
	Profiles              *profiles.ProfilesClient
	RealUserMetrics       *realusermetrics.RealUserMetricsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoints client: %+v", err)
	}
	configureFunc(endpointsClient.Client)

	geographicHierarchiesClient, err := geographichierarchies.NewGeographicHierarchiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GeographicHierarchies client: %+v", err)
	}
	configureFunc(geographicHierarchiesClient.Client)

	heatMapsClient, err := heatmaps.NewHeatMapsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HeatMaps client: %+v", err)
	}
	configureFunc(heatMapsClient.Client)

	profilesClient, err := profiles.NewProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Profiles client: %+v", err)
	}
	configureFunc(profilesClient.Client)

	realUserMetricsClient, err := realusermetrics.NewRealUserMetricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RealUserMetrics client: %+v", err)
	}
	configureFunc(realUserMetricsClient.Client)

	return &Client{
		Endpoints:             endpointsClient,
		GeographicHierarchies: geographicHierarchiesClient,
		HeatMaps:              heatMapsClient,
		Profiles:              profilesClient,
		RealUserMetrics:       realUserMetricsClient,
	}, nil
}
