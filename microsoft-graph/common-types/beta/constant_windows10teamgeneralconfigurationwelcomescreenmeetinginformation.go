package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation string

const (
	Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_ShowOrganizerAndTimeAndSubject Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation = "showOrganizerAndTimeAndSubject"
	Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_ShowOrganizerAndTimeOnly       Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation = "showOrganizerAndTimeOnly"
	Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_UserDefined                    Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation = "userDefined"
)

func PossibleValuesForWindows10TeamGeneralConfigurationWelcomeScreenMeetingInformation() []string {
	return []string{
		string(Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_ShowOrganizerAndTimeAndSubject),
		string(Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_ShowOrganizerAndTimeOnly),
		string(Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_UserDefined),
	}
}

func (s *Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10TeamGeneralConfigurationWelcomeScreenMeetingInformation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10TeamGeneralConfigurationWelcomeScreenMeetingInformation(input string) (*Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation, error) {
	vals := map[string]Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation{
		"showorganizerandtimeandsubject": Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_ShowOrganizerAndTimeAndSubject,
		"showorganizerandtimeonly":       Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_ShowOrganizerAndTimeOnly,
		"userdefined":                    Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation(input)
	return &out, nil
}
