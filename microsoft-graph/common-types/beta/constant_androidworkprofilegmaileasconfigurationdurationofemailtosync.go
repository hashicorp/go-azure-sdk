package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync string

const (
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneDay      AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "oneDay"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneMonth    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "oneMonth"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneWeek     AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "oneWeek"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_ThreeDays   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "threeDays"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_TwoWeeks    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "twoWeeks"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_Unlimited   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "unlimited"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_UserDefined AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneDay),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneMonth),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneWeek),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_ThreeDays),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_TwoWeeks),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_Unlimited),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync(input string) (*AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
