package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationDurationOfEmailToSync string

const (
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneDay      AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "oneDay"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneMonth    AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "oneMonth"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneWeek     AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "oneWeek"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_ThreeDays   AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "threeDays"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_TwoWeeks    AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "twoWeeks"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_Unlimited   AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "unlimited"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_UserDefined AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneDay),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneMonth),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneWeek),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_ThreeDays),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_TwoWeeks),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_Unlimited),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidForWorkGmailEasConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGmailEasConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGmailEasConfigurationDurationOfEmailToSync(input string) (*AndroidForWorkGmailEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidForWorkGmailEasConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
