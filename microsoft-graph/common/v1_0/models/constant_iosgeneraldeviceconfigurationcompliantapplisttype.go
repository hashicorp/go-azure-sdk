package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationCompliantAppListType string

const (
	IosGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant    IosGeneralDeviceConfigurationCompliantAppListType = "AppsInListCompliant"
	IosGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant IosGeneralDeviceConfigurationCompliantAppListType = "AppsNotInListCompliant"
	IosGeneralDeviceConfigurationCompliantAppListTypenone                   IosGeneralDeviceConfigurationCompliantAppListType = "None"
)

func PossibleValuesForIosGeneralDeviceConfigurationCompliantAppListType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant),
		string(IosGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant),
		string(IosGeneralDeviceConfigurationCompliantAppListTypenone),
	}
}

func parseIosGeneralDeviceConfigurationCompliantAppListType(input string) (*IosGeneralDeviceConfigurationCompliantAppListType, error) {
	vals := map[string]IosGeneralDeviceConfigurationCompliantAppListType{
		"appsinlistcompliant":    IosGeneralDeviceConfigurationCompliantAppListTypeappsInListCompliant,
		"appsnotinlistcompliant": IosGeneralDeviceConfigurationCompliantAppListTypeappsNotInListCompliant,
		"none":                   IosGeneralDeviceConfigurationCompliantAppListTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationCompliantAppListType(input)
	return &out, nil
}
