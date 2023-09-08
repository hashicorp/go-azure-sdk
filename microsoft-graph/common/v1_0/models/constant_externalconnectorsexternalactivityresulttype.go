package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalActivityResultType string

const (
	ExternalConnectorsExternalActivityResultTypecommented ExternalConnectorsExternalActivityResultType = "Commented"
	ExternalConnectorsExternalActivityResultTypecreated   ExternalConnectorsExternalActivityResultType = "Created"
	ExternalConnectorsExternalActivityResultTypemodified  ExternalConnectorsExternalActivityResultType = "Modified"
	ExternalConnectorsExternalActivityResultTypeviewed    ExternalConnectorsExternalActivityResultType = "Viewed"
)

func PossibleValuesForExternalConnectorsExternalActivityResultType() []string {
	return []string{
		string(ExternalConnectorsExternalActivityResultTypecommented),
		string(ExternalConnectorsExternalActivityResultTypecreated),
		string(ExternalConnectorsExternalActivityResultTypemodified),
		string(ExternalConnectorsExternalActivityResultTypeviewed),
	}
}

func parseExternalConnectorsExternalActivityResultType(input string) (*ExternalConnectorsExternalActivityResultType, error) {
	vals := map[string]ExternalConnectorsExternalActivityResultType{
		"commented": ExternalConnectorsExternalActivityResultTypecommented,
		"created":   ExternalConnectorsExternalActivityResultTypecreated,
		"modified":  ExternalConnectorsExternalActivityResultTypemodified,
		"viewed":    ExternalConnectorsExternalActivityResultTypeviewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalActivityResultType(input)
	return &out, nil
}
