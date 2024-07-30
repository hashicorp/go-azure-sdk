package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationEmailSyncSchedule string

const (
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive Windows10EasEmailProfileConfigurationEmailSyncSchedule = "asMessagesArrive"
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage   Windows10EasEmailProfileConfigurationEmailSyncSchedule = "basedOnMyUsage"
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes   Windows10EasEmailProfileConfigurationEmailSyncSchedule = "fifteenMinutes"
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_Manual           Windows10EasEmailProfileConfigurationEmailSyncSchedule = "manual"
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes     Windows10EasEmailProfileConfigurationEmailSyncSchedule = "sixtyMinutes"
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes    Windows10EasEmailProfileConfigurationEmailSyncSchedule = "thirtyMinutes"
	Windows10EasEmailProfileConfigurationEmailSyncSchedule_UserDefined      Windows10EasEmailProfileConfigurationEmailSyncSchedule = "userDefined"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationEmailSyncSchedule() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_Manual),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedule_UserDefined),
	}
}

func (s *Windows10EasEmailProfileConfigurationEmailSyncSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EasEmailProfileConfigurationEmailSyncSchedule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EasEmailProfileConfigurationEmailSyncSchedule(input string) (*Windows10EasEmailProfileConfigurationEmailSyncSchedule, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationEmailSyncSchedule{
		"asmessagesarrive": Windows10EasEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive,
		"basedonmyusage":   Windows10EasEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage,
		"fifteenminutes":   Windows10EasEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes,
		"manual":           Windows10EasEmailProfileConfigurationEmailSyncSchedule_Manual,
		"sixtyminutes":     Windows10EasEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes,
		"thirtyminutes":    Windows10EasEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes,
		"userdefined":      Windows10EasEmailProfileConfigurationEmailSyncSchedule_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationEmailSyncSchedule(input)
	return &out, nil
}
