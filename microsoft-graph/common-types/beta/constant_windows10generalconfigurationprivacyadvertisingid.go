package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPrivacyAdvertisingId string

const (
	Windows10GeneralConfigurationPrivacyAdvertisingId_Allowed       Windows10GeneralConfigurationPrivacyAdvertisingId = "allowed"
	Windows10GeneralConfigurationPrivacyAdvertisingId_Blocked       Windows10GeneralConfigurationPrivacyAdvertisingId = "blocked"
	Windows10GeneralConfigurationPrivacyAdvertisingId_NotConfigured Windows10GeneralConfigurationPrivacyAdvertisingId = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationPrivacyAdvertisingId() []string {
	return []string{
		string(Windows10GeneralConfigurationPrivacyAdvertisingId_Allowed),
		string(Windows10GeneralConfigurationPrivacyAdvertisingId_Blocked),
		string(Windows10GeneralConfigurationPrivacyAdvertisingId_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationPrivacyAdvertisingId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPrivacyAdvertisingId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPrivacyAdvertisingId(input string) (*Windows10GeneralConfigurationPrivacyAdvertisingId, error) {
	vals := map[string]Windows10GeneralConfigurationPrivacyAdvertisingId{
		"allowed":       Windows10GeneralConfigurationPrivacyAdvertisingId_Allowed,
		"blocked":       Windows10GeneralConfigurationPrivacyAdvertisingId_Blocked,
		"notconfigured": Windows10GeneralConfigurationPrivacyAdvertisingId_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPrivacyAdvertisingId(input)
	return &out, nil
}
