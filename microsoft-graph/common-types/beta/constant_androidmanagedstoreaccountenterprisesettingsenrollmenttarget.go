package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget string

const (
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_All                              AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "all"
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_None                             AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "none"
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_Targeted                         AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "targeted"
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_TargetedAsEnrollmentRestrictions AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "targetedAsEnrollmentRestrictions"
)

func PossibleValuesForAndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget() []string {
	return []string{
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_All),
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_None),
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_Targeted),
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_TargetedAsEnrollmentRestrictions),
	}
}

func (s *AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget(input string) (*AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget, error) {
	vals := map[string]AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget{
		"all":                              AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_All,
		"none":                             AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_None,
		"targeted":                         AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_Targeted,
		"targetedasenrollmentrestrictions": AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget_TargetedAsEnrollmentRestrictions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget(input)
	return &out, nil
}
