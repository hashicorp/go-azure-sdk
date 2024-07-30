package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/attributeset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthoritycertificatebasedapplicationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/customsecurityattributedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/customsecurityattributedefinitionallowedvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/deleteditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/devicelocalcredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/featurerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/featurerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/federationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/impactedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/inboundshareduserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/onpremisessynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/outboundshareduserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/outboundshareduserprofiletenant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationimpactedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/sharedemaildomain"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/subscription"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdministrativeUnit                                                                      *administrativeunit.AdministrativeUnitClient
	AdministrativeUnitExtension                                                             *administrativeunitextension.AdministrativeUnitExtensionClient
	AdministrativeUnitMember                                                                *administrativeunitmember.AdministrativeUnitMemberClient
	AdministrativeUnitScopedRoleMember                                                      *administrativeunitscopedrolemember.AdministrativeUnitScopedRoleMemberClient
	AttributeSet                                                                            *attributeset.AttributeSetClient
	CertificateAuthority                                                                    *certificateauthority.CertificateAuthorityClient
	CertificateAuthorityCertificateBasedApplicationConfiguration                            *certificateauthoritycertificatebasedapplicationconfiguration.CertificateAuthorityCertificateBasedApplicationConfigurationClient
	CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority *certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority.CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient
	CustomSecurityAttributeDefinition                                                       *customsecurityattributedefinition.CustomSecurityAttributeDefinitionClient
	CustomSecurityAttributeDefinitionAllowedValue                                           *customsecurityattributedefinitionallowedvalue.CustomSecurityAttributeDefinitionAllowedValueClient
	DeletedItem                                                                             *deleteditem.DeletedItemClient
	DeviceLocalCredential                                                                   *devicelocalcredential.DeviceLocalCredentialClient
	Directory                                                                               *directory.DirectoryClient
	FeatureRolloutPolicy                                                                    *featurerolloutpolicy.FeatureRolloutPolicyClient
	FeatureRolloutPolicyAppliesTo                                                           *featurerolloutpolicyappliesto.FeatureRolloutPolicyAppliesToClient
	FederationConfiguration                                                                 *federationconfiguration.FederationConfigurationClient
	ImpactedResource                                                                        *impactedresource.ImpactedResourceClient
	InboundSharedUserProfile                                                                *inboundshareduserprofile.InboundSharedUserProfileClient
	OnPremisesSynchronization                                                               *onpremisessynchronization.OnPremisesSynchronizationClient
	OutboundSharedUserProfile                                                               *outboundshareduserprofile.OutboundSharedUserProfileClient
	OutboundSharedUserProfileTenant                                                         *outboundshareduserprofiletenant.OutboundSharedUserProfileTenantClient
	Recommendation                                                                          *recommendation.RecommendationClient
	RecommendationImpactedResource                                                          *recommendationimpactedresource.RecommendationImpactedResourceClient
	SharedEmailDomain                                                                       *sharedemaildomain.SharedEmailDomainClient
	Subscription                                                                            *subscription.SubscriptionClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	administrativeUnitClient, err := administrativeunit.NewAdministrativeUnitClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnit client: %+v", err)
	}
	configureFunc(administrativeUnitClient.Client)

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

	attributeSetClient, err := attributeset.NewAttributeSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AttributeSet client: %+v", err)
	}
	configureFunc(attributeSetClient.Client)

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

	recommendationClient, err := recommendation.NewRecommendationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Recommendation client: %+v", err)
	}
	configureFunc(recommendationClient.Client)

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

	return &Client{
		AdministrativeUnit:                 administrativeUnitClient,
		AdministrativeUnitExtension:        administrativeUnitExtensionClient,
		AdministrativeUnitMember:           administrativeUnitMemberClient,
		AdministrativeUnitScopedRoleMember: administrativeUnitScopedRoleMemberClient,
		AttributeSet:                       attributeSetClient,
		CertificateAuthority:               certificateAuthorityClient,
		CertificateAuthorityCertificateBasedApplicationConfiguration:                            certificateAuthorityCertificateBasedApplicationConfigurationClient,
		CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority: certificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient,
		CustomSecurityAttributeDefinition:                                                       customSecurityAttributeDefinitionClient,
		CustomSecurityAttributeDefinitionAllowedValue:                                           customSecurityAttributeDefinitionAllowedValueClient,
		DeletedItem:                     deletedItemClient,
		DeviceLocalCredential:           deviceLocalCredentialClient,
		Directory:                       directoryClient,
		FeatureRolloutPolicy:            featureRolloutPolicyClient,
		FeatureRolloutPolicyAppliesTo:   featureRolloutPolicyAppliesToClient,
		FederationConfiguration:         federationConfigurationClient,
		ImpactedResource:                impactedResourceClient,
		InboundSharedUserProfile:        inboundSharedUserProfileClient,
		OnPremisesSynchronization:       onPremisesSynchronizationClient,
		OutboundSharedUserProfile:       outboundSharedUserProfileClient,
		OutboundSharedUserProfileTenant: outboundSharedUserProfileTenantClient,
		Recommendation:                  recommendationClient,
		RecommendationImpactedResource:  recommendationImpactedResourceClient,
		SharedEmailDomain:               sharedEmailDomainClient,
		Subscription:                    subscriptionClient,
	}, nil
}
