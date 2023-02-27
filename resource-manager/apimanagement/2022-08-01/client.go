// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/api"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apidiagnostic"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiissue"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiissueattachment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiissuecomment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apimanagementservice"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apimanagementserviceskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apioperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apioperationpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apioperationsbytag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apioperationtag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apipolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiproduct"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apirelease"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apirevision"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apisbytag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apischema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apitag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apitagdescription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiversionset"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiversionsets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiwiki"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorization"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationconfirmconsentcode"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationloginlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationprovider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationserver"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/backend"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/backendreconnect"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/cache"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/certificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/contenttype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/contenttypecontentitem"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/delegationsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/deletedservice"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/diagnostic"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/documentationresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/emailtemplate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/emailtemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gateway"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewayapi"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewaycertificateauthority"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewaygeneratetoken"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewayhostnameconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewaylistkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewayregeneratekey"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/graphqlapiresolver"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/graphqlapiresolverpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/group"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/groupuser"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/identityprovider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/issue"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/logger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/namedvalue"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/networkstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/notification"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/notificationrecipientemail"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/notificationrecipientuser"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/openidconnectprovider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/performconnectivitycheck"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/policy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/policydescription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/policyfragment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/portalconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/portalrevision"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/portalsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/product"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productapi"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productsbytag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productsubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/producttag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productwiki"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/quotabycounterkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/quotabyperiodkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/region"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/reports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/schema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/signinsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/signupsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/subscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tag"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tagresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantaccess"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantaccessgit"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantconfigurationsyncstate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/user"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/userconfirmationpasswordsend"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/usergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/useridentity"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/users"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/usersubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/usertoken"
)

