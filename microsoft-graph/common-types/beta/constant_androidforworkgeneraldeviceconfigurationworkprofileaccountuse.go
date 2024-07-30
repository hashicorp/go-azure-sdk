package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_AllowAll                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse = "allowAll"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_AllowAllExceptGoogleAccounts AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse = "allowAllExceptGoogleAccounts"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_BlockAll                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse = "blockAll"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_AllowAll),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_AllowAllExceptGoogleAccounts),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_BlockAll),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse{
		"allowall":                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_AllowAll,
		"allowallexceptgoogleaccounts": AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_AllowAllExceptGoogleAccounts,
		"blockall":                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse_BlockAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse(input)
	return &out, nil
}
