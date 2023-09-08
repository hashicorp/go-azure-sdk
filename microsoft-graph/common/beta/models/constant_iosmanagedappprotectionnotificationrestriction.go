package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionNotificationRestriction string

const (
	IosManagedAppProtectionNotificationRestrictionallow                   IosManagedAppProtectionNotificationRestriction = "Allow"
	IosManagedAppProtectionNotificationRestrictionblock                   IosManagedAppProtectionNotificationRestriction = "Block"
	IosManagedAppProtectionNotificationRestrictionblockOrganizationalData IosManagedAppProtectionNotificationRestriction = "BlockOrganizationalData"
)

func PossibleValuesForIosManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(IosManagedAppProtectionNotificationRestrictionallow),
		string(IosManagedAppProtectionNotificationRestrictionblock),
		string(IosManagedAppProtectionNotificationRestrictionblockOrganizationalData),
	}
}

func parseIosManagedAppProtectionNotificationRestriction(input string) (*IosManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]IosManagedAppProtectionNotificationRestriction{
		"allow":                   IosManagedAppProtectionNotificationRestrictionallow,
		"block":                   IosManagedAppProtectionNotificationRestrictionblock,
		"blockorganizationaldata": IosManagedAppProtectionNotificationRestrictionblockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
