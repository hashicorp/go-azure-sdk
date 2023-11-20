package applicationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGroupPolicyType string

const (
	ApplicationGroupPolicyTypeThrottlingPolicy ApplicationGroupPolicyType = "ThrottlingPolicy"
)

func PossibleValuesForApplicationGroupPolicyType() []string {
	return []string{
		string(ApplicationGroupPolicyTypeThrottlingPolicy),
	}
}

func (s *ApplicationGroupPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationGroupPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationGroupPolicyType(input string) (*ApplicationGroupPolicyType, error) {
	vals := map[string]ApplicationGroupPolicyType{
		"throttlingpolicy": ApplicationGroupPolicyTypeThrottlingPolicy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationGroupPolicyType(input)
	return &out, nil
}

type MetricId string

const (
	MetricIdIncomingBytes    MetricId = "IncomingBytes"
	MetricIdIncomingMessages MetricId = "IncomingMessages"
	MetricIdOutgoingBytes    MetricId = "OutgoingBytes"
	MetricIdOutgoingMessages MetricId = "OutgoingMessages"
)

func PossibleValuesForMetricId() []string {
	return []string{
		string(MetricIdIncomingBytes),
		string(MetricIdIncomingMessages),
		string(MetricIdOutgoingBytes),
		string(MetricIdOutgoingMessages),
	}
}

func (s *MetricId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricId(input string) (*MetricId, error) {
	vals := map[string]MetricId{
		"incomingbytes":    MetricIdIncomingBytes,
		"incomingmessages": MetricIdIncomingMessages,
		"outgoingbytes":    MetricIdOutgoingBytes,
		"outgoingmessages": MetricIdOutgoingMessages,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricId(input)
	return &out, nil
}
