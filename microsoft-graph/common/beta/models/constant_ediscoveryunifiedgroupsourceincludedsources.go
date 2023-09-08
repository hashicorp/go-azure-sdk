package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUnifiedGroupSourceIncludedSources string

const (
	EdiscoveryUnifiedGroupSourceIncludedSourcesmailbox EdiscoveryUnifiedGroupSourceIncludedSources = "Mailbox"
	EdiscoveryUnifiedGroupSourceIncludedSourcessite    EdiscoveryUnifiedGroupSourceIncludedSources = "Site"
)

func PossibleValuesForEdiscoveryUnifiedGroupSourceIncludedSources() []string {
	return []string{
		string(EdiscoveryUnifiedGroupSourceIncludedSourcesmailbox),
		string(EdiscoveryUnifiedGroupSourceIncludedSourcessite),
	}
}

func parseEdiscoveryUnifiedGroupSourceIncludedSources(input string) (*EdiscoveryUnifiedGroupSourceIncludedSources, error) {
	vals := map[string]EdiscoveryUnifiedGroupSourceIncludedSources{
		"mailbox": EdiscoveryUnifiedGroupSourceIncludedSourcesmailbox,
		"site":    EdiscoveryUnifiedGroupSourceIncludedSourcessite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUnifiedGroupSourceIncludedSources(input)
	return &out, nil
}
