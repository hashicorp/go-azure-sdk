package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationCompliantAppListType string

const (
	AndroidGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant    AndroidGeneralDeviceConfigurationCompliantAppListType = "AppsInListCompliant"
	AndroidGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant AndroidGeneralDeviceConfigurationCompliantAppListType = "AppsNotInListCompliant"
	AndroidGeneralDeviceConfigurationCompliantAppListTypenone                   AndroidGeneralDeviceConfigurationCompliantAppListType = "None"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationCompliantAppListType() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant),
		string(AndroidGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant),
		string(AndroidGeneralDeviceConfigurationCompliantAppListTypenone),
	}
}

func parseAndroidGeneralDeviceConfigurationCompliantAppListType(input string) (*AndroidGeneralDeviceConfigurationCompliantAppListType, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationCompliantAppListType{
		"appsinlistcompliant":    AndroidGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant,
		"appsnotinlistcompliant": AndroidGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant,
		"none":                   AndroidGeneralDeviceConfigurationCompliantAppListTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationCompliantAppListType(input)
	return &out, nil
}
