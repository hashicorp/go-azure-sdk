package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessClientApplications struct {
	ExcludeServicePrincipals *[]string                `json:"excludeServicePrincipals,omitempty"`
	IncludeServicePrincipals *[]string                `json:"includeServicePrincipals,omitempty"`
	ODataType                *string                  `json:"@odata.type,omitempty"`
	ServicePrincipalFilter   *ConditionalAccessFilter `json:"servicePrincipalFilter,omitempty"`
}
