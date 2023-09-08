package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostImportance string

const (
	PostImportancehigh   PostImportance = "High"
	PostImportancelow    PostImportance = "Low"
	PostImportancenormal PostImportance = "Normal"
)

func PossibleValuesForPostImportance() []string {
	return []string{
		string(PostImportancehigh),
		string(PostImportancelow),
		string(PostImportancenormal),
	}
}

func parsePostImportance(input string) (*PostImportance, error) {
	vals := map[string]PostImportance{
		"high":   PostImportancehigh,
		"low":    PostImportancelow,
		"normal": PostImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PostImportance(input)
	return &out, nil
}
