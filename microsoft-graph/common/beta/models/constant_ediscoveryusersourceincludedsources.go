package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUserSourceIncludedSources string

const (
	EdiscoveryUserSourceIncludedSourcesmailbox EdiscoveryUserSourceIncludedSources = "Mailbox"
	EdiscoveryUserSourceIncludedSourcessite    EdiscoveryUserSourceIncludedSources = "Site"
)

func PossibleValuesForEdiscoveryUserSourceIncludedSources() []string {
	return []string{
		string(EdiscoveryUserSourceIncludedSourcesmailbox),
		string(EdiscoveryUserSourceIncludedSourcessite),
	}
}

func parseEdiscoveryUserSourceIncludedSources(input string) (*EdiscoveryUserSourceIncludedSources, error) {
	vals := map[string]EdiscoveryUserSourceIncludedSources{
		"mailbox": EdiscoveryUserSourceIncludedSourcesmailbox,
		"site":    EdiscoveryUserSourceIncludedSourcessite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUserSourceIncludedSources(input)
	return &out, nil
}
