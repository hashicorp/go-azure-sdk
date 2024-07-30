package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/attributeset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/customsecurityattributedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/customsecurityattributedefinitionallowedvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/deleteditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/devicelocalcredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/federationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/onpremisessynchronization"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AdministrativeUnit                            *administrativeunit.AdministrativeUnitClient
	AdministrativeUnitExtension                   *administrativeunitextension.AdministrativeUnitExtensionClient
	AdministrativeUnitMember                      *administrativeunitmember.AdministrativeUnitMemberClient
	AdministrativeUnitScopedRoleMember            *administrativeunitscopedrolemember.AdministrativeUnitScopedRoleMemberClient
	AttributeSet                                  *attributeset.AttributeSetClient
	CustomSecurityAttributeDefinition             *customsecurityattributedefinition.CustomSecurityAttributeDefinitionClient
	CustomSecurityAttributeDefinitionAllowedValue *customsecurityattributedefinitionallowedvalue.CustomSecurityAttributeDefinitionAllowedValueClient
	DeletedItem                                   *deleteditem.DeletedItemClient
	DeviceLocalCredential                         *devicelocalcredential.DeviceLocalCredentialClient
	Directory                                     *directory.DirectoryClient
	FederationConfiguration                       *federationconfiguration.FederationConfigurationClient
	OnPremisesSynchronization                     *onpremisessynchronization.OnPremisesSynchronizationClient
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

	federationConfigurationClient, err := federationconfiguration.NewFederationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FederationConfiguration client: %+v", err)
	}
	configureFunc(federationConfigurationClient.Client)

	onPremisesSynchronizationClient, err := onpremisessynchronization.NewOnPremisesSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnPremisesSynchronization client: %+v", err)
	}
	configureFunc(onPremisesSynchronizationClient.Client)

	return &Client{
		AdministrativeUnit:                            administrativeUnitClient,
		AdministrativeUnitExtension:                   administrativeUnitExtensionClient,
		AdministrativeUnitMember:                      administrativeUnitMemberClient,
		AdministrativeUnitScopedRoleMember:            administrativeUnitScopedRoleMemberClient,
		AttributeSet:                                  attributeSetClient,
		CustomSecurityAttributeDefinition:             customSecurityAttributeDefinitionClient,
		CustomSecurityAttributeDefinitionAllowedValue: customSecurityAttributeDefinitionAllowedValueClient,
		DeletedItem:                                   deletedItemClient,
		DeviceLocalCredential:                         deviceLocalCredentialClient,
		Directory:                                     directoryClient,
		FederationConfiguration:                       federationConfigurationClient,
		OnPremisesSynchronization:                     onPremisesSynchronizationClient,
	}, nil
}
