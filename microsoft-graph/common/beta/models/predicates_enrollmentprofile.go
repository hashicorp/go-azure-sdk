package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentProfileOperationPredicate struct {
	ConfigurationEndpointUrl                            *string
	Description                                         *string
	DisplayName                                         *string
	EnableAuthenticationViaCompanyPortal                *bool
	Id                                                  *string
	ODataType                                           *string
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool
	RequiresUserAuthentication                          *bool
}

func (p EnrollmentProfileOperationPredicate) Matches(input EnrollmentProfile) bool {

	if p.ConfigurationEndpointUrl != nil && (input.ConfigurationEndpointUrl == nil || *p.ConfigurationEndpointUrl != *input.ConfigurationEndpointUrl) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnableAuthenticationViaCompanyPortal != nil && (input.EnableAuthenticationViaCompanyPortal == nil || *p.EnableAuthenticationViaCompanyPortal != *input.EnableAuthenticationViaCompanyPortal) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequireCompanyPortalOnSetupAssistantEnrolledDevices != nil && (input.RequireCompanyPortalOnSetupAssistantEnrolledDevices == nil || *p.RequireCompanyPortalOnSetupAssistantEnrolledDevices != *input.RequireCompanyPortalOnSetupAssistantEnrolledDevices) {
		return false
	}

	if p.RequiresUserAuthentication != nil && (input.RequiresUserAuthentication == nil || *p.RequiresUserAuthentication != *input.RequiresUserAuthentication) {
		return false
	}

	return true
}
