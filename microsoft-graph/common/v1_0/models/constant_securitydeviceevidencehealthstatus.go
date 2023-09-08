package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceHealthStatus string

const (
	SecurityDeviceEvidenceHealthStatusactive                            SecurityDeviceEvidenceHealthStatus = "Active"
	SecurityDeviceEvidenceHealthStatusimpairedCommunication             SecurityDeviceEvidenceHealthStatus = "ImpairedCommunication"
	SecurityDeviceEvidenceHealthStatusinactive                          SecurityDeviceEvidenceHealthStatus = "Inactive"
	SecurityDeviceEvidenceHealthStatusnoSensorData                      SecurityDeviceEvidenceHealthStatus = "NoSensorData"
	SecurityDeviceEvidenceHealthStatusnoSensorDataImpairedCommunication SecurityDeviceEvidenceHealthStatus = "NoSensorDataImpairedCommunication"
	SecurityDeviceEvidenceHealthStatusunknown                           SecurityDeviceEvidenceHealthStatus = "Unknown"
)

func PossibleValuesForSecurityDeviceEvidenceHealthStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceHealthStatusactive),
		string(SecurityDeviceEvidenceHealthStatusimpairedCommunication),
		string(SecurityDeviceEvidenceHealthStatusinactive),
		string(SecurityDeviceEvidenceHealthStatusnoSensorData),
		string(SecurityDeviceEvidenceHealthStatusnoSensorDataImpairedCommunication),
		string(SecurityDeviceEvidenceHealthStatusunknown),
	}
}

func parseSecurityDeviceEvidenceHealthStatus(input string) (*SecurityDeviceEvidenceHealthStatus, error) {
	vals := map[string]SecurityDeviceEvidenceHealthStatus{
		"active":                            SecurityDeviceEvidenceHealthStatusactive,
		"impairedcommunication":             SecurityDeviceEvidenceHealthStatusimpairedCommunication,
		"inactive":                          SecurityDeviceEvidenceHealthStatusinactive,
		"nosensordata":                      SecurityDeviceEvidenceHealthStatusnoSensorData,
		"nosensordataimpairedcommunication": SecurityDeviceEvidenceHealthStatusnoSensorDataImpairedCommunication,
		"unknown":                           SecurityDeviceEvidenceHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceHealthStatus(input)
	return &out, nil
}
