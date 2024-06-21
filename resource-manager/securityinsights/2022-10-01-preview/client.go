package v2022_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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
		Actions:                        actionsClient,
		AlertRuleTemplates:             alertRuleTemplatesClient,
		AlertRules:                     alertRulesClient,
		AutomationRules:                automationRulesClient,
		Bookmark:                       bookmarkClient,
		BookmarkRelations:              bookmarkRelationsClient,
		Bookmarks:                      bookmarksClient,
		CheckDataConnectorRequirements: checkDataConnectorRequirementsClient,
		DataConnectors:                 dataConnectorsClient,
		DataConnectorsConnect:          dataConnectorsConnectClient,
		DataConnectorsDisconnect:       dataConnectorsDisconnectClient,
		Enrichment:                     enrichmentClient,
		Entities:                       entitiesClient,
		EntityQueries:                  entityQueriesClient,
		EntityRelations:                entityRelationsClient,
		FileImports:                    fileImportsClient,
		IncidentAlerts:                 incidentAlertsClient,
		IncidentBookmarks:              incidentBookmarksClient,
		IncidentComments:               incidentCommentsClient,
		IncidentEntities:               incidentEntitiesClient,
		IncidentRelations:              incidentRelationsClient,
		IncidentTeam:                   incidentTeamClient,
		Incidents:                      incidentsClient,
		ManualTrigger:                  manualTriggerClient,
		Metadata:                       metadataClient,
		OfficeConsents:                 officeConsentsClient,
		Repositories:                   repositoriesClient,
		SecurityMLAnalyticsSettings:    securityMLAnalyticsSettingsClient,
		SentinelOnboardingStates:       sentinelOnboardingStatesClient,
		Settings:                       settingsClient,
		SourceControls:                 sourceControlsClient,
		ThreatIntelligence:             threatIntelligenceClient,
		WatchlistItems:                 watchlistItemsClient,
		Watchlists:                     watchlistsClient,
	}, nil
}
