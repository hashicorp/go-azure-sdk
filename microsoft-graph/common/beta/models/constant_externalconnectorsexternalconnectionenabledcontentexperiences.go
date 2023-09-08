package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnectionEnabledContentExperiences string

const (
	ExternalConnectorsExternalConnectionEnabledContentExperiencescompliance ExternalConnectorsExternalConnectionEnabledContentExperiences = "Compliance"
	ExternalConnectorsExternalConnectionEnabledContentExperiencessearch     ExternalConnectorsExternalConnectionEnabledContentExperiences = "Search"
)

func PossibleValuesForExternalConnectorsExternalConnectionEnabledContentExperiences() []string {
	return []string{
		string(ExternalConnectorsExternalConnectionEnabledContentExperiencescompliance),
		string(ExternalConnectorsExternalConnectionEnabledContentExperiencessearch),
	}
}

func parseExternalConnectorsExternalConnectionEnabledContentExperiences(input string) (*ExternalConnectorsExternalConnectionEnabledContentExperiences, error) {
	vals := map[string]ExternalConnectorsExternalConnectionEnabledContentExperiences{
		"compliance": ExternalConnectorsExternalConnectionEnabledContentExperiencescompliance,
		"search":     ExternalConnectorsExternalConnectionEnabledContentExperiencessearch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalConnectionEnabledContentExperiences(input)
	return &out, nil
}
