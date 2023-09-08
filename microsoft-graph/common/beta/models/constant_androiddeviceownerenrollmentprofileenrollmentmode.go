package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileEnrollmentMode string

const (
	AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedAOSPUserAssociatedDevice AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "CorporateOwnedAOSPUserAssociatedDevice"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedAOSPUserlessDevice       AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "CorporateOwnedAOSPUserlessDevice"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedDedicatedDevice          AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "CorporateOwnedDedicatedDevice"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedFullyManaged             AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "CorporateOwnedFullyManaged"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedWorkProfile              AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "CorporateOwnedWorkProfile"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileEnrollmentMode() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedAOSPUserAssociatedDevice),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedAOSPUserlessDevice),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedDedicatedDevice),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedFullyManaged),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedWorkProfile),
	}
}

func parseAndroidDeviceOwnerEnrollmentProfileEnrollmentMode(input string) (*AndroidDeviceOwnerEnrollmentProfileEnrollmentMode, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileEnrollmentMode{
		"corporateownedaospuserassociateddevice": AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedAOSPUserAssociatedDevice,
		"corporateownedaospuserlessdevice":       AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedAOSPUserlessDevice,
		"corporateowneddedicateddevice":          AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedDedicatedDevice,
		"corporateownedfullymanaged":             AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedFullyManaged,
		"corporateownedworkprofile":              AndroidDeviceOwnerEnrollmentProfileEnrollmentModecorporateOwnedWorkProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileEnrollmentMode(input)
	return &out, nil
}
