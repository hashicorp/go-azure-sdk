package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessIntegrityLevel string

const (
	ProcessIntegrityLevelhigh      ProcessIntegrityLevel = "High"
	ProcessIntegrityLevellow       ProcessIntegrityLevel = "Low"
	ProcessIntegrityLevelmedium    ProcessIntegrityLevel = "Medium"
	ProcessIntegrityLevelsystem    ProcessIntegrityLevel = "System"
	ProcessIntegrityLevelunknown   ProcessIntegrityLevel = "Unknown"
	ProcessIntegrityLeveluntrusted ProcessIntegrityLevel = "Untrusted"
)

func PossibleValuesForProcessIntegrityLevel() []string {
	return []string{
		string(ProcessIntegrityLevelhigh),
		string(ProcessIntegrityLevellow),
		string(ProcessIntegrityLevelmedium),
		string(ProcessIntegrityLevelsystem),
		string(ProcessIntegrityLevelunknown),
		string(ProcessIntegrityLeveluntrusted),
	}
}

func parseProcessIntegrityLevel(input string) (*ProcessIntegrityLevel, error) {
	vals := map[string]ProcessIntegrityLevel{
		"high":      ProcessIntegrityLevelhigh,
		"low":       ProcessIntegrityLevellow,
		"medium":    ProcessIntegrityLevelmedium,
		"system":    ProcessIntegrityLevelsystem,
		"unknown":   ProcessIntegrityLevelunknown,
		"untrusted": ProcessIntegrityLeveluntrusted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProcessIntegrityLevel(input)
	return &out, nil
}
