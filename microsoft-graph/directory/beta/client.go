package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryadministrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryadministrativeunitextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryadministrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryadministrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryattributeset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorycertificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorycertificateauthoritycertificatebasedapplicationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorycustomsecurityattributedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorycustomsecurityattributedefinitionallowedvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorydeleteditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryfeaturerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryfeaturerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryfederationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryimpactedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryinboundshareduserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryonpremisessynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryoutboundshareduserprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryoutboundshareduserprofiletenant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryrecommendation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directoryrecommendationimpactedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorysharedemaildomain"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/directorysubscription"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Directory                                                                                        *directory.DirectoryClient
	DirectoryAdministrativeUnit                                                                      *directoryadministrativeunit.DirectoryAdministrativeUnitClient
	DirectoryAdministrativeUnitExtension                                                             *directoryadministrativeunitextension.DirectoryAdministrativeUnitExtensionClient
	DirectoryAdministrativeUnitMember                                                                *directoryadministrativeunitmember.DirectoryAdministrativeUnitMemberClient
	DirectoryAdministrativeUnitScopedRoleMember                                                      *directoryadministrativeunitscopedrolemember.DirectoryAdministrativeUnitScopedRoleMemberClient
	DirectoryAttributeSet                                                                            *directoryattributeset.DirectoryAttributeSetClient
	DirectoryCertificateAuthority                                                                    *directorycertificateauthority.DirectoryCertificateAuthorityClient
	DirectoryCertificateAuthorityCertificateBasedApplicationConfiguration                            *directorycertificateauthoritycertificatebasedapplicationconfiguration.DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient
	DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority *directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority.DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient
	DirectoryCustomSecurityAttributeDefinition                                                       *directorycustomsecurityattributedefinition.DirectoryCustomSecurityAttributeDefinitionClient
	DirectoryCustomSecurityAttributeDefinitionAllowedValue                                           *directorycustomsecurityattributedefinitionallowedvalue.DirectoryCustomSecurityAttributeDefinitionAllowedValueClient
	DirectoryDeletedItem                                                                             *directorydeleteditem.DirectoryDeletedItemClient
	DirectoryFeatureRolloutPolicy                                                                    *directoryfeaturerolloutpolicy.DirectoryFeatureRolloutPolicyClient
	DirectoryFeatureRolloutPolicyAppliesTo                                                           *directoryfeaturerolloutpolicyappliesto.DirectoryFeatureRolloutPolicyAppliesToClient
	DirectoryFederationConfiguration                                                                 *directoryfederationconfiguration.DirectoryFederationConfigurationClient
	DirectoryImpactedResource                                                                        *directoryimpactedresource.DirectoryImpactedResourceClient
	DirectoryInboundSharedUserProfile                                                                *directoryinboundshareduserprofile.DirectoryInboundSharedUserProfileClient
	DirectoryOnPremisesSynchronization                                                               *directoryonpremisessynchronization.DirectoryOnPremisesSynchronizationClient
	DirectoryOutboundSharedUserProfile                                                               *directoryoutboundshareduserprofile.DirectoryOutboundSharedUserProfileClient
	DirectoryOutboundSharedUserProfileTenant                                                         *directoryoutboundshareduserprofiletenant.DirectoryOutboundSharedUserProfileTenantClient
	DirectoryRecommendation                                                                          *directoryrecommendation.DirectoryRecommendationClient
	DirectoryRecommendationImpactedResource                                                          *directoryrecommendationimpactedresource.DirectoryRecommendationImpactedResourceClient
	DirectorySharedEmailDomain                                                                       *directorysharedemaildomain.DirectorySharedEmailDomainClient
	DirectorySubscription                                                                            *directorysubscription.DirectorySubscriptionClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	directoryAdministrativeUnitClient, err := directoryadministrativeunit.NewDirectoryAdministrativeUnitClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAdministrativeUnit client: %+v", err)
	}
	configureFunc(directoryAdministrativeUnitClient.Client)

	directoryAdministrativeUnitExtensionClient, err := directoryadministrativeunitextension.NewDirectoryAdministrativeUnitExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAdministrativeUnitExtension client: %+v", err)
	}
	configureFunc(directoryAdministrativeUnitExtensionClient.Client)

	directoryAdministrativeUnitMemberClient, err := directoryadministrativeunitmember.NewDirectoryAdministrativeUnitMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAdministrativeUnitMember client: %+v", err)
	}
	configureFunc(directoryAdministrativeUnitMemberClient.Client)

	directoryAdministrativeUnitScopedRoleMemberClient, err := directoryadministrativeunitscopedrolemember.NewDirectoryAdministrativeUnitScopedRoleMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAdministrativeUnitScopedRoleMember client: %+v", err)
	}
	configureFunc(directoryAdministrativeUnitScopedRoleMemberClient.Client)

	directoryAttributeSetClient, err := directoryattributeset.NewDirectoryAttributeSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAttributeSet client: %+v", err)
	}
	configureFunc(directoryAttributeSetClient.Client)

	directoryCertificateAuthorityCertificateBasedApplicationConfigurationClient, err := directorycertificateauthoritycertificatebasedapplicationconfiguration.NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryCertificateAuthorityCertificateBasedApplicationConfiguration client: %+v", err)
	}
	configureFunc(directoryCertificateAuthorityCertificateBasedApplicationConfigurationClient.Client)

	directoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient, err := directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority.NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority client: %+v", err)
	}
	configureFunc(directoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient.Client)

	directoryCertificateAuthorityClient, err := directorycertificateauthority.NewDirectoryCertificateAuthorityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryCertificateAuthority client: %+v", err)
	}
	configureFunc(directoryCertificateAuthorityClient.Client)

	directoryClient, err := directory.NewDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Directory client: %+v", err)
	}
	configureFunc(directoryClient.Client)

	directoryCustomSecurityAttributeDefinitionAllowedValueClient, err := directorycustomsecurityattributedefinitionallowedvalue.NewDirectoryCustomSecurityAttributeDefinitionAllowedValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryCustomSecurityAttributeDefinitionAllowedValue client: %+v", err)
	}
	configureFunc(directoryCustomSecurityAttributeDefinitionAllowedValueClient.Client)

	directoryCustomSecurityAttributeDefinitionClient, err := directorycustomsecurityattributedefinition.NewDirectoryCustomSecurityAttributeDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryCustomSecurityAttributeDefinition client: %+v", err)
	}
	configureFunc(directoryCustomSecurityAttributeDefinitionClient.Client)

	directoryDeletedItemClient, err := directorydeleteditem.NewDirectoryDeletedItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryDeletedItem client: %+v", err)
	}
	configureFunc(directoryDeletedItemClient.Client)

	directoryFeatureRolloutPolicyAppliesToClient, err := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryFeatureRolloutPolicyAppliesTo client: %+v", err)
	}
	configureFunc(directoryFeatureRolloutPolicyAppliesToClient.Client)

	directoryFeatureRolloutPolicyClient, err := directoryfeaturerolloutpolicy.NewDirectoryFeatureRolloutPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryFeatureRolloutPolicy client: %+v", err)
	}
	configureFunc(directoryFeatureRolloutPolicyClient.Client)

	directoryFederationConfigurationClient, err := directoryfederationconfiguration.NewDirectoryFederationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryFederationConfiguration client: %+v", err)
	}
	configureFunc(directoryFederationConfigurationClient.Client)

	directoryImpactedResourceClient, err := directoryimpactedresource.NewDirectoryImpactedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryImpactedResource client: %+v", err)
	}
	configureFunc(directoryImpactedResourceClient.Client)

	directoryInboundSharedUserProfileClient, err := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryInboundSharedUserProfile client: %+v", err)
	}
	configureFunc(directoryInboundSharedUserProfileClient.Client)

	directoryOnPremisesSynchronizationClient, err := directoryonpremisessynchronization.NewDirectoryOnPremisesSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryOnPremisesSynchronization client: %+v", err)
	}
	configureFunc(directoryOnPremisesSynchronizationClient.Client)

	directoryOutboundSharedUserProfileClient, err := directoryoutboundshareduserprofile.NewDirectoryOutboundSharedUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryOutboundSharedUserProfile client: %+v", err)
	}
	configureFunc(directoryOutboundSharedUserProfileClient.Client)

	directoryOutboundSharedUserProfileTenantClient, err := directoryoutboundshareduserprofiletenant.NewDirectoryOutboundSharedUserProfileTenantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryOutboundSharedUserProfileTenant client: %+v", err)
	}
	configureFunc(directoryOutboundSharedUserProfileTenantClient.Client)

	directoryRecommendationClient, err := directoryrecommendation.NewDirectoryRecommendationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRecommendation client: %+v", err)
	}
	configureFunc(directoryRecommendationClient.Client)

	directoryRecommendationImpactedResourceClient, err := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRecommendationImpactedResource client: %+v", err)
	}
	configureFunc(directoryRecommendationImpactedResourceClient.Client)

	directorySharedEmailDomainClient, err := directorysharedemaildomain.NewDirectorySharedEmailDomainClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectorySharedEmailDomain client: %+v", err)
	}
	configureFunc(directorySharedEmailDomainClient.Client)

	directorySubscriptionClient, err := directorysubscription.NewDirectorySubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectorySubscription client: %+v", err)
	}
	configureFunc(directorySubscriptionClient.Client)

	return &Client{
		Directory:                                   directoryClient,
		DirectoryAdministrativeUnit:                 directoryAdministrativeUnitClient,
		DirectoryAdministrativeUnitExtension:        directoryAdministrativeUnitExtensionClient,
		DirectoryAdministrativeUnitMember:           directoryAdministrativeUnitMemberClient,
		DirectoryAdministrativeUnitScopedRoleMember: directoryAdministrativeUnitScopedRoleMemberClient,
		DirectoryAttributeSet:                       directoryAttributeSetClient,
		DirectoryCertificateAuthority:               directoryCertificateAuthorityClient,
		DirectoryCertificateAuthorityCertificateBasedApplicationConfiguration:                            directoryCertificateAuthorityCertificateBasedApplicationConfigurationClient,
		DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority: directoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient,
		DirectoryCustomSecurityAttributeDefinition:                                                       directoryCustomSecurityAttributeDefinitionClient,
		DirectoryCustomSecurityAttributeDefinitionAllowedValue:                                           directoryCustomSecurityAttributeDefinitionAllowedValueClient,
		DirectoryDeletedItem:                     directoryDeletedItemClient,
		DirectoryFeatureRolloutPolicy:            directoryFeatureRolloutPolicyClient,
		DirectoryFeatureRolloutPolicyAppliesTo:   directoryFeatureRolloutPolicyAppliesToClient,
		DirectoryFederationConfiguration:         directoryFederationConfigurationClient,
		DirectoryImpactedResource:                directoryImpactedResourceClient,
		DirectoryInboundSharedUserProfile:        directoryInboundSharedUserProfileClient,
		DirectoryOnPremisesSynchronization:       directoryOnPremisesSynchronizationClient,
		DirectoryOutboundSharedUserProfile:       directoryOutboundSharedUserProfileClient,
		DirectoryOutboundSharedUserProfileTenant: directoryOutboundSharedUserProfileTenantClient,
		DirectoryRecommendation:                  directoryRecommendationClient,
		DirectoryRecommendationImpactedResource:  directoryRecommendationImpactedResourceClient,
		DirectorySharedEmailDomain:               directorySharedEmailDomainClient,
		DirectorySubscription:                    directorySubscriptionClient,
	}, nil
}
