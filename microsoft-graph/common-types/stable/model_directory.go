package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Directory struct {
	AdministrativeUnits                *[]AdministrativeUnit                 `json:"administrativeUnits,omitempty"`
	AttributeSets                      *[]AttributeSet                       `json:"attributeSets,omitempty"`
	CustomSecurityAttributeDefinitions *[]CustomSecurityAttributeDefinition  `json:"customSecurityAttributeDefinitions,omitempty"`
	DeletedItems                       *[]DirectoryObject                    `json:"deletedItems,omitempty"`
	DeviceLocalCredentials             *[]DeviceLocalCredentialInfo          `json:"deviceLocalCredentials,omitempty"`
	FederationConfigurations           *[]IdentityProviderBase               `json:"federationConfigurations,omitempty"`
	Id                                 *string                               `json:"id,omitempty"`
	ODataType                          *string                               `json:"@odata.type,omitempty"`
	OnPremisesSynchronization          *[]OnPremisesDirectorySynchronization `json:"onPremisesSynchronization,omitempty"`
}
