package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceChassisType string

const (
	ManagedDeviceChassisTypedesktop          ManagedDeviceChassisType = "Desktop"
	ManagedDeviceChassisTypeenterpriseServer ManagedDeviceChassisType = "EnterpriseServer"
	ManagedDeviceChassisTypelaptop           ManagedDeviceChassisType = "Laptop"
	ManagedDeviceChassisTypemobileOther      ManagedDeviceChassisType = "MobileOther"
	ManagedDeviceChassisTypemobileUnknown    ManagedDeviceChassisType = "MobileUnknown"
	ManagedDeviceChassisTypephone            ManagedDeviceChassisType = "Phone"
	ManagedDeviceChassisTypetablet           ManagedDeviceChassisType = "Tablet"
	ManagedDeviceChassisTypeunknown          ManagedDeviceChassisType = "Unknown"
	ManagedDeviceChassisTypeworksWorkstation ManagedDeviceChassisType = "WorksWorkstation"
)

func PossibleValuesForManagedDeviceChassisType() []string {
	return []string{
		string(ManagedDeviceChassisTypedesktop),
		string(ManagedDeviceChassisTypeenterpriseServer),
		string(ManagedDeviceChassisTypelaptop),
		string(ManagedDeviceChassisTypemobileOther),
		string(ManagedDeviceChassisTypemobileUnknown),
		string(ManagedDeviceChassisTypephone),
		string(ManagedDeviceChassisTypetablet),
		string(ManagedDeviceChassisTypeunknown),
		string(ManagedDeviceChassisTypeworksWorkstation),
	}
}

func parseManagedDeviceChassisType(input string) (*ManagedDeviceChassisType, error) {
	vals := map[string]ManagedDeviceChassisType{
		"desktop":          ManagedDeviceChassisTypedesktop,
		"enterpriseserver": ManagedDeviceChassisTypeenterpriseServer,
		"laptop":           ManagedDeviceChassisTypelaptop,
		"mobileother":      ManagedDeviceChassisTypemobileOther,
		"mobileunknown":    ManagedDeviceChassisTypemobileUnknown,
		"phone":            ManagedDeviceChassisTypephone,
		"tablet":           ManagedDeviceChassisTypetablet,
		"unknown":          ManagedDeviceChassisTypeunknown,
		"worksworkstation": ManagedDeviceChassisTypeworksWorkstation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceChassisType(input)
	return &out, nil
}
