package v2023_08_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/arcsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/cluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/deploymentsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/edgedevices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/offers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/publishers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/skuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/updateruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/updates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-08-01-preview/updatesummaries"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ArcSettings        *arcsettings.ArcSettingsClient
	Cluster            *cluster.ClusterClient
	Clusters           *clusters.ClustersClient
	DeploymentSettings *deploymentsettings.DeploymentSettingsClient
	EdgeDevices        *edgedevices.EdgeDevicesClient
	Extensions         *extensions.ExtensionsClient
	Offers             *offers.OffersClient
	Publishers         *publishers.PublishersClient
	Skuses             *skuses.SkusesClient
	UpdateRuns         *updateruns.UpdateRunsClient
	UpdateSummaries    *updatesummaries.UpdateSummariesClient
	Updates            *updates.UpdatesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	arcSettingsClient, err := arcsettings.NewArcSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ArcSettings client: %+v", err)
	}
	configureFunc(arcSettingsClient.Client)

	clusterClient, err := cluster.NewClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %+v", err)
	}
	configureFunc(clusterClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	deploymentSettingsClient, err := deploymentsettings.NewDeploymentSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentSettings client: %+v", err)
	}
	configureFunc(deploymentSettingsClient.Client)

	edgeDevicesClient, err := edgedevices.NewEdgeDevicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EdgeDevices client: %+v", err)
	}
	configureFunc(edgeDevicesClient.Client)

	extensionsClient, err := extensions.NewExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extensions client: %+v", err)
	}
	configureFunc(extensionsClient.Client)

	offersClient, err := offers.NewOffersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Offers client: %+v", err)
	}
	configureFunc(offersClient.Client)

	publishersClient, err := publishers.NewPublishersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Publishers client: %+v", err)
	}
	configureFunc(publishersClient.Client)

	skusesClient, err := skuses.NewSkusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skuses client: %+v", err)
	}
	configureFunc(skusesClient.Client)

	updateRunsClient, err := updateruns.NewUpdateRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateRuns client: %+v", err)
	}
	configureFunc(updateRunsClient.Client)

	updateSummariesClient, err := updatesummaries.NewUpdateSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateSummaries client: %+v", err)
	}
	configureFunc(updateSummariesClient.Client)

	updatesClient, err := updates.NewUpdatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Updates client: %+v", err)
	}
	configureFunc(updatesClient.Client)

	return &Client{
		ArcSettings:        arcSettingsClient,
		Cluster:            clusterClient,
		Clusters:           clustersClient,
		DeploymentSettings: deploymentSettingsClient,
		EdgeDevices:        edgeDevicesClient,
		Extensions:         extensionsClient,
		Offers:             offersClient,
		Publishers:         publishersClient,
		Skuses:             skusesClient,
		UpdateRuns:         updateRunsClient,
		UpdateSummaries:    updateSummariesClient,
		Updates:            updatesClient,
	}, nil
}
