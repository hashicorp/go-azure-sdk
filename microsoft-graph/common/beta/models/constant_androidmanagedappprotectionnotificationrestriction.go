package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionNotificationRestriction string

const (
	AndroidManagedAppProtectionNotificationRestrictionallow                   AndroidManagedAppProtectionNotificationRestriction = "Allow"
	AndroidManagedAppProtectionNotificationRestrictionblock                   AndroidManagedAppProtectionNotificationRestriction = "Block"
	AndroidManagedAppProtectionNotificationRestrictionblockOrganizationalData AndroidManagedAppProtectionNotificationRestriction = "BlockOrganizationalData"
)

func PossibleValuesForAndroidManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(AndroidManagedAppProtectionNotificationRestrictionallow),
		string(AndroidManagedAppProtectionNotificationRestrictionblock),
		string(AndroidManagedAppProtectionNotificationRestrictionblockOrganizationalData),
	}
}

func parseAndroidManagedAppProtectionNotificationRestriction(input string) (*AndroidManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]AndroidManagedAppProtectionNotificationRestriction{
		"allow":                   AndroidManagedAppProtectionNotificationRestrictionallow,
		"block":                   AndroidManagedAppProtectionNotificationRestrictionblock,
		"blockorganizationaldata": AndroidManagedAppProtectionNotificationRestrictionblockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
