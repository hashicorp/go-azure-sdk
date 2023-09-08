package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceChassisType string

const (
	WindowsManagedDeviceChassisTypedesktop          WindowsManagedDeviceChassisType = "Desktop"
	WindowsManagedDeviceChassisTypeenterpriseServer WindowsManagedDeviceChassisType = "EnterpriseServer"
	WindowsManagedDeviceChassisTypelaptop           WindowsManagedDeviceChassisType = "Laptop"
	WindowsManagedDeviceChassisTypemobileOther      WindowsManagedDeviceChassisType = "MobileOther"
	WindowsManagedDeviceChassisTypemobileUnknown    WindowsManagedDeviceChassisType = "MobileUnknown"
	WindowsManagedDeviceChassisTypephone            WindowsManagedDeviceChassisType = "Phone"
	WindowsManagedDeviceChassisTypetablet           WindowsManagedDeviceChassisType = "Tablet"
	WindowsManagedDeviceChassisTypeunknown          WindowsManagedDeviceChassisType = "Unknown"
	WindowsManagedDeviceChassisTypeworksWorkstation WindowsManagedDeviceChassisType = "WorksWorkstation"
)

func PossibleValuesForWindowsManagedDeviceChassisType() []string {
	return []string{
		string(WindowsManagedDeviceChassisTypedesktop),
		string(WindowsManagedDeviceChassisTypeenterpriseServer),
		string(WindowsManagedDeviceChassisTypelaptop),
		string(WindowsManagedDeviceChassisTypemobileOther),
		string(WindowsManagedDeviceChassisTypemobileUnknown),
		string(WindowsManagedDeviceChassisTypephone),
		string(WindowsManagedDeviceChassisTypetablet),
		string(WindowsManagedDeviceChassisTypeunknown),
		string(WindowsManagedDeviceChassisTypeworksWorkstation),
	}
}

func parseWindowsManagedDeviceChassisType(input string) (*WindowsManagedDeviceChassisType, error) {
	vals := map[string]WindowsManagedDeviceChassisType{
		"desktop":          WindowsManagedDeviceChassisTypedesktop,
		"enterpriseserver": WindowsManagedDeviceChassisTypeenterpriseServer,
		"laptop":           WindowsManagedDeviceChassisTypelaptop,
		"mobileother":      WindowsManagedDeviceChassisTypemobileOther,
		"mobileunknown":    WindowsManagedDeviceChassisTypemobileUnknown,
		"phone":            WindowsManagedDeviceChassisTypephone,
		"tablet":           WindowsManagedDeviceChassisTypetablet,
		"unknown":          WindowsManagedDeviceChassisTypeunknown,
		"worksworkstation": WindowsManagedDeviceChassisTypeworksWorkstation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceChassisType(input)
	return &out, nil
}
