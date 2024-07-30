package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtension struct {
	AuthenticationConfiguration *CustomExtensionAuthenticationConfiguration `json:"authenticationConfiguration,omitempty"`
	CallbackConfiguration       *CustomExtensionCallbackConfiguration       `json:"callbackConfiguration,omitempty"`
	ClientConfiguration         *CustomExtensionClientConfiguration         `json:"clientConfiguration,omitempty"`
	CreatedBy                   *User                                       `json:"createdBy,omitempty"`
	CreatedDateTime             *string                                     `json:"createdDateTime,omitempty"`
	Description                 *string                                     `json:"description,omitempty"`
	DisplayName                 *string                                     `json:"displayName,omitempty"`
	EndpointConfiguration       *CustomExtensionEndpointConfiguration       `json:"endpointConfiguration,omitempty"`
	Id                          *string                                     `json:"id,omitempty"`
	LastModifiedBy              *User                                       `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime        *string                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType                   *string                                     `json:"@odata.type,omitempty"`
}
