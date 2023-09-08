package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssueClassification string

const (
	ServiceHealthIssueClassificationadvisory ServiceHealthIssueClassification = "Advisory"
	ServiceHealthIssueClassificationincident ServiceHealthIssueClassification = "Incident"
)

func PossibleValuesForServiceHealthIssueClassification() []string {
	return []string{
		string(ServiceHealthIssueClassificationadvisory),
		string(ServiceHealthIssueClassificationincident),
	}
}

func parseServiceHealthIssueClassification(input string) (*ServiceHealthIssueClassification, error) {
	vals := map[string]ServiceHealthIssueClassification{
		"advisory": ServiceHealthIssueClassificationadvisory,
		"incident": ServiceHealthIssueClassificationincident,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssueClassification(input)
	return &out, nil
}
