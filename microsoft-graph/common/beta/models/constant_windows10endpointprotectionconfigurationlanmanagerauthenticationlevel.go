package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel string

const (
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNltm             Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "LmAndNltm"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNtlmOnly         Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "LmAndNtlmOnly"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNtlmV2           Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "LmAndNtlmV2"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmAndNtlmV2       Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "LmNtlmAndNtlmV2"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmV2AndNotLm      Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "LmNtlmV2AndNotLm"
	Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmV2AndNotLmOrNtm Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel = "LmNtlmV2AndNotLmOrNtm"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLanManagerAuthenticationLevel() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNltm),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNtlmOnly),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNtlmV2),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmAndNtlmV2),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmV2AndNotLm),
		string(Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmV2AndNotLmOrNtm),
	}
}

func parseWindows10EndpointProtectionConfigurationLanManagerAuthenticationLevel(input string) (*Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel{
		"lmandnltm":             Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNltm,
		"lmandntlmonly":         Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNtlmOnly,
		"lmandntlmv2":           Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmAndNtlmV2,
		"lmntlmandntlmv2":       Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmAndNtlmV2,
		"lmntlmv2andnotlm":      Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmV2AndNotLm,
		"lmntlmv2andnotlmorntm": Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevellmNtlmV2AndNotLmOrNtm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLanManagerAuthenticationLevel(input)
	return &out, nil
}
