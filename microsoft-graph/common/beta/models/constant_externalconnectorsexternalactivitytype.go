package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalActivityType string

const (
	ExternalConnectorsExternalActivityTypecommented ExternalConnectorsExternalActivityType = "Commented"
	ExternalConnectorsExternalActivityTypecreated   ExternalConnectorsExternalActivityType = "Created"
	ExternalConnectorsExternalActivityTypemodified  ExternalConnectorsExternalActivityType = "Modified"
	ExternalConnectorsExternalActivityTypeviewed    ExternalConnectorsExternalActivityType = "Viewed"
)

func PossibleValuesForExternalConnectorsExternalActivityType() []string {
	return []string{
		string(ExternalConnectorsExternalActivityTypecommented),
		string(ExternalConnectorsExternalActivityTypecreated),
		string(ExternalConnectorsExternalActivityTypemodified),
		string(ExternalConnectorsExternalActivityTypeviewed),
	}
}

func parseExternalConnectorsExternalActivityType(input string) (*ExternalConnectorsExternalActivityType, error) {
	vals := map[string]ExternalConnectorsExternalActivityType{
		"commented": ExternalConnectorsExternalActivityTypecommented,
		"created":   ExternalConnectorsExternalActivityTypecreated,
		"modified":  ExternalConnectorsExternalActivityTypemodified,
		"viewed":    ExternalConnectorsExternalActivityTypeviewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalActivityType(input)
	return &out, nil
}
