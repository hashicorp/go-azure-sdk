package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryadministrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryadministrativeunitextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryadministrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryadministrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryattributeset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directorycustomsecurityattributedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directorycustomsecurityattributedefinitionallowedvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directorydeleteditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryfederationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/v1_0/directoryonpremisessynchronization"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Directory                                              *directory.DirectoryClient
	DirectoryAdministrativeUnit                            *directoryadministrativeunit.DirectoryAdministrativeUnitClient
	DirectoryAdministrativeUnitExtension                   *directoryadministrativeunitextension.DirectoryAdministrativeUnitExtensionClient
	DirectoryAdministrativeUnitMember                      *directoryadministrativeunitmember.DirectoryAdministrativeUnitMemberClient
	DirectoryAdministrativeUnitScopedRoleMember            *directoryadministrativeunitscopedrolemember.DirectoryAdministrativeUnitScopedRoleMemberClient
	DirectoryAttributeSet                                  *directoryattributeset.DirectoryAttributeSetClient
	DirectoryCustomSecurityAttributeDefinition             *directorycustomsecurityattributedefinition.DirectoryCustomSecurityAttributeDefinitionClient
	DirectoryCustomSecurityAttributeDefinitionAllowedValue *directorycustomsecurityattributedefinitionallowedvalue.DirectoryCustomSecurityAttributeDefinitionAllowedValueClient
	DirectoryDeletedItem                                   *directorydeleteditem.DirectoryDeletedItemClient
	DirectoryFederationConfiguration                       *directoryfederationconfiguration.DirectoryFederationConfigurationClient
	DirectoryOnPremisesSynchronization                     *directoryonpremisessynchronization.DirectoryOnPremisesSynchronizationClient
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

	directoryFederationConfigurationClient, err := directoryfederationconfiguration.NewDirectoryFederationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryFederationConfiguration client: %+v", err)
	}
	configureFunc(directoryFederationConfigurationClient.Client)

	directoryOnPremisesSynchronizationClient, err := directoryonpremisessynchronization.NewDirectoryOnPremisesSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryOnPremisesSynchronization client: %+v", err)
	}
	configureFunc(directoryOnPremisesSynchronizationClient.Client)

	return &Client{
		Directory:                                              directoryClient,
		DirectoryAdministrativeUnit:                            directoryAdministrativeUnitClient,
		DirectoryAdministrativeUnitExtension:                   directoryAdministrativeUnitExtensionClient,
		DirectoryAdministrativeUnitMember:                      directoryAdministrativeUnitMemberClient,
		DirectoryAdministrativeUnitScopedRoleMember:            directoryAdministrativeUnitScopedRoleMemberClient,
		DirectoryAttributeSet:                                  directoryAttributeSetClient,
		DirectoryCustomSecurityAttributeDefinition:             directoryCustomSecurityAttributeDefinitionClient,
		DirectoryCustomSecurityAttributeDefinitionAllowedValue: directoryCustomSecurityAttributeDefinitionAllowedValueClient,
		DirectoryDeletedItem:                                   directoryDeletedItemClient,
		DirectoryFederationConfiguration:                       directoryFederationConfigurationClient,
		DirectoryOnPremisesSynchronization:                     directoryOnPremisesSynchronizationClient,
	}, nil
}
