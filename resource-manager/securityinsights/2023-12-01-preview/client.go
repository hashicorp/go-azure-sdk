package v2023_12_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/actions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/alertruletemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/automationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/billingstatistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/bookmark"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/bookmarkrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/bookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/checkdataconnectorrequirements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/connectordefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/contentpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/contentproductpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/contentproducttemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/contenttemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/dataconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/dataconnectorsconnect"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/dataconnectorsdisconnect"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/enrichment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entityqueries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entityrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/fileimports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/huntcomments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/huntrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/hunts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentbookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentcomments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidententities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidenttasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentteam"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/manualtrigger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/officeconsents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/recommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/repositories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/securitymlanalyticssettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/sentinelonboardingstates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/settings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/sourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/threatintelligence"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeranalyticsrulerun"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeredanalyticsrulerun"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeredanalyticsruleruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/watchlistitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/watchlists"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagerassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagerconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagermember"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Actions                        *actions.ActionsClient
	AlertRuleTemplates             *alertruletemplates.AlertRuleTemplatesClient
	AlertRules                     *alertrules.AlertRulesClient
	AutomationRules                *automationrules.AutomationRulesClient
	BillingStatistics              *billingstatistics.BillingStatisticsClient
	Bookmark                       *bookmark.BookmarkClient
	BookmarkRelations              *bookmarkrelations.BookmarkRelationsClient
	Bookmarks                      *bookmarks.BookmarksClient
	CheckDataConnectorRequirements *checkdataconnectorrequirements.CheckDataConnectorRequirementsClient
	ConnectorDefinitions           *connectordefinitions.ConnectorDefinitionsClient
	ContentPackages                *contentpackages.ContentPackagesClient
	ContentProductPackages         *contentproductpackages.ContentProductPackagesClient
	ContentProductTemplates        *contentproducttemplates.ContentProductTemplatesClient
	ContentTemplates               *contenttemplates.ContentTemplatesClient
	DataConnectors                 *dataconnectors.DataConnectorsClient
	DataConnectorsConnect          *dataconnectorsconnect.DataConnectorsConnectClient
	DataConnectorsDisconnect       *dataconnectorsdisconnect.DataConnectorsDisconnectClient
	Enrichment                     *enrichment.EnrichmentClient
	Entities                       *entities.EntitiesClient
	EntityQueries                  *entityqueries.EntityQueriesClient
	EntityRelations                *entityrelations.EntityRelationsClient
	FileImports                    *fileimports.FileImportsClient
	HuntComments                   *huntcomments.HuntCommentsClient
	HuntRelations                  *huntrelations.HuntRelationsClient
	Hunts                          *hunts.HuntsClient
	IncidentAlerts                 *incidentalerts.IncidentAlertsClient
	IncidentBookmarks              *incidentbookmarks.IncidentBookmarksClient
	IncidentComments               *incidentcomments.IncidentCommentsClient
	IncidentEntities               *incidententities.IncidentEntitiesClient
	IncidentRelations              *incidentrelations.IncidentRelationsClient
	IncidentTasks                  *incidenttasks.IncidentTasksClient
	IncidentTeam                   *incidentteam.IncidentTeamClient
	Incidents                      *incidents.IncidentsClient
	ManualTrigger                  *manualtrigger.ManualTriggerClient
	Metadata                       *metadata.MetadataClient
	OfficeConsents                 *officeconsents.OfficeConsentsClient
	Recommendations                *recommendations.RecommendationsClient
	Repositories                   *repositories.RepositoriesClient
	SecurityMLAnalyticsSettings    *securitymlanalyticssettings.SecurityMLAnalyticsSettingsClient
	SentinelOnboardingStates       *sentinelonboardingstates.SentinelOnboardingStatesClient
	Settings                       *settings.SettingsClient
	SourceControls                 *sourcecontrols.SourceControlsClient
	ThreatIntelligence             *threatintelligence.ThreatIntelligenceClient
	TriggerAnalyticsRuleRun        *triggeranalyticsrulerun.TriggerAnalyticsRuleRunClient
	TriggeredAnalyticsRuleRun      *triggeredanalyticsrulerun.TriggeredAnalyticsRuleRunClient
	TriggeredAnalyticsRuleRuns     *triggeredanalyticsruleruns.TriggeredAnalyticsRuleRunsClient
	WatchlistItems                 *watchlistitems.WatchlistItemsClient
	Watchlists                     *watchlists.WatchlistsClient
	WorkspaceManagerAssignments    *workspacemanagerassignments.WorkspaceManagerAssignmentsClient
	WorkspaceManagerConfigurations *workspacemanagerconfigurations.WorkspaceManagerConfigurationsClient
	WorkspaceManagerGroups         *workspacemanagergroups.WorkspaceManagerGroupsClient
	WorkspaceManagerMember         *workspacemanagermember.WorkspaceManagerMemberClient
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

	billingStatisticsClient, err := billingstatistics.NewBillingStatisticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingStatistics client: %+v", err)
	}
	configureFunc(billingStatisticsClient.Client)

	bookmarkClient, err := bookmark.NewBookmarkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Bookmark client: %+v", err)
	}
	configureFunc(bookmarkClient.Client)

	bookmarkRelationsClient, err := bookmarkrelations.NewBookmarkRelationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BookmarkRelations client: %+v", err)
	}
	configureFunc(bookmarkRelationsClient.Client)

	bookmarksClient, err := bookmarks.NewBookmarksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Bookmarks client: %+v", err)
	}
	configureFunc(bookmarksClient.Client)

	checkDataConnectorRequirementsClient, err := checkdataconnectorrequirements.NewCheckDataConnectorRequirementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckDataConnectorRequirements client: %+v", err)
	}
	configureFunc(checkDataConnectorRequirementsClient.Client)

	connectorDefinitionsClient, err := connectordefinitions.NewConnectorDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectorDefinitions client: %+v", err)
	}
	configureFunc(connectorDefinitionsClient.Client)

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

	dataConnectorsConnectClient, err := dataconnectorsconnect.NewDataConnectorsConnectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataConnectorsConnect client: %+v", err)
	}
	configureFunc(dataConnectorsConnectClient.Client)

	dataConnectorsDisconnectClient, err := dataconnectorsdisconnect.NewDataConnectorsDisconnectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataConnectorsDisconnect client: %+v", err)
	}
	configureFunc(dataConnectorsDisconnectClient.Client)

	enrichmentClient, err := enrichment.NewEnrichmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Enrichment client: %+v", err)
	}
	configureFunc(enrichmentClient.Client)

	entitiesClient, err := entities.NewEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Entities client: %+v", err)
	}
	configureFunc(entitiesClient.Client)

	entityQueriesClient, err := entityqueries.NewEntityQueriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntityQueries client: %+v", err)
	}
	configureFunc(entityQueriesClient.Client)

	entityRelationsClient, err := entityrelations.NewEntityRelationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntityRelations client: %+v", err)
	}
	configureFunc(entityRelationsClient.Client)

	fileImportsClient, err := fileimports.NewFileImportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileImports client: %+v", err)
	}
	configureFunc(fileImportsClient.Client)

	huntCommentsClient, err := huntcomments.NewHuntCommentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HuntComments client: %+v", err)
	}
	configureFunc(huntCommentsClient.Client)

	huntRelationsClient, err := huntrelations.NewHuntRelationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HuntRelations client: %+v", err)
	}
	configureFunc(huntRelationsClient.Client)

	huntsClient, err := hunts.NewHuntsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Hunts client: %+v", err)
	}
	configureFunc(huntsClient.Client)

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

	incidentTeamClient, err := incidentteam.NewIncidentTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentTeam client: %+v", err)
	}
	configureFunc(incidentTeamClient.Client)

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

	officeConsentsClient, err := officeconsents.NewOfficeConsentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OfficeConsents client: %+v", err)
	}
	configureFunc(officeConsentsClient.Client)

	recommendationsClient, err := recommendations.NewRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Recommendations client: %+v", err)
	}
	configureFunc(recommendationsClient.Client)

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

	settingsClient, err := settings.NewSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %+v", err)
	}
	configureFunc(settingsClient.Client)

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

	triggerAnalyticsRuleRunClient, err := triggeranalyticsrulerun.NewTriggerAnalyticsRuleRunClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggerAnalyticsRuleRun client: %+v", err)
	}
	configureFunc(triggerAnalyticsRuleRunClient.Client)

	triggeredAnalyticsRuleRunClient, err := triggeredanalyticsrulerun.NewTriggeredAnalyticsRuleRunClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggeredAnalyticsRuleRun client: %+v", err)
	}
	configureFunc(triggeredAnalyticsRuleRunClient.Client)

	triggeredAnalyticsRuleRunsClient, err := triggeredanalyticsruleruns.NewTriggeredAnalyticsRuleRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggeredAnalyticsRuleRuns client: %+v", err)
	}
	configureFunc(triggeredAnalyticsRuleRunsClient.Client)

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

	workspaceManagerAssignmentsClient, err := workspacemanagerassignments.NewWorkspaceManagerAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagerAssignments client: %+v", err)
	}
	configureFunc(workspaceManagerAssignmentsClient.Client)

	workspaceManagerConfigurationsClient, err := workspacemanagerconfigurations.NewWorkspaceManagerConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagerConfigurations client: %+v", err)
	}
	configureFunc(workspaceManagerConfigurationsClient.Client)

	workspaceManagerGroupsClient, err := workspacemanagergroups.NewWorkspaceManagerGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagerGroups client: %+v", err)
	}
	configureFunc(workspaceManagerGroupsClient.Client)

	workspaceManagerMemberClient, err := workspacemanagermember.NewWorkspaceManagerMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagerMember client: %+v", err)
	}
	configureFunc(workspaceManagerMemberClient.Client)

	return &Client{
		Actions:                        actionsClient,
		AlertRuleTemplates:             alertRuleTemplatesClient,
		AlertRules:                     alertRulesClient,
		AutomationRules:                automationRulesClient,
		BillingStatistics:              billingStatisticsClient,
		Bookmark:                       bookmarkClient,
		BookmarkRelations:              bookmarkRelationsClient,
		Bookmarks:                      bookmarksClient,
		CheckDataConnectorRequirements: checkDataConnectorRequirementsClient,
		ConnectorDefinitions:           connectorDefinitionsClient,
		ContentPackages:                contentPackagesClient,
		ContentProductPackages:         contentProductPackagesClient,
		ContentProductTemplates:        contentProductTemplatesClient,
		ContentTemplates:               contentTemplatesClient,
		DataConnectors:                 dataConnectorsClient,
		DataConnectorsConnect:          dataConnectorsConnectClient,
		DataConnectorsDisconnect:       dataConnectorsDisconnectClient,
		Enrichment:                     enrichmentClient,
		Entities:                       entitiesClient,
		EntityQueries:                  entityQueriesClient,
		EntityRelations:                entityRelationsClient,
		FileImports:                    fileImportsClient,
		HuntComments:                   huntCommentsClient,
		HuntRelations:                  huntRelationsClient,
		Hunts:                          huntsClient,
		IncidentAlerts:                 incidentAlertsClient,
		IncidentBookmarks:              incidentBookmarksClient,
		IncidentComments:               incidentCommentsClient,
		IncidentEntities:               incidentEntitiesClient,
		IncidentRelations:              incidentRelationsClient,
		IncidentTasks:                  incidentTasksClient,
		IncidentTeam:                   incidentTeamClient,
		Incidents:                      incidentsClient,
		ManualTrigger:                  manualTriggerClient,
		Metadata:                       metadataClient,
		OfficeConsents:                 officeConsentsClient,
		Recommendations:                recommendationsClient,
		Repositories:                   repositoriesClient,
		SecurityMLAnalyticsSettings:    securityMLAnalyticsSettingsClient,
		SentinelOnboardingStates:       sentinelOnboardingStatesClient,
		Settings:                       settingsClient,
		SourceControls:                 sourceControlsClient,
		ThreatIntelligence:             threatIntelligenceClient,
		TriggerAnalyticsRuleRun:        triggerAnalyticsRuleRunClient,
		TriggeredAnalyticsRuleRun:      triggeredAnalyticsRuleRunClient,
		TriggeredAnalyticsRuleRuns:     triggeredAnalyticsRuleRunsClient,
		WatchlistItems:                 watchlistItemsClient,
		Watchlists:                     watchlistsClient,
		WorkspaceManagerAssignments:    workspaceManagerAssignmentsClient,
		WorkspaceManagerConfigurations: workspaceManagerConfigurationsClient,
		WorkspaceManagerGroups:         workspaceManagerGroupsClient,
		WorkspaceManagerMember:         workspaceManagerMemberClient,
	}, nil
}
