package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUnifiedGroupSourceIncludedSources string

const (
	SecurityUnifiedGroupSourceIncludedSourcesmailbox SecurityUnifiedGroupSourceIncludedSources = "Mailbox"
	SecurityUnifiedGroupSourceIncludedSourcessite    SecurityUnifiedGroupSourceIncludedSources = "Site"
)

func PossibleValuesForSecurityUnifiedGroupSourceIncludedSources() []string {
	return []string{
		string(SecurityUnifiedGroupSourceIncludedSourcesmailbox),
		string(SecurityUnifiedGroupSourceIncludedSourcessite),
	}
}

func parseSecurityUnifiedGroupSourceIncludedSources(input string) (*SecurityUnifiedGroupSourceIncludedSources, error) {
	vals := map[string]SecurityUnifiedGroupSourceIncludedSources{
		"mailbox": SecurityUnifiedGroupSourceIncludedSourcesmailbox,
		"site":    SecurityUnifiedGroupSourceIncludedSourcessite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUnifiedGroupSourceIncludedSources(input)
	return &out, nil
}
