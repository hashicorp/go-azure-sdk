package v2022_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/actions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/alertruletemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/automationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/bookmark"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/bookmarkrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/bookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/checkdataconnectorrequirements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/dataconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/dataconnectorsconnect"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/dataconnectorsdisconnect"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/enrichment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/entities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/entityqueries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/entityrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/entitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/fileimports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidentalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidentbookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidentcomments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidententities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidentrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/incidentteam"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/manualtrigger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/officeconsents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/repositories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/securitymlanalyticssettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/sentinelonboardingstates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/settings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/sourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/threatintelligence"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/watchlistitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/watchlists"
)

type Client struct {
	Actions                        *actions.ActionsClient
	AlertRuleTemplates             *alertruletemplates.AlertRuleTemplatesClient
	AlertRules                     *alertrules.AlertRulesClient
	AutomationRules                *automationrules.AutomationRulesClient
	Bookmark                       *bookmark.BookmarkClient
	BookmarkRelations              *bookmarkrelations.BookmarkRelationsClient
	Bookmarks                      *bookmarks.BookmarksClient
	CheckDataConnectorRequirements *checkdataconnectorrequirements.CheckDataConnectorRequirementsClient
	DataConnectors                 *dataconnectors.DataConnectorsClient
	DataConnectorsConnect          *dataconnectorsconnect.DataConnectorsConnectClient
	DataConnectorsDisconnect       *dataconnectorsdisconnect.DataConnectorsDisconnectClient
	Enrichment                     *enrichment.EnrichmentClient
	Entities                       *entities.EntitiesClient
	EntityQueries                  *entityqueries.EntityQueriesClient
	EntityRelations                *entityrelations.EntityRelationsClient
	EntityTypes                    *entitytypes.EntityTypesClient
	FileImports                    *fileimports.FileImportsClient
	IncidentAlerts                 *incidentalerts.IncidentAlertsClient
	IncidentBookmarks              *incidentbookmarks.IncidentBookmarksClient
	IncidentComments               *incidentcomments.IncidentCommentsClient
	IncidentEntities               *incidententities.IncidentEntitiesClient
	IncidentRelations              *incidentrelations.IncidentRelationsClient
	IncidentTeam                   *incidentteam.IncidentTeamClient
	Incidents                      *incidents.IncidentsClient
	ManualTrigger                  *manualtrigger.ManualTriggerClient
	Metadata                       *metadata.MetadataClient
	OfficeConsents                 *officeconsents.OfficeConsentsClient
	Repositories                   *repositories.RepositoriesClient
	SecurityMLAnalyticsSettings    *securitymlanalyticssettings.SecurityMLAnalyticsSettingsClient
	SentinelOnboardingStates       *sentinelonboardingstates.SentinelOnboardingStatesClient
	Settings                       *settings.SettingsClient
	SourceControls                 *sourcecontrols.SourceControlsClient
	ThreatIntelligence             *threatintelligence.ThreatIntelligenceClient
	WatchlistItems                 *watchlistitems.WatchlistItemsClient
	Watchlists                     *watchlists.WatchlistsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionsClient := actions.NewActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionsClient.Client)

	alertRuleTemplatesClient := alertruletemplates.NewAlertRuleTemplatesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRuleTemplatesClient.Client)

	alertRulesClient := alertrules.NewAlertRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRulesClient.Client)

	automationRulesClient := automationrules.NewAutomationRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&automationRulesClient.Client)

	bookmarkClient := bookmark.NewBookmarkClientWithBaseURI(endpoint)
	configureAuthFunc(&bookmarkClient.Client)

	bookmarkRelationsClient := bookmarkrelations.NewBookmarkRelationsClientWithBaseURI(endpoint)
	configureAuthFunc(&bookmarkRelationsClient.Client)

	bookmarksClient := bookmarks.NewBookmarksClientWithBaseURI(endpoint)
	configureAuthFunc(&bookmarksClient.Client)

	checkDataConnectorRequirementsClient := checkdataconnectorrequirements.NewCheckDataConnectorRequirementsClientWithBaseURI(endpoint)
	configureAuthFunc(&checkDataConnectorRequirementsClient.Client)

	dataConnectorsClient := dataconnectors.NewDataConnectorsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataConnectorsClient.Client)

	dataConnectorsConnectClient := dataconnectorsconnect.NewDataConnectorsConnectClientWithBaseURI(endpoint)
	configureAuthFunc(&dataConnectorsConnectClient.Client)

	dataConnectorsDisconnectClient := dataconnectorsdisconnect.NewDataConnectorsDisconnectClientWithBaseURI(endpoint)
	configureAuthFunc(&dataConnectorsDisconnectClient.Client)

	enrichmentClient := enrichment.NewEnrichmentClientWithBaseURI(endpoint)
	configureAuthFunc(&enrichmentClient.Client)

	entitiesClient := entities.NewEntitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&entitiesClient.Client)

	entityQueriesClient := entityqueries.NewEntityQueriesClientWithBaseURI(endpoint)
	configureAuthFunc(&entityQueriesClient.Client)

	entityRelationsClient := entityrelations.NewEntityRelationsClientWithBaseURI(endpoint)
	configureAuthFunc(&entityRelationsClient.Client)

	entityTypesClient := entitytypes.NewEntityTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&entityTypesClient.Client)

	fileImportsClient := fileimports.NewFileImportsClientWithBaseURI(endpoint)
	configureAuthFunc(&fileImportsClient.Client)

	incidentAlertsClient := incidentalerts.NewIncidentAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentAlertsClient.Client)

	incidentBookmarksClient := incidentbookmarks.NewIncidentBookmarksClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentBookmarksClient.Client)

	incidentCommentsClient := incidentcomments.NewIncidentCommentsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentCommentsClient.Client)

	incidentEntitiesClient := incidententities.NewIncidentEntitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentEntitiesClient.Client)

	incidentRelationsClient := incidentrelations.NewIncidentRelationsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentRelationsClient.Client)

	incidentTeamClient := incidentteam.NewIncidentTeamClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentTeamClient.Client)

	incidentsClient := incidents.NewIncidentsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentsClient.Client)

	manualTriggerClient := manualtrigger.NewManualTriggerClientWithBaseURI(endpoint)
	configureAuthFunc(&manualTriggerClient.Client)

	metadataClient := metadata.NewMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&metadataClient.Client)

	officeConsentsClient := officeconsents.NewOfficeConsentsClientWithBaseURI(endpoint)
	configureAuthFunc(&officeConsentsClient.Client)

	repositoriesClient := repositories.NewRepositoriesClientWithBaseURI(endpoint)
	configureAuthFunc(&repositoriesClient.Client)

	securityMLAnalyticsSettingsClient := securitymlanalyticssettings.NewSecurityMLAnalyticsSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&securityMLAnalyticsSettingsClient.Client)

	sentinelOnboardingStatesClient := sentinelonboardingstates.NewSentinelOnboardingStatesClientWithBaseURI(endpoint)
	configureAuthFunc(&sentinelOnboardingStatesClient.Client)

	settingsClient := settings.NewSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&settingsClient.Client)

	sourceControlsClient := sourcecontrols.NewSourceControlsClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlsClient.Client)

	threatIntelligenceClient := threatintelligence.NewThreatIntelligenceClientWithBaseURI(endpoint)
	configureAuthFunc(&threatIntelligenceClient.Client)

	watchlistItemsClient := watchlistitems.NewWatchlistItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&watchlistItemsClient.Client)

	watchlistsClient := watchlists.NewWatchlistsClientWithBaseURI(endpoint)
	configureAuthFunc(&watchlistsClient.Client)

	return Client{
		Actions:                        &actionsClient,
		AlertRuleTemplates:             &alertRuleTemplatesClient,
		AlertRules:                     &alertRulesClient,
		AutomationRules:                &automationRulesClient,
		Bookmark:                       &bookmarkClient,
		BookmarkRelations:              &bookmarkRelationsClient,
		Bookmarks:                      &bookmarksClient,
		CheckDataConnectorRequirements: &checkDataConnectorRequirementsClient,
		DataConnectors:                 &dataConnectorsClient,
		DataConnectorsConnect:          &dataConnectorsConnectClient,
		DataConnectorsDisconnect:       &dataConnectorsDisconnectClient,
		Enrichment:                     &enrichmentClient,
		Entities:                       &entitiesClient,
		EntityQueries:                  &entityQueriesClient,
		EntityRelations:                &entityRelationsClient,
		EntityTypes:                    &entityTypesClient,
		FileImports:                    &fileImportsClient,
		IncidentAlerts:                 &incidentAlertsClient,
		IncidentBookmarks:              &incidentBookmarksClient,
		IncidentComments:               &incidentCommentsClient,
		IncidentEntities:               &incidentEntitiesClient,
		IncidentRelations:              &incidentRelationsClient,
		IncidentTeam:                   &incidentTeamClient,
		Incidents:                      &incidentsClient,
		ManualTrigger:                  &manualTriggerClient,
		Metadata:                       &metadataClient,
		OfficeConsents:                 &officeConsentsClient,
		Repositories:                   &repositoriesClient,
		SecurityMLAnalyticsSettings:    &securityMLAnalyticsSettingsClient,
		SentinelOnboardingStates:       &sentinelOnboardingStatesClient,
		Settings:                       &settingsClient,
		SourceControls:                 &sourceControlsClient,
		ThreatIntelligence:             &threatIntelligenceClient,
		WatchlistItems:                 &watchlistItemsClient,
		Watchlists:                     &watchlistsClient,
	}
}
