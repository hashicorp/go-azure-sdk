package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleSummaryStatus string

const (
	PrivilegedRoleSummaryStatus_Bad PrivilegedRoleSummaryStatus = "bad"
	PrivilegedRoleSummaryStatus_Ok  PrivilegedRoleSummaryStatus = "ok"
)

func PossibleValuesForPrivilegedRoleSummaryStatus() []string {
	return []string{
		string(PrivilegedRoleSummaryStatus_Bad),
		string(PrivilegedRoleSummaryStatus_Ok),
	}
}

func (s *PrivilegedRoleSummaryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedRoleSummaryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedRoleSummaryStatus(input string) (*PrivilegedRoleSummaryStatus, error) {
	vals := map[string]PrivilegedRoleSummaryStatus{
		"bad": PrivilegedRoleSummaryStatus_Bad,
		"ok":  PrivilegedRoleSummaryStatus_Ok,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedRoleSummaryStatus(input)
	return &out, nil
}
