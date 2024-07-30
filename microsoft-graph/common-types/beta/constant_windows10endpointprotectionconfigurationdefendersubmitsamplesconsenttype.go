package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType string

const (
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_AlwaysPrompt                 Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "alwaysPrompt"
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_NeverSend                    Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "neverSend"
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_SendAllSamplesAutomatically  Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "sendAllSamplesAutomatically"
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "sendSafeSamplesAutomatically"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_AlwaysPrompt),
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_NeverSend),
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_SendAllSamplesAutomatically),
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType(input string) (*Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType{
		"alwaysprompt":                 Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_AlwaysPrompt,
		"neversend":                    Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_NeverSend,
		"sendallsamplesautomatically":  Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_SendAllSamplesAutomatically,
		"sendsafesamplesautomatically": Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType(input)
	return &out, nil
}
