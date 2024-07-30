package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_AllowAll                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse = "allowAll"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_AllowAllExceptGoogleAccounts AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse = "allowAllExceptGoogleAccounts"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_BlockAll                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse = "blockAll"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_AllowAll),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_AllowAllExceptGoogleAccounts),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_BlockAll),
	}
}

func (s *AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse{
		"allowall":                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_AllowAll,
		"allowallexceptgoogleaccounts": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_AllowAllExceptGoogleAccounts,
		"blockall":                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse_BlockAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse(input)
	return &out, nil
}
