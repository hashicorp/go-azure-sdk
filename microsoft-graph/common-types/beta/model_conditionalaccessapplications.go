package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessApplications struct {
	ApplicationFilter                           *ConditionalAccessFilter `json:"applicationFilter,omitempty"`
	ExcludeApplications                         *[]string                `json:"excludeApplications,omitempty"`
	IncludeApplications                         *[]string                `json:"includeApplications,omitempty"`
	IncludeAuthenticationContextClassReferences *[]string                `json:"includeAuthenticationContextClassReferences,omitempty"`
	IncludeUserActions                          *[]string                `json:"includeUserActions,omitempty"`
	ODataType                                   *string                  `json:"@odata.type,omitempty"`
}
