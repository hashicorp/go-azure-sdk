package v2025_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/addressresponses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/analysisdefinitionoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/analysisdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceenvironmentresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appsettingkeyvaultreference"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appsettingkeyvaultreferenceslot"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/asev3networkingconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/backupitemoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/backupitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/bysiteslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/certificateoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/continuouswebjoboperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/continuouswebjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmdeploymentstatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmdeploymentstatusoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentityftpallowedslot"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentityoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentityscmallowedslot"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/customdnssuffixconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/databaseconnectionoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/databaseconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deletedsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deploymentoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectordefinitionresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectordefinitionresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectorresponseoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectorresponses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnosticcategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnosticcategoryoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/functionenvelopeoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/functionenvelopes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/global"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hostnamebindingoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hostnamebindings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionlimitsoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/identifieroperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/identifiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instancemsdeploystatusoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instanceprocessmoduleslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instanceprocessslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/kubeenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/migratemysqlstatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/migratemysqlstatusoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatusoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatusslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/networkfeaturesoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/networkfeaturesslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/premieraddonoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/premieraddons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/privateaccesses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/privateaccessoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/privateendpointconnectionslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processinfooperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processinfos"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleinfooperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleinfos"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processslotoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/publiccertificateoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/publiccertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/recommendationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/recommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/relayserviceconnectionentities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/relayserviceconnectionentityoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/remoteprivateendpointconnectionarmresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/remoteprivateendpointconnectionarmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/requesthistories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/resourcehealthmetadataoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteauthsettingsv2operationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteauthsettingsv2s"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitecertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigslotresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigsnapshotslotresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconnectionstringkeyvaultreference"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconnectionstringkeyvaultreferenceslot"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitecontaineroperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitecontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteextensioninfooperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteextensioninfos"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitelogsconfigoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitelogsconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitesourcecontroloperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitesourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/slotconfignamesresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitearmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitebasicauthpropertiesarmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitebuildarmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitecustomdomainoverviewarmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitelinkedbackendarmresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitelinkedbackendarmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsiteuserprovidedfunctionapparmresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsiteuserprovidedfunctionapparmresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/swiftvirtualnetworkoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/swiftvirtualnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredjobhistories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredjobhistoryoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredwebjoboperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredwebjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/users"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetconnectiongatewayoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetconnectionoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetgatewayoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetinforesourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetinforesources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetroutes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/webapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/webjoboperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/webjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/websiteinstancestatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/websiteinstancestatusoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workerpoolresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowenvelopeoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowenvelopes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactionrepetitiondefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactionscoperepetitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowtriggerhistories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowtriggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowversions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AddressResponses                                           *addressresponses.AddressResponsesClient
	AnalysisDefinitionOperationGroup                           *analysisdefinitionoperationgroup.AnalysisDefinitionOperationGroupClient
	AnalysisDefinitions                                        *analysisdefinitions.AnalysisDefinitionsClient
	AppServiceEnvironmentResources                             *appserviceenvironmentresources.AppServiceEnvironmentResourcesClient
	AppServiceEnvironments                                     *appserviceenvironments.AppServiceEnvironmentsClient
	AppServicePlans                                            *appserviceplans.AppServicePlansClient
	AppSettingKeyVaultReference                                *appsettingkeyvaultreference.AppSettingKeyVaultReferenceClient
	AppSettingKeyVaultReferenceSlot                            *appsettingkeyvaultreferenceslot.AppSettingKeyVaultReferenceSlotClient
	AseV3NetworkingConfigurations                              *asev3networkingconfigurations.AseV3NetworkingConfigurationsClient
	BackupItemOperationGroup                                   *backupitemoperationgroup.BackupItemOperationGroupClient
	BackupItems                                                *backupitems.BackupItemsClient
	BySiteSlotOperationGroup                                   *bysiteslotoperationgroup.BySiteSlotOperationGroupClient
	CertificateOperationGroup                                  *certificateoperationgroup.CertificateOperationGroupClient
	Certificates                                               *certificates.CertificatesClient
	ContinuousWebJobOperationGroup                             *continuouswebjoboperationgroup.ContinuousWebJobOperationGroupClient
	ContinuousWebJobs                                          *continuouswebjobs.ContinuousWebJobsClient
	CsmDeploymentStatusOperationGroup                          *csmdeploymentstatusoperationgroup.CsmDeploymentStatusOperationGroupClient
	CsmDeploymentStatuses                                      *csmdeploymentstatuses.CsmDeploymentStatusesClient
	CsmPublishingCredentialsPoliciesEntities                   *csmpublishingcredentialspoliciesentities.CsmPublishingCredentialsPoliciesEntitiesClient
	CsmPublishingCredentialsPoliciesEntityFtpAllowedSlot       *csmpublishingcredentialspoliciesentityftpallowedslot.CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient
	CsmPublishingCredentialsPoliciesEntityOperationGroup       *csmpublishingcredentialspoliciesentityoperationgroup.CsmPublishingCredentialsPoliciesEntityOperationGroupClient
	CsmPublishingCredentialsPoliciesEntityScmAllowedSlot       *csmpublishingcredentialspoliciesentityscmallowedslot.CsmPublishingCredentialsPoliciesEntityScmAllowedSlotClient
	CustomDnsSuffixConfigurations                              *customdnssuffixconfigurations.CustomDnsSuffixConfigurationsClient
	DatabaseConnectionOperationGroup                           *databaseconnectionoperationgroup.DatabaseConnectionOperationGroupClient
	DatabaseConnections                                        *databaseconnections.DatabaseConnectionsClient
	DeletedSites                                               *deletedsites.DeletedSitesClient
	DeploymentOperationGroup                                   *deploymentoperationgroup.DeploymentOperationGroupClient
	Deployments                                                *deployments.DeploymentsClient
	DetectorDefinitionResourceOperationGroup                   *detectordefinitionresourceoperationgroup.DetectorDefinitionResourceOperationGroupClient
	DetectorDefinitionResources                                *detectordefinitionresources.DetectorDefinitionResourcesClient
	DetectorResponseOperationGroup                             *detectorresponseoperationgroup.DetectorResponseOperationGroupClient
	DetectorResponses                                          *detectorresponses.DetectorResponsesClient
	DiagnosticCategories                                       *diagnosticcategories.DiagnosticCategoriesClient
	DiagnosticCategoryOperationGroup                           *diagnosticcategoryoperationgroup.DiagnosticCategoryOperationGroupClient
	Diagnostics                                                *diagnostics.DiagnosticsClient
	FunctionEnvelopeOperationGroup                             *functionenvelopeoperationgroup.FunctionEnvelopeOperationGroupClient
	FunctionEnvelopes                                          *functionenvelopes.FunctionEnvelopesClient
	Global                                                     *global.GlobalClient
	HostNameBindingOperationGroup                              *hostnamebindingoperationgroup.HostNameBindingOperationGroupClient
	HostNameBindings                                           *hostnamebindings.HostNameBindingsClient
	HybridConnectionLimitsOperationGroup                       *hybridconnectionlimitsoperationgroup.HybridConnectionLimitsOperationGroupClient
	HybridConnectionOperationGroup                             *hybridconnectionoperationgroup.HybridConnectionOperationGroupClient
	HybridConnectionSlotOperationGroup                         *hybridconnectionslotoperationgroup.HybridConnectionSlotOperationGroupClient
	HybridConnections                                          *hybridconnections.HybridConnectionsClient
	IdentifierOperationGroup                                   *identifieroperationgroup.IdentifierOperationGroupClient
	Identifiers                                                *identifiers.IdentifiersClient
	InstanceMSDeployStatusOperationGroup                       *instancemsdeploystatusoperationgroup.InstanceMSDeployStatusOperationGroupClient
	InstanceProcessModuleSlotOperationGroup                    *instanceprocessmoduleslotoperationgroup.InstanceProcessModuleSlotOperationGroupClient
	InstanceProcessSlotOperationGroup                          *instanceprocessslotoperationgroup.InstanceProcessSlotOperationGroupClient
	KubeEnvironments                                           *kubeenvironments.KubeEnvironmentsClient
	MSDeployStatusOperationGroup                               *msdeploystatusoperationgroup.MSDeployStatusOperationGroupClient
	MSDeployStatusSlotOperationGroup                           *msdeploystatusslotoperationgroup.MSDeployStatusSlotOperationGroupClient
	MSDeployStatuses                                           *msdeploystatuses.MSDeployStatusesClient
	MigrateMySqlStatusOperationGroup                           *migratemysqlstatusoperationgroup.MigrateMySqlStatusOperationGroupClient
	MigrateMySqlStatuses                                       *migratemysqlstatuses.MigrateMySqlStatusesClient
	NetworkFeaturesOperationGroup                              *networkfeaturesoperationgroup.NetworkFeaturesOperationGroupClient
	NetworkFeaturesSlotOperationGroup                          *networkfeaturesslotoperationgroup.NetworkFeaturesSlotOperationGroupClient
	Openapis                                                   *openapis.OpenapisClient
	PremierAddOnOperationGroup                                 *premieraddonoperationgroup.PremierAddOnOperationGroupClient
	PremierAddons                                              *premieraddons.PremierAddonsClient
	PrivateAccessOperationGroup                                *privateaccessoperationgroup.PrivateAccessOperationGroupClient
	PrivateAccesses                                            *privateaccesses.PrivateAccessesClient
	PrivateEndpointConnectionSlotOperationGroup                *privateendpointconnectionslotoperationgroup.PrivateEndpointConnectionSlotOperationGroupClient
	ProcessInfoOperationGroup                                  *processinfooperationgroup.ProcessInfoOperationGroupClient
	ProcessInfos                                               *processinfos.ProcessInfosClient
	ProcessModuleInfoOperationGroup                            *processmoduleinfooperationgroup.ProcessModuleInfoOperationGroupClient
	ProcessModuleInfos                                         *processmoduleinfos.ProcessModuleInfosClient
	ProcessModuleSlotOperationGroup                            *processmoduleslotoperationgroup.ProcessModuleSlotOperationGroupClient
	ProcessSlotOperationGroup                                  *processslotoperationgroup.ProcessSlotOperationGroupClient
	PublicCertificateOperationGroup                            *publiccertificateoperationgroup.PublicCertificateOperationGroupClient
	PublicCertificates                                         *publiccertificates.PublicCertificatesClient
	RecommendationRules                                        *recommendationrules.RecommendationRulesClient
	Recommendations                                            *recommendations.RecommendationsClient
	RelayServiceConnectionEntities                             *relayserviceconnectionentities.RelayServiceConnectionEntitiesClient
	RelayServiceConnectionEntityOperationGroup                 *relayserviceconnectionentityoperationgroup.RelayServiceConnectionEntityOperationGroupClient
	RemotePrivateEndpointConnectionARMResourceOperationGroup   *remoteprivateendpointconnectionarmresourceoperationgroup.RemotePrivateEndpointConnectionARMResourceOperationGroupClient
	RemotePrivateEndpointConnectionARMResources                *remoteprivateendpointconnectionarmresources.RemotePrivateEndpointConnectionARMResourcesClient
	RequestHistories                                           *requesthistories.RequestHistoriesClient
	ResourceHealthMetadataOperationGroup                       *resourcehealthmetadataoperationgroup.ResourceHealthMetadataOperationGroupClient
	SiteAuthSettingsV2OperationGroup                           *siteauthsettingsv2operationgroup.SiteAuthSettingsV2OperationGroupClient
	SiteAuthSettingsV2s                                        *siteauthsettingsv2s.SiteAuthSettingsV2sClient
	SiteCertificates                                           *sitecertificates.SiteCertificatesClient
	SiteConfigResourceOperationGroup                           *siteconfigresourceoperationgroup.SiteConfigResourceOperationGroupClient
	SiteConfigResources                                        *siteconfigresources.SiteConfigResourcesClient
	SiteConfigSlotResourceOperationGroup                       *siteconfigslotresourceoperationgroup.SiteConfigSlotResourceOperationGroupClient
	SiteConfigSnapshotSlotResourceOperationGroup               *siteconfigsnapshotslotresourceoperationgroup.SiteConfigSnapshotSlotResourceOperationGroupClient
	SiteConnectionStringKeyVaultReference                      *siteconnectionstringkeyvaultreference.SiteConnectionStringKeyVaultReferenceClient
	SiteConnectionStringKeyVaultReferenceSlot                  *siteconnectionstringkeyvaultreferenceslot.SiteConnectionStringKeyVaultReferenceSlotClient
	SiteContainerOperationGroup                                *sitecontaineroperationgroup.SiteContainerOperationGroupClient
	SiteContainers                                             *sitecontainers.SiteContainersClient
	SiteExtensionInfoOperationGroup                            *siteextensioninfooperationgroup.SiteExtensionInfoOperationGroupClient
	SiteExtensionInfos                                         *siteextensioninfos.SiteExtensionInfosClient
	SiteLogsConfigOperationGroup                               *sitelogsconfigoperationgroup.SiteLogsConfigOperationGroupClient
	SiteLogsConfigs                                            *sitelogsconfigs.SiteLogsConfigsClient
	SiteSourceControlOperationGroup                            *sitesourcecontroloperationgroup.SiteSourceControlOperationGroupClient
	SiteSourceControls                                         *sitesourcecontrols.SiteSourceControlsClient
	Sites                                                      *sites.SitesClient
	SlotConfigNamesResources                                   *slotconfignamesresources.SlotConfigNamesResourcesClient
	SourceControls                                             *sourcecontrols.SourceControlsClient
	StaticSiteARMResources                                     *staticsitearmresources.StaticSiteARMResourcesClient
	StaticSiteBasicAuthPropertiesARMResources                  *staticsitebasicauthpropertiesarmresources.StaticSiteBasicAuthPropertiesARMResourcesClient
	StaticSiteBuildARMResources                                *staticsitebuildarmresources.StaticSiteBuildARMResourcesClient
	StaticSiteCustomDomainOverviewARMResources                 *staticsitecustomdomainoverviewarmresources.StaticSiteCustomDomainOverviewARMResourcesClient
	StaticSiteLinkedBackendARMResourceOperationGroup           *staticsitelinkedbackendarmresourceoperationgroup.StaticSiteLinkedBackendARMResourceOperationGroupClient
	StaticSiteLinkedBackendARMResources                        *staticsitelinkedbackendarmresources.StaticSiteLinkedBackendARMResourcesClient
	StaticSiteUserProvidedFunctionAppARMResourceOperationGroup *staticsiteuserprovidedfunctionapparmresourceoperationgroup.StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient
	StaticSiteUserProvidedFunctionAppARMResources              *staticsiteuserprovidedfunctionapparmresources.StaticSiteUserProvidedFunctionAppARMResourcesClient
	StaticSites                                                *staticsites.StaticSitesClient
	SwiftVirtualNetworkOperationGroup                          *swiftvirtualnetworkoperationgroup.SwiftVirtualNetworkOperationGroupClient
	SwiftVirtualNetworks                                       *swiftvirtualnetworks.SwiftVirtualNetworksClient
	TriggeredJobHistories                                      *triggeredjobhistories.TriggeredJobHistoriesClient
	TriggeredJobHistoryOperationGroup                          *triggeredjobhistoryoperationgroup.TriggeredJobHistoryOperationGroupClient
	TriggeredWebJobOperationGroup                              *triggeredwebjoboperationgroup.TriggeredWebJobOperationGroupClient
	TriggeredWebJobs                                           *triggeredwebjobs.TriggeredWebJobsClient
	Users                                                      *users.UsersClient
	VnetConnectionGatewayOperationGroup                        *vnetconnectiongatewayoperationgroup.VnetConnectionGatewayOperationGroupClient
	VnetConnectionOperationGroup                               *vnetconnectionoperationgroup.VnetConnectionOperationGroupClient
	VnetGatewayOperationGroup                                  *vnetgatewayoperationgroup.VnetGatewayOperationGroupClient
	VnetGateways                                               *vnetgateways.VnetGatewaysClient
	VnetInfoResourceOperationGroup                             *vnetinforesourceoperationgroup.VnetInfoResourceOperationGroupClient
	VnetInfoResources                                          *vnetinforesources.VnetInfoResourcesClient
	VnetRoutes                                                 *vnetroutes.VnetRoutesClient
	WebApps                                                    *webapps.WebAppsClient
	WebJobOperationGroup                                       *webjoboperationgroup.WebJobOperationGroupClient
	WebJobs                                                    *webjobs.WebJobsClient
	WebSiteInstanceStatusOperationGroup                        *websiteinstancestatusoperationgroup.WebSiteInstanceStatusOperationGroupClient
	WebSiteInstanceStatuses                                    *websiteinstancestatuses.WebSiteInstanceStatusesClient
	WorkerPoolResources                                        *workerpoolresources.WorkerPoolResourcesClient
	WorkflowEnvelopeOperationGroup                             *workflowenvelopeoperationgroup.WorkflowEnvelopeOperationGroupClient
	WorkflowEnvelopes                                          *workflowenvelopes.WorkflowEnvelopesClient
	WorkflowRunActionRepetitionDefinitions                     *workflowrunactionrepetitiondefinitions.WorkflowRunActionRepetitionDefinitionsClient
	WorkflowRunActionScopeRepetitions                          *workflowrunactionscoperepetitions.WorkflowRunActionScopeRepetitionsClient
	WorkflowRunActions                                         *workflowrunactions.WorkflowRunActionsClient
	WorkflowRuns                                               *workflowruns.WorkflowRunsClient
	WorkflowTriggerHistories                                   *workflowtriggerhistories.WorkflowTriggerHistoriesClient
	WorkflowTriggers                                           *workflowtriggers.WorkflowTriggersClient
	WorkflowVersions                                           *workflowversions.WorkflowVersionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	addressResponsesClient, err := addressresponses.NewAddressResponsesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AddressResponses client: %+v", err)
	}
	configureFunc(addressResponsesClient.Client)

	analysisDefinitionOperationGroupClient, err := analysisdefinitionoperationgroup.NewAnalysisDefinitionOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AnalysisDefinitionOperationGroup client: %+v", err)
	}
	configureFunc(analysisDefinitionOperationGroupClient.Client)

	analysisDefinitionsClient, err := analysisdefinitions.NewAnalysisDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AnalysisDefinitions client: %+v", err)
	}
	configureFunc(analysisDefinitionsClient.Client)

	appServiceEnvironmentResourcesClient, err := appserviceenvironmentresources.NewAppServiceEnvironmentResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServiceEnvironmentResources client: %+v", err)
	}
	configureFunc(appServiceEnvironmentResourcesClient.Client)

	appServiceEnvironmentsClient, err := appserviceenvironments.NewAppServiceEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServiceEnvironments client: %+v", err)
	}
	configureFunc(appServiceEnvironmentsClient.Client)

	appServicePlansClient, err := appserviceplans.NewAppServicePlansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServicePlans client: %+v", err)
	}
	configureFunc(appServicePlansClient.Client)

	appSettingKeyVaultReferenceClient, err := appsettingkeyvaultreference.NewAppSettingKeyVaultReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppSettingKeyVaultReference client: %+v", err)
	}
	configureFunc(appSettingKeyVaultReferenceClient.Client)

	appSettingKeyVaultReferenceSlotClient, err := appsettingkeyvaultreferenceslot.NewAppSettingKeyVaultReferenceSlotClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppSettingKeyVaultReferenceSlot client: %+v", err)
	}
	configureFunc(appSettingKeyVaultReferenceSlotClient.Client)

	aseV3NetworkingConfigurationsClient, err := asev3networkingconfigurations.NewAseV3NetworkingConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AseV3NetworkingConfigurations client: %+v", err)
	}
	configureFunc(aseV3NetworkingConfigurationsClient.Client)

	backupItemOperationGroupClient, err := backupitemoperationgroup.NewBackupItemOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupItemOperationGroup client: %+v", err)
	}
	configureFunc(backupItemOperationGroupClient.Client)

	backupItemsClient, err := backupitems.NewBackupItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupItems client: %+v", err)
	}
	configureFunc(backupItemsClient.Client)

	bySiteSlotOperationGroupClient, err := bysiteslotoperationgroup.NewBySiteSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BySiteSlotOperationGroup client: %+v", err)
	}
	configureFunc(bySiteSlotOperationGroupClient.Client)

	certificateOperationGroupClient, err := certificateoperationgroup.NewCertificateOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateOperationGroup client: %+v", err)
	}
	configureFunc(certificateOperationGroupClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	continuousWebJobOperationGroupClient, err := continuouswebjoboperationgroup.NewContinuousWebJobOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContinuousWebJobOperationGroup client: %+v", err)
	}
	configureFunc(continuousWebJobOperationGroupClient.Client)

	continuousWebJobsClient, err := continuouswebjobs.NewContinuousWebJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContinuousWebJobs client: %+v", err)
	}
	configureFunc(continuousWebJobsClient.Client)

	csmDeploymentStatusOperationGroupClient, err := csmdeploymentstatusoperationgroup.NewCsmDeploymentStatusOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CsmDeploymentStatusOperationGroup client: %+v", err)
	}
	configureFunc(csmDeploymentStatusOperationGroupClient.Client)

	csmDeploymentStatusesClient, err := csmdeploymentstatuses.NewCsmDeploymentStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CsmDeploymentStatuses client: %+v", err)
	}
	configureFunc(csmDeploymentStatusesClient.Client)

	csmPublishingCredentialsPoliciesEntitiesClient, err := csmpublishingcredentialspoliciesentities.NewCsmPublishingCredentialsPoliciesEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CsmPublishingCredentialsPoliciesEntities client: %+v", err)
	}
	configureFunc(csmPublishingCredentialsPoliciesEntitiesClient.Client)

	csmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient, err := csmpublishingcredentialspoliciesentityftpallowedslot.NewCsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CsmPublishingCredentialsPoliciesEntityFtpAllowedSlot client: %+v", err)
	}
	configureFunc(csmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient.Client)

	csmPublishingCredentialsPoliciesEntityOperationGroupClient, err := csmpublishingcredentialspoliciesentityoperationgroup.NewCsmPublishingCredentialsPoliciesEntityOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CsmPublishingCredentialsPoliciesEntityOperationGroup client: %+v", err)
	}
	configureFunc(csmPublishingCredentialsPoliciesEntityOperationGroupClient.Client)

	csmPublishingCredentialsPoliciesEntityScmAllowedSlotClient, err := csmpublishingcredentialspoliciesentityscmallowedslot.NewCsmPublishingCredentialsPoliciesEntityScmAllowedSlotClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CsmPublishingCredentialsPoliciesEntityScmAllowedSlot client: %+v", err)
	}
	configureFunc(csmPublishingCredentialsPoliciesEntityScmAllowedSlotClient.Client)

	customDnsSuffixConfigurationsClient, err := customdnssuffixconfigurations.NewCustomDnsSuffixConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomDnsSuffixConfigurations client: %+v", err)
	}
	configureFunc(customDnsSuffixConfigurationsClient.Client)

	databaseConnectionOperationGroupClient, err := databaseconnectionoperationgroup.NewDatabaseConnectionOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseConnectionOperationGroup client: %+v", err)
	}
	configureFunc(databaseConnectionOperationGroupClient.Client)

	databaseConnectionsClient, err := databaseconnections.NewDatabaseConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseConnections client: %+v", err)
	}
	configureFunc(databaseConnectionsClient.Client)

	deletedSitesClient, err := deletedsites.NewDeletedSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedSites client: %+v", err)
	}
	configureFunc(deletedSitesClient.Client)

	deploymentOperationGroupClient, err := deploymentoperationgroup.NewDeploymentOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentOperationGroup client: %+v", err)
	}
	configureFunc(deploymentOperationGroupClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	detectorDefinitionResourceOperationGroupClient, err := detectordefinitionresourceoperationgroup.NewDetectorDefinitionResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectorDefinitionResourceOperationGroup client: %+v", err)
	}
	configureFunc(detectorDefinitionResourceOperationGroupClient.Client)

	detectorDefinitionResourcesClient, err := detectordefinitionresources.NewDetectorDefinitionResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectorDefinitionResources client: %+v", err)
	}
	configureFunc(detectorDefinitionResourcesClient.Client)

	detectorResponseOperationGroupClient, err := detectorresponseoperationgroup.NewDetectorResponseOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectorResponseOperationGroup client: %+v", err)
	}
	configureFunc(detectorResponseOperationGroupClient.Client)

	detectorResponsesClient, err := detectorresponses.NewDetectorResponsesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectorResponses client: %+v", err)
	}
	configureFunc(detectorResponsesClient.Client)

	diagnosticCategoriesClient, err := diagnosticcategories.NewDiagnosticCategoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticCategories client: %+v", err)
	}
	configureFunc(diagnosticCategoriesClient.Client)

	diagnosticCategoryOperationGroupClient, err := diagnosticcategoryoperationgroup.NewDiagnosticCategoryOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticCategoryOperationGroup client: %+v", err)
	}
	configureFunc(diagnosticCategoryOperationGroupClient.Client)

	diagnosticsClient, err := diagnostics.NewDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Diagnostics client: %+v", err)
	}
	configureFunc(diagnosticsClient.Client)

	functionEnvelopeOperationGroupClient, err := functionenvelopeoperationgroup.NewFunctionEnvelopeOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FunctionEnvelopeOperationGroup client: %+v", err)
	}
	configureFunc(functionEnvelopeOperationGroupClient.Client)

	functionEnvelopesClient, err := functionenvelopes.NewFunctionEnvelopesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FunctionEnvelopes client: %+v", err)
	}
	configureFunc(functionEnvelopesClient.Client)

	globalClient, err := global.NewGlobalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Global client: %+v", err)
	}
	configureFunc(globalClient.Client)

	hostNameBindingOperationGroupClient, err := hostnamebindingoperationgroup.NewHostNameBindingOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HostNameBindingOperationGroup client: %+v", err)
	}
	configureFunc(hostNameBindingOperationGroupClient.Client)

	hostNameBindingsClient, err := hostnamebindings.NewHostNameBindingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HostNameBindings client: %+v", err)
	}
	configureFunc(hostNameBindingsClient.Client)

	hybridConnectionLimitsOperationGroupClient, err := hybridconnectionlimitsoperationgroup.NewHybridConnectionLimitsOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnectionLimitsOperationGroup client: %+v", err)
	}
	configureFunc(hybridConnectionLimitsOperationGroupClient.Client)

	hybridConnectionOperationGroupClient, err := hybridconnectionoperationgroup.NewHybridConnectionOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnectionOperationGroup client: %+v", err)
	}
	configureFunc(hybridConnectionOperationGroupClient.Client)

	hybridConnectionSlotOperationGroupClient, err := hybridconnectionslotoperationgroup.NewHybridConnectionSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnectionSlotOperationGroup client: %+v", err)
	}
	configureFunc(hybridConnectionSlotOperationGroupClient.Client)

	hybridConnectionsClient, err := hybridconnections.NewHybridConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnections client: %+v", err)
	}
	configureFunc(hybridConnectionsClient.Client)

	identifierOperationGroupClient, err := identifieroperationgroup.NewIdentifierOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentifierOperationGroup client: %+v", err)
	}
	configureFunc(identifierOperationGroupClient.Client)

	identifiersClient, err := identifiers.NewIdentifiersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Identifiers client: %+v", err)
	}
	configureFunc(identifiersClient.Client)

	instanceMSDeployStatusOperationGroupClient, err := instancemsdeploystatusoperationgroup.NewInstanceMSDeployStatusOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InstanceMSDeployStatusOperationGroup client: %+v", err)
	}
	configureFunc(instanceMSDeployStatusOperationGroupClient.Client)

	instanceProcessModuleSlotOperationGroupClient, err := instanceprocessmoduleslotoperationgroup.NewInstanceProcessModuleSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InstanceProcessModuleSlotOperationGroup client: %+v", err)
	}
	configureFunc(instanceProcessModuleSlotOperationGroupClient.Client)

	instanceProcessSlotOperationGroupClient, err := instanceprocessslotoperationgroup.NewInstanceProcessSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InstanceProcessSlotOperationGroup client: %+v", err)
	}
	configureFunc(instanceProcessSlotOperationGroupClient.Client)

	kubeEnvironmentsClient, err := kubeenvironments.NewKubeEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building KubeEnvironments client: %+v", err)
	}
	configureFunc(kubeEnvironmentsClient.Client)

	mSDeployStatusOperationGroupClient, err := msdeploystatusoperationgroup.NewMSDeployStatusOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MSDeployStatusOperationGroup client: %+v", err)
	}
	configureFunc(mSDeployStatusOperationGroupClient.Client)

	mSDeployStatusSlotOperationGroupClient, err := msdeploystatusslotoperationgroup.NewMSDeployStatusSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MSDeployStatusSlotOperationGroup client: %+v", err)
	}
	configureFunc(mSDeployStatusSlotOperationGroupClient.Client)

	mSDeployStatusesClient, err := msdeploystatuses.NewMSDeployStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MSDeployStatuses client: %+v", err)
	}
	configureFunc(mSDeployStatusesClient.Client)

	migrateMySqlStatusOperationGroupClient, err := migratemysqlstatusoperationgroup.NewMigrateMySqlStatusOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateMySqlStatusOperationGroup client: %+v", err)
	}
	configureFunc(migrateMySqlStatusOperationGroupClient.Client)

	migrateMySqlStatusesClient, err := migratemysqlstatuses.NewMigrateMySqlStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateMySqlStatuses client: %+v", err)
	}
	configureFunc(migrateMySqlStatusesClient.Client)

	networkFeaturesOperationGroupClient, err := networkfeaturesoperationgroup.NewNetworkFeaturesOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkFeaturesOperationGroup client: %+v", err)
	}
	configureFunc(networkFeaturesOperationGroupClient.Client)

	networkFeaturesSlotOperationGroupClient, err := networkfeaturesslotoperationgroup.NewNetworkFeaturesSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkFeaturesSlotOperationGroup client: %+v", err)
	}
	configureFunc(networkFeaturesSlotOperationGroupClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	premierAddOnOperationGroupClient, err := premieraddonoperationgroup.NewPremierAddOnOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PremierAddOnOperationGroup client: %+v", err)
	}
	configureFunc(premierAddOnOperationGroupClient.Client)

	premierAddonsClient, err := premieraddons.NewPremierAddonsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PremierAddons client: %+v", err)
	}
	configureFunc(premierAddonsClient.Client)

	privateAccessOperationGroupClient, err := privateaccessoperationgroup.NewPrivateAccessOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateAccessOperationGroup client: %+v", err)
	}
	configureFunc(privateAccessOperationGroupClient.Client)

	privateAccessesClient, err := privateaccesses.NewPrivateAccessesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateAccesses client: %+v", err)
	}
	configureFunc(privateAccessesClient.Client)

	privateEndpointConnectionSlotOperationGroupClient, err := privateendpointconnectionslotoperationgroup.NewPrivateEndpointConnectionSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnectionSlotOperationGroup client: %+v", err)
	}
	configureFunc(privateEndpointConnectionSlotOperationGroupClient.Client)

	processInfoOperationGroupClient, err := processinfooperationgroup.NewProcessInfoOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProcessInfoOperationGroup client: %+v", err)
	}
	configureFunc(processInfoOperationGroupClient.Client)

	processInfosClient, err := processinfos.NewProcessInfosClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProcessInfos client: %+v", err)
	}
	configureFunc(processInfosClient.Client)

	processModuleInfoOperationGroupClient, err := processmoduleinfooperationgroup.NewProcessModuleInfoOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProcessModuleInfoOperationGroup client: %+v", err)
	}
	configureFunc(processModuleInfoOperationGroupClient.Client)

	processModuleInfosClient, err := processmoduleinfos.NewProcessModuleInfosClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProcessModuleInfos client: %+v", err)
	}
	configureFunc(processModuleInfosClient.Client)

	processModuleSlotOperationGroupClient, err := processmoduleslotoperationgroup.NewProcessModuleSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProcessModuleSlotOperationGroup client: %+v", err)
	}
	configureFunc(processModuleSlotOperationGroupClient.Client)

	processSlotOperationGroupClient, err := processslotoperationgroup.NewProcessSlotOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProcessSlotOperationGroup client: %+v", err)
	}
	configureFunc(processSlotOperationGroupClient.Client)

	publicCertificateOperationGroupClient, err := publiccertificateoperationgroup.NewPublicCertificateOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicCertificateOperationGroup client: %+v", err)
	}
	configureFunc(publicCertificateOperationGroupClient.Client)

	publicCertificatesClient, err := publiccertificates.NewPublicCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicCertificates client: %+v", err)
	}
	configureFunc(publicCertificatesClient.Client)

	recommendationRulesClient, err := recommendationrules.NewRecommendationRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecommendationRules client: %+v", err)
	}
	configureFunc(recommendationRulesClient.Client)

	recommendationsClient, err := recommendations.NewRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Recommendations client: %+v", err)
	}
	configureFunc(recommendationsClient.Client)

	relayServiceConnectionEntitiesClient, err := relayserviceconnectionentities.NewRelayServiceConnectionEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RelayServiceConnectionEntities client: %+v", err)
	}
	configureFunc(relayServiceConnectionEntitiesClient.Client)

	relayServiceConnectionEntityOperationGroupClient, err := relayserviceconnectionentityoperationgroup.NewRelayServiceConnectionEntityOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RelayServiceConnectionEntityOperationGroup client: %+v", err)
	}
	configureFunc(relayServiceConnectionEntityOperationGroupClient.Client)

	remotePrivateEndpointConnectionARMResourceOperationGroupClient, err := remoteprivateendpointconnectionarmresourceoperationgroup.NewRemotePrivateEndpointConnectionARMResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemotePrivateEndpointConnectionARMResourceOperationGroup client: %+v", err)
	}
	configureFunc(remotePrivateEndpointConnectionARMResourceOperationGroupClient.Client)

	remotePrivateEndpointConnectionARMResourcesClient, err := remoteprivateendpointconnectionarmresources.NewRemotePrivateEndpointConnectionARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemotePrivateEndpointConnectionARMResources client: %+v", err)
	}
	configureFunc(remotePrivateEndpointConnectionARMResourcesClient.Client)

	requestHistoriesClient, err := requesthistories.NewRequestHistoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RequestHistories client: %+v", err)
	}
	configureFunc(requestHistoriesClient.Client)

	resourceHealthMetadataOperationGroupClient, err := resourcehealthmetadataoperationgroup.NewResourceHealthMetadataOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceHealthMetadataOperationGroup client: %+v", err)
	}
	configureFunc(resourceHealthMetadataOperationGroupClient.Client)

	siteAuthSettingsV2OperationGroupClient, err := siteauthsettingsv2operationgroup.NewSiteAuthSettingsV2OperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAuthSettingsV2OperationGroup client: %+v", err)
	}
	configureFunc(siteAuthSettingsV2OperationGroupClient.Client)

	siteAuthSettingsV2sClient, err := siteauthsettingsv2s.NewSiteAuthSettingsV2sClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAuthSettingsV2s client: %+v", err)
	}
	configureFunc(siteAuthSettingsV2sClient.Client)

	siteCertificatesClient, err := sitecertificates.NewSiteCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteCertificates client: %+v", err)
	}
	configureFunc(siteCertificatesClient.Client)

	siteConfigResourceOperationGroupClient, err := siteconfigresourceoperationgroup.NewSiteConfigResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteConfigResourceOperationGroup client: %+v", err)
	}
	configureFunc(siteConfigResourceOperationGroupClient.Client)

	siteConfigResourcesClient, err := siteconfigresources.NewSiteConfigResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteConfigResources client: %+v", err)
	}
	configureFunc(siteConfigResourcesClient.Client)

	siteConfigSlotResourceOperationGroupClient, err := siteconfigslotresourceoperationgroup.NewSiteConfigSlotResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteConfigSlotResourceOperationGroup client: %+v", err)
	}
	configureFunc(siteConfigSlotResourceOperationGroupClient.Client)

	siteConfigSnapshotSlotResourceOperationGroupClient, err := siteconfigsnapshotslotresourceoperationgroup.NewSiteConfigSnapshotSlotResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteConfigSnapshotSlotResourceOperationGroup client: %+v", err)
	}
	configureFunc(siteConfigSnapshotSlotResourceOperationGroupClient.Client)

	siteConnectionStringKeyVaultReferenceClient, err := siteconnectionstringkeyvaultreference.NewSiteConnectionStringKeyVaultReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteConnectionStringKeyVaultReference client: %+v", err)
	}
	configureFunc(siteConnectionStringKeyVaultReferenceClient.Client)

	siteConnectionStringKeyVaultReferenceSlotClient, err := siteconnectionstringkeyvaultreferenceslot.NewSiteConnectionStringKeyVaultReferenceSlotClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteConnectionStringKeyVaultReferenceSlot client: %+v", err)
	}
	configureFunc(siteConnectionStringKeyVaultReferenceSlotClient.Client)

	siteContainerOperationGroupClient, err := sitecontaineroperationgroup.NewSiteContainerOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContainerOperationGroup client: %+v", err)
	}
	configureFunc(siteContainerOperationGroupClient.Client)

	siteContainersClient, err := sitecontainers.NewSiteContainersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContainers client: %+v", err)
	}
	configureFunc(siteContainersClient.Client)

	siteExtensionInfoOperationGroupClient, err := siteextensioninfooperationgroup.NewSiteExtensionInfoOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteExtensionInfoOperationGroup client: %+v", err)
	}
	configureFunc(siteExtensionInfoOperationGroupClient.Client)

	siteExtensionInfosClient, err := siteextensioninfos.NewSiteExtensionInfosClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteExtensionInfos client: %+v", err)
	}
	configureFunc(siteExtensionInfosClient.Client)

	siteLogsConfigOperationGroupClient, err := sitelogsconfigoperationgroup.NewSiteLogsConfigOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteLogsConfigOperationGroup client: %+v", err)
	}
	configureFunc(siteLogsConfigOperationGroupClient.Client)

	siteLogsConfigsClient, err := sitelogsconfigs.NewSiteLogsConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteLogsConfigs client: %+v", err)
	}
	configureFunc(siteLogsConfigsClient.Client)

	siteSourceControlOperationGroupClient, err := sitesourcecontroloperationgroup.NewSiteSourceControlOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteSourceControlOperationGroup client: %+v", err)
	}
	configureFunc(siteSourceControlOperationGroupClient.Client)

	siteSourceControlsClient, err := sitesourcecontrols.NewSiteSourceControlsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteSourceControls client: %+v", err)
	}
	configureFunc(siteSourceControlsClient.Client)

	sitesClient, err := sites.NewSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Sites client: %+v", err)
	}
	configureFunc(sitesClient.Client)

	slotConfigNamesResourcesClient, err := slotconfignamesresources.NewSlotConfigNamesResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SlotConfigNamesResources client: %+v", err)
	}
	configureFunc(slotConfigNamesResourcesClient.Client)

	sourceControlsClient, err := sourcecontrols.NewSourceControlsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SourceControls client: %+v", err)
	}
	configureFunc(sourceControlsClient.Client)

	staticSiteARMResourcesClient, err := staticsitearmresources.NewStaticSiteARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteARMResources client: %+v", err)
	}
	configureFunc(staticSiteARMResourcesClient.Client)

	staticSiteBasicAuthPropertiesARMResourcesClient, err := staticsitebasicauthpropertiesarmresources.NewStaticSiteBasicAuthPropertiesARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteBasicAuthPropertiesARMResources client: %+v", err)
	}
	configureFunc(staticSiteBasicAuthPropertiesARMResourcesClient.Client)

	staticSiteBuildARMResourcesClient, err := staticsitebuildarmresources.NewStaticSiteBuildARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteBuildARMResources client: %+v", err)
	}
	configureFunc(staticSiteBuildARMResourcesClient.Client)

	staticSiteCustomDomainOverviewARMResourcesClient, err := staticsitecustomdomainoverviewarmresources.NewStaticSiteCustomDomainOverviewARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteCustomDomainOverviewARMResources client: %+v", err)
	}
	configureFunc(staticSiteCustomDomainOverviewARMResourcesClient.Client)

	staticSiteLinkedBackendARMResourceOperationGroupClient, err := staticsitelinkedbackendarmresourceoperationgroup.NewStaticSiteLinkedBackendARMResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteLinkedBackendARMResourceOperationGroup client: %+v", err)
	}
	configureFunc(staticSiteLinkedBackendARMResourceOperationGroupClient.Client)

	staticSiteLinkedBackendARMResourcesClient, err := staticsitelinkedbackendarmresources.NewStaticSiteLinkedBackendARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteLinkedBackendARMResources client: %+v", err)
	}
	configureFunc(staticSiteLinkedBackendARMResourcesClient.Client)

	staticSiteUserProvidedFunctionAppARMResourceOperationGroupClient, err := staticsiteuserprovidedfunctionapparmresourceoperationgroup.NewStaticSiteUserProvidedFunctionAppARMResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteUserProvidedFunctionAppARMResourceOperationGroup client: %+v", err)
	}
	configureFunc(staticSiteUserProvidedFunctionAppARMResourceOperationGroupClient.Client)

	staticSiteUserProvidedFunctionAppARMResourcesClient, err := staticsiteuserprovidedfunctionapparmresources.NewStaticSiteUserProvidedFunctionAppARMResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSiteUserProvidedFunctionAppARMResources client: %+v", err)
	}
	configureFunc(staticSiteUserProvidedFunctionAppARMResourcesClient.Client)

	staticSitesClient, err := staticsites.NewStaticSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSites client: %+v", err)
	}
	configureFunc(staticSitesClient.Client)

	swiftVirtualNetworkOperationGroupClient, err := swiftvirtualnetworkoperationgroup.NewSwiftVirtualNetworkOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SwiftVirtualNetworkOperationGroup client: %+v", err)
	}
	configureFunc(swiftVirtualNetworkOperationGroupClient.Client)

	swiftVirtualNetworksClient, err := swiftvirtualnetworks.NewSwiftVirtualNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SwiftVirtualNetworks client: %+v", err)
	}
	configureFunc(swiftVirtualNetworksClient.Client)

	triggeredJobHistoriesClient, err := triggeredjobhistories.NewTriggeredJobHistoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggeredJobHistories client: %+v", err)
	}
	configureFunc(triggeredJobHistoriesClient.Client)

	triggeredJobHistoryOperationGroupClient, err := triggeredjobhistoryoperationgroup.NewTriggeredJobHistoryOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggeredJobHistoryOperationGroup client: %+v", err)
	}
	configureFunc(triggeredJobHistoryOperationGroupClient.Client)

	triggeredWebJobOperationGroupClient, err := triggeredwebjoboperationgroup.NewTriggeredWebJobOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggeredWebJobOperationGroup client: %+v", err)
	}
	configureFunc(triggeredWebJobOperationGroupClient.Client)

	triggeredWebJobsClient, err := triggeredwebjobs.NewTriggeredWebJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TriggeredWebJobs client: %+v", err)
	}
	configureFunc(triggeredWebJobsClient.Client)

	usersClient, err := users.NewUsersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Users client: %+v", err)
	}
	configureFunc(usersClient.Client)

	vnetConnectionGatewayOperationGroupClient, err := vnetconnectiongatewayoperationgroup.NewVnetConnectionGatewayOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetConnectionGatewayOperationGroup client: %+v", err)
	}
	configureFunc(vnetConnectionGatewayOperationGroupClient.Client)

	vnetConnectionOperationGroupClient, err := vnetconnectionoperationgroup.NewVnetConnectionOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetConnectionOperationGroup client: %+v", err)
	}
	configureFunc(vnetConnectionOperationGroupClient.Client)

	vnetGatewayOperationGroupClient, err := vnetgatewayoperationgroup.NewVnetGatewayOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetGatewayOperationGroup client: %+v", err)
	}
	configureFunc(vnetGatewayOperationGroupClient.Client)

	vnetGatewaysClient, err := vnetgateways.NewVnetGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetGateways client: %+v", err)
	}
	configureFunc(vnetGatewaysClient.Client)

	vnetInfoResourceOperationGroupClient, err := vnetinforesourceoperationgroup.NewVnetInfoResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetInfoResourceOperationGroup client: %+v", err)
	}
	configureFunc(vnetInfoResourceOperationGroupClient.Client)

	vnetInfoResourcesClient, err := vnetinforesources.NewVnetInfoResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetInfoResources client: %+v", err)
	}
	configureFunc(vnetInfoResourcesClient.Client)

	vnetRoutesClient, err := vnetroutes.NewVnetRoutesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VnetRoutes client: %+v", err)
	}
	configureFunc(vnetRoutesClient.Client)

	webAppsClient, err := webapps.NewWebAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebApps client: %+v", err)
	}
	configureFunc(webAppsClient.Client)

	webJobOperationGroupClient, err := webjoboperationgroup.NewWebJobOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebJobOperationGroup client: %+v", err)
	}
	configureFunc(webJobOperationGroupClient.Client)

	webJobsClient, err := webjobs.NewWebJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebJobs client: %+v", err)
	}
	configureFunc(webJobsClient.Client)

	webSiteInstanceStatusOperationGroupClient, err := websiteinstancestatusoperationgroup.NewWebSiteInstanceStatusOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebSiteInstanceStatusOperationGroup client: %+v", err)
	}
	configureFunc(webSiteInstanceStatusOperationGroupClient.Client)

	webSiteInstanceStatusesClient, err := websiteinstancestatuses.NewWebSiteInstanceStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebSiteInstanceStatuses client: %+v", err)
	}
	configureFunc(webSiteInstanceStatusesClient.Client)

	workerPoolResourcesClient, err := workerpoolresources.NewWorkerPoolResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkerPoolResources client: %+v", err)
	}
	configureFunc(workerPoolResourcesClient.Client)

	workflowEnvelopeOperationGroupClient, err := workflowenvelopeoperationgroup.NewWorkflowEnvelopeOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowEnvelopeOperationGroup client: %+v", err)
	}
	configureFunc(workflowEnvelopeOperationGroupClient.Client)

	workflowEnvelopesClient, err := workflowenvelopes.NewWorkflowEnvelopesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowEnvelopes client: %+v", err)
	}
	configureFunc(workflowEnvelopesClient.Client)

	workflowRunActionRepetitionDefinitionsClient, err := workflowrunactionrepetitiondefinitions.NewWorkflowRunActionRepetitionDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRunActionRepetitionDefinitions client: %+v", err)
	}
	configureFunc(workflowRunActionRepetitionDefinitionsClient.Client)

	workflowRunActionScopeRepetitionsClient, err := workflowrunactionscoperepetitions.NewWorkflowRunActionScopeRepetitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRunActionScopeRepetitions client: %+v", err)
	}
	configureFunc(workflowRunActionScopeRepetitionsClient.Client)

	workflowRunActionsClient, err := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRunActions client: %+v", err)
	}
	configureFunc(workflowRunActionsClient.Client)

	workflowRunsClient, err := workflowruns.NewWorkflowRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRuns client: %+v", err)
	}
	configureFunc(workflowRunsClient.Client)

	workflowTriggerHistoriesClient, err := workflowtriggerhistories.NewWorkflowTriggerHistoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowTriggerHistories client: %+v", err)
	}
	configureFunc(workflowTriggerHistoriesClient.Client)

	workflowTriggersClient, err := workflowtriggers.NewWorkflowTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowTriggers client: %+v", err)
	}
	configureFunc(workflowTriggersClient.Client)

	workflowVersionsClient, err := workflowversions.NewWorkflowVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowVersions client: %+v", err)
	}
	configureFunc(workflowVersionsClient.Client)

	return &Client{
		AddressResponses:                                           addressResponsesClient,
		AnalysisDefinitionOperationGroup:                           analysisDefinitionOperationGroupClient,
		AnalysisDefinitions:                                        analysisDefinitionsClient,
		AppServiceEnvironmentResources:                             appServiceEnvironmentResourcesClient,
		AppServiceEnvironments:                                     appServiceEnvironmentsClient,
		AppServicePlans:                                            appServicePlansClient,
		AppSettingKeyVaultReference:                                appSettingKeyVaultReferenceClient,
		AppSettingKeyVaultReferenceSlot:                            appSettingKeyVaultReferenceSlotClient,
		AseV3NetworkingConfigurations:                              aseV3NetworkingConfigurationsClient,
		BackupItemOperationGroup:                                   backupItemOperationGroupClient,
		BackupItems:                                                backupItemsClient,
		BySiteSlotOperationGroup:                                   bySiteSlotOperationGroupClient,
		CertificateOperationGroup:                                  certificateOperationGroupClient,
		Certificates:                                               certificatesClient,
		ContinuousWebJobOperationGroup:                             continuousWebJobOperationGroupClient,
		ContinuousWebJobs:                                          continuousWebJobsClient,
		CsmDeploymentStatusOperationGroup:                          csmDeploymentStatusOperationGroupClient,
		CsmDeploymentStatuses:                                      csmDeploymentStatusesClient,
		CsmPublishingCredentialsPoliciesEntities:                   csmPublishingCredentialsPoliciesEntitiesClient,
		CsmPublishingCredentialsPoliciesEntityFtpAllowedSlot:       csmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient,
		CsmPublishingCredentialsPoliciesEntityOperationGroup:       csmPublishingCredentialsPoliciesEntityOperationGroupClient,
		CsmPublishingCredentialsPoliciesEntityScmAllowedSlot:       csmPublishingCredentialsPoliciesEntityScmAllowedSlotClient,
		CustomDnsSuffixConfigurations:                              customDnsSuffixConfigurationsClient,
		DatabaseConnectionOperationGroup:                           databaseConnectionOperationGroupClient,
		DatabaseConnections:                                        databaseConnectionsClient,
		DeletedSites:                                               deletedSitesClient,
		DeploymentOperationGroup:                                   deploymentOperationGroupClient,
		Deployments:                                                deploymentsClient,
		DetectorDefinitionResourceOperationGroup:                   detectorDefinitionResourceOperationGroupClient,
		DetectorDefinitionResources:                                detectorDefinitionResourcesClient,
		DetectorResponseOperationGroup:                             detectorResponseOperationGroupClient,
		DetectorResponses:                                          detectorResponsesClient,
		DiagnosticCategories:                                       diagnosticCategoriesClient,
		DiagnosticCategoryOperationGroup:                           diagnosticCategoryOperationGroupClient,
		Diagnostics:                                                diagnosticsClient,
		FunctionEnvelopeOperationGroup:                             functionEnvelopeOperationGroupClient,
		FunctionEnvelopes:                                          functionEnvelopesClient,
		Global:                                                     globalClient,
		HostNameBindingOperationGroup:                              hostNameBindingOperationGroupClient,
		HostNameBindings:                                           hostNameBindingsClient,
		HybridConnectionLimitsOperationGroup:                       hybridConnectionLimitsOperationGroupClient,
		HybridConnectionOperationGroup:                             hybridConnectionOperationGroupClient,
		HybridConnectionSlotOperationGroup:                         hybridConnectionSlotOperationGroupClient,
		HybridConnections:                                          hybridConnectionsClient,
		IdentifierOperationGroup:                                   identifierOperationGroupClient,
		Identifiers:                                                identifiersClient,
		InstanceMSDeployStatusOperationGroup:                       instanceMSDeployStatusOperationGroupClient,
		InstanceProcessModuleSlotOperationGroup:                    instanceProcessModuleSlotOperationGroupClient,
		InstanceProcessSlotOperationGroup:                          instanceProcessSlotOperationGroupClient,
		KubeEnvironments:                                           kubeEnvironmentsClient,
		MSDeployStatusOperationGroup:                               mSDeployStatusOperationGroupClient,
		MSDeployStatusSlotOperationGroup:                           mSDeployStatusSlotOperationGroupClient,
		MSDeployStatuses:                                           mSDeployStatusesClient,
		MigrateMySqlStatusOperationGroup:                           migrateMySqlStatusOperationGroupClient,
		MigrateMySqlStatuses:                                       migrateMySqlStatusesClient,
		NetworkFeaturesOperationGroup:                              networkFeaturesOperationGroupClient,
		NetworkFeaturesSlotOperationGroup:                          networkFeaturesSlotOperationGroupClient,
		Openapis:                                                   openapisClient,
		PremierAddOnOperationGroup:                                 premierAddOnOperationGroupClient,
		PremierAddons:                                              premierAddonsClient,
		PrivateAccessOperationGroup:                                privateAccessOperationGroupClient,
		PrivateAccesses:                                            privateAccessesClient,
		PrivateEndpointConnectionSlotOperationGroup:                privateEndpointConnectionSlotOperationGroupClient,
		ProcessInfoOperationGroup:                                  processInfoOperationGroupClient,
		ProcessInfos:                                               processInfosClient,
		ProcessModuleInfoOperationGroup:                            processModuleInfoOperationGroupClient,
		ProcessModuleInfos:                                         processModuleInfosClient,
		ProcessModuleSlotOperationGroup:                            processModuleSlotOperationGroupClient,
		ProcessSlotOperationGroup:                                  processSlotOperationGroupClient,
		PublicCertificateOperationGroup:                            publicCertificateOperationGroupClient,
		PublicCertificates:                                         publicCertificatesClient,
		RecommendationRules:                                        recommendationRulesClient,
		Recommendations:                                            recommendationsClient,
		RelayServiceConnectionEntities:                             relayServiceConnectionEntitiesClient,
		RelayServiceConnectionEntityOperationGroup:                 relayServiceConnectionEntityOperationGroupClient,
		RemotePrivateEndpointConnectionARMResourceOperationGroup:   remotePrivateEndpointConnectionARMResourceOperationGroupClient,
		RemotePrivateEndpointConnectionARMResources:                remotePrivateEndpointConnectionARMResourcesClient,
		RequestHistories:                                           requestHistoriesClient,
		ResourceHealthMetadataOperationGroup:                       resourceHealthMetadataOperationGroupClient,
		SiteAuthSettingsV2OperationGroup:                           siteAuthSettingsV2OperationGroupClient,
		SiteAuthSettingsV2s:                                        siteAuthSettingsV2sClient,
		SiteCertificates:                                           siteCertificatesClient,
		SiteConfigResourceOperationGroup:                           siteConfigResourceOperationGroupClient,
		SiteConfigResources:                                        siteConfigResourcesClient,
		SiteConfigSlotResourceOperationGroup:                       siteConfigSlotResourceOperationGroupClient,
		SiteConfigSnapshotSlotResourceOperationGroup:               siteConfigSnapshotSlotResourceOperationGroupClient,
		SiteConnectionStringKeyVaultReference:                      siteConnectionStringKeyVaultReferenceClient,
		SiteConnectionStringKeyVaultReferenceSlot:                  siteConnectionStringKeyVaultReferenceSlotClient,
		SiteContainerOperationGroup:                                siteContainerOperationGroupClient,
		SiteContainers:                                             siteContainersClient,
		SiteExtensionInfoOperationGroup:                            siteExtensionInfoOperationGroupClient,
		SiteExtensionInfos:                                         siteExtensionInfosClient,
		SiteLogsConfigOperationGroup:                               siteLogsConfigOperationGroupClient,
		SiteLogsConfigs:                                            siteLogsConfigsClient,
		SiteSourceControlOperationGroup:                            siteSourceControlOperationGroupClient,
		SiteSourceControls:                                         siteSourceControlsClient,
		Sites:                                                      sitesClient,
		SlotConfigNamesResources:                                   slotConfigNamesResourcesClient,
		SourceControls:                                             sourceControlsClient,
		StaticSiteARMResources:                                     staticSiteARMResourcesClient,
		StaticSiteBasicAuthPropertiesARMResources:                  staticSiteBasicAuthPropertiesARMResourcesClient,
		StaticSiteBuildARMResources:                                staticSiteBuildARMResourcesClient,
		StaticSiteCustomDomainOverviewARMResources:                 staticSiteCustomDomainOverviewARMResourcesClient,
		StaticSiteLinkedBackendARMResourceOperationGroup:           staticSiteLinkedBackendARMResourceOperationGroupClient,
		StaticSiteLinkedBackendARMResources:                        staticSiteLinkedBackendARMResourcesClient,
		StaticSiteUserProvidedFunctionAppARMResourceOperationGroup: staticSiteUserProvidedFunctionAppARMResourceOperationGroupClient,
		StaticSiteUserProvidedFunctionAppARMResources:              staticSiteUserProvidedFunctionAppARMResourcesClient,
		StaticSites:                                                staticSitesClient,
		SwiftVirtualNetworkOperationGroup:                          swiftVirtualNetworkOperationGroupClient,
		SwiftVirtualNetworks:                                       swiftVirtualNetworksClient,
		TriggeredJobHistories:                                      triggeredJobHistoriesClient,
		TriggeredJobHistoryOperationGroup:                          triggeredJobHistoryOperationGroupClient,
		TriggeredWebJobOperationGroup:                              triggeredWebJobOperationGroupClient,
		TriggeredWebJobs:                                           triggeredWebJobsClient,
		Users:                                                      usersClient,
		VnetConnectionGatewayOperationGroup:                        vnetConnectionGatewayOperationGroupClient,
		VnetConnectionOperationGroup:                               vnetConnectionOperationGroupClient,
		VnetGatewayOperationGroup:                                  vnetGatewayOperationGroupClient,
		VnetGateways:                                               vnetGatewaysClient,
		VnetInfoResourceOperationGroup:                             vnetInfoResourceOperationGroupClient,
		VnetInfoResources:                                          vnetInfoResourcesClient,
		VnetRoutes:                                                 vnetRoutesClient,
		WebApps:                                                    webAppsClient,
		WebJobOperationGroup:                                       webJobOperationGroupClient,
		WebJobs:                                                    webJobsClient,
		WebSiteInstanceStatusOperationGroup:                        webSiteInstanceStatusOperationGroupClient,
		WebSiteInstanceStatuses:                                    webSiteInstanceStatusesClient,
		WorkerPoolResources:                                        workerPoolResourcesClient,
		WorkflowEnvelopeOperationGroup:                             workflowEnvelopeOperationGroupClient,
		WorkflowEnvelopes:                                          workflowEnvelopesClient,
		WorkflowRunActionRepetitionDefinitions:                     workflowRunActionRepetitionDefinitionsClient,
		WorkflowRunActionScopeRepetitions:                          workflowRunActionScopeRepetitionsClient,
		WorkflowRunActions:                                         workflowRunActionsClient,
		WorkflowRuns:                                               workflowRunsClient,
		WorkflowTriggerHistories:                                   workflowTriggerHistoriesClient,
		WorkflowTriggers:                                           workflowTriggersClient,
		WorkflowVersions:                                           workflowVersionsClient,
	}, nil
}
