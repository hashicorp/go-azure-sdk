package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserSourceIncludedSources string

const (
	SecurityUserSourceIncludedSourcesmailbox SecurityUserSourceIncludedSources = "Mailbox"
	SecurityUserSourceIncludedSourcessite    SecurityUserSourceIncludedSources = "Site"
)

func PossibleValuesForSecurityUserSourceIncludedSources() []string {
	return []string{
		string(SecurityUserSourceIncludedSourcesmailbox),
		string(SecurityUserSourceIncludedSourcessite),
	}
}

func parseSecurityUserSourceIncludedSources(input string) (*SecurityUserSourceIncludedSources, error) {
	vals := map[string]SecurityUserSourceIncludedSources{
		"mailbox": SecurityUserSourceIncludedSourcesmailbox,
		"site":    SecurityUserSourceIncludedSourcessite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserSourceIncludedSources(input)
	return &out, nil
}
