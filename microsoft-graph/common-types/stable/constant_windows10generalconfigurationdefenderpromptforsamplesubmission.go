package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderPromptForSampleSubmission string

const (
	Windows10GeneralConfigurationDefenderPromptForSampleSubmission_AlwaysPrompt                    Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "alwaysPrompt"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmission_NeverSendData                   Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "neverSendData"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmission_PromptBeforeSendingPersonalData Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "promptBeforeSendingPersonalData"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmission_SendAllDataWithoutPrompting     Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "sendAllDataWithoutPrompting"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmission_UserDefined                     Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderPromptForSampleSubmission() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmission_AlwaysPrompt),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmission_NeverSendData),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmission_PromptBeforeSendingPersonalData),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmission_SendAllDataWithoutPrompting),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmission_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationDefenderPromptForSampleSubmission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderPromptForSampleSubmission(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderPromptForSampleSubmission(input string) (*Windows10GeneralConfigurationDefenderPromptForSampleSubmission, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderPromptForSampleSubmission{
		"alwaysprompt":                    Windows10GeneralConfigurationDefenderPromptForSampleSubmission_AlwaysPrompt,
		"neversenddata":                   Windows10GeneralConfigurationDefenderPromptForSampleSubmission_NeverSendData,
		"promptbeforesendingpersonaldata": Windows10GeneralConfigurationDefenderPromptForSampleSubmission_PromptBeforeSendingPersonalData,
		"sendalldatawithoutprompting":     Windows10GeneralConfigurationDefenderPromptForSampleSubmission_SendAllDataWithoutPrompting,
		"userdefined":                     Windows10GeneralConfigurationDefenderPromptForSampleSubmission_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderPromptForSampleSubmission(input)
	return &out, nil
}
