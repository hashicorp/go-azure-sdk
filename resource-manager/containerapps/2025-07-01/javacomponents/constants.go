package javacomponents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JavaComponentProvisioningState string

const (
	JavaComponentProvisioningStateCanceled   JavaComponentProvisioningState = "Canceled"
	JavaComponentProvisioningStateDeleting   JavaComponentProvisioningState = "Deleting"
	JavaComponentProvisioningStateFailed     JavaComponentProvisioningState = "Failed"
	JavaComponentProvisioningStateInProgress JavaComponentProvisioningState = "InProgress"
	JavaComponentProvisioningStateSucceeded  JavaComponentProvisioningState = "Succeeded"
)

func PossibleValuesForJavaComponentProvisioningState() []string {
	return []string{
		string(JavaComponentProvisioningStateCanceled),
		string(JavaComponentProvisioningStateDeleting),
		string(JavaComponentProvisioningStateFailed),
		string(JavaComponentProvisioningStateInProgress),
		string(JavaComponentProvisioningStateSucceeded),
	}
}

func (s *JavaComponentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJavaComponentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJavaComponentProvisioningState(input string) (*JavaComponentProvisioningState, error) {
	vals := map[string]JavaComponentProvisioningState{
		"canceled":   JavaComponentProvisioningStateCanceled,
		"deleting":   JavaComponentProvisioningStateDeleting,
		"failed":     JavaComponentProvisioningStateFailed,
		"inprogress": JavaComponentProvisioningStateInProgress,
		"succeeded":  JavaComponentProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JavaComponentProvisioningState(input)
	return &out, nil
}

type JavaComponentType string

const (
	JavaComponentTypeSpringBootAdmin   JavaComponentType = "SpringBootAdmin"
	JavaComponentTypeSpringCloudConfig JavaComponentType = "SpringCloudConfig"
	JavaComponentTypeSpringCloudEureka JavaComponentType = "SpringCloudEureka"
)

func PossibleValuesForJavaComponentType() []string {
	return []string{
		string(JavaComponentTypeSpringBootAdmin),
		string(JavaComponentTypeSpringCloudConfig),
		string(JavaComponentTypeSpringCloudEureka),
	}
}

func (s *JavaComponentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJavaComponentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJavaComponentType(input string) (*JavaComponentType, error) {
	vals := map[string]JavaComponentType{
		"springbootadmin":   JavaComponentTypeSpringBootAdmin,
		"springcloudconfig": JavaComponentTypeSpringCloudConfig,
		"springcloudeureka": JavaComponentTypeSpringCloudEureka,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JavaComponentType(input)
	return &out, nil
}
