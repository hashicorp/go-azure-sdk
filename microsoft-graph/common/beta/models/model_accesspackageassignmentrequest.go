package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequest struct {
	AccessPackage                   *AccessPackage                    `json:"accessPackage,omitempty"`
	AccessPackageAssignment         *AccessPackageAssignment          `json:"accessPackageAssignment,omitempty"`
	Answers                         *[]AccessPackageAnswer            `json:"answers,omitempty"`
	CompletedDate                   *string                           `json:"completedDate,omitempty"`
	CreatedDateTime                 *string                           `json:"createdDateTime,omitempty"`
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`
	CustomExtensionHandlerInstances *[]CustomExtensionHandlerInstance `json:"customExtensionHandlerInstances,omitempty"`
	ExpirationDateTime              *string                           `json:"expirationDateTime,omitempty"`
	Id                              *string                           `json:"id,omitempty"`
	IsValidationOnly                *bool                             `json:"isValidationOnly,omitempty"`
	Justification                   *string                           `json:"justification,omitempty"`
	ODataType                       *string                           `json:"@odata.type,omitempty"`
	RequestState                    *string                           `json:"requestState,omitempty"`
	RequestStatus                   *string                           `json:"requestStatus,omitempty"`
	RequestType                     *string                           `json:"requestType,omitempty"`
	Requestor                       *AccessPackageSubject             `json:"requestor,omitempty"`
	Schedule                        *RequestSchedule                  `json:"schedule,omitempty"`
	VerifiedCredentialsData         *[]VerifiedCredentialData         `json:"verifiedCredentialsData,omitempty"`
}
