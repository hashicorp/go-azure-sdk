package batchmanagements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NameAvailabilityReason string

const (
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = "AlreadyExists"
	NameAvailabilityReasonInvalid       NameAvailabilityReason = "Invalid"
)

func PossibleValuesForNameAvailabilityReason() []string {
	return []string{
		string(NameAvailabilityReasonAlreadyExists),
		string(NameAvailabilityReasonInvalid),
	}
}

type ResourceType string

const (
	ResourceTypeMicrosoftPointBatchBatchAccounts ResourceType = "Microsoft.Batch/batchAccounts"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeMicrosoftPointBatchBatchAccounts),
	}
}
