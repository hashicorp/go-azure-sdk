package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNowConnectionServiceNowConnectionStatus string

const (
	ServiceNowConnectionServiceNowConnectionStatusdisabled ServiceNowConnectionServiceNowConnectionStatus = "Disabled"
	ServiceNowConnectionServiceNowConnectionStatusenabled  ServiceNowConnectionServiceNowConnectionStatus = "Enabled"
)

func PossibleValuesForServiceNowConnectionServiceNowConnectionStatus() []string {
	return []string{
		string(ServiceNowConnectionServiceNowConnectionStatusdisabled),
		string(ServiceNowConnectionServiceNowConnectionStatusenabled),
	}
}

func parseServiceNowConnectionServiceNowConnectionStatus(input string) (*ServiceNowConnectionServiceNowConnectionStatus, error) {
	vals := map[string]ServiceNowConnectionServiceNowConnectionStatus{
		"disabled": ServiceNowConnectionServiceNowConnectionStatusdisabled,
		"enabled":  ServiceNowConnectionServiceNowConnectionStatusenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceNowConnectionServiceNowConnectionStatus(input)
	return &out, nil
}
