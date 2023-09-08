package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPrivacyAdvertisingId string

const (
	Windows10GeneralConfigurationPrivacyAdvertisingIdallowed       Windows10GeneralConfigurationPrivacyAdvertisingId = "Allowed"
	Windows10GeneralConfigurationPrivacyAdvertisingIdblocked       Windows10GeneralConfigurationPrivacyAdvertisingId = "Blocked"
	Windows10GeneralConfigurationPrivacyAdvertisingIdnotConfigured Windows10GeneralConfigurationPrivacyAdvertisingId = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationPrivacyAdvertisingId() []string {
	return []string{
		string(Windows10GeneralConfigurationPrivacyAdvertisingIdallowed),
		string(Windows10GeneralConfigurationPrivacyAdvertisingIdblocked),
		string(Windows10GeneralConfigurationPrivacyAdvertisingIdnotConfigured),
	}
}

func parseWindows10GeneralConfigurationPrivacyAdvertisingId(input string) (*Windows10GeneralConfigurationPrivacyAdvertisingId, error) {
	vals := map[string]Windows10GeneralConfigurationPrivacyAdvertisingId{
		"allowed":       Windows10GeneralConfigurationPrivacyAdvertisingIdallowed,
		"blocked":       Windows10GeneralConfigurationPrivacyAdvertisingIdblocked,
		"notconfigured": Windows10GeneralConfigurationPrivacyAdvertisingIdnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPrivacyAdvertisingId(input)
	return &out, nil
}
