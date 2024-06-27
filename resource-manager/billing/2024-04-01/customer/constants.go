package customer

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerStatus string

const (
	CustomerStatusActive      CustomerStatus = "Active"
	CustomerStatusDeleted     CustomerStatus = "Deleted"
	CustomerStatusDisabled    CustomerStatus = "Disabled"
	CustomerStatusOther       CustomerStatus = "Other"
	CustomerStatusPending     CustomerStatus = "Pending"
	CustomerStatusUnderReview CustomerStatus = "UnderReview"
	CustomerStatusWarned      CustomerStatus = "Warned"
)

func PossibleValuesForCustomerStatus() []string {
	return []string{
		string(CustomerStatusActive),
		string(CustomerStatusDeleted),
		string(CustomerStatusDisabled),
		string(CustomerStatusOther),
		string(CustomerStatusPending),
		string(CustomerStatusUnderReview),
		string(CustomerStatusWarned),
	}
}

func (s *CustomerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomerStatus(input string) (*CustomerStatus, error) {
	vals := map[string]CustomerStatus{
		"active":      CustomerStatusActive,
		"deleted":     CustomerStatusDeleted,
		"disabled":    CustomerStatusDisabled,
		"other":       CustomerStatusOther,
		"pending":     CustomerStatusPending,
		"underreview": CustomerStatusUnderReview,
		"warned":      CustomerStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomerStatus(input)
	return &out, nil
}