type Client struct {
	Api                                  *api.ApiClient
	ApiDiagnostic                        *apidiagnostic.ApiDiagnosticClient
	ApiIssue                             *apiissue.ApiIssueClient
	ApiIssueAttachment                   *apiissueattachment.ApiIssueAttachmentClient
	ApiIssueComment                      *apiissuecomment.ApiIssueCommentClient
	ApiManagementService                 *apimanagementservice.ApiManagementServiceClient
	ApiManagementServiceSkus             *apimanagementserviceskus.ApiManagementServiceSkusClient
	ApiOperation                         *apioperation.ApiOperationClient
	ApiOperationPolicy                   *apioperationpolicy.ApiOperationPolicyClient
	ApiOperationTag                      *apioperationtag.ApiOperationTagClient
	ApiOperationsByTag                   *apioperationsbytag.ApiOperationsByTagClient
	ApiPolicy                            *apipolicy.ApiPolicyClient
	ApiProduct                           *apiproduct.ApiProductClient
	ApiRelease                           *apirelease.ApiReleaseClient
	ApiRevision                          *apirevision.ApiRevisionClient
	ApiSchema                            *apischema.ApiSchemaClient
	ApiTag                               *apitag.ApiTagClient
	ApiTagDescription                    *apitagdescription.ApiTagDescriptionClient
	ApiVersionSet                        *apiversionset.ApiVersionSetClient
	ApiVersionSets                       *apiversionsets.ApiVersionSetsClient
	ApiWiki                              *apiwiki.ApiWikiClient
	ApisByTag                            *apisbytag.ApisByTagClient
	Authorization                        *authorization.AuthorizationClient
	AuthorizationAccessPolicy            *authorizationaccesspolicy.AuthorizationAccessPolicyClient
	AuthorizationConfirmConsentCode      *authorizationconfirmconsentcode.AuthorizationConfirmConsentCodeClient
	AuthorizationLoginLinks              *authorizationloginlinks.AuthorizationLoginLinksClient
	AuthorizationProvider                *authorizationprovider.AuthorizationProviderClient
	AuthorizationServer                  *authorizationserver.AuthorizationServerClient
	Authorizations                       *authorizations.AuthorizationsClient
	Backend                              *backend.BackendClient
	BackendReconnect                     *backendreconnect.BackendReconnectClient
	Cache                                *cache.CacheClient
	Certificate                          *certificate.CertificateClient
	ContentType                          *contenttype.ContentTypeClient
	ContentTypeContentItem               *contenttypecontentitem.ContentTypeContentItemClient
	DelegationSettings                   *delegationsettings.DelegationSettingsClient
	DeletedService                       *deletedservice.DeletedServiceClient
	Diagnostic                           *diagnostic.DiagnosticClient
	DocumentationResource                *documentationresource.DocumentationResourceClient
	EmailTemplate                        *emailtemplate.EmailTemplateClient
	EmailTemplates                       *emailtemplates.EmailTemplatesClient
	Gateway                              *gateway.GatewayClient
	GatewayApi                           *gatewayapi.GatewayApiClient
	GatewayCertificateAuthority          *gatewaycertificateauthority.GatewayCertificateAuthorityClient
	GatewayGenerateToken                 *gatewaygeneratetoken.GatewayGenerateTokenClient
	GatewayHostnameConfiguration         *gatewayhostnameconfiguration.GatewayHostnameConfigurationClient
	GatewayListKeys                      *gatewaylistkeys.GatewayListKeysClient
	GatewayRegenerateKey                 *gatewayregeneratekey.GatewayRegenerateKeyClient
	GraphQLApiResolver                   *graphqlapiresolver.GraphQLApiResolverClient
	GraphQLApiResolverPolicy             *graphqlapiresolverpolicy.GraphQLApiResolverPolicyClient
	Group                                *group.GroupClient
	GroupUser                            *groupuser.GroupUserClient
	IdentityProvider                     *identityprovider.IdentityProviderClient
	Issue                                *issue.IssueClient
	Logger                               *logger.LoggerClient
	NamedValue                           *namedvalue.NamedValueClient
	NetworkStatus                        *networkstatus.NetworkStatusClient
	Notification                         *notification.NotificationClient
	NotificationRecipientEmail           *notificationrecipientemail.NotificationRecipientEmailClient
	NotificationRecipientUser            *notificationrecipientuser.NotificationRecipientUserClient
	OpenidConnectProvider                *openidconnectprovider.OpenidConnectProviderClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PerformConnectivityCheck             *performconnectivitycheck.PerformConnectivityCheckClient
	Policy                               *policy.PolicyClient
	PolicyDescription                    *policydescription.PolicyDescriptionClient
	PolicyFragment                       *policyfragment.PolicyFragmentClient
	PortalConfig                         *portalconfig.PortalConfigClient
	PortalRevision                       *portalrevision.PortalRevisionClient
	PortalSettings                       *portalsettings.PortalSettingsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	Product                              *product.ProductClient
	ProductApi                           *productapi.ProductApiClient
	ProductGroup                         *productgroup.ProductGroupClient
	ProductPolicy                        *productpolicy.ProductPolicyClient
	ProductSubscription                  *productsubscription.ProductSubscriptionClient
	ProductTag                           *producttag.ProductTagClient
	ProductWiki                          *productwiki.ProductWikiClient
	ProductsByTag                        *productsbytag.ProductsByTagClient
	QuotaByCounterKeys                   *quotabycounterkeys.QuotaByCounterKeysClient
	QuotaByPeriodKeys                    *quotabyperiodkeys.QuotaByPeriodKeysClient
	Region                               *region.RegionClient
	Reports                              *reports.ReportsClient
	Schema                               *schema.SchemaClient
	SignInSettings                       *signinsettings.SignInSettingsClient
	SignUpSettings                       *signupsettings.SignUpSettingsClient
	Skus                                 *skus.SkusClient
	Subscription                         *subscription.SubscriptionClient
	Tag                                  *tag.TagClient
	TagResource                          *tagresource.TagResourceClient
	TenantAccess                         *tenantaccess.TenantAccessClient
	TenantAccessGit                      *tenantaccessgit.TenantAccessGitClient
	TenantConfiguration                  *tenantconfiguration.TenantConfigurationClient
	TenantConfigurationSyncState         *tenantconfigurationsyncstate.TenantConfigurationSyncStateClient
	TenantSettings                       *tenantsettings.TenantSettingsClient
	User                                 *user.UserClient
	UserConfirmationPasswordSend         *userconfirmationpasswordsend.UserConfirmationPasswordSendClient
	UserGroup                            *usergroup.UserGroupClient
	UserIdentity                         *useridentity.UserIdentityClient
	UserSubscription                     *usersubscription.UserSubscriptionClient
	UserToken                            *usertoken.UserTokenClient
	Users                                *users.UsersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	apiClient := api.NewApiClientWithBaseURI(endpoint)
	configureAuthFunc(&apiClient.Client)

	apiDiagnosticClient := apidiagnostic.NewApiDiagnosticClientWithBaseURI(endpoint)
	configureAuthFunc(&apiDiagnosticClient.Client)

	apiIssueAttachmentClient := apiissueattachment.NewApiIssueAttachmentClientWithBaseURI(endpoint)
	configureAuthFunc(&apiIssueAttachmentClient.Client)

	apiIssueClient := apiissue.NewApiIssueClientWithBaseURI(endpoint)
	configureAuthFunc(&apiIssueClient.Client)

	apiIssueCommentClient := apiissuecomment.NewApiIssueCommentClientWithBaseURI(endpoint)
	configureAuthFunc(&apiIssueCommentClient.Client)

	apiManagementServiceClient := apimanagementservice.NewApiManagementServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&apiManagementServiceClient.Client)

	apiManagementServiceSkusClient := apimanagementserviceskus.NewApiManagementServiceSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&apiManagementServiceSkusClient.Client)

	apiOperationClient := apioperation.NewApiOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&apiOperationClient.Client)

	apiOperationPolicyClient := apioperationpolicy.NewApiOperationPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&apiOperationPolicyClient.Client)

	apiOperationTagClient := apioperationtag.NewApiOperationTagClientWithBaseURI(endpoint)
	configureAuthFunc(&apiOperationTagClient.Client)

	apiOperationsByTagClient := apioperationsbytag.NewApiOperationsByTagClientWithBaseURI(endpoint)
	configureAuthFunc(&apiOperationsByTagClient.Client)

	apiPolicyClient := apipolicy.NewApiPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&apiPolicyClient.Client)

	apiProductClient := apiproduct.NewApiProductClientWithBaseURI(endpoint)
	configureAuthFunc(&apiProductClient.Client)

	apiReleaseClient := apirelease.NewApiReleaseClientWithBaseURI(endpoint)
	configureAuthFunc(&apiReleaseClient.Client)

	apiRevisionClient := apirevision.NewApiRevisionClientWithBaseURI(endpoint)
	configureAuthFunc(&apiRevisionClient.Client)

	apiSchemaClient := apischema.NewApiSchemaClientWithBaseURI(endpoint)
	configureAuthFunc(&apiSchemaClient.Client)

	apiTagClient := apitag.NewApiTagClientWithBaseURI(endpoint)
	configureAuthFunc(&apiTagClient.Client)

	apiTagDescriptionClient := apitagdescription.NewApiTagDescriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&apiTagDescriptionClient.Client)

	apiVersionSetClient := apiversionset.NewApiVersionSetClientWithBaseURI(endpoint)
	configureAuthFunc(&apiVersionSetClient.Client)

	apiVersionSetsClient := apiversionsets.NewApiVersionSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&apiVersionSetsClient.Client)

	apiWikiClient := apiwiki.NewApiWikiClientWithBaseURI(endpoint)
	configureAuthFunc(&apiWikiClient.Client)

	apisByTagClient := apisbytag.NewApisByTagClientWithBaseURI(endpoint)
	configureAuthFunc(&apisByTagClient.Client)

	authorizationAccessPolicyClient := authorizationaccesspolicy.NewAuthorizationAccessPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationAccessPolicyClient.Client)

	authorizationClient := authorization.NewAuthorizationClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationClient.Client)

	authorizationConfirmConsentCodeClient := authorizationconfirmconsentcode.NewAuthorizationConfirmConsentCodeClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationConfirmConsentCodeClient.Client)

	authorizationLoginLinksClient := authorizationloginlinks.NewAuthorizationLoginLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationLoginLinksClient.Client)

	authorizationProviderClient := authorizationprovider.NewAuthorizationProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationProviderClient.Client)

	authorizationServerClient := authorizationserver.NewAuthorizationServerClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationServerClient.Client)

	authorizationsClient := authorizations.NewAuthorizationsClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationsClient.Client)

	backendClient := backend.NewBackendClientWithBaseURI(endpoint)
	configureAuthFunc(&backendClient.Client)

	backendReconnectClient := backendreconnect.NewBackendReconnectClientWithBaseURI(endpoint)
	configureAuthFunc(&backendReconnectClient.Client)

	cacheClient := cache.NewCacheClientWithBaseURI(endpoint)
	configureAuthFunc(&cacheClient.Client)

	certificateClient := certificate.NewCertificateClientWithBaseURI(endpoint)
	configureAuthFunc(&certificateClient.Client)

	contentTypeClient := contenttype.NewContentTypeClientWithBaseURI(endpoint)
	configureAuthFunc(&contentTypeClient.Client)

	contentTypeContentItemClient := contenttypecontentitem.NewContentTypeContentItemClientWithBaseURI(endpoint)
	configureAuthFunc(&contentTypeContentItemClient.Client)

	delegationSettingsClient := delegationsettings.NewDelegationSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&delegationSettingsClient.Client)

	deletedServiceClient := deletedservice.NewDeletedServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedServiceClient.Client)

	diagnosticClient := diagnostic.NewDiagnosticClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticClient.Client)

	documentationResourceClient := documentationresource.NewDocumentationResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&documentationResourceClient.Client)

	emailTemplateClient := emailtemplate.NewEmailTemplateClientWithBaseURI(endpoint)
	configureAuthFunc(&emailTemplateClient.Client)

	emailTemplatesClient := emailtemplates.NewEmailTemplatesClientWithBaseURI(endpoint)
	configureAuthFunc(&emailTemplatesClient.Client)

	gatewayApiClient := gatewayapi.NewGatewayApiClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayApiClient.Client)

	gatewayCertificateAuthorityClient := gatewaycertificateauthority.NewGatewayCertificateAuthorityClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayCertificateAuthorityClient.Client)

	gatewayClient := gateway.NewGatewayClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayClient.Client)

	gatewayGenerateTokenClient := gatewaygeneratetoken.NewGatewayGenerateTokenClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayGenerateTokenClient.Client)

	gatewayHostnameConfigurationClient := gatewayhostnameconfiguration.NewGatewayHostnameConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayHostnameConfigurationClient.Client)

	gatewayListKeysClient := gatewaylistkeys.NewGatewayListKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayListKeysClient.Client)

	gatewayRegenerateKeyClient := gatewayregeneratekey.NewGatewayRegenerateKeyClientWithBaseURI(endpoint)
	configureAuthFunc(&gatewayRegenerateKeyClient.Client)

	graphQLApiResolverClient := graphqlapiresolver.NewGraphQLApiResolverClientWithBaseURI(endpoint)
	configureAuthFunc(&graphQLApiResolverClient.Client)

	graphQLApiResolverPolicyClient := graphqlapiresolverpolicy.NewGraphQLApiResolverPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&graphQLApiResolverPolicyClient.Client)

	groupClient := group.NewGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&groupClient.Client)

	groupUserClient := groupuser.NewGroupUserClientWithBaseURI(endpoint)
	configureAuthFunc(&groupUserClient.Client)

	identityProviderClient := identityprovider.NewIdentityProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&identityProviderClient.Client)

	issueClient := issue.NewIssueClientWithBaseURI(endpoint)
	configureAuthFunc(&issueClient.Client)

	loggerClient := logger.NewLoggerClientWithBaseURI(endpoint)
	configureAuthFunc(&loggerClient.Client)

	namedValueClient := namedvalue.NewNamedValueClientWithBaseURI(endpoint)
	configureAuthFunc(&namedValueClient.Client)

	networkStatusClient := networkstatus.NewNetworkStatusClientWithBaseURI(endpoint)
	configureAuthFunc(&networkStatusClient.Client)

	notificationClient := notification.NewNotificationClientWithBaseURI(endpoint)
	configureAuthFunc(&notificationClient.Client)

	notificationRecipientEmailClient := notificationrecipientemail.NewNotificationRecipientEmailClientWithBaseURI(endpoint)
	configureAuthFunc(&notificationRecipientEmailClient.Client)

	notificationRecipientUserClient := notificationrecipientuser.NewNotificationRecipientUserClientWithBaseURI(endpoint)
	configureAuthFunc(&notificationRecipientUserClient.Client)

	openidConnectProviderClient := openidconnectprovider.NewOpenidConnectProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&openidConnectProviderClient.Client)

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	performConnectivityCheckClient := performconnectivitycheck.NewPerformConnectivityCheckClientWithBaseURI(endpoint)
	configureAuthFunc(&performConnectivityCheckClient.Client)

	policyClient := policy.NewPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&policyClient.Client)

	policyDescriptionClient := policydescription.NewPolicyDescriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&policyDescriptionClient.Client)

	policyFragmentClient := policyfragment.NewPolicyFragmentClientWithBaseURI(endpoint)
	configureAuthFunc(&policyFragmentClient.Client)

	portalConfigClient := portalconfig.NewPortalConfigClientWithBaseURI(endpoint)
	configureAuthFunc(&portalConfigClient.Client)

	portalRevisionClient := portalrevision.NewPortalRevisionClientWithBaseURI(endpoint)
	configureAuthFunc(&portalRevisionClient.Client)

	portalSettingsClient := portalsettings.NewPortalSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&portalSettingsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	productApiClient := productapi.NewProductApiClientWithBaseURI(endpoint)
	configureAuthFunc(&productApiClient.Client)

	productClient := product.NewProductClientWithBaseURI(endpoint)
	configureAuthFunc(&productClient.Client)

	productGroupClient := productgroup.NewProductGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&productGroupClient.Client)

	productPolicyClient := productpolicy.NewProductPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&productPolicyClient.Client)

	productSubscriptionClient := productsubscription.NewProductSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&productSubscriptionClient.Client)

	productTagClient := producttag.NewProductTagClientWithBaseURI(endpoint)
	configureAuthFunc(&productTagClient.Client)

	productWikiClient := productwiki.NewProductWikiClientWithBaseURI(endpoint)
	configureAuthFunc(&productWikiClient.Client)

	productsByTagClient := productsbytag.NewProductsByTagClientWithBaseURI(endpoint)
	configureAuthFunc(&productsByTagClient.Client)

	quotaByCounterKeysClient := quotabycounterkeys.NewQuotaByCounterKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&quotaByCounterKeysClient.Client)

	quotaByPeriodKeysClient := quotabyperiodkeys.NewQuotaByPeriodKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&quotaByPeriodKeysClient.Client)

	regionClient := region.NewRegionClientWithBaseURI(endpoint)
	configureAuthFunc(&regionClient.Client)

	reportsClient := reports.NewReportsClientWithBaseURI(endpoint)
	configureAuthFunc(&reportsClient.Client)

	schemaClient := schema.NewSchemaClientWithBaseURI(endpoint)
	configureAuthFunc(&schemaClient.Client)

	signInSettingsClient := signinsettings.NewSignInSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&signInSettingsClient.Client)

	signUpSettingsClient := signupsettings.NewSignUpSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&signUpSettingsClient.Client)

	skusClient := skus.NewSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&skusClient.Client)

	subscriptionClient := subscription.NewSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionClient.Client)

	tagClient := tag.NewTagClientWithBaseURI(endpoint)
	configureAuthFunc(&tagClient.Client)

	tagResourceClient := tagresource.NewTagResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&tagResourceClient.Client)

	tenantAccessClient := tenantaccess.NewTenantAccessClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantAccessClient.Client)

	tenantAccessGitClient := tenantaccessgit.NewTenantAccessGitClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantAccessGitClient.Client)

	tenantConfigurationClient := tenantconfiguration.NewTenantConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantConfigurationClient.Client)

	tenantConfigurationSyncStateClient := tenantconfigurationsyncstate.NewTenantConfigurationSyncStateClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantConfigurationSyncStateClient.Client)

	tenantSettingsClient := tenantsettings.NewTenantSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantSettingsClient.Client)

	userClient := user.NewUserClientWithBaseURI(endpoint)
	configureAuthFunc(&userClient.Client)

	userConfirmationPasswordSendClient := userconfirmationpasswordsend.NewUserConfirmationPasswordSendClientWithBaseURI(endpoint)
	configureAuthFunc(&userConfirmationPasswordSendClient.Client)

	userGroupClient := usergroup.NewUserGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&userGroupClient.Client)

	userIdentityClient := useridentity.NewUserIdentityClientWithBaseURI(endpoint)
	configureAuthFunc(&userIdentityClient.Client)

	userSubscriptionClient := usersubscription.NewUserSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&userSubscriptionClient.Client)

	userTokenClient := usertoken.NewUserTokenClientWithBaseURI(endpoint)
	configureAuthFunc(&userTokenClient.Client)

	usersClient := users.NewUsersClientWithBaseURI(endpoint)
	configureAuthFunc(&usersClient.Client)

	return Client{
		Api:                                  &apiClient,
		ApiDiagnostic:                        &apiDiagnosticClient,
		ApiIssue:                             &apiIssueClient,
		ApiIssueAttachment:                   &apiIssueAttachmentClient,
		ApiIssueComment:                      &apiIssueCommentClient,
		ApiManagementService:                 &apiManagementServiceClient,
		ApiManagementServiceSkus:             &apiManagementServiceSkusClient,
		ApiOperation:                         &apiOperationClient,
		ApiOperationPolicy:                   &apiOperationPolicyClient,
		ApiOperationTag:                      &apiOperationTagClient,
		ApiOperationsByTag:                   &apiOperationsByTagClient,
		ApiPolicy:                            &apiPolicyClient,
		ApiProduct:                           &apiProductClient,
		ApiRelease:                           &apiReleaseClient,
		ApiRevision:                          &apiRevisionClient,
		ApiSchema:                            &apiSchemaClient,
		ApiTag:                               &apiTagClient,
		ApiTagDescription:                    &apiTagDescriptionClient,
		ApiVersionSet:                        &apiVersionSetClient,
		ApiVersionSets:                       &apiVersionSetsClient,
		ApiWiki:                              &apiWikiClient,
		ApisByTag:                            &apisByTagClient,
		Authorization:                        &authorizationClient,
		AuthorizationAccessPolicy:            &authorizationAccessPolicyClient,
		AuthorizationConfirmConsentCode:      &authorizationConfirmConsentCodeClient,
		AuthorizationLoginLinks:              &authorizationLoginLinksClient,
		AuthorizationProvider:                &authorizationProviderClient,
		AuthorizationServer:                  &authorizationServerClient,
		Authorizations:                       &authorizationsClient,
		Backend:                              &backendClient,
		BackendReconnect:                     &backendReconnectClient,
		Cache:                                &cacheClient,
		Certificate:                          &certificateClient,
		ContentType:                          &contentTypeClient,
		ContentTypeContentItem:               &contentTypeContentItemClient,
		DelegationSettings:                   &delegationSettingsClient,
		DeletedService:                       &deletedServiceClient,
		Diagnostic:                           &diagnosticClient,
		DocumentationResource:                &documentationResourceClient,
		EmailTemplate:                        &emailTemplateClient,
		EmailTemplates:                       &emailTemplatesClient,
		Gateway:                              &gatewayClient,
		GatewayApi:                           &gatewayApiClient,
		GatewayCertificateAuthority:          &gatewayCertificateAuthorityClient,
		GatewayGenerateToken:                 &gatewayGenerateTokenClient,
		GatewayHostnameConfiguration:         &gatewayHostnameConfigurationClient,
		GatewayListKeys:                      &gatewayListKeysClient,
		GatewayRegenerateKey:                 &gatewayRegenerateKeyClient,
		GraphQLApiResolver:                   &graphQLApiResolverClient,
		GraphQLApiResolverPolicy:             &graphQLApiResolverPolicyClient,
		Group:                                &groupClient,
		GroupUser:                            &groupUserClient,
		IdentityProvider:                     &identityProviderClient,
		Issue:                                &issueClient,
		Logger:                               &loggerClient,
		NamedValue:                           &namedValueClient,
		NetworkStatus:                        &networkStatusClient,
		Notification:                         &notificationClient,
		NotificationRecipientEmail:           &notificationRecipientEmailClient,
		NotificationRecipientUser:            &notificationRecipientUserClient,
		OpenidConnectProvider:                &openidConnectProviderClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PerformConnectivityCheck:             &performConnectivityCheckClient,
		Policy:                               &policyClient,
		PolicyDescription:                    &policyDescriptionClient,
		PolicyFragment:                       &policyFragmentClient,
		PortalConfig:                         &portalConfigClient,
		PortalRevision:                       &portalRevisionClient,
		PortalSettings:                       &portalSettingsClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		Product:                              &productClient,
		ProductApi:                           &productApiClient,
		ProductGroup:                         &productGroupClient,
		ProductPolicy:                        &productPolicyClient,
		ProductSubscription:                  &productSubscriptionClient,
		ProductTag:                           &productTagClient,
		ProductWiki:                          &productWikiClient,
		ProductsByTag:                        &productsByTagClient,
		QuotaByCounterKeys:                   &quotaByCounterKeysClient,
		QuotaByPeriodKeys:                    &quotaByPeriodKeysClient,
		Region:                               &regionClient,
		Reports:                              &reportsClient,
		Schema:                               &schemaClient,
		SignInSettings:                       &signInSettingsClient,
		SignUpSettings:                       &signUpSettingsClient,
		Skus:                                 &skusClient,
		Subscription:                         &subscriptionClient,
		Tag:                                  &tagClient,
		TagResource:                          &tagResourceClient,
		TenantAccess:                         &tenantAccessClient,
		TenantAccessGit:                      &tenantAccessGitClient,
		TenantConfiguration:                  &tenantConfigurationClient,
		TenantConfigurationSyncState:         &tenantConfigurationSyncStateClient,
		TenantSettings:                       &tenantSettingsClient,
		User:                                 &userClient,
		UserConfirmationPasswordSend:         &userConfirmationPasswordSendClient,
		UserGroup:                            &userGroupClient,
		UserIdentity:                         &userIdentityClient,
		UserSubscription:                     &userSubscriptionClient,
		UserToken:                            &userTokenClient,
		Users:                                &usersClient,
	}
}
