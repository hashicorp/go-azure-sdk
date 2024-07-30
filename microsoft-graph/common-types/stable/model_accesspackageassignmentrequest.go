package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequest struct {
	AccessPackage                   *AccessPackage                             `json:"accessPackage,omitempty"`
	Answers                         *[]AccessPackageAnswer                     `json:"answers,omitempty"`
	Assignment                      *AccessPackageAssignment                   `json:"assignment,omitempty"`
	CompletedDateTime               *string                                    `json:"completedDateTime,omitempty"`
	CreatedDateTime                 *string                                    `json:"createdDateTime,omitempty"`
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance          `json:"customExtensionCalloutInstances,omitempty"`
	Id                              *string                                    `json:"id,omitempty"`
	ODataType                       *string                                    `json:"@odata.type,omitempty"`
	RequestType                     *AccessPackageAssignmentRequestRequestType `json:"requestType,omitempty"`
	Requestor                       *AccessPackageSubject                      `json:"requestor,omitempty"`
	Schedule                        *EntitlementManagementSchedule             `json:"schedule,omitempty"`
	State                           *AccessPackageAssignmentRequestState       `json:"state,omitempty"`
	Status                          *string                                    `json:"status,omitempty"`
}
