package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionNotificationRestriction string

const (
	DefaultManagedAppProtectionNotificationRestrictionallow                   DefaultManagedAppProtectionNotificationRestriction = "Allow"
	DefaultManagedAppProtectionNotificationRestrictionblock                   DefaultManagedAppProtectionNotificationRestriction = "Block"
	DefaultManagedAppProtectionNotificationRestrictionblockOrganizationalData DefaultManagedAppProtectionNotificationRestriction = "BlockOrganizationalData"
)

func PossibleValuesForDefaultManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(DefaultManagedAppProtectionNotificationRestrictionallow),
		string(DefaultManagedAppProtectionNotificationRestrictionblock),
		string(DefaultManagedAppProtectionNotificationRestrictionblockOrganizationalData),
	}
}

func parseDefaultManagedAppProtectionNotificationRestriction(input string) (*DefaultManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]DefaultManagedAppProtectionNotificationRestriction{
		"allow":                   DefaultManagedAppProtectionNotificationRestrictionallow,
		"block":                   DefaultManagedAppProtectionNotificationRestrictionblock,
		"blockorganizationaldata": DefaultManagedAppProtectionNotificationRestrictionblockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
