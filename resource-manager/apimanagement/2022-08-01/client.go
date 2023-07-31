package v2022_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	apiClient, err := api.NewApiClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Api client: %+v", err)
	}
	configureFunc(apiClient.Client)

	apiDiagnosticClient, err := apidiagnostic.NewApiDiagnosticClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiDiagnostic client: %+v", err)
	}
	configureFunc(apiDiagnosticClient.Client)

	apiIssueAttachmentClient, err := apiissueattachment.NewApiIssueAttachmentClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiIssueAttachment client: %+v", err)
	}
	configureFunc(apiIssueAttachmentClient.Client)

	apiIssueClient, err := apiissue.NewApiIssueClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiIssue client: %+v", err)
	}
	configureFunc(apiIssueClient.Client)

	apiIssueCommentClient, err := apiissuecomment.NewApiIssueCommentClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiIssueComment client: %+v", err)
	}
	configureFunc(apiIssueCommentClient.Client)

	apiManagementServiceClient, err := apimanagementservice.NewApiManagementServiceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiManagementService client: %+v", err)
	}
	configureFunc(apiManagementServiceClient.Client)

	apiManagementServiceSkusClient, err := apimanagementserviceskus.NewApiManagementServiceSkusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiManagementServiceSkus client: %+v", err)
	}
	configureFunc(apiManagementServiceSkusClient.Client)

	apiOperationClient, err := apioperation.NewApiOperationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiOperation client: %+v", err)
	}
	configureFunc(apiOperationClient.Client)

	apiOperationPolicyClient, err := apioperationpolicy.NewApiOperationPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiOperationPolicy client: %+v", err)
	}
	configureFunc(apiOperationPolicyClient.Client)

	apiOperationTagClient, err := apioperationtag.NewApiOperationTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiOperationTag client: %+v", err)
	}
	configureFunc(apiOperationTagClient.Client)

	apiOperationsByTagClient, err := apioperationsbytag.NewApiOperationsByTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiOperationsByTag client: %+v", err)
	}
	configureFunc(apiOperationsByTagClient.Client)

	apiPolicyClient, err := apipolicy.NewApiPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiPolicy client: %+v", err)
	}
	configureFunc(apiPolicyClient.Client)

	apiProductClient, err := apiproduct.NewApiProductClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiProduct client: %+v", err)
	}
	configureFunc(apiProductClient.Client)

	apiReleaseClient, err := apirelease.NewApiReleaseClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiRelease client: %+v", err)
	}
	configureFunc(apiReleaseClient.Client)

	apiRevisionClient, err := apirevision.NewApiRevisionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiRevision client: %+v", err)
	}
	configureFunc(apiRevisionClient.Client)

	apiSchemaClient, err := apischema.NewApiSchemaClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiSchema client: %+v", err)
	}
	configureFunc(apiSchemaClient.Client)

	apiTagClient, err := apitag.NewApiTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiTag client: %+v", err)
	}
	configureFunc(apiTagClient.Client)

	apiTagDescriptionClient, err := apitagdescription.NewApiTagDescriptionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiTagDescription client: %+v", err)
	}
	configureFunc(apiTagDescriptionClient.Client)

	apiVersionSetClient, err := apiversionset.NewApiVersionSetClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiVersionSet client: %+v", err)
	}
	configureFunc(apiVersionSetClient.Client)

	apiVersionSetsClient, err := apiversionsets.NewApiVersionSetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiVersionSets client: %+v", err)
	}
	configureFunc(apiVersionSetsClient.Client)

	apiWikiClient, err := apiwiki.NewApiWikiClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiWiki client: %+v", err)
	}
	configureFunc(apiWikiClient.Client)

	apisByTagClient, err := apisbytag.NewApisByTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApisByTag client: %+v", err)
	}
	configureFunc(apisByTagClient.Client)

	authorizationAccessPolicyClient, err := authorizationaccesspolicy.NewAuthorizationAccessPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationAccessPolicy client: %+v", err)
	}
	configureFunc(authorizationAccessPolicyClient.Client)

	authorizationClient, err := authorization.NewAuthorizationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Authorization client: %+v", err)
	}
	configureFunc(authorizationClient.Client)

	authorizationConfirmConsentCodeClient, err := authorizationconfirmconsentcode.NewAuthorizationConfirmConsentCodeClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationConfirmConsentCode client: %+v", err)
	}
	configureFunc(authorizationConfirmConsentCodeClient.Client)

	authorizationLoginLinksClient, err := authorizationloginlinks.NewAuthorizationLoginLinksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationLoginLinks client: %+v", err)
	}
	configureFunc(authorizationLoginLinksClient.Client)

	authorizationProviderClient, err := authorizationprovider.NewAuthorizationProviderClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationProvider client: %+v", err)
	}
	configureFunc(authorizationProviderClient.Client)

	authorizationServerClient, err := authorizationserver.NewAuthorizationServerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationServer client: %+v", err)
	}
	configureFunc(authorizationServerClient.Client)

	authorizationsClient, err := authorizations.NewAuthorizationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Authorizations client: %+v", err)
	}
	configureFunc(authorizationsClient.Client)

	backendClient, err := backend.NewBackendClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Backend client: %+v", err)
	}
	configureFunc(backendClient.Client)

	backendReconnectClient, err := backendreconnect.NewBackendReconnectClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BackendReconnect client: %+v", err)
	}
	configureFunc(backendReconnectClient.Client)

	cacheClient, err := cache.NewCacheClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Cache client: %+v", err)
	}
	configureFunc(cacheClient.Client)

	certificateClient, err := certificate.NewCertificateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Certificate client: %+v", err)
	}
	configureFunc(certificateClient.Client)

	contentTypeClient, err := contenttype.NewContentTypeClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ContentType client: %+v", err)
	}
	configureFunc(contentTypeClient.Client)

	contentTypeContentItemClient, err := contenttypecontentitem.NewContentTypeContentItemClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ContentTypeContentItem client: %+v", err)
	}
	configureFunc(contentTypeContentItemClient.Client)

	delegationSettingsClient, err := delegationsettings.NewDelegationSettingsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DelegationSettings client: %+v", err)
	}
	configureFunc(delegationSettingsClient.Client)

	deletedServiceClient, err := deletedservice.NewDeletedServiceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DeletedService client: %+v", err)
	}
	configureFunc(deletedServiceClient.Client)

	diagnosticClient, err := diagnostic.NewDiagnosticClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Diagnostic client: %+v", err)
	}
	configureFunc(diagnosticClient.Client)

	documentationResourceClient, err := documentationresource.NewDocumentationResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DocumentationResource client: %+v", err)
	}
	configureFunc(documentationResourceClient.Client)

	emailTemplateClient, err := emailtemplate.NewEmailTemplateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EmailTemplate client: %+v", err)
	}
	configureFunc(emailTemplateClient.Client)

	emailTemplatesClient, err := emailtemplates.NewEmailTemplatesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EmailTemplates client: %+v", err)
	}
	configureFunc(emailTemplatesClient.Client)

	gatewayApiClient, err := gatewayapi.NewGatewayApiClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GatewayApi client: %+v", err)
	}
	configureFunc(gatewayApiClient.Client)

	gatewayCertificateAuthorityClient, err := gatewaycertificateauthority.NewGatewayCertificateAuthorityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GatewayCertificateAuthority client: %+v", err)
	}
	configureFunc(gatewayCertificateAuthorityClient.Client)

	gatewayClient, err := gateway.NewGatewayClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Gateway client: %+v", err)
	}
	configureFunc(gatewayClient.Client)

	gatewayGenerateTokenClient, err := gatewaygeneratetoken.NewGatewayGenerateTokenClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GatewayGenerateToken client: %+v", err)
	}
	configureFunc(gatewayGenerateTokenClient.Client)

	gatewayHostnameConfigurationClient, err := gatewayhostnameconfiguration.NewGatewayHostnameConfigurationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GatewayHostnameConfiguration client: %+v", err)
	}
	configureFunc(gatewayHostnameConfigurationClient.Client)

	gatewayListKeysClient, err := gatewaylistkeys.NewGatewayListKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GatewayListKeys client: %+v", err)
	}
	configureFunc(gatewayListKeysClient.Client)

	gatewayRegenerateKeyClient, err := gatewayregeneratekey.NewGatewayRegenerateKeyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GatewayRegenerateKey client: %+v", err)
	}
	configureFunc(gatewayRegenerateKeyClient.Client)

	graphQLApiResolverClient, err := graphqlapiresolver.NewGraphQLApiResolverClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GraphQLApiResolver client: %+v", err)
	}
	configureFunc(graphQLApiResolverClient.Client)

	graphQLApiResolverPolicyClient, err := graphqlapiresolverpolicy.NewGraphQLApiResolverPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GraphQLApiResolverPolicy client: %+v", err)
	}
	configureFunc(graphQLApiResolverPolicyClient.Client)

	groupClient, err := group.NewGroupClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Group client: %+v", err)
	}
	configureFunc(groupClient.Client)

	groupUserClient, err := groupuser.NewGroupUserClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GroupUser client: %+v", err)
	}
	configureFunc(groupUserClient.Client)

	identityProviderClient, err := identityprovider.NewIdentityProviderClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building IdentityProvider client: %+v", err)
	}
	configureFunc(identityProviderClient.Client)

	issueClient, err := issue.NewIssueClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Issue client: %+v", err)
	}
	configureFunc(issueClient.Client)

	loggerClient, err := logger.NewLoggerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Logger client: %+v", err)
	}
	configureFunc(loggerClient.Client)

	namedValueClient, err := namedvalue.NewNamedValueClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NamedValue client: %+v", err)
	}
	configureFunc(namedValueClient.Client)

	networkStatusClient, err := networkstatus.NewNetworkStatusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NetworkStatus client: %+v", err)
	}
	configureFunc(networkStatusClient.Client)

	notificationClient, err := notification.NewNotificationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Notification client: %+v", err)
	}
	configureFunc(notificationClient.Client)

	notificationRecipientEmailClient, err := notificationrecipientemail.NewNotificationRecipientEmailClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NotificationRecipientEmail client: %+v", err)
	}
	configureFunc(notificationRecipientEmailClient.Client)

	notificationRecipientUserClient, err := notificationrecipientuser.NewNotificationRecipientUserClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NotificationRecipientUser client: %+v", err)
	}
	configureFunc(notificationRecipientUserClient.Client)

	openidConnectProviderClient, err := openidconnectprovider.NewOpenidConnectProviderClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OpenidConnectProvider client: %+v", err)
	}
	configureFunc(openidConnectProviderClient.Client)

	outboundNetworkDependenciesEndpointsClient, err := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OutboundNetworkDependenciesEndpoints client: %+v", err)
	}
	configureFunc(outboundNetworkDependenciesEndpointsClient.Client)

	performConnectivityCheckClient, err := performconnectivitycheck.NewPerformConnectivityCheckClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PerformConnectivityCheck client: %+v", err)
	}
	configureFunc(performConnectivityCheckClient.Client)

	policyClient, err := policy.NewPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %+v", err)
	}
	configureFunc(policyClient.Client)

	policyDescriptionClient, err := policydescription.NewPolicyDescriptionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDescription client: %+v", err)
	}
	configureFunc(policyDescriptionClient.Client)

	policyFragmentClient, err := policyfragment.NewPolicyFragmentClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PolicyFragment client: %+v", err)
	}
	configureFunc(policyFragmentClient.Client)

	portalConfigClient, err := portalconfig.NewPortalConfigClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PortalConfig client: %+v", err)
	}
	configureFunc(portalConfigClient.Client)

	portalRevisionClient, err := portalrevision.NewPortalRevisionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PortalRevision client: %+v", err)
	}
	configureFunc(portalRevisionClient.Client)

	portalSettingsClient, err := portalsettings.NewPortalSettingsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PortalSettings client: %+v", err)
	}
	configureFunc(portalSettingsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	productApiClient, err := productapi.NewProductApiClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductApi client: %+v", err)
	}
	configureFunc(productApiClient.Client)

	productClient, err := product.NewProductClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Product client: %+v", err)
	}
	configureFunc(productClient.Client)

	productGroupClient, err := productgroup.NewProductGroupClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductGroup client: %+v", err)
	}
	configureFunc(productGroupClient.Client)

	productPolicyClient, err := productpolicy.NewProductPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductPolicy client: %+v", err)
	}
	configureFunc(productPolicyClient.Client)

	productSubscriptionClient, err := productsubscription.NewProductSubscriptionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductSubscription client: %+v", err)
	}
	configureFunc(productSubscriptionClient.Client)

	productTagClient, err := producttag.NewProductTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductTag client: %+v", err)
	}
	configureFunc(productTagClient.Client)

	productWikiClient, err := productwiki.NewProductWikiClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductWiki client: %+v", err)
	}
	configureFunc(productWikiClient.Client)

	productsByTagClient, err := productsbytag.NewProductsByTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProductsByTag client: %+v", err)
	}
	configureFunc(productsByTagClient.Client)

	quotaByCounterKeysClient, err := quotabycounterkeys.NewQuotaByCounterKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building QuotaByCounterKeys client: %+v", err)
	}
	configureFunc(quotaByCounterKeysClient.Client)

	quotaByPeriodKeysClient, err := quotabyperiodkeys.NewQuotaByPeriodKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building QuotaByPeriodKeys client: %+v", err)
	}
	configureFunc(quotaByPeriodKeysClient.Client)

	regionClient, err := region.NewRegionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Region client: %+v", err)
	}
	configureFunc(regionClient.Client)

	reportsClient, err := reports.NewReportsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Reports client: %+v", err)
	}
	configureFunc(reportsClient.Client)

	schemaClient, err := schema.NewSchemaClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Schema client: %+v", err)
	}
	configureFunc(schemaClient.Client)

	signInSettingsClient, err := signinsettings.NewSignInSettingsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SignInSettings client: %+v", err)
	}
	configureFunc(signInSettingsClient.Client)

	signUpSettingsClient, err := signupsettings.NewSignUpSettingsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SignUpSettings client: %+v", err)
	}
	configureFunc(signUpSettingsClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	subscriptionClient, err := subscription.NewSubscriptionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Subscription client: %+v", err)
	}
	configureFunc(subscriptionClient.Client)

	tagClient, err := tag.NewTagClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Tag client: %+v", err)
	}
	configureFunc(tagClient.Client)

	tagResourceClient, err := tagresource.NewTagResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TagResource client: %+v", err)
	}
	configureFunc(tagResourceClient.Client)

	tenantAccessClient, err := tenantaccess.NewTenantAccessClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TenantAccess client: %+v", err)
	}
	configureFunc(tenantAccessClient.Client)

	tenantAccessGitClient, err := tenantaccessgit.NewTenantAccessGitClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TenantAccessGit client: %+v", err)
	}
	configureFunc(tenantAccessGitClient.Client)

	tenantConfigurationClient, err := tenantconfiguration.NewTenantConfigurationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TenantConfiguration client: %+v", err)
	}
	configureFunc(tenantConfigurationClient.Client)

	tenantConfigurationSyncStateClient, err := tenantconfigurationsyncstate.NewTenantConfigurationSyncStateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TenantConfigurationSyncState client: %+v", err)
	}
	configureFunc(tenantConfigurationSyncStateClient.Client)

	tenantSettingsClient, err := tenantsettings.NewTenantSettingsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TenantSettings client: %+v", err)
	}
	configureFunc(tenantSettingsClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building User client: %+v", err)
	}
	configureFunc(userClient.Client)

	userConfirmationPasswordSendClient, err := userconfirmationpasswordsend.NewUserConfirmationPasswordSendClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UserConfirmationPasswordSend client: %+v", err)
	}
	configureFunc(userConfirmationPasswordSendClient.Client)

	userGroupClient, err := usergroup.NewUserGroupClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UserGroup client: %+v", err)
	}
	configureFunc(userGroupClient.Client)

	userIdentityClient, err := useridentity.NewUserIdentityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UserIdentity client: %+v", err)
	}
	configureFunc(userIdentityClient.Client)

	userSubscriptionClient, err := usersubscription.NewUserSubscriptionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UserSubscription client: %+v", err)
	}
	configureFunc(userSubscriptionClient.Client)

	userTokenClient, err := usertoken.NewUserTokenClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UserToken client: %+v", err)
	}
	configureFunc(userTokenClient.Client)

	usersClient, err := users.NewUsersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Users client: %+v", err)
	}
	configureFunc(usersClient.Client)

	return &Client{
		Api:                                  apiClient,
		ApiDiagnostic:                        apiDiagnosticClient,
		ApiIssue:                             apiIssueClient,
		ApiIssueAttachment:                   apiIssueAttachmentClient,
		ApiIssueComment:                      apiIssueCommentClient,
		ApiManagementService:                 apiManagementServiceClient,
		ApiManagementServiceSkus:             apiManagementServiceSkusClient,
		ApiOperation:                         apiOperationClient,
		ApiOperationPolicy:                   apiOperationPolicyClient,
		ApiOperationTag:                      apiOperationTagClient,
		ApiOperationsByTag:                   apiOperationsByTagClient,
		ApiPolicy:                            apiPolicyClient,
		ApiProduct:                           apiProductClient,
		ApiRelease:                           apiReleaseClient,
		ApiRevision:                          apiRevisionClient,
		ApiSchema:                            apiSchemaClient,
		ApiTag:                               apiTagClient,
		ApiTagDescription:                    apiTagDescriptionClient,
		ApiVersionSet:                        apiVersionSetClient,
		ApiVersionSets:                       apiVersionSetsClient,
		ApiWiki:                              apiWikiClient,
		ApisByTag:                            apisByTagClient,
		Authorization:                        authorizationClient,
		AuthorizationAccessPolicy:            authorizationAccessPolicyClient,
		AuthorizationConfirmConsentCode:      authorizationConfirmConsentCodeClient,
		AuthorizationLoginLinks:              authorizationLoginLinksClient,
		AuthorizationProvider:                authorizationProviderClient,
		AuthorizationServer:                  authorizationServerClient,
		Authorizations:                       authorizationsClient,
		Backend:                              backendClient,
		BackendReconnect:                     backendReconnectClient,
		Cache:                                cacheClient,
		Certificate:                          certificateClient,
		ContentType:                          contentTypeClient,
		ContentTypeContentItem:               contentTypeContentItemClient,
		DelegationSettings:                   delegationSettingsClient,
		DeletedService:                       deletedServiceClient,
		Diagnostic:                           diagnosticClient,
		DocumentationResource:                documentationResourceClient,
		EmailTemplate:                        emailTemplateClient,
		EmailTemplates:                       emailTemplatesClient,
		Gateway:                              gatewayClient,
		GatewayApi:                           gatewayApiClient,
		GatewayCertificateAuthority:          gatewayCertificateAuthorityClient,
		GatewayGenerateToken:                 gatewayGenerateTokenClient,
		GatewayHostnameConfiguration:         gatewayHostnameConfigurationClient,
		GatewayListKeys:                      gatewayListKeysClient,
		GatewayRegenerateKey:                 gatewayRegenerateKeyClient,
		GraphQLApiResolver:                   graphQLApiResolverClient,
		GraphQLApiResolverPolicy:             graphQLApiResolverPolicyClient,
		Group:                                groupClient,
		GroupUser:                            groupUserClient,
		IdentityProvider:                     identityProviderClient,
		Issue:                                issueClient,
		Logger:                               loggerClient,
		NamedValue:                           namedValueClient,
		NetworkStatus:                        networkStatusClient,
		Notification:                         notificationClient,
		NotificationRecipientEmail:           notificationRecipientEmailClient,
		NotificationRecipientUser:            notificationRecipientUserClient,
		OpenidConnectProvider:                openidConnectProviderClient,
		OutboundNetworkDependenciesEndpoints: outboundNetworkDependenciesEndpointsClient,
		PerformConnectivityCheck:             performConnectivityCheckClient,
		Policy:                               policyClient,
		PolicyDescription:                    policyDescriptionClient,
		PolicyFragment:                       policyFragmentClient,
		PortalConfig:                         portalConfigClient,
		PortalRevision:                       portalRevisionClient,
		PortalSettings:                       portalSettingsClient,
		PrivateEndpointConnections:           privateEndpointConnectionsClient,
		Product:                              productClient,
		ProductApi:                           productApiClient,
		ProductGroup:                         productGroupClient,
		ProductPolicy:                        productPolicyClient,
		ProductSubscription:                  productSubscriptionClient,
		ProductTag:                           productTagClient,
		ProductWiki:                          productWikiClient,
		ProductsByTag:                        productsByTagClient,
		QuotaByCounterKeys:                   quotaByCounterKeysClient,
		QuotaByPeriodKeys:                    quotaByPeriodKeysClient,
		Region:                               regionClient,
		Reports:                              reportsClient,
		Schema:                               schemaClient,
		SignInSettings:                       signInSettingsClient,
		SignUpSettings:                       signUpSettingsClient,
		Skus:                                 skusClient,
		Subscription:                         subscriptionClient,
		Tag:                                  tagClient,
		TagResource:                          tagResourceClient,
		TenantAccess:                         tenantAccessClient,
		TenantAccessGit:                      tenantAccessGitClient,
		TenantConfiguration:                  tenantConfigurationClient,
		TenantConfigurationSyncState:         tenantConfigurationSyncStateClient,
		TenantSettings:                       tenantSettingsClient,
		User:                                 userClient,
		UserConfirmationPasswordSend:         userConfirmationPasswordSendClient,
		UserGroup:                            userGroupClient,
		UserIdentity:                         userIdentityClient,
		UserSubscription:                     userSubscriptionClient,
		UserToken:                            userTokenClient,
		Users:                                usersClient,
	}, nil
}
