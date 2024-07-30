package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationDurationOfEmailToSync string

const (
	IosEasEmailProfileConfigurationDurationOfEmailToSync_OneDay      IosEasEmailProfileConfigurationDurationOfEmailToSync = "oneDay"
	IosEasEmailProfileConfigurationDurationOfEmailToSync_OneMonth    IosEasEmailProfileConfigurationDurationOfEmailToSync = "oneMonth"
	IosEasEmailProfileConfigurationDurationOfEmailToSync_OneWeek     IosEasEmailProfileConfigurationDurationOfEmailToSync = "oneWeek"
	IosEasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays   IosEasEmailProfileConfigurationDurationOfEmailToSync = "threeDays"
	IosEasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks    IosEasEmailProfileConfigurationDurationOfEmailToSync = "twoWeeks"
	IosEasEmailProfileConfigurationDurationOfEmailToSync_Unlimited   IosEasEmailProfileConfigurationDurationOfEmailToSync = "unlimited"
	IosEasEmailProfileConfigurationDurationOfEmailToSync_UserDefined IosEasEmailProfileConfigurationDurationOfEmailToSync = "userDefined"
)

func PossibleValuesForIosEasEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_OneDay),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_OneMonth),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_OneWeek),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_Unlimited),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSync_UserDefined),
	}
}

func (s *IosEasEmailProfileConfigurationDurationOfEmailToSync) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationDurationOfEmailToSync(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationDurationOfEmailToSync(input string) (*IosEasEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]IosEasEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      IosEasEmailProfileConfigurationDurationOfEmailToSync_OneDay,
		"onemonth":    IosEasEmailProfileConfigurationDurationOfEmailToSync_OneMonth,
		"oneweek":     IosEasEmailProfileConfigurationDurationOfEmailToSync_OneWeek,
		"threedays":   IosEasEmailProfileConfigurationDurationOfEmailToSync_ThreeDays,
		"twoweeks":    IosEasEmailProfileConfigurationDurationOfEmailToSync_TwoWeeks,
		"unlimited":   IosEasEmailProfileConfigurationDurationOfEmailToSync_Unlimited,
		"userdefined": IosEasEmailProfileConfigurationDurationOfEmailToSync_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
