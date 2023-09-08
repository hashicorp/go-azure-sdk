package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAccessPackageWorkflowExtension struct {
	AuthenticationConfiguration *CustomExtensionAuthenticationConfiguration `json:"authenticationConfiguration,omitempty"`
	ClientConfiguration         *CustomExtensionClientConfiguration         `json:"clientConfiguration,omitempty"`
	CreatedDateTime             *string                                     `json:"createdDateTime,omitempty"`
	Description                 *string                                     `json:"description,omitempty"`
	DisplayName                 *string                                     `json:"displayName,omitempty"`
	EndpointConfiguration       *CustomExtensionEndpointConfiguration       `json:"endpointConfiguration,omitempty"`
	Id                          *string                                     `json:"id,omitempty"`
	LastModifiedDateTime        *string                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType                   *string                                     `json:"@odata.type,omitempty"`
}
