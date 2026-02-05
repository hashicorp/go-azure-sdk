package certificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateOperationPredicate struct {
	PreviousStateTransitionTime *string
	PublicData                  *string
	StateTransitionTime         *string
	Thumbprint                  *string
	ThumbprintAlgorithm         *string
	Url                         *string
}

func (p CertificateOperationPredicate) Matches(input Certificate) bool {

	if p.PreviousStateTransitionTime != nil && (input.PreviousStateTransitionTime == nil || *p.PreviousStateTransitionTime != *input.PreviousStateTransitionTime) {
		return false
	}

	if p.PublicData != nil && (input.PublicData == nil || *p.PublicData != *input.PublicData) {
		return false
	}

	if p.StateTransitionTime != nil && (input.StateTransitionTime == nil || *p.StateTransitionTime != *input.StateTransitionTime) {
		return false
	}

	if p.Thumbprint != nil && (input.Thumbprint == nil || *p.Thumbprint != *input.Thumbprint) {
		return false
	}

	if p.ThumbprintAlgorithm != nil && (input.ThumbprintAlgorithm == nil || *p.ThumbprintAlgorithm != *input.ThumbprintAlgorithm) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
