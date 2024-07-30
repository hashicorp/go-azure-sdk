package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderSubmitSamplesConsentType string

const (
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_AlwaysPrompt                 Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "alwaysPrompt"
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_NeverSend                    Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "neverSend"
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_SendAllSamplesAutomatically  Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "sendAllSamplesAutomatically"
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "sendSafeSamplesAutomatically"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderSubmitSamplesConsentType() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_AlwaysPrompt),
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_NeverSend),
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_SendAllSamplesAutomatically),
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically),
	}
}

func (s *Windows10GeneralConfigurationDefenderSubmitSamplesConsentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderSubmitSamplesConsentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderSubmitSamplesConsentType(input string) (*Windows10GeneralConfigurationDefenderSubmitSamplesConsentType, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderSubmitSamplesConsentType{
		"alwaysprompt":                 Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_AlwaysPrompt,
		"neversend":                    Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_NeverSend,
		"sendallsamplesautomatically":  Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_SendAllSamplesAutomatically,
		"sendsafesamplesautomatically": Windows10GeneralConfigurationDefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderSubmitSamplesConsentType(input)
	return &out, nil
}
