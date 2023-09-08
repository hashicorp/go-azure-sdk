package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81GeneralConfigurationCompliantAppListType string

const (
	WindowsPhone81GeneralConfigurationCompliantAppListTypeappsInListCompliant    WindowsPhone81GeneralConfigurationCompliantAppListType = "AppsInListCompliant"
	WindowsPhone81GeneralConfigurationCompliantAppListTypeappsNotInListCompliant WindowsPhone81GeneralConfigurationCompliantAppListType = "AppsNotInListCompliant"
	WindowsPhone81GeneralConfigurationCompliantAppListTypenone                   WindowsPhone81GeneralConfigurationCompliantAppListType = "None"
)

func PossibleValuesForWindowsPhone81GeneralConfigurationCompliantAppListType() []string {
	return []string{
		string(WindowsPhone81GeneralConfigurationCompliantAppListTypeappsInListCompliant),
		string(WindowsPhone81GeneralConfigurationCompliantAppListTypeappsNotInListCompliant),
		string(WindowsPhone81GeneralConfigurationCompliantAppListTypenone),
	}
}

func parseWindowsPhone81GeneralConfigurationCompliantAppListType(input string) (*WindowsPhone81GeneralConfigurationCompliantAppListType, error) {
	vals := map[string]WindowsPhone81GeneralConfigurationCompliantAppListType{
		"appsinlistcompliant":    WindowsPhone81GeneralConfigurationCompliantAppListTypeappsInListCompliant,
		"appsnotinlistcompliant": WindowsPhone81GeneralConfigurationCompliantAppListTypeappsNotInListCompliant,
		"none":                   WindowsPhone81GeneralConfigurationCompliantAppListTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81GeneralConfigurationCompliantAppListType(input)
	return &out, nil
}
