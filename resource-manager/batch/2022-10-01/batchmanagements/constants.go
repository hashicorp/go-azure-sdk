package batchmanagements

import "strings"

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

func parseNameAvailabilityReason(input string) (*NameAvailabilityReason, error) {
	vals := map[string]NameAvailabilityReason{
		"alreadyexists": NameAvailabilityReasonAlreadyExists,
		"invalid":       NameAvailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NameAvailabilityReason(input)
	return &out, nil
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

func parseResourceType(input string) (*ResourceType, error) {
	vals := map[string]ResourceType{
		"microsoft.batch/batchaccounts": ResourceTypeMicrosoftPointBatchBatchAccounts,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceType(input)
	return &out, nil
}
