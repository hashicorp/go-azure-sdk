package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventActivityOperationType string

const (
	CloudPCAuditEventActivityOperationType_Create CloudPCAuditEventActivityOperationType = "create"
	CloudPCAuditEventActivityOperationType_Delete CloudPCAuditEventActivityOperationType = "delete"
	CloudPCAuditEventActivityOperationType_Other  CloudPCAuditEventActivityOperationType = "other"
	CloudPCAuditEventActivityOperationType_Patch  CloudPCAuditEventActivityOperationType = "patch"
)

func PossibleValuesForCloudPCAuditEventActivityOperationType() []string {
	return []string{
		string(CloudPCAuditEventActivityOperationType_Create),
		string(CloudPCAuditEventActivityOperationType_Delete),
		string(CloudPCAuditEventActivityOperationType_Other),
		string(CloudPCAuditEventActivityOperationType_Patch),
	}
}

func (s *CloudPCAuditEventActivityOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditEventActivityOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditEventActivityOperationType(input string) (*CloudPCAuditEventActivityOperationType, error) {
	vals := map[string]CloudPCAuditEventActivityOperationType{
		"create": CloudPCAuditEventActivityOperationType_Create,
		"delete": CloudPCAuditEventActivityOperationType_Delete,
		"other":  CloudPCAuditEventActivityOperationType_Other,
		"patch":  CloudPCAuditEventActivityOperationType_Patch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditEventActivityOperationType(input)
	return &out, nil
}
