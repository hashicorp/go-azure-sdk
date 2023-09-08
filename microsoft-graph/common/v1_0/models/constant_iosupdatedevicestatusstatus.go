package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateDeviceStatusStatus string

const (
	IosUpdateDeviceStatusStatuscompliant     IosUpdateDeviceStatusStatus = "Compliant"
	IosUpdateDeviceStatusStatusconflict      IosUpdateDeviceStatusStatus = "Conflict"
	IosUpdateDeviceStatusStatuserror         IosUpdateDeviceStatusStatus = "Error"
	IosUpdateDeviceStatusStatusnonCompliant  IosUpdateDeviceStatusStatus = "NonCompliant"
	IosUpdateDeviceStatusStatusnotApplicable IosUpdateDeviceStatusStatus = "NotApplicable"
	IosUpdateDeviceStatusStatusnotAssigned   IosUpdateDeviceStatusStatus = "NotAssigned"
	IosUpdateDeviceStatusStatusremediated    IosUpdateDeviceStatusStatus = "Remediated"
	IosUpdateDeviceStatusStatusunknown       IosUpdateDeviceStatusStatus = "Unknown"
)

func PossibleValuesForIosUpdateDeviceStatusStatus() []string {
	return []string{
		string(IosUpdateDeviceStatusStatuscompliant),
		string(IosUpdateDeviceStatusStatusconflict),
		string(IosUpdateDeviceStatusStatuserror),
		string(IosUpdateDeviceStatusStatusnonCompliant),
		string(IosUpdateDeviceStatusStatusnotApplicable),
		string(IosUpdateDeviceStatusStatusnotAssigned),
		string(IosUpdateDeviceStatusStatusremediated),
		string(IosUpdateDeviceStatusStatusunknown),
	}
}

func parseIosUpdateDeviceStatusStatus(input string) (*IosUpdateDeviceStatusStatus, error) {
	vals := map[string]IosUpdateDeviceStatusStatus{
		"compliant":     IosUpdateDeviceStatusStatuscompliant,
		"conflict":      IosUpdateDeviceStatusStatusconflict,
		"error":         IosUpdateDeviceStatusStatuserror,
		"noncompliant":  IosUpdateDeviceStatusStatusnonCompliant,
		"notapplicable": IosUpdateDeviceStatusStatusnotApplicable,
		"notassigned":   IosUpdateDeviceStatusStatusnotAssigned,
		"remediated":    IosUpdateDeviceStatusStatusremediated,
		"unknown":       IosUpdateDeviceStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateDeviceStatusStatus(input)
	return &out, nil
}
