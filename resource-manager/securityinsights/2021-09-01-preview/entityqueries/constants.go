package entityqueries

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomEntityQueryKind string

const (
	CustomEntityQueryKindActivity CustomEntityQueryKind = "Activity"
)

func PossibleValuesForCustomEntityQueryKind() []string {
	return []string{
		string(CustomEntityQueryKindActivity),
	}
}

func parseCustomEntityQueryKind(input string) (*CustomEntityQueryKind, error) {
	vals := map[string]CustomEntityQueryKind{
		"activity": CustomEntityQueryKindActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomEntityQueryKind(input)
	return &out, nil
}

type EntityQueryKind string

const (
	EntityQueryKindActivity  EntityQueryKind = "Activity"
	EntityQueryKindExpansion EntityQueryKind = "Expansion"
	EntityQueryKindInsight   EntityQueryKind = "Insight"
)

func PossibleValuesForEntityQueryKind() []string {
	return []string{
		string(EntityQueryKindActivity),
		string(EntityQueryKindExpansion),
		string(EntityQueryKindInsight),
	}
}

func parseEntityQueryKind(input string) (*EntityQueryKind, error) {
	vals := map[string]EntityQueryKind{
		"activity":  EntityQueryKindActivity,
		"expansion": EntityQueryKindExpansion,
		"insight":   EntityQueryKindInsight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityQueryKind(input)
	return &out, nil
}

type EntityQueryTemplateKind string

const (
	EntityQueryTemplateKindActivity EntityQueryTemplateKind = "Activity"
)

func PossibleValuesForEntityQueryTemplateKind() []string {
	return []string{
		string(EntityQueryTemplateKindActivity),
	}
}

func parseEntityQueryTemplateKind(input string) (*EntityQueryTemplateKind, error) {
	vals := map[string]EntityQueryTemplateKind{
		"activity": EntityQueryTemplateKindActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityQueryTemplateKind(input)
	return &out, nil
}

type Kind string

const (
	KindActivity  Kind = "Activity"
	KindExpansion Kind = "Expansion"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindActivity),
		string(KindExpansion),
	}
}

func parseKind(input string) (*Kind, error) {
	vals := map[string]Kind{
		"activity":  KindActivity,
		"expansion": KindExpansion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Kind(input)
	return &out, nil
}
