package v2024_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/actions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/alertruletemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/automationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/bookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/contentpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/contentproductpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/contentproducttemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/contenttemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/dataconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidentalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidentbookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidentcomments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidententities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidentrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidenttasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/manualtrigger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/repositories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/securitymlanalyticssettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/sentinelonboardingstates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/sourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/threatintelligence"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/watchlistitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/watchlists"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Actions                     *actions.ActionsClient
	AlertRuleTemplates          *alertruletemplates.AlertRuleTemplatesClient
	AlertRules                  *alertrules.AlertRulesClient
	AutomationRules             *automationrules.AutomationRulesClient
	Bookmarks                   *bookmarks.BookmarksClient
	ContentPackages             *contentpackages.ContentPackagesClient
	ContentProductPackages      *contentproductpackages.ContentProductPackagesClient
	ContentProductTemplates     *contentproducttemplates.ContentProductTemplatesClient
	ContentTemplates            *contenttemplates.ContentTemplatesClient
	DataConnectors              *dataconnectors.DataConnectorsClient
	IncidentAlerts              *incidentalerts.IncidentAlertsClient
	IncidentBookmarks           *incidentbookmarks.IncidentBookmarksClient
	IncidentComments            *incidentcomments.IncidentCommentsClient
	IncidentEntities            *incidententities.IncidentEntitiesClient
	IncidentRelations           *incidentrelations.IncidentRelationsClient
	IncidentTasks               *incidenttasks.IncidentTasksClient
	Incidents                   *incidents.IncidentsClient
	ManualTrigger               *manualtrigger.ManualTriggerClient
	Metadata                    *metadata.MetadataClient
	Repositories                *repositories.RepositoriesClient
	SecurityMLAnalyticsSettings *securitymlanalyticssettings.SecurityMLAnalyticsSettingsClient
	SentinelOnboardingStates    *sentinelonboardingstates.SentinelOnboardingStatesClient
	SourceControls              *sourcecontrols.SourceControlsClient
	ThreatIntelligence          *threatintelligence.ThreatIntelligenceClient
	WatchlistItems              *watchlistitems.WatchlistItemsClient
	Watchlists                  *watchlists.WatchlistsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionsClient, err := actions.NewActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Actions client: %+v", err)
	}
	configureFunc(actionsClient.Client)

	alertRuleTemplatesClient, err := alertruletemplates.NewAlertRuleTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRuleTemplates client: %+v", err)
	}
	configureFunc(alertRuleTemplatesClient.Client)

	alertRulesClient, err := alertrules.NewAlertRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRules client: %+v", err)
	}
	configureFunc(alertRulesClient.Client)

	automationRulesClient, err := automationrules.NewAutomationRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutomationRules client: %+v", err)
	}
	configureFunc(automationRulesClient.Client)

	bookmarksClient, err := bookmarks.NewBookmarksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Bookmarks client: %+v", err)
	}
	configureFunc(bookmarksClient.Client)

	contentPackagesClient, err := contentpackages.NewContentPackagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContentPackages client: %+v", err)
	}
	configureFunc(contentPackagesClient.Client)

	contentProductPackagesClient, err := contentproductpackages.NewContentProductPackagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContentProductPackages client: %+v", err)
	}
	configureFunc(contentProductPackagesClient.Client)

	contentProductTemplatesClient, err := contentproducttemplates.NewContentProductTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContentProductTemplates client: %+v", err)
	}
	configureFunc(contentProductTemplatesClient.Client)

	contentTemplatesClient, err := contenttemplates.NewContentTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContentTemplates client: %+v", err)
	}
	configureFunc(contentTemplatesClient.Client)

	dataConnectorsClient, err := dataconnectors.NewDataConnectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataConnectors client: %+v", err)
	}
	configureFunc(dataConnectorsClient.Client)

	incidentAlertsClient, err := incidentalerts.NewIncidentAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentAlerts client: %+v", err)
	}
	configureFunc(incidentAlertsClient.Client)

	incidentBookmarksClient, err := incidentbookmarks.NewIncidentBookmarksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentBookmarks client: %+v", err)
	}
	configureFunc(incidentBookmarksClient.Client)

	incidentCommentsClient, err := incidentcomments.NewIncidentCommentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentComments client: %+v", err)
	}
	configureFunc(incidentCommentsClient.Client)

	incidentEntitiesClient, err := incidententities.NewIncidentEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentEntities client: %+v", err)
	}
	configureFunc(incidentEntitiesClient.Client)

	incidentRelationsClient, err := incidentrelations.NewIncidentRelationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentRelations client: %+v", err)
	}
	configureFunc(incidentRelationsClient.Client)

	incidentTasksClient, err := incidenttasks.NewIncidentTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentTasks client: %+v", err)
	}
	configureFunc(incidentTasksClient.Client)

	incidentsClient, err := incidents.NewIncidentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Incidents client: %+v", err)
	}
	configureFunc(incidentsClient.Client)

	manualTriggerClient, err := manualtrigger.NewManualTriggerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManualTrigger client: %+v", err)
	}
	configureFunc(manualTriggerClient.Client)

	metadataClient, err := metadata.NewMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Metadata client: %+v", err)
	}
	configureFunc(metadataClient.Client)

	repositoriesClient, err := repositories.NewRepositoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Repositories client: %+v", err)
	}
	configureFunc(repositoriesClient.Client)

	securityMLAnalyticsSettingsClient, err := securitymlanalyticssettings.NewSecurityMLAnalyticsSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityMLAnalyticsSettings client: %+v", err)
	}
	configureFunc(securityMLAnalyticsSettingsClient.Client)

	sentinelOnboardingStatesClient, err := sentinelonboardingstates.NewSentinelOnboardingStatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SentinelOnboardingStates client: %+v", err)
	}
	configureFunc(sentinelOnboardingStatesClient.Client)

	sourceControlsClient, err := sourcecontrols.NewSourceControlsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SourceControls client: %+v", err)
	}
	configureFunc(sourceControlsClient.Client)

	threatIntelligenceClient, err := threatintelligence.NewThreatIntelligenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreatIntelligence client: %+v", err)
	}
	configureFunc(threatIntelligenceClient.Client)

	watchlistItemsClient, err := watchlistitems.NewWatchlistItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WatchlistItems client: %+v", err)
	}
	configureFunc(watchlistItemsClient.Client)

	watchlistsClient, err := watchlists.NewWatchlistsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Watchlists client: %+v", err)
	}
	configureFunc(watchlistsClient.Client)

	return &Client{
		Actions:                     actionsClient,
		AlertRuleTemplates:          alertRuleTemplatesClient,
		AlertRules:                  alertRulesClient,
		AutomationRules:             automationRulesClient,
		Bookmarks:                   bookmarksClient,
		ContentPackages:             contentPackagesClient,
		ContentProductPackages:      contentProductPackagesClient,
		ContentProductTemplates:     contentProductTemplatesClient,
		ContentTemplates:            contentTemplatesClient,
		DataConnectors:              dataConnectorsClient,
		IncidentAlerts:              incidentAlertsClient,
		IncidentBookmarks:           incidentBookmarksClient,
		IncidentComments:            incidentCommentsClient,
		IncidentEntities:            incidentEntitiesClient,
		IncidentRelations:           incidentRelationsClient,
		IncidentTasks:               incidentTasksClient,
		Incidents:                   incidentsClient,
		ManualTrigger:               manualTriggerClient,
		Metadata:                    metadataClient,
		Repositories:                repositoriesClient,
		SecurityMLAnalyticsSettings: securityMLAnalyticsSettingsClient,
		SentinelOnboardingStates:    sentinelOnboardingStatesClient,
		SourceControls:              sourceControlsClient,
		ThreatIntelligence:          threatIntelligenceClient,
		WatchlistItems:              watchlistItemsClient,
		Watchlists:                  watchlistsClient,
	}, nil
}
