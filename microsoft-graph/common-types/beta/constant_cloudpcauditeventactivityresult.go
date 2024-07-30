package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventActivityResult string

const (
	CloudPCAuditEventActivityResult_ClientError CloudPCAuditEventActivityResult = "clientError"
	CloudPCAuditEventActivityResult_Failure     CloudPCAuditEventActivityResult = "failure"
	CloudPCAuditEventActivityResult_Other       CloudPCAuditEventActivityResult = "other"
	CloudPCAuditEventActivityResult_Success     CloudPCAuditEventActivityResult = "success"
	CloudPCAuditEventActivityResult_Timeout     CloudPCAuditEventActivityResult = "timeout"
)

func PossibleValuesForCloudPCAuditEventActivityResult() []string {
	return []string{
		string(CloudPCAuditEventActivityResult_ClientError),
		string(CloudPCAuditEventActivityResult_Failure),
		string(CloudPCAuditEventActivityResult_Other),
		string(CloudPCAuditEventActivityResult_Success),
		string(CloudPCAuditEventActivityResult_Timeout),
	}
}

func (s *CloudPCAuditEventActivityResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditEventActivityResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditEventActivityResult(input string) (*CloudPCAuditEventActivityResult, error) {
	vals := map[string]CloudPCAuditEventActivityResult{
		"clienterror": CloudPCAuditEventActivityResult_ClientError,
		"failure":     CloudPCAuditEventActivityResult_Failure,
		"other":       CloudPCAuditEventActivityResult_Other,
		"success":     CloudPCAuditEventActivityResult_Success,
		"timeout":     CloudPCAuditEventActivityResult_Timeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditEventActivityResult(input)
	return &out, nil
}
