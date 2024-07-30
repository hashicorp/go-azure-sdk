package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceHealthStatus string

const (
	SecurityDeviceEvidenceHealthStatus_Active                            SecurityDeviceEvidenceHealthStatus = "active"
	SecurityDeviceEvidenceHealthStatus_ImpairedCommunication             SecurityDeviceEvidenceHealthStatus = "impairedCommunication"
	SecurityDeviceEvidenceHealthStatus_Inactive                          SecurityDeviceEvidenceHealthStatus = "inactive"
	SecurityDeviceEvidenceHealthStatus_NoSensorData                      SecurityDeviceEvidenceHealthStatus = "noSensorData"
	SecurityDeviceEvidenceHealthStatus_NoSensorDataImpairedCommunication SecurityDeviceEvidenceHealthStatus = "noSensorDataImpairedCommunication"
	SecurityDeviceEvidenceHealthStatus_Unknown                           SecurityDeviceEvidenceHealthStatus = "unknown"
)

func PossibleValuesForSecurityDeviceEvidenceHealthStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceHealthStatus_Active),
		string(SecurityDeviceEvidenceHealthStatus_ImpairedCommunication),
		string(SecurityDeviceEvidenceHealthStatus_Inactive),
		string(SecurityDeviceEvidenceHealthStatus_NoSensorData),
		string(SecurityDeviceEvidenceHealthStatus_NoSensorDataImpairedCommunication),
		string(SecurityDeviceEvidenceHealthStatus_Unknown),
	}
}

func (s *SecurityDeviceEvidenceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceHealthStatus(input string) (*SecurityDeviceEvidenceHealthStatus, error) {
	vals := map[string]SecurityDeviceEvidenceHealthStatus{
		"active":                            SecurityDeviceEvidenceHealthStatus_Active,
		"impairedcommunication":             SecurityDeviceEvidenceHealthStatus_ImpairedCommunication,
		"inactive":                          SecurityDeviceEvidenceHealthStatus_Inactive,
		"nosensordata":                      SecurityDeviceEvidenceHealthStatus_NoSensorData,
		"nosensordataimpairedcommunication": SecurityDeviceEvidenceHealthStatus_NoSensorDataImpairedCommunication,
		"unknown":                           SecurityDeviceEvidenceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceHealthStatus(input)
	return &out, nil
}
