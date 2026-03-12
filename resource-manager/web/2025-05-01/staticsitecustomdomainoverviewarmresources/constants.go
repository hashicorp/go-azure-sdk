package staticsitecustomdomainoverviewarmresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDomainStatus string

const (
	CustomDomainStatusAdding                    CustomDomainStatus = "Adding"
	CustomDomainStatusDeleting                  CustomDomainStatus = "Deleting"
	CustomDomainStatusFailed                    CustomDomainStatus = "Failed"
	CustomDomainStatusReady                     CustomDomainStatus = "Ready"
	CustomDomainStatusRetrievingValidationToken CustomDomainStatus = "RetrievingValidationToken"
	CustomDomainStatusUnhealthy                 CustomDomainStatus = "Unhealthy"
	CustomDomainStatusValidating                CustomDomainStatus = "Validating"
)

func PossibleValuesForCustomDomainStatus() []string {
	return []string{
		string(CustomDomainStatusAdding),
		string(CustomDomainStatusDeleting),
		string(CustomDomainStatusFailed),
		string(CustomDomainStatusReady),
		string(CustomDomainStatusRetrievingValidationToken),
		string(CustomDomainStatusUnhealthy),
		string(CustomDomainStatusValidating),
	}
}

func (s *CustomDomainStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomDomainStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomDomainStatus(input string) (*CustomDomainStatus, error) {
	vals := map[string]CustomDomainStatus{
		"adding":                    CustomDomainStatusAdding,
		"deleting":                  CustomDomainStatusDeleting,
		"failed":                    CustomDomainStatusFailed,
		"ready":                     CustomDomainStatusReady,
		"retrievingvalidationtoken": CustomDomainStatusRetrievingValidationToken,
		"unhealthy":                 CustomDomainStatusUnhealthy,
		"validating":                CustomDomainStatusValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomDomainStatus(input)
	return &out, nil
}
