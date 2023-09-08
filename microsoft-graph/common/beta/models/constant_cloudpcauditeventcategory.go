package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventCategory string

const (
	CloudPCAuditEventCategorycloudPC CloudPCAuditEventCategory = "CloudPC"
	CloudPCAuditEventCategoryother   CloudPCAuditEventCategory = "Other"
)

func PossibleValuesForCloudPCAuditEventCategory() []string {
	return []string{
		string(CloudPCAuditEventCategorycloudPC),
		string(CloudPCAuditEventCategoryother),
	}
}

func parseCloudPCAuditEventCategory(input string) (*CloudPCAuditEventCategory, error) {
	vals := map[string]CloudPCAuditEventCategory{
		"cloudpc": CloudPCAuditEventCategorycloudPC,
		"other":   CloudPCAuditEventCategoryother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditEventCategory(input)
	return &out, nil
}
