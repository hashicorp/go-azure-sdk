package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionNotificationRestriction string

const (
	TargetedManagedAppProtectionNotificationRestrictionallow                   TargetedManagedAppProtectionNotificationRestriction = "Allow"
	TargetedManagedAppProtectionNotificationRestrictionblock                   TargetedManagedAppProtectionNotificationRestriction = "Block"
	TargetedManagedAppProtectionNotificationRestrictionblockOrganizationalData TargetedManagedAppProtectionNotificationRestriction = "BlockOrganizationalData"
)

func PossibleValuesForTargetedManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(TargetedManagedAppProtectionNotificationRestrictionallow),
		string(TargetedManagedAppProtectionNotificationRestrictionblock),
		string(TargetedManagedAppProtectionNotificationRestrictionblockOrganizationalData),
	}
}

func parseTargetedManagedAppProtectionNotificationRestriction(input string) (*TargetedManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]TargetedManagedAppProtectionNotificationRestriction{
		"allow":                   TargetedManagedAppProtectionNotificationRestrictionallow,
		"block":                   TargetedManagedAppProtectionNotificationRestrictionblock,
		"blockorganizationaldata": TargetedManagedAppProtectionNotificationRestrictionblockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
