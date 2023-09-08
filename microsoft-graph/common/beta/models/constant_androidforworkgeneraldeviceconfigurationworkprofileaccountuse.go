package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseallowAll                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse = "AllowAll"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseallowAllExceptGoogleAccounts AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse = "AllowAllExceptGoogleAccounts"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseblockAll                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse = "BlockAll"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseallowAll),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseallowAllExceptGoogleAccounts),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseblockAll),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse{
		"allowall":                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseallowAll,
		"allowallexceptgoogleaccounts": AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseallowAllExceptGoogleAccounts,
		"blockall":                     AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUseblockAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse(input)
	return &out, nil
}
