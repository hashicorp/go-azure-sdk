package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationPriority string

const (
	MacOSSoftwareUpdateConfigurationPriorityhigh MacOSSoftwareUpdateConfigurationPriority = "High"
	MacOSSoftwareUpdateConfigurationPrioritylow  MacOSSoftwareUpdateConfigurationPriority = "Low"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationPriority() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationPriorityhigh),
		string(MacOSSoftwareUpdateConfigurationPrioritylow),
	}
}

func parseMacOSSoftwareUpdateConfigurationPriority(input string) (*MacOSSoftwareUpdateConfigurationPriority, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationPriority{
		"high": MacOSSoftwareUpdateConfigurationPriorityhigh,
		"low":  MacOSSoftwareUpdateConfigurationPrioritylow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationPriority(input)
	return &out, nil
}
