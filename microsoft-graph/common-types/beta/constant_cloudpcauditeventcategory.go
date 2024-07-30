package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventCategory string

const (
	CloudPCAuditEventCategory_CloudPC CloudPCAuditEventCategory = "cloudPC"
	CloudPCAuditEventCategory_Other   CloudPCAuditEventCategory = "other"
)

func PossibleValuesForCloudPCAuditEventCategory() []string {
	return []string{
		string(CloudPCAuditEventCategory_CloudPC),
		string(CloudPCAuditEventCategory_Other),
	}
}

func (s *CloudPCAuditEventCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditEventCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditEventCategory(input string) (*CloudPCAuditEventCategory, error) {
	vals := map[string]CloudPCAuditEventCategory{
		"cloudpc": CloudPCAuditEventCategory_CloudPC,
		"other":   CloudPCAuditEventCategory_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditEventCategory(input)
	return &out, nil
}
