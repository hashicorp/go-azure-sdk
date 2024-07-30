package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule string

const (
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "asMessagesArrive"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "basedOnMyUsage"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "fifteenMinutes"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_Manual           WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "manual"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes     WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "sixtyMinutes"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes    WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "thirtyMinutes"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_UserDefined      WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "userDefined"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_Manual),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_UserDefined),
	}
}

func (s *WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule(input string) (*WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule{
		"asmessagesarrive": WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_AsMessagesArrive,
		"basedonmyusage":   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_BasedOnMyUsage,
		"fifteenminutes":   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_FifteenMinutes,
		"manual":           WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_Manual,
		"sixtyminutes":     WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_SixtyMinutes,
		"thirtyminutes":    WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_ThirtyMinutes,
		"userdefined":      WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule(input)
	return &out, nil
}
