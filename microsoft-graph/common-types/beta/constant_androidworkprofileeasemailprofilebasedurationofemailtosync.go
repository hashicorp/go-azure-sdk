package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync string

const (
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneDay      AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "oneDay"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneMonth    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "oneMonth"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneWeek     AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "oneWeek"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_ThreeDays   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "threeDays"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_TwoWeeks    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "twoWeeks"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_Unlimited   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "unlimited"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_UserDefined AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneDay),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneMonth),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneWeek),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_ThreeDays),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_TwoWeeks),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_Unlimited),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync(input string) (*AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync{
		"oneday":      AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync(input)
	return &out, nil
}
