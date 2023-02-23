package metadata

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Operator string

const (
	OperatorAND Operator = "AND"
	OperatorOR  Operator = "OR"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(OperatorAND),
		string(OperatorOR),
	}
}

func parseOperator(input string) (*Operator, error) {
	vals := map[string]Operator{
		"and": OperatorAND,
		"or":  OperatorOR,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Operator(input)
	return &out, nil
}

type SourceKind string

const (
	SourceKindCommunity        SourceKind = "Community"
	SourceKindLocalWorkspace   SourceKind = "LocalWorkspace"
	SourceKindSolution         SourceKind = "Solution"
	SourceKindSourceRepository SourceKind = "SourceRepository"
)

func PossibleValuesForSourceKind() []string {
	return []string{
		string(SourceKindCommunity),
		string(SourceKindLocalWorkspace),
		string(SourceKindSolution),
		string(SourceKindSourceRepository),
	}
}

func parseSourceKind(input string) (*SourceKind, error) {
	vals := map[string]SourceKind{
		"community":        SourceKindCommunity,
		"localworkspace":   SourceKindLocalWorkspace,
		"solution":         SourceKindSolution,
		"sourcerepository": SourceKindSourceRepository,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceKind(input)
	return &out, nil
}

type SupportTier string

const (
	SupportTierCommunity SupportTier = "Community"
	SupportTierMicrosoft SupportTier = "Microsoft"
	SupportTierPartner   SupportTier = "Partner"
)

func PossibleValuesForSupportTier() []string {
	return []string{
		string(SupportTierCommunity),
		string(SupportTierMicrosoft),
		string(SupportTierPartner),
	}
}

func parseSupportTier(input string) (*SupportTier, error) {
	vals := map[string]SupportTier{
		"community": SupportTierCommunity,
		"microsoft": SupportTierMicrosoft,
		"partner":   SupportTierPartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportTier(input)
	return &out, nil
}
