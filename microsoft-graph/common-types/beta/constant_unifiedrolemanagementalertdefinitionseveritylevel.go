package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementAlertDefinitionSeverityLevel string

const (
	UnifiedRoleManagementAlertDefinitionSeverityLevel_High          UnifiedRoleManagementAlertDefinitionSeverityLevel = "high"
	UnifiedRoleManagementAlertDefinitionSeverityLevel_Informational UnifiedRoleManagementAlertDefinitionSeverityLevel = "informational"
	UnifiedRoleManagementAlertDefinitionSeverityLevel_Low           UnifiedRoleManagementAlertDefinitionSeverityLevel = "low"
	UnifiedRoleManagementAlertDefinitionSeverityLevel_Medium        UnifiedRoleManagementAlertDefinitionSeverityLevel = "medium"
	UnifiedRoleManagementAlertDefinitionSeverityLevel_Unknown       UnifiedRoleManagementAlertDefinitionSeverityLevel = "unknown"
)

func PossibleValuesForUnifiedRoleManagementAlertDefinitionSeverityLevel() []string {
	return []string{
		string(UnifiedRoleManagementAlertDefinitionSeverityLevel_High),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevel_Informational),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevel_Low),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevel_Medium),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevel_Unknown),
	}
}

func (s *UnifiedRoleManagementAlertDefinitionSeverityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnifiedRoleManagementAlertDefinitionSeverityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnifiedRoleManagementAlertDefinitionSeverityLevel(input string) (*UnifiedRoleManagementAlertDefinitionSeverityLevel, error) {
	vals := map[string]UnifiedRoleManagementAlertDefinitionSeverityLevel{
		"high":          UnifiedRoleManagementAlertDefinitionSeverityLevel_High,
		"informational": UnifiedRoleManagementAlertDefinitionSeverityLevel_Informational,
		"low":           UnifiedRoleManagementAlertDefinitionSeverityLevel_Low,
		"medium":        UnifiedRoleManagementAlertDefinitionSeverityLevel_Medium,
		"unknown":       UnifiedRoleManagementAlertDefinitionSeverityLevel_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleManagementAlertDefinitionSeverityLevel(input)
	return &out, nil
}
