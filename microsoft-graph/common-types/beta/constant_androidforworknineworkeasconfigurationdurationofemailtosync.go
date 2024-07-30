package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync string

const (
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneDay      AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "oneDay"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneMonth    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "oneMonth"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneWeek     AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "oneWeek"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_ThreeDays   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "threeDays"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_TwoWeeks    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "twoWeeks"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_Unlimited   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "unlimited"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_UserDefined AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneDay),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneMonth),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneWeek),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_ThreeDays),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_TwoWeeks),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_Unlimited),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync(input string) (*AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
