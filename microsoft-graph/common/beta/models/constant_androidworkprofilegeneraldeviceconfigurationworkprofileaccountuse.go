package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseallowAll                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse = "AllowAll"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseallowAllExceptGoogleAccounts AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse = "AllowAllExceptGoogleAccounts"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseblockAll                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse = "BlockAll"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseallowAll),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseallowAllExceptGoogleAccounts),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseblockAll),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse{
		"allowall":                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseallowAll,
		"allowallexceptgoogleaccounts": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseallowAllExceptGoogleAccounts,
		"blockall":                     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUseblockAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileAccountUse(input)
	return &out, nil
}
