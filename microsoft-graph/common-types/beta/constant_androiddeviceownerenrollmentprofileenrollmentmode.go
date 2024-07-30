package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileEnrollmentMode string

const (
	AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedAOSPUserAssociatedDevice AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "corporateOwnedAOSPUserAssociatedDevice"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedAOSPUserlessDevice       AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "corporateOwnedAOSPUserlessDevice"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedDedicatedDevice          AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "corporateOwnedDedicatedDevice"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedFullyManaged             AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "corporateOwnedFullyManaged"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedWorkProfile              AndroidDeviceOwnerEnrollmentProfileEnrollmentMode = "corporateOwnedWorkProfile"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileEnrollmentMode() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedAOSPUserAssociatedDevice),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedAOSPUserlessDevice),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedDedicatedDevice),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedFullyManaged),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedWorkProfile),
	}
}

func (s *AndroidDeviceOwnerEnrollmentProfileEnrollmentMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnrollmentProfileEnrollmentMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnrollmentProfileEnrollmentMode(input string) (*AndroidDeviceOwnerEnrollmentProfileEnrollmentMode, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileEnrollmentMode{
		"corporateownedaospuserassociateddevice": AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedAOSPUserAssociatedDevice,
		"corporateownedaospuserlessdevice":       AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedAOSPUserlessDevice,
		"corporateowneddedicateddevice":          AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedDedicatedDevice,
		"corporateownedfullymanaged":             AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedFullyManaged,
		"corporateownedworkprofile":              AndroidDeviceOwnerEnrollmentProfileEnrollmentMode_CorporateOwnedWorkProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileEnrollmentMode(input)
	return &out, nil
}
