package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationAppsVisibilityListType string

const (
	IosGeneralDeviceConfigurationAppsVisibilityListTypeappsInListCompliant    IosGeneralDeviceConfigurationAppsVisibilityListType = "AppsInListCompliant"
	IosGeneralDeviceConfigurationAppsVisibilityListTypeappsNotInListCompliant IosGeneralDeviceConfigurationAppsVisibilityListType = "AppsNotInListCompliant"
	IosGeneralDeviceConfigurationAppsVisibilityListTypenone                   IosGeneralDeviceConfigurationAppsVisibilityListType = "None"
)

func PossibleValuesForIosGeneralDeviceConfigurationAppsVisibilityListType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationAppsVisibilityListTypeappsInListCompliant),
		string(IosGeneralDeviceConfigurationAppsVisibilityListTypeappsNotInListCompliant),
		string(IosGeneralDeviceConfigurationAppsVisibilityListTypenone),
	}
}

func parseIosGeneralDeviceConfigurationAppsVisibilityListType(input string) (*IosGeneralDeviceConfigurationAppsVisibilityListType, error) {
	vals := map[string]IosGeneralDeviceConfigurationAppsVisibilityListType{
		"appsinlistcompliant":    IosGeneralDeviceConfigurationAppsVisibilityListTypeappsInListCompliant,
		"appsnotinlistcompliant": IosGeneralDeviceConfigurationAppsVisibilityListTypeappsNotInListCompliant,
		"none":                   IosGeneralDeviceConfigurationAppsVisibilityListTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationAppsVisibilityListType(input)
	return &out, nil
}
