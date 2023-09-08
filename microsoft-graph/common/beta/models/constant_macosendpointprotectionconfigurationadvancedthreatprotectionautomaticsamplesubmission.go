package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissiondisabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission = "Disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissionenabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission = "Enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissionnotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission = "NotConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissiondisabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissionenabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissionnotConfigured),
	}
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissiondisabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissionenabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmissionnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission(input)
	return &out, nil
}
