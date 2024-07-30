package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReportMigrationReadiness string

const (
	GroupPolicyMigrationReportMigrationReadiness_Complete      GroupPolicyMigrationReportMigrationReadiness = "complete"
	GroupPolicyMigrationReportMigrationReadiness_Error         GroupPolicyMigrationReportMigrationReadiness = "error"
	GroupPolicyMigrationReportMigrationReadiness_None          GroupPolicyMigrationReportMigrationReadiness = "none"
	GroupPolicyMigrationReportMigrationReadiness_NotApplicable GroupPolicyMigrationReportMigrationReadiness = "notApplicable"
	GroupPolicyMigrationReportMigrationReadiness_Partial       GroupPolicyMigrationReportMigrationReadiness = "partial"
)

func PossibleValuesForGroupPolicyMigrationReportMigrationReadiness() []string {
	return []string{
		string(GroupPolicyMigrationReportMigrationReadiness_Complete),
		string(GroupPolicyMigrationReportMigrationReadiness_Error),
		string(GroupPolicyMigrationReportMigrationReadiness_None),
		string(GroupPolicyMigrationReportMigrationReadiness_NotApplicable),
		string(GroupPolicyMigrationReportMigrationReadiness_Partial),
	}
}

func (s *GroupPolicyMigrationReportMigrationReadiness) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyMigrationReportMigrationReadiness(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyMigrationReportMigrationReadiness(input string) (*GroupPolicyMigrationReportMigrationReadiness, error) {
	vals := map[string]GroupPolicyMigrationReportMigrationReadiness{
		"complete":      GroupPolicyMigrationReportMigrationReadiness_Complete,
		"error":         GroupPolicyMigrationReportMigrationReadiness_Error,
		"none":          GroupPolicyMigrationReportMigrationReadiness_None,
		"notapplicable": GroupPolicyMigrationReportMigrationReadiness_NotApplicable,
		"partial":       GroupPolicyMigrationReportMigrationReadiness_Partial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyMigrationReportMigrationReadiness(input)
	return &out, nil
}
