package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfigurationCompliantAppListType string

const (
	MacOSGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant    MacOSGeneralDeviceConfigurationCompliantAppListType = "AppsInListCompliant"
	MacOSGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant MacOSGeneralDeviceConfigurationCompliantAppListType = "AppsNotInListCompliant"
	MacOSGeneralDeviceConfigurationCompliantAppListTypenone                   MacOSGeneralDeviceConfigurationCompliantAppListType = "None"
)

func PossibleValuesForMacOSGeneralDeviceConfigurationCompliantAppListType() []string {
	return []string{
		string(MacOSGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant),
		string(MacOSGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant),
		string(MacOSGeneralDeviceConfigurationCompliantAppListTypenone),
	}
}

func parseMacOSGeneralDeviceConfigurationCompliantAppListType(input string) (*MacOSGeneralDeviceConfigurationCompliantAppListType, error) {
	vals := map[string]MacOSGeneralDeviceConfigurationCompliantAppListType{
		"appsinlistcompliant":    MacOSGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant,
		"appsnotinlistcompliant": MacOSGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant,
		"none":                   MacOSGeneralDeviceConfigurationCompliantAppListTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGeneralDeviceConfigurationCompliantAppListType(input)
	return &out, nil
}
