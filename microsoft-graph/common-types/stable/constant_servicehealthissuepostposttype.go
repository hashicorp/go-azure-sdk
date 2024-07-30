package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssuePostPostType string

const (
	ServiceHealthIssuePostPostType_Quick     ServiceHealthIssuePostPostType = "quick"
	ServiceHealthIssuePostPostType_Regular   ServiceHealthIssuePostPostType = "regular"
	ServiceHealthIssuePostPostType_Strategic ServiceHealthIssuePostPostType = "strategic"
)

func PossibleValuesForServiceHealthIssuePostPostType() []string {
	return []string{
		string(ServiceHealthIssuePostPostType_Quick),
		string(ServiceHealthIssuePostPostType_Regular),
		string(ServiceHealthIssuePostPostType_Strategic),
	}
}

func (s *ServiceHealthIssuePostPostType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthIssuePostPostType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthIssuePostPostType(input string) (*ServiceHealthIssuePostPostType, error) {
	vals := map[string]ServiceHealthIssuePostPostType{
		"quick":     ServiceHealthIssuePostPostType_Quick,
		"regular":   ServiceHealthIssuePostPostType_Regular,
		"strategic": ServiceHealthIssuePostPostType_Strategic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssuePostPostType(input)
	return &out, nil
}
