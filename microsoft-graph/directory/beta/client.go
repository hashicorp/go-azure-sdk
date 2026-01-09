package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitdeletedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/assigndirectoryauthenticationmethoddevicehardwareoathdeviceto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/attributeset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/authenticationmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/authenticationmethoddevicehardwareoathdevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/authenticationmethoddevicehardwareoathdeviceassigntomailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/authenticationmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthoritycertificatebasedapplicationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthoritymutualtlsoauthconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/customsecurityattributedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/customsecurityattributedefinitionallowedvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/deleteditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/devicelocalcredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/externaluserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/featurerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/featurerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/federationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/impactedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/inboundshareduserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/onpremisessynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/outboundshareduserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/outboundshareduserprofiletenant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/pendingexternaluserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/publickeyinfrastructure"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/publickeyinfrastructurecertificatebasedauthconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/publickeyinfrastructurecertificatebasedauthconfigurationcertificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationimpactedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/sharedemaildomain"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/subscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/template"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/templatedevicetemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/templatedevicetemplatedeviceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/templatedevicetemplateowner"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdministrativeUnit                                                                      *administrativeunit.AdministrativeUnitClient
	AdministrativeUnitDeletedMember                                                         *administrativeunitdeletedmember.AdministrativeUnitDeletedMemberClient
	AdministrativeUnitExtension                                                             *administrativeunitextension.AdministrativeUnitExtensionClient
	AdministrativeUnitMember                                                                *administrativeunitmember.AdministrativeUnitMemberClient
	AdministrativeUnitScopedRoleMember                                                      *administrativeunitscopedrolemember.AdministrativeUnitScopedRoleMemberClient
	AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceTo                           *assigndirectoryauthenticationmethoddevicehardwareoathdeviceto.AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient
	AttributeSet                                                                            *attributeset.AttributeSetClient
	AuthenticationMethodDevice                                                              *authenticationmethoddevice.AuthenticationMethodDeviceClient
	AuthenticationMethodDeviceHardwareOathDevice                                            *authenticationmethoddevicehardwareoathdevice.AuthenticationMethodDeviceHardwareOathDeviceClient
	AuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSetting                      *authenticationmethoddevicehardwareoathdeviceassigntomailboxsetting.AuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient
	AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningError            *authenticationmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror.AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient
	CertificateAuthority                                                                    *certificateauthority.CertificateAuthorityClient
	CertificateAuthorityCertificateBasedApplicationConfiguration                            *certificateauthoritycertificatebasedapplicationconfiguration.CertificateAuthorityCertificateBasedApplicationConfigurationClient
	CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority *certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority.CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient
	CertificateAuthorityMutualTlsOauthConfiguration                                         *certificateauthoritymutualtlsoauthconfiguration.CertificateAuthorityMutualTlsOauthConfigurationClient
	CustomSecurityAttributeDefinition                                                       *customsecurityattributedefinition.CustomSecurityAttributeDefinitionClient
	CustomSecurityAttributeDefinitionAllowedValue                                           *customsecurityattributedefinitionallowedvalue.CustomSecurityAttributeDefinitionAllowedValueClient
	DeletedItem                                                                             *deleteditem.DeletedItemClient
	DeviceLocalCredential                                                                   *devicelocalcredential.DeviceLocalCredentialClient
	Directory                                                                               *directory.DirectoryClient
	ExternalUserProfile                                                                     *externaluserprofile.ExternalUserProfileClient
	FeatureRolloutPolicy                                                                    *featurerolloutpolicy.FeatureRolloutPolicyClient
	FeatureRolloutPolicyAppliesTo                                                           *featurerolloutpolicyappliesto.FeatureRolloutPolicyAppliesToClient
	FederationConfiguration                                                                 *federationconfiguration.FederationConfigurationClient
	ImpactedResource                                                                        *impactedresource.ImpactedResourceClient
	InboundSharedUserProfile                                                                *inboundshareduserprofile.InboundSharedUserProfileClient
	OnPremisesSynchronization                                                               *onpremisessynchronization.OnPremisesSynchronizationClient
	OutboundSharedUserProfile                                                               *outboundshareduserprofile.OutboundSharedUserProfileClient
	OutboundSharedUserProfileTenant                                                         *outboundshareduserprofiletenant.OutboundSharedUserProfileTenantClient
	PendingExternalUserProfile                                                              *pendingexternaluserprofile.PendingExternalUserProfileClient
	PublicKeyInfrastructure                                                                 *publickeyinfrastructure.PublicKeyInfrastructureClient
	PublicKeyInfrastructureCertificateBasedAuthConfiguration                                *publickeyinfrastructurecertificatebasedauthconfiguration.PublicKeyInfrastructureCertificateBasedAuthConfigurationClient
	PublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthority            *publickeyinfrastructurecertificatebasedauthconfigurationcertificateauthority.PublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityClient
	Recommendation                                                                          *recommendation.RecommendationClient
	RecommendationConfiguration                                                             *recommendationconfiguration.RecommendationConfigurationClient
	RecommendationImpactedResource                                                          *recommendationimpactedresource.RecommendationImpactedResourceClient
	SharedEmailDomain                                                                       *sharedemaildomain.SharedEmailDomainClient
	Subscription                                                                            *subscription.SubscriptionClient
	Template                                                                                *template.TemplateClient
	TemplateDeviceTemplate                                                                  *templatedevicetemplate.TemplateDeviceTemplateClient
	TemplateDeviceTemplateDeviceInstance                                                    *templatedevicetemplatedeviceinstance.TemplateDeviceTemplateDeviceInstanceClient
	TemplateDeviceTemplateOwner                                                             *templatedevicetemplateowner.TemplateDeviceTemplateOwnerClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	administrativeUnitClient, err := administrativeunit.NewAdministrativeUnitClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnit client: %+v", err)
	}
	configureFunc(administrativeUnitClient.Client)

	administrativeUnitDeletedMemberClient, err := administrativeunitdeletedmember.NewAdministrativeUnitDeletedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitDeletedMember client: %+v", err)
	}
	configureFunc(administrativeUnitDeletedMemberClient.Client)

	administrativeUnitExtensionClient, err := administrativeunitextension.NewAdministrativeUnitExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitExtension client: %+v", err)
	}
	configureFunc(administrativeUnitExtensionClient.Client)

	administrativeUnitMemberClient, err := administrativeunitmember.NewAdministrativeUnitMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitMember client: %+v", err)
	}
	configureFunc(administrativeUnitMemberClient.Client)

	administrativeUnitScopedRoleMemberClient, err := administrativeunitscopedrolemember.NewAdministrativeUnitScopedRoleMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitScopedRoleMember client: %+v", err)
	}
	configureFunc(administrativeUnitScopedRoleMemberClient.Client)

	assignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient, err := assigndirectoryauthenticationmethoddevicehardwareoathdeviceto.NewAssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceTo client: %+v", err)
	}
	configureFunc(assignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient.Client)

	attributeSetClient, err := attributeset.NewAttributeSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AttributeSet client: %+v", err)
	}
	configureFunc(attributeSetClient.Client)

	authenticationMethodDeviceClient, err := authenticationmethoddevice.NewAuthenticationMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodDevice client: %+v", err)
	}
	configureFunc(authenticationMethodDeviceClient.Client)

	authenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient, err := authenticationmethoddevicehardwareoathdeviceassigntomailboxsetting.NewAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSetting client: %+v", err)
	}
	configureFunc(authenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient.Client)

	authenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient, err := authenticationmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror.NewAuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningError client: %+v", err)
	}
	configureFunc(authenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient.Client)

	authenticationMethodDeviceHardwareOathDeviceClient, err := authenticationmethoddevicehardwareoathdevice.NewAuthenticationMethodDeviceHardwareOathDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodDeviceHardwareOathDevice client: %+v", err)
	}
	configureFunc(authenticationMethodDeviceHardwareOathDeviceClient.Client)

	certificateAuthorityCertificateBasedApplicationConfigurationClient, err := certificateauthoritycertificatebasedapplicationconfiguration.NewCertificateAuthorityCertificateBasedApplicationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateAuthorityCertificateBasedApplicationConfiguration client: %+v", err)
	}
	configureFunc(certificateAuthorityCertificateBasedApplicationConfigurationClient.Client)

	certificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient, err := certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority.NewCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority client: %+v", err)
	}
	configureFunc(certificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient.Client)

	certificateAuthorityClient, err := certificateauthority.NewCertificateAuthorityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateAuthority client: %+v", err)
	}
	configureFunc(certificateAuthorityClient.Client)

	certificateAuthorityMutualTlsOauthConfigurationClient, err := certificateauthoritymutualtlsoauthconfiguration.NewCertificateAuthorityMutualTlsOauthConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateAuthorityMutualTlsOauthConfiguration client: %+v", err)
	}
	configureFunc(certificateAuthorityMutualTlsOauthConfigurationClient.Client)

	customSecurityAttributeDefinitionAllowedValueClient, err := customsecurityattributedefinitionallowedvalue.NewCustomSecurityAttributeDefinitionAllowedValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomSecurityAttributeDefinitionAllowedValue client: %+v", err)
	}
	configureFunc(customSecurityAttributeDefinitionAllowedValueClient.Client)

	customSecurityAttributeDefinitionClient, err := customsecurityattributedefinition.NewCustomSecurityAttributeDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomSecurityAttributeDefinition client: %+v", err)
	}
	configureFunc(customSecurityAttributeDefinitionClient.Client)

	deletedItemClient, err := deleteditem.NewDeletedItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedItem client: %+v", err)
	}
	configureFunc(deletedItemClient.Client)

	deviceLocalCredentialClient, err := devicelocalcredential.NewDeviceLocalCredentialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceLocalCredential client: %+v", err)
	}
	configureFunc(deviceLocalCredentialClient.Client)

	directoryClient, err := directory.NewDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Directory client: %+v", err)
	}
	configureFunc(directoryClient.Client)

	externalUserProfileClient, err := externaluserprofile.NewExternalUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExternalUserProfile client: %+v", err)
	}
	configureFunc(externalUserProfileClient.Client)

	featureRolloutPolicyAppliesToClient, err := featurerolloutpolicyappliesto.NewFeatureRolloutPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeatureRolloutPolicyAppliesTo client: %+v", err)
	}
	configureFunc(featureRolloutPolicyAppliesToClient.Client)

	featureRolloutPolicyClient, err := featurerolloutpolicy.NewFeatureRolloutPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeatureRolloutPolicy client: %+v", err)
	}
	configureFunc(featureRolloutPolicyClient.Client)

	federationConfigurationClient, err := federationconfiguration.NewFederationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FederationConfiguration client: %+v", err)
	}
	configureFunc(federationConfigurationClient.Client)

	impactedResourceClient, err := impactedresource.NewImpactedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ImpactedResource client: %+v", err)
	}
	configureFunc(impactedResourceClient.Client)

	inboundSharedUserProfileClient, err := inboundshareduserprofile.NewInboundSharedUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InboundSharedUserProfile client: %+v", err)
	}
	configureFunc(inboundSharedUserProfileClient.Client)

	onPremisesSynchronizationClient, err := onpremisessynchronization.NewOnPremisesSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnPremisesSynchronization client: %+v", err)
	}
	configureFunc(onPremisesSynchronizationClient.Client)

	outboundSharedUserProfileClient, err := outboundshareduserprofile.NewOutboundSharedUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundSharedUserProfile client: %+v", err)
	}
	configureFunc(outboundSharedUserProfileClient.Client)

	outboundSharedUserProfileTenantClient, err := outboundshareduserprofiletenant.NewOutboundSharedUserProfileTenantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundSharedUserProfileTenant client: %+v", err)
	}
	configureFunc(outboundSharedUserProfileTenantClient.Client)

	pendingExternalUserProfileClient, err := pendingexternaluserprofile.NewPendingExternalUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingExternalUserProfile client: %+v", err)
	}
	configureFunc(pendingExternalUserProfileClient.Client)

	publicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityClient, err := publickeyinfrastructurecertificatebasedauthconfigurationcertificateauthority.NewPublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthority client: %+v", err)
	}
	configureFunc(publicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityClient.Client)

	publicKeyInfrastructureCertificateBasedAuthConfigurationClient, err := publickeyinfrastructurecertificatebasedauthconfiguration.NewPublicKeyInfrastructureCertificateBasedAuthConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicKeyInfrastructureCertificateBasedAuthConfiguration client: %+v", err)
	}
	configureFunc(publicKeyInfrastructureCertificateBasedAuthConfigurationClient.Client)

	publicKeyInfrastructureClient, err := publickeyinfrastructure.NewPublicKeyInfrastructureClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicKeyInfrastructure client: %+v", err)
	}
	configureFunc(publicKeyInfrastructureClient.Client)

	recommendationClient, err := recommendation.NewRecommendationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Recommendation client: %+v", err)
	}
	configureFunc(recommendationClient.Client)

	recommendationConfigurationClient, err := recommendationconfiguration.NewRecommendationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecommendationConfiguration client: %+v", err)
	}
	configureFunc(recommendationConfigurationClient.Client)

	recommendationImpactedResourceClient, err := recommendationimpactedresource.NewRecommendationImpactedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecommendationImpactedResource client: %+v", err)
	}
	configureFunc(recommendationImpactedResourceClient.Client)

	sharedEmailDomainClient, err := sharedemaildomain.NewSharedEmailDomainClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedEmailDomain client: %+v", err)
	}
	configureFunc(sharedEmailDomainClient.Client)

	subscriptionClient, err := subscription.NewSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscription client: %+v", err)
	}
	configureFunc(subscriptionClient.Client)

	templateClient, err := template.NewTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Template client: %+v", err)
	}
	configureFunc(templateClient.Client)

	templateDeviceTemplateClient, err := templatedevicetemplate.NewTemplateDeviceTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateDeviceTemplate client: %+v", err)
	}
	configureFunc(templateDeviceTemplateClient.Client)

	templateDeviceTemplateDeviceInstanceClient, err := templatedevicetemplatedeviceinstance.NewTemplateDeviceTemplateDeviceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateDeviceTemplateDeviceInstance client: %+v", err)
	}
	configureFunc(templateDeviceTemplateDeviceInstanceClient.Client)

	templateDeviceTemplateOwnerClient, err := templatedevicetemplateowner.NewTemplateDeviceTemplateOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateDeviceTemplateOwner client: %+v", err)
	}
	configureFunc(templateDeviceTemplateOwnerClient.Client)

	return &Client{
		AdministrativeUnit:                                            administrativeUnitClient,
		AdministrativeUnitDeletedMember:                               administrativeUnitDeletedMemberClient,
		AdministrativeUnitExtension:                                   administrativeUnitExtensionClient,
		AdministrativeUnitMember:                                      administrativeUnitMemberClient,
		AdministrativeUnitScopedRoleMember:                            administrativeUnitScopedRoleMemberClient,
		AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceTo: assignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient,
		AttributeSet:                                 attributeSetClient,
		AuthenticationMethodDevice:                   authenticationMethodDeviceClient,
		AuthenticationMethodDeviceHardwareOathDevice: authenticationMethodDeviceHardwareOathDeviceClient,
		AuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSetting:           authenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient,
		AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningError: authenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient,
		CertificateAuthority: certificateAuthorityClient,
		CertificateAuthorityCertificateBasedApplicationConfiguration:                            certificateAuthorityCertificateBasedApplicationConfigurationClient,
		CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority: certificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient,
		CertificateAuthorityMutualTlsOauthConfiguration:                                         certificateAuthorityMutualTlsOauthConfigurationClient,
		CustomSecurityAttributeDefinition:                                                       customSecurityAttributeDefinitionClient,
		CustomSecurityAttributeDefinitionAllowedValue:                                           customSecurityAttributeDefinitionAllowedValueClient,
		DeletedItem:                     deletedItemClient,
		DeviceLocalCredential:           deviceLocalCredentialClient,
		Directory:                       directoryClient,
		ExternalUserProfile:             externalUserProfileClient,
		FeatureRolloutPolicy:            featureRolloutPolicyClient,
		FeatureRolloutPolicyAppliesTo:   featureRolloutPolicyAppliesToClient,
		FederationConfiguration:         federationConfigurationClient,
		ImpactedResource:                impactedResourceClient,
		InboundSharedUserProfile:        inboundSharedUserProfileClient,
		OnPremisesSynchronization:       onPremisesSynchronizationClient,
		OutboundSharedUserProfile:       outboundSharedUserProfileClient,
		OutboundSharedUserProfileTenant: outboundSharedUserProfileTenantClient,
		PendingExternalUserProfile:      pendingExternalUserProfileClient,
		PublicKeyInfrastructure:         publicKeyInfrastructureClient,
		PublicKeyInfrastructureCertificateBasedAuthConfiguration:                     publicKeyInfrastructureCertificateBasedAuthConfigurationClient,
		PublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthority: publicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityClient,
		Recommendation:                       recommendationClient,
		RecommendationConfiguration:          recommendationConfigurationClient,
		RecommendationImpactedResource:       recommendationImpactedResourceClient,
		SharedEmailDomain:                    sharedEmailDomainClient,
		Subscription:                         subscriptionClient,
		Template:                             templateClient,
		TemplateDeviceTemplate:               templateDeviceTemplateClient,
		TemplateDeviceTemplateDeviceInstance: templateDeviceTemplateDeviceInstanceClient,
		TemplateDeviceTemplateOwner:          templateDeviceTemplateOwnerClient,
	}, nil
}
