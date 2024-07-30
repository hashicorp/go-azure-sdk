package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationsIdentitySet struct {
	Application                    *Identity                              `json:"application,omitempty"`
	ApplicationInstance            *Identity                              `json:"applicationInstance,omitempty"`
	AssertedIdentity               *Identity                              `json:"assertedIdentity,omitempty"`
	AzureCommunicationServicesUser *Identity                              `json:"azureCommunicationServicesUser,omitempty"`
	Device                         *Identity                              `json:"device,omitempty"`
	Encrypted                      *Identity                              `json:"encrypted,omitempty"`
	EndpointType                   *CommunicationsIdentitySetEndpointType `json:"endpointType,omitempty"`
	Guest                          *Identity                              `json:"guest,omitempty"`
	ODataType                      *string                                `json:"@odata.type,omitempty"`
	OnPremises                     *Identity                              `json:"onPremises,omitempty"`
	Phone                          *Identity                              `json:"phone,omitempty"`
	User                           *Identity                              `json:"user,omitempty"`
}
