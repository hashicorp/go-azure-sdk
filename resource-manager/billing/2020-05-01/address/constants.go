package address

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressValidationStatus string

const (
	AddressValidationStatusInvalid AddressValidationStatus = "Invalid"
	AddressValidationStatusValid   AddressValidationStatus = "Valid"
)

func PossibleValuesForAddressValidationStatus() []string {
	return []string{
		string(AddressValidationStatusInvalid),
		string(AddressValidationStatusValid),
	}
}

func (s *AddressValidationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddressValidationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddressValidationStatus(input string) (*AddressValidationStatus, error) {
	vals := map[string]AddressValidationStatus{
		"invalid": AddressValidationStatusInvalid,
		"valid":   AddressValidationStatusValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddressValidationStatus(input)
	return &out, nil
}
