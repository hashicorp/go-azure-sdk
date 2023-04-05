package voiceservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}
