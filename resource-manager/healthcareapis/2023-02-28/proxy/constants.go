package proxy

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNameUnavailabilityReason string

const (
	ServiceNameUnavailabilityReasonAlreadyExists ServiceNameUnavailabilityReason = "AlreadyExists"
	ServiceNameUnavailabilityReasonInvalid       ServiceNameUnavailabilityReason = "Invalid"
)

func PossibleValuesForServiceNameUnavailabilityReason() []string {
	return []string{
		string(ServiceNameUnavailabilityReasonAlreadyExists),
		string(ServiceNameUnavailabilityReasonInvalid),
	}
}

func parseServiceNameUnavailabilityReason(input string) (*ServiceNameUnavailabilityReason, error) {
	vals := map[string]ServiceNameUnavailabilityReason{
		"alreadyexists": ServiceNameUnavailabilityReasonAlreadyExists,
		"invalid":       ServiceNameUnavailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceNameUnavailabilityReason(input)
	return &out, nil
}
