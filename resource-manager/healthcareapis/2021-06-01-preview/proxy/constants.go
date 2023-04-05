package proxy

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
