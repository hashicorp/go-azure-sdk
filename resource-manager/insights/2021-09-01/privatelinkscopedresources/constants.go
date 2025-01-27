package privatelinkscopedresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopedResourceProvisioningState string

const (
	ScopedResourceProvisioningStateCanceled     ScopedResourceProvisioningState = "Canceled"
	ScopedResourceProvisioningStateFailed       ScopedResourceProvisioningState = "Failed"
	ScopedResourceProvisioningStateProvisioning ScopedResourceProvisioningState = "Provisioning"
	ScopedResourceProvisioningStateSucceeded    ScopedResourceProvisioningState = "Succeeded"
)

func PossibleValuesForScopedResourceProvisioningState() []string {
	return []string{
		string(ScopedResourceProvisioningStateCanceled),
		string(ScopedResourceProvisioningStateFailed),
		string(ScopedResourceProvisioningStateProvisioning),
		string(ScopedResourceProvisioningStateSucceeded),
	}
}

func (s *ScopedResourceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScopedResourceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScopedResourceProvisioningState(input string) (*ScopedResourceProvisioningState, error) {
	vals := map[string]ScopedResourceProvisioningState{
		"canceled":     ScopedResourceProvisioningStateCanceled,
		"failed":       ScopedResourceProvisioningStateFailed,
		"provisioning": ScopedResourceProvisioningStateProvisioning,
		"succeeded":    ScopedResourceProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScopedResourceProvisioningState(input)
	return &out, nil
}
