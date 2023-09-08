package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadComplexity string

const (
	PayloadComplexityhigh    PayloadComplexity = "High"
	PayloadComplexitylow     PayloadComplexity = "Low"
	PayloadComplexitymedium  PayloadComplexity = "Medium"
	PayloadComplexityunknown PayloadComplexity = "Unknown"
)

func PossibleValuesForPayloadComplexity() []string {
	return []string{
		string(PayloadComplexityhigh),
		string(PayloadComplexitylow),
		string(PayloadComplexitymedium),
		string(PayloadComplexityunknown),
	}
}

func parsePayloadComplexity(input string) (*PayloadComplexity, error) {
	vals := map[string]PayloadComplexity{
		"high":    PayloadComplexityhigh,
		"low":     PayloadComplexitylow,
		"medium":  PayloadComplexitymedium,
		"unknown": PayloadComplexityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadComplexity(input)
	return &out, nil
}
