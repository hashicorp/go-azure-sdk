package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningErrorInfoErrorCategory string

const (
	ProvisioningErrorInfoErrorCategory_Failure           ProvisioningErrorInfoErrorCategory = "failure"
	ProvisioningErrorInfoErrorCategory_NonServiceFailure ProvisioningErrorInfoErrorCategory = "nonServiceFailure"
	ProvisioningErrorInfoErrorCategory_Success           ProvisioningErrorInfoErrorCategory = "success"
)

func PossibleValuesForProvisioningErrorInfoErrorCategory() []string {
	return []string{
		string(ProvisioningErrorInfoErrorCategory_Failure),
		string(ProvisioningErrorInfoErrorCategory_NonServiceFailure),
		string(ProvisioningErrorInfoErrorCategory_Success),
	}
}

func (s *ProvisioningErrorInfoErrorCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningErrorInfoErrorCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningErrorInfoErrorCategory(input string) (*ProvisioningErrorInfoErrorCategory, error) {
	vals := map[string]ProvisioningErrorInfoErrorCategory{
		"failure":           ProvisioningErrorInfoErrorCategory_Failure,
		"nonservicefailure": ProvisioningErrorInfoErrorCategory_NonServiceFailure,
		"success":           ProvisioningErrorInfoErrorCategory_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningErrorInfoErrorCategory(input)
	return &out, nil
}
