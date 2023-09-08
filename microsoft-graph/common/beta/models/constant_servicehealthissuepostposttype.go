package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssuePostPostType string

const (
	ServiceHealthIssuePostPostTypequick     ServiceHealthIssuePostPostType = "Quick"
	ServiceHealthIssuePostPostTyperegular   ServiceHealthIssuePostPostType = "Regular"
	ServiceHealthIssuePostPostTypestrategic ServiceHealthIssuePostPostType = "Strategic"
)

func PossibleValuesForServiceHealthIssuePostPostType() []string {
	return []string{
		string(ServiceHealthIssuePostPostTypequick),
		string(ServiceHealthIssuePostPostTyperegular),
		string(ServiceHealthIssuePostPostTypestrategic),
	}
}

func parseServiceHealthIssuePostPostType(input string) (*ServiceHealthIssuePostPostType, error) {
	vals := map[string]ServiceHealthIssuePostPostType{
		"quick":     ServiceHealthIssuePostPostTypequick,
		"regular":   ServiceHealthIssuePostPostTyperegular,
		"strategic": ServiceHealthIssuePostPostTypestrategic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssuePostPostType(input)
	return &out, nil
}
