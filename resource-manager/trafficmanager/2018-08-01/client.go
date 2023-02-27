// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2018_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/geographichierarchies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/heatmaps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/profiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/realusermetrics"
)

type Client struct {
	Endpoints             *endpoints.EndpointsClient
	GeographicHierarchies *geographichierarchies.GeographicHierarchiesClient
	HeatMaps              *heatmaps.HeatMapsClient
	Profiles              *profiles.ProfilesClient
	RealUserMetrics       *realusermetrics.RealUserMetricsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	endpointsClient := endpoints.NewEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&endpointsClient.Client)

	geographicHierarchiesClient := geographichierarchies.NewGeographicHierarchiesClientWithBaseURI(endpoint)
	configureAuthFunc(&geographicHierarchiesClient.Client)

	heatMapsClient := heatmaps.NewHeatMapsClientWithBaseURI(endpoint)
	configureAuthFunc(&heatMapsClient.Client)

	profilesClient := profiles.NewProfilesClientWithBaseURI(endpoint)
	configureAuthFunc(&profilesClient.Client)

	realUserMetricsClient := realusermetrics.NewRealUserMetricsClientWithBaseURI(endpoint)
	configureAuthFunc(&realUserMetricsClient.Client)

	return Client{
		Endpoints:             &endpointsClient,
		GeographicHierarchies: &geographicHierarchiesClient,
		HeatMaps:              &heatMapsClient,
		Profiles:              &profilesClient,
		RealUserMetrics:       &realUserMetricsClient,
	}
}
