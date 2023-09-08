package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantConditionSet struct {
	CertifiedClientApplicationsOnly             *bool                                      `json:"certifiedClientApplicationsOnly,omitempty"`
	ClientApplicationIds                        *[]string                                  `json:"clientApplicationIds,omitempty"`
	ClientApplicationPublisherIds               *[]string                                  `json:"clientApplicationPublisherIds,omitempty"`
	ClientApplicationTenantIds                  *[]string                                  `json:"clientApplicationTenantIds,omitempty"`
	ClientApplicationsFromVerifiedPublisherOnly *bool                                      `json:"clientApplicationsFromVerifiedPublisherOnly,omitempty"`
	Id                                          *string                                    `json:"id,omitempty"`
	ODataType                                   *string                                    `json:"@odata.type,omitempty"`
	PermissionClassification                    *string                                    `json:"permissionClassification,omitempty"`
	PermissionType                              *PermissionGrantConditionSetPermissionType `json:"permissionType,omitempty"`
	Permissions                                 *[]string                                  `json:"permissions,omitempty"`
	ResourceApplication                         *string                                    `json:"resourceApplication,omitempty"`
}
