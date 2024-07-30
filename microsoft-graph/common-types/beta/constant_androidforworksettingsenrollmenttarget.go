package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingsEnrollmentTarget string

const (
	AndroidForWorkSettingsEnrollmentTarget_All                              AndroidForWorkSettingsEnrollmentTarget = "all"
	AndroidForWorkSettingsEnrollmentTarget_None                             AndroidForWorkSettingsEnrollmentTarget = "none"
	AndroidForWorkSettingsEnrollmentTarget_Targeted                         AndroidForWorkSettingsEnrollmentTarget = "targeted"
	AndroidForWorkSettingsEnrollmentTarget_TargetedAsEnrollmentRestrictions AndroidForWorkSettingsEnrollmentTarget = "targetedAsEnrollmentRestrictions"
)

func PossibleValuesForAndroidForWorkSettingsEnrollmentTarget() []string {
	return []string{
		string(AndroidForWorkSettingsEnrollmentTarget_All),
		string(AndroidForWorkSettingsEnrollmentTarget_None),
		string(AndroidForWorkSettingsEnrollmentTarget_Targeted),
		string(AndroidForWorkSettingsEnrollmentTarget_TargetedAsEnrollmentRestrictions),
	}
}

func (s *AndroidForWorkSettingsEnrollmentTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkSettingsEnrollmentTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkSettingsEnrollmentTarget(input string) (*AndroidForWorkSettingsEnrollmentTarget, error) {
	vals := map[string]AndroidForWorkSettingsEnrollmentTarget{
		"all":                              AndroidForWorkSettingsEnrollmentTarget_All,
		"none":                             AndroidForWorkSettingsEnrollmentTarget_None,
		"targeted":                         AndroidForWorkSettingsEnrollmentTarget_Targeted,
		"targetedasenrollmentrestrictions": AndroidForWorkSettingsEnrollmentTarget_TargetedAsEnrollmentRestrictions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSettingsEnrollmentTarget(input)
	return &out, nil
}
