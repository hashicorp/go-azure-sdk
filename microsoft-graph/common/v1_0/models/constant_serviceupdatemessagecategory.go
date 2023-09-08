package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageCategory string

const (
	ServiceUpdateMessageCategoryplanForChange     ServiceUpdateMessageCategory = "PlanForChange"
	ServiceUpdateMessageCategorypreventOrFixIssue ServiceUpdateMessageCategory = "PreventOrFixIssue"
	ServiceUpdateMessageCategorystayInformed      ServiceUpdateMessageCategory = "StayInformed"
)

func PossibleValuesForServiceUpdateMessageCategory() []string {
	return []string{
		string(ServiceUpdateMessageCategoryplanForChange),
		string(ServiceUpdateMessageCategorypreventOrFixIssue),
		string(ServiceUpdateMessageCategorystayInformed),
	}
}

func parseServiceUpdateMessageCategory(input string) (*ServiceUpdateMessageCategory, error) {
	vals := map[string]ServiceUpdateMessageCategory{
		"planforchange":     ServiceUpdateMessageCategoryplanForChange,
		"preventorfixissue": ServiceUpdateMessageCategorypreventOrFixIssue,
		"stayinformed":      ServiceUpdateMessageCategorystayInformed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceUpdateMessageCategory(input)
	return &out, nil
}
