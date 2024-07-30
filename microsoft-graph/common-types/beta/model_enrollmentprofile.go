package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentProfile struct {
	ConfigurationEndpointUrl                            *string `json:"configurationEndpointUrl,omitempty"`
	Description                                         *string `json:"description,omitempty"`
	DisplayName                                         *string `json:"displayName,omitempty"`
	EnableAuthenticationViaCompanyPortal                *bool   `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	Id                                                  *string `json:"id,omitempty"`
	ODataType                                           *string `json:"@odata.type,omitempty"`
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool   `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
	RequiresUserAuthentication                          *bool   `json:"requiresUserAuthentication,omitempty"`
}
