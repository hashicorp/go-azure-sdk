package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequirements struct {
	ExistingAnswers                       *[]AccessPackageAnswer                 `json:"existingAnswers,omitempty"`
	IsApprovalRequired                    *bool                                  `json:"isApprovalRequired,omitempty"`
	IsApprovalRequiredForExtension        *bool                                  `json:"isApprovalRequiredForExtension,omitempty"`
	IsCustomAssignmentScheduleAllowed     *bool                                  `json:"isCustomAssignmentScheduleAllowed,omitempty"`
	IsRequestorJustificationRequired      *bool                                  `json:"isRequestorJustificationRequired,omitempty"`
	ODataType                             *string                                `json:"@odata.type,omitempty"`
	PolicyDescription                     *string                                `json:"policyDescription,omitempty"`
	PolicyDisplayName                     *string                                `json:"policyDisplayName,omitempty"`
	PolicyId                              *string                                `json:"policyId,omitempty"`
	Questions                             *[]AccessPackageQuestion               `json:"questions,omitempty"`
	Schedule                              *RequestSchedule                       `json:"schedule,omitempty"`
	VerifiableCredentialRequirementStatus *VerifiableCredentialRequirementStatus `json:"verifiableCredentialRequirementStatus,omitempty"`
}
