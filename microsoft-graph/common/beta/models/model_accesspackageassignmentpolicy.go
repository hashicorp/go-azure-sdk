package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentPolicy struct {
	AccessPackage                *AccessPackage                 `json:"accessPackage,omitempty"`
	AccessPackageCatalog         *AccessPackageCatalog          `json:"accessPackageCatalog,omitempty"`
	AccessPackageId              *string                        `json:"accessPackageId,omitempty"`
	AccessReviewSettings         *AssignmentReviewSettings      `json:"accessReviewSettings,omitempty"`
	CanExtend                    *bool                          `json:"canExtend,omitempty"`
	CreatedBy                    *string                        `json:"createdBy,omitempty"`
	CreatedDateTime              *string                        `json:"createdDateTime,omitempty"`
	CustomExtensionHandlers      *[]CustomExtensionHandler      `json:"customExtensionHandlers,omitempty"`
	CustomExtensionStageSettings *[]CustomExtensionStageSetting `json:"customExtensionStageSettings,omitempty"`
	Description                  *string                        `json:"description,omitempty"`
	DisplayName                  *string                        `json:"displayName,omitempty"`
	DurationInDays               *int64                         `json:"durationInDays,omitempty"`
	ExpirationDateTime           *string                        `json:"expirationDateTime,omitempty"`
	Id                           *string                        `json:"id,omitempty"`
	ModifiedBy                   *string                        `json:"modifiedBy,omitempty"`
	ModifiedDateTime             *string                        `json:"modifiedDateTime,omitempty"`
	ODataType                    *string                        `json:"@odata.type,omitempty"`
	Questions                    *[]AccessPackageQuestion       `json:"questions,omitempty"`
	RequestApprovalSettings      *ApprovalSettings              `json:"requestApprovalSettings,omitempty"`
	RequestorSettings            *RequestorSettings             `json:"requestorSettings,omitempty"`
	VerifiableCredentialSettings *VerifiableCredentialSettings  `json:"verifiableCredentialSettings,omitempty"`
}
