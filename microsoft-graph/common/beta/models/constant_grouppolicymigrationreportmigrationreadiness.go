package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReportMigrationReadiness string

const (
	GroupPolicyMigrationReportMigrationReadinesscomplete      GroupPolicyMigrationReportMigrationReadiness = "Complete"
	GroupPolicyMigrationReportMigrationReadinesserror         GroupPolicyMigrationReportMigrationReadiness = "Error"
	GroupPolicyMigrationReportMigrationReadinessnone          GroupPolicyMigrationReportMigrationReadiness = "None"
	GroupPolicyMigrationReportMigrationReadinessnotApplicable GroupPolicyMigrationReportMigrationReadiness = "NotApplicable"
	GroupPolicyMigrationReportMigrationReadinesspartial       GroupPolicyMigrationReportMigrationReadiness = "Partial"
)

func PossibleValuesForGroupPolicyMigrationReportMigrationReadiness() []string {
	return []string{
		string(GroupPolicyMigrationReportMigrationReadinesscomplete),
		string(GroupPolicyMigrationReportMigrationReadinesserror),
		string(GroupPolicyMigrationReportMigrationReadinessnone),
		string(GroupPolicyMigrationReportMigrationReadinessnotApplicable),
		string(GroupPolicyMigrationReportMigrationReadinesspartial),
	}
}

func parseGroupPolicyMigrationReportMigrationReadiness(input string) (*GroupPolicyMigrationReportMigrationReadiness, error) {
	vals := map[string]GroupPolicyMigrationReportMigrationReadiness{
		"complete":      GroupPolicyMigrationReportMigrationReadinesscomplete,
		"error":         GroupPolicyMigrationReportMigrationReadinesserror,
		"none":          GroupPolicyMigrationReportMigrationReadinessnone,
		"notapplicable": GroupPolicyMigrationReportMigrationReadinessnotApplicable,
		"partial":       GroupPolicyMigrationReportMigrationReadinesspartial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyMigrationReportMigrationReadiness(input)
	return &out, nil
}
