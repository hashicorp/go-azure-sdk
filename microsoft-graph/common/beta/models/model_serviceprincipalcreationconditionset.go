package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalCreationConditionSet struct {
	ApplicationIds                        *[]string `json:"applicationIds,omitempty"`
	ApplicationPublisherIds               *[]string `json:"applicationPublisherIds,omitempty"`
	ApplicationTenantIds                  *[]string `json:"applicationTenantIds,omitempty"`
	ApplicationsFromVerifiedPublisherOnly *bool     `json:"applicationsFromVerifiedPublisherOnly,omitempty"`
	CertifiedApplicationsOnly             *bool     `json:"certifiedApplicationsOnly,omitempty"`
	Id                                    *string   `json:"id,omitempty"`
	ODataType                             *string   `json:"@odata.type,omitempty"`
}
