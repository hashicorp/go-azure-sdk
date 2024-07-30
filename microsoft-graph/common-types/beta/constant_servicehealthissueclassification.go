package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssueClassification string

const (
	ServiceHealthIssueClassification_Advisory ServiceHealthIssueClassification = "advisory"
	ServiceHealthIssueClassification_Incident ServiceHealthIssueClassification = "incident"
)

func PossibleValuesForServiceHealthIssueClassification() []string {
	return []string{
		string(ServiceHealthIssueClassification_Advisory),
		string(ServiceHealthIssueClassification_Incident),
	}
}

func (s *ServiceHealthIssueClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthIssueClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthIssueClassification(input string) (*ServiceHealthIssueClassification, error) {
	vals := map[string]ServiceHealthIssueClassification{
		"advisory": ServiceHealthIssueClassification_Advisory,
		"incident": ServiceHealthIssueClassification_Incident,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssueClassification(input)
	return &out, nil
}
