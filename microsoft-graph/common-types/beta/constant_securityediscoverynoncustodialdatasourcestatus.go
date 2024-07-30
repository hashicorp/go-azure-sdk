package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryNoncustodialDataSourceStatus string

const (
	SecurityEdiscoveryNoncustodialDataSourceStatus_Active   SecurityEdiscoveryNoncustodialDataSourceStatus = "active"
	SecurityEdiscoveryNoncustodialDataSourceStatus_Released SecurityEdiscoveryNoncustodialDataSourceStatus = "released"
)

func PossibleValuesForSecurityEdiscoveryNoncustodialDataSourceStatus() []string {
	return []string{
		string(SecurityEdiscoveryNoncustodialDataSourceStatus_Active),
		string(SecurityEdiscoveryNoncustodialDataSourceStatus_Released),
	}
}

func (s *SecurityEdiscoveryNoncustodialDataSourceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryNoncustodialDataSourceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryNoncustodialDataSourceStatus(input string) (*SecurityEdiscoveryNoncustodialDataSourceStatus, error) {
	vals := map[string]SecurityEdiscoveryNoncustodialDataSourceStatus{
		"active":   SecurityEdiscoveryNoncustodialDataSourceStatus_Active,
		"released": SecurityEdiscoveryNoncustodialDataSourceStatus_Released,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryNoncustodialDataSourceStatus(input)
	return &out, nil
}
