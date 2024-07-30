package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssueOrigin string

const (
	ServiceHealthIssueOrigin_Customer   ServiceHealthIssueOrigin = "customer"
	ServiceHealthIssueOrigin_Microsoft  ServiceHealthIssueOrigin = "microsoft"
	ServiceHealthIssueOrigin_ThirdParty ServiceHealthIssueOrigin = "thirdParty"
)

func PossibleValuesForServiceHealthIssueOrigin() []string {
	return []string{
		string(ServiceHealthIssueOrigin_Customer),
		string(ServiceHealthIssueOrigin_Microsoft),
		string(ServiceHealthIssueOrigin_ThirdParty),
	}
}

func (s *ServiceHealthIssueOrigin) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthIssueOrigin(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthIssueOrigin(input string) (*ServiceHealthIssueOrigin, error) {
	vals := map[string]ServiceHealthIssueOrigin{
		"customer":   ServiceHealthIssueOrigin_Customer,
		"microsoft":  ServiceHealthIssueOrigin_Microsoft,
		"thirdparty": ServiceHealthIssueOrigin_ThirdParty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssueOrigin(input)
	return &out, nil
}
