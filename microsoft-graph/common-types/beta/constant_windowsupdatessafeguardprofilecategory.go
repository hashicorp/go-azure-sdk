package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSafeguardProfileCategory string

const (
	WindowsUpdatesSafeguardProfileCategory_LikelyIssues WindowsUpdatesSafeguardProfileCategory = "likelyIssues"
)

func PossibleValuesForWindowsUpdatesSafeguardProfileCategory() []string {
	return []string{
		string(WindowsUpdatesSafeguardProfileCategory_LikelyIssues),
	}
}

func (s *WindowsUpdatesSafeguardProfileCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesSafeguardProfileCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesSafeguardProfileCategory(input string) (*WindowsUpdatesSafeguardProfileCategory, error) {
	vals := map[string]WindowsUpdatesSafeguardProfileCategory{
		"likelyissues": WindowsUpdatesSafeguardProfileCategory_LikelyIssues,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesSafeguardProfileCategory(input)
	return &out, nil
}
