package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageApprovalStage struct {
	DurationBeforeAutomaticDenial   *string       `json:"durationBeforeAutomaticDenial,omitempty"`
	DurationBeforeEscalation        *string       `json:"durationBeforeEscalation,omitempty"`
	EscalationApprovers             *[]SubjectSet `json:"escalationApprovers,omitempty"`
	FallbackEscalationApprovers     *[]SubjectSet `json:"fallbackEscalationApprovers,omitempty"`
	FallbackPrimaryApprovers        *[]SubjectSet `json:"fallbackPrimaryApprovers,omitempty"`
	IsApproverJustificationRequired *bool         `json:"isApproverJustificationRequired,omitempty"`
	IsEscalationEnabled             *bool         `json:"isEscalationEnabled,omitempty"`
	ODataType                       *string       `json:"@odata.type,omitempty"`
	PrimaryApprovers                *[]SubjectSet `json:"primaryApprovers,omitempty"`
}
