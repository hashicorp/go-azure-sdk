package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync string

const (
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneDay      WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "oneDay"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneMonth    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "oneMonth"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneWeek     WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "oneWeek"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_ThreeDays   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "threeDays"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "twoWeeks"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_Unlimited   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "unlimited"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_UserDefined WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneDay),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneMonth),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneWeek),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_ThreeDays),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_Unlimited),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync(input string) (*WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
