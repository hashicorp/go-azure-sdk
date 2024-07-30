package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationDurationOfEmailToSync string

const (
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneDay      Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "oneDay"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneMonth    Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "oneMonth"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneWeek     Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "oneWeek"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays   Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "threeDays"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks    Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "twoWeeks"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_Unlimited   Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "unlimited"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSync_UserDefined Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneDay),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneMonth),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneWeek),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_Unlimited),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *Windows10EasEmailProfileConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EasEmailProfileConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EasEmailProfileConfigurationDurationOfEmailToSync(input string) (*Windows10EasEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     Windows10EasEmailProfileConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   Windows10EasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    Windows10EasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   Windows10EasEmailProfileConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": Windows10EasEmailProfileConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
