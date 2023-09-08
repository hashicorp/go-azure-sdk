package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectiondisabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection = "Disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectionenabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection = "Enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectionnotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection = "NotConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectiondisabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectionenabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectionnotConfigured),
	}
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectiondisabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectionenabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollectionnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection(input)
	return &out, nil
}
