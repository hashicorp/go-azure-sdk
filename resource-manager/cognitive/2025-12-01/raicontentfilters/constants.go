package raicontentfilters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiPolicyContentSource string

const (
	RaiPolicyContentSourceCompletion   RaiPolicyContentSource = "Completion"
	RaiPolicyContentSourcePostRun      RaiPolicyContentSource = "PostRun"
	RaiPolicyContentSourcePostToolCall RaiPolicyContentSource = "PostToolCall"
	RaiPolicyContentSourcePreRun       RaiPolicyContentSource = "PreRun"
	RaiPolicyContentSourcePreToolCall  RaiPolicyContentSource = "PreToolCall"
	RaiPolicyContentSourcePrompt       RaiPolicyContentSource = "Prompt"
)

func PossibleValuesForRaiPolicyContentSource() []string {
	return []string{
		string(RaiPolicyContentSourceCompletion),
		string(RaiPolicyContentSourcePostRun),
		string(RaiPolicyContentSourcePostToolCall),
		string(RaiPolicyContentSourcePreRun),
		string(RaiPolicyContentSourcePreToolCall),
		string(RaiPolicyContentSourcePrompt),
	}
}

func (s *RaiPolicyContentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyContentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyContentSource(input string) (*RaiPolicyContentSource, error) {
	vals := map[string]RaiPolicyContentSource{
		"completion":   RaiPolicyContentSourceCompletion,
		"postrun":      RaiPolicyContentSourcePostRun,
		"posttoolcall": RaiPolicyContentSourcePostToolCall,
		"prerun":       RaiPolicyContentSourcePreRun,
		"pretoolcall":  RaiPolicyContentSourcePreToolCall,
		"prompt":       RaiPolicyContentSourcePrompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyContentSource(input)
	return &out, nil
}
