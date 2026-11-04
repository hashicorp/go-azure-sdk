package virtualclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DNSRefreshOperationStatus string

const (
	DNSRefreshOperationStatusFailed     DNSRefreshOperationStatus = "Failed"
	DNSRefreshOperationStatusInProgress DNSRefreshOperationStatus = "InProgress"
	DNSRefreshOperationStatusSucceeded  DNSRefreshOperationStatus = "Succeeded"
)

func PossibleValuesForDNSRefreshOperationStatus() []string {
	return []string{
		string(DNSRefreshOperationStatusFailed),
		string(DNSRefreshOperationStatusInProgress),
		string(DNSRefreshOperationStatusSucceeded),
	}
}

func (s *DNSRefreshOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDNSRefreshOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDNSRefreshOperationStatus(input string) (*DNSRefreshOperationStatus, error) {
	vals := map[string]DNSRefreshOperationStatus{
		"failed":     DNSRefreshOperationStatusFailed,
		"inprogress": DNSRefreshOperationStatusInProgress,
		"succeeded":  DNSRefreshOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DNSRefreshOperationStatus(input)
	return &out, nil
}
