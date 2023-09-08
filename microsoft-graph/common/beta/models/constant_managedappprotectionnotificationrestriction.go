package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionNotificationRestriction string

const (
	ManagedAppProtectionNotificationRestrictionallow                   ManagedAppProtectionNotificationRestriction = "Allow"
	ManagedAppProtectionNotificationRestrictionblock                   ManagedAppProtectionNotificationRestriction = "Block"
	ManagedAppProtectionNotificationRestrictionblockOrganizationalData ManagedAppProtectionNotificationRestriction = "BlockOrganizationalData"
)

func PossibleValuesForManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(ManagedAppProtectionNotificationRestrictionallow),
		string(ManagedAppProtectionNotificationRestrictionblock),
		string(ManagedAppProtectionNotificationRestrictionblockOrganizationalData),
	}
}

func parseManagedAppProtectionNotificationRestriction(input string) (*ManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]ManagedAppProtectionNotificationRestriction{
		"allow":                   ManagedAppProtectionNotificationRestrictionallow,
		"block":                   ManagedAppProtectionNotificationRestrictionblock,
		"blockorganizationaldata": ManagedAppProtectionNotificationRestrictionblockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
