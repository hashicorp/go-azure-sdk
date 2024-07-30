package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationEmailSyncSchedule string

const (
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive AndroidEasEmailProfileConfigurationEmailSyncSchedule = "asMessagesArrive"
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage   AndroidEasEmailProfileConfigurationEmailSyncSchedule = "basedOnMyUsage"
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes   AndroidEasEmailProfileConfigurationEmailSyncSchedule = "fifteenMinutes"
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_Manual           AndroidEasEmailProfileConfigurationEmailSyncSchedule = "manual"
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes     AndroidEasEmailProfileConfigurationEmailSyncSchedule = "sixtyMinutes"
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes    AndroidEasEmailProfileConfigurationEmailSyncSchedule = "thirtyMinutes"
	AndroidEasEmailProfileConfigurationEmailSyncSchedule_UserDefined      AndroidEasEmailProfileConfigurationEmailSyncSchedule = "userDefined"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationEmailSyncSchedule() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_Manual),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedule_UserDefined),
	}
}

func (s *AndroidEasEmailProfileConfigurationEmailSyncSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEasEmailProfileConfigurationEmailSyncSchedule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEasEmailProfileConfigurationEmailSyncSchedule(input string) (*AndroidEasEmailProfileConfigurationEmailSyncSchedule, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationEmailSyncSchedule{
		"asmessagesarrive": AndroidEasEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive,
		"basedonmyusage":   AndroidEasEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage,
		"fifteenminutes":   AndroidEasEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes,
		"manual":           AndroidEasEmailProfileConfigurationEmailSyncSchedule_Manual,
		"sixtyminutes":     AndroidEasEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes,
		"thirtyminutes":    AndroidEasEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes,
		"userdefined":      AndroidEasEmailProfileConfigurationEmailSyncSchedule_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationEmailSyncSchedule(input)
	return &out, nil
}
