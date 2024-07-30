package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel string

const (
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNltm             Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "lmAndNltm"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNtlmOnly         Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "lmAndNtlmOnly"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNtlmV2           Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "lmAndNtlmV2"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmAndNtlmV2       Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "lmNtlmAndNtlmV2"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmV2AndNotLm      Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "lmNtlmV2AndNotLm"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmV2AndNotLmOrNtm Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "lmNtlmV2AndNotLmOrNtm"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLanManagerAuthenticationLevel() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNltm),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNtlmOnly),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNtlmV2),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmAndNtlmV2),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmV2AndNotLm),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmV2AndNotLmOrNtm),
	}
}

func (s *Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLanManagerAuthenticationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLanManagerAuthenticationLevel(input string) (*Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel{
		"lmandnltm":             Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNltm,
		"lmandntlmonly":         Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNtlmOnly,
		"lmandntlmv2":           Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmAndNtlmV2,
		"lmntlmandntlmv2":       Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmAndNtlmV2,
		"lmntlmv2andnotlm":      Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmV2AndNotLm,
		"lmntlmv2andnotlmorntm": Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel_LmNtlmV2AndNotLmOrNtm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel(input)
	return &out, nil
}
