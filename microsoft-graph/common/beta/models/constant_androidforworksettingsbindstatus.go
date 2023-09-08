package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingsBindStatus string

const (
	AndroidForWorkSettingsBindStatusbound             AndroidForWorkSettingsBindStatus = "Bound"
	AndroidForWorkSettingsBindStatusboundAndValidated AndroidForWorkSettingsBindStatus = "BoundAndValidated"
	AndroidForWorkSettingsBindStatusnotBound          AndroidForWorkSettingsBindStatus = "NotBound"
	AndroidForWorkSettingsBindStatusunbinding         AndroidForWorkSettingsBindStatus = "Unbinding"
)

func PossibleValuesForAndroidForWorkSettingsBindStatus() []string {
	return []string{
		string(AndroidForWorkSettingsBindStatusbound),
		string(AndroidForWorkSettingsBindStatusboundAndValidated),
		string(AndroidForWorkSettingsBindStatusnotBound),
		string(AndroidForWorkSettingsBindStatusunbinding),
	}
}

func parseAndroidForWorkSettingsBindStatus(input string) (*AndroidForWorkSettingsBindStatus, error) {
	vals := map[string]AndroidForWorkSettingsBindStatus{
		"bound":             AndroidForWorkSettingsBindStatusbound,
		"boundandvalidated": AndroidForWorkSettingsBindStatusboundAndValidated,
		"notbound":          AndroidForWorkSettingsBindStatusnotBound,
		"unbinding":         AndroidForWorkSettingsBindStatusunbinding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSettingsBindStatus(input)
	return &out, nil
}
