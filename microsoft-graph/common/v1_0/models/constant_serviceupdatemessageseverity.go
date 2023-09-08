package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageSeverity string

const (
	ServiceUpdateMessageSeveritycritical ServiceUpdateMessageSeverity = "Critical"
	ServiceUpdateMessageSeverityhigh     ServiceUpdateMessageSeverity = "High"
	ServiceUpdateMessageSeveritynormal   ServiceUpdateMessageSeverity = "Normal"
)

func PossibleValuesForServiceUpdateMessageSeverity() []string {
	return []string{
		string(ServiceUpdateMessageSeveritycritical),
		string(ServiceUpdateMessageSeverityhigh),
		string(ServiceUpdateMessageSeveritynormal),
	}
}

func parseServiceUpdateMessageSeverity(input string) (*ServiceUpdateMessageSeverity, error) {
	vals := map[string]ServiceUpdateMessageSeverity{
		"critical": ServiceUpdateMessageSeveritycritical,
		"high":     ServiceUpdateMessageSeverityhigh,
		"normal":   ServiceUpdateMessageSeveritynormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceUpdateMessageSeverity(input)
	return &out, nil
}
