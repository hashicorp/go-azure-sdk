package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryNoncustodialDataSourceStatus string

const (
	EdiscoveryNoncustodialDataSourceStatus_Active   EdiscoveryNoncustodialDataSourceStatus = "Active"
	EdiscoveryNoncustodialDataSourceStatus_Released EdiscoveryNoncustodialDataSourceStatus = "Released"
)

func PossibleValuesForEdiscoveryNoncustodialDataSourceStatus() []string {
	return []string{
		string(EdiscoveryNoncustodialDataSourceStatus_Active),
		string(EdiscoveryNoncustodialDataSourceStatus_Released),
	}
}

func (s *EdiscoveryNoncustodialDataSourceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryNoncustodialDataSourceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryNoncustodialDataSourceStatus(input string) (*EdiscoveryNoncustodialDataSourceStatus, error) {
	vals := map[string]EdiscoveryNoncustodialDataSourceStatus{
		"active":   EdiscoveryNoncustodialDataSourceStatus_Active,
		"released": EdiscoveryNoncustodialDataSourceStatus_Released,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryNoncustodialDataSourceStatus(input)
	return &out, nil
}
