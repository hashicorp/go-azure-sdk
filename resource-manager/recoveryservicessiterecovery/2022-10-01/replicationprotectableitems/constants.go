package replicationprotectableitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

func PossibleValuesForHealthErrorCustomerResolvability() []string {
	return []string{
		string(HealthErrorCustomerResolvabilityAllowed),
		string(HealthErrorCustomerResolvabilityNotAllowed),
	}
}

type PresenceStatus string

const (
	PresenceStatusNotPresent PresenceStatus = "NotPresent"
	PresenceStatusPresent    PresenceStatus = "Present"
	PresenceStatusUnknown    PresenceStatus = "Unknown"
)

func PossibleValuesForPresenceStatus() []string {
	return []string{
		string(PresenceStatusNotPresent),
		string(PresenceStatusPresent),
		string(PresenceStatusUnknown),
	}
}
