package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync string

const (
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneDay      AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "oneDay"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneMonth    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "oneMonth"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneWeek     AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "oneWeek"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_ThreeDays   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "threeDays"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_TwoWeeks    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "twoWeeks"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_Unlimited   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "unlimited"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_UserDefined AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseDurationOfEmailToSync() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneDay),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneMonth),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneWeek),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_ThreeDays),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_TwoWeeks),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_Unlimited),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEasEmailProfileBaseDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEasEmailProfileBaseDurationOfEmailToSync(input string) (*AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync{
		"oneday":      AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync(input)
	return &out, nil
}
