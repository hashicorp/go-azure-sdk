package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssueOrigin string

const (
	ServiceHealthIssueOrigincustomer   ServiceHealthIssueOrigin = "Customer"
	ServiceHealthIssueOriginmicrosoft  ServiceHealthIssueOrigin = "Microsoft"
	ServiceHealthIssueOriginthirdParty ServiceHealthIssueOrigin = "ThirdParty"
)

func PossibleValuesForServiceHealthIssueOrigin() []string {
	return []string{
		string(ServiceHealthIssueOrigincustomer),
		string(ServiceHealthIssueOriginmicrosoft),
		string(ServiceHealthIssueOriginthirdParty),
	}
}

func parseServiceHealthIssueOrigin(input string) (*ServiceHealthIssueOrigin, error) {
	vals := map[string]ServiceHealthIssueOrigin{
		"customer":   ServiceHealthIssueOrigincustomer,
		"microsoft":  ServiceHealthIssueOriginmicrosoft,
		"thirdparty": ServiceHealthIssueOriginthirdParty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssueOrigin(input)
	return &out, nil
}
