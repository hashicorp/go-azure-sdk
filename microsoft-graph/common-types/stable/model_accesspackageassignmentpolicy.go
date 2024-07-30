package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentPolicy struct {
	AccessPackage                *AccessPackage                                   `json:"accessPackage,omitempty"`
	AllowedTargetScope           *AccessPackageAssignmentPolicyAllowedTargetScope `json:"allowedTargetScope,omitempty"`
	AutomaticRequestSettings     *AccessPackageAutomaticRequestSettings           `json:"automaticRequestSettings,omitempty"`
	Catalog                      *AccessPackageCatalog                            `json:"catalog,omitempty"`
	CreatedDateTime              *string                                          `json:"createdDateTime,omitempty"`
	CustomExtensionStageSettings *[]CustomExtensionStageSetting                   `json:"customExtensionStageSettings,omitempty"`
	Description                  *string                                          `json:"description,omitempty"`
	DisplayName                  *string                                          `json:"displayName,omitempty"`
	Expiration                   *ExpirationPattern                               `json:"expiration,omitempty"`
	Id                           *string                                          `json:"id,omitempty"`
	ModifiedDateTime             *string                                          `json:"modifiedDateTime,omitempty"`
	ODataType                    *string                                          `json:"@odata.type,omitempty"`
	Questions                    *[]AccessPackageQuestion                         `json:"questions,omitempty"`
	RequestApprovalSettings      *AccessPackageAssignmentApprovalSettings         `json:"requestApprovalSettings,omitempty"`
	RequestorSettings            *AccessPackageAssignmentRequestorSettings        `json:"requestorSettings,omitempty"`
	ReviewSettings               *AccessPackageAssignmentReviewSettings           `json:"reviewSettings,omitempty"`
	SpecificAllowedTargets       *[]SubjectSet                                    `json:"specificAllowedTargets,omitempty"`
}
