package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync string

const (
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneDay      AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "oneDay"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneMonth    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "oneMonth"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneWeek     AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "oneWeek"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_ThreeDays   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "threeDays"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_TwoWeeks    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "twoWeeks"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_Unlimited   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "unlimited"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_UserDefined AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneDay),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneMonth),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneWeek),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_ThreeDays),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_TwoWeeks),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_Unlimited),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync(input string) (*AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
