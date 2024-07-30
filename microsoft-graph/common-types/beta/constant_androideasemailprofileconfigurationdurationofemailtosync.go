package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationDurationOfEmailToSync string

const (
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneDay      AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "oneDay"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneMonth    AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "oneMonth"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneWeek     AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "oneWeek"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays   AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "threeDays"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks    AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "twoWeeks"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_Unlimited   AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "unlimited"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSync_UserDefined AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneDay),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneMonth),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneWeek),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_Unlimited),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidEasEmailProfileConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEasEmailProfileConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEasEmailProfileConfigurationDurationOfEmailToSync(input string) (*AndroidEasEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidEasEmailProfileConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidEasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidEasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidEasEmailProfileConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidEasEmailProfileConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
