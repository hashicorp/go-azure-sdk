package subscriptionfeatureregistrations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionFeatureRegistrationApprovalType string

const (
	SubscriptionFeatureRegistrationApprovalTypeApprovalRequired SubscriptionFeatureRegistrationApprovalType = "ApprovalRequired"
	SubscriptionFeatureRegistrationApprovalTypeAutoApproval     SubscriptionFeatureRegistrationApprovalType = "AutoApproval"
	SubscriptionFeatureRegistrationApprovalTypeNotSpecified     SubscriptionFeatureRegistrationApprovalType = "NotSpecified"
)

func PossibleValuesForSubscriptionFeatureRegistrationApprovalType() []string {
	return []string{
		string(SubscriptionFeatureRegistrationApprovalTypeApprovalRequired),
		string(SubscriptionFeatureRegistrationApprovalTypeAutoApproval),
		string(SubscriptionFeatureRegistrationApprovalTypeNotSpecified),
	}
}

func (s *SubscriptionFeatureRegistrationApprovalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionFeatureRegistrationApprovalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionFeatureRegistrationApprovalType(input string) (*SubscriptionFeatureRegistrationApprovalType, error) {
	vals := map[string]SubscriptionFeatureRegistrationApprovalType{
		"approvalrequired": SubscriptionFeatureRegistrationApprovalTypeApprovalRequired,
		"autoapproval":     SubscriptionFeatureRegistrationApprovalTypeAutoApproval,
		"notspecified":     SubscriptionFeatureRegistrationApprovalTypeNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionFeatureRegistrationApprovalType(input)
	return &out, nil
}

type SubscriptionFeatureRegistrationState string

const (
	SubscriptionFeatureRegistrationStateNotRegistered SubscriptionFeatureRegistrationState = "NotRegistered"
	SubscriptionFeatureRegistrationStateNotSpecified  SubscriptionFeatureRegistrationState = "NotSpecified"
	SubscriptionFeatureRegistrationStatePending       SubscriptionFeatureRegistrationState = "Pending"
	SubscriptionFeatureRegistrationStateRegistered    SubscriptionFeatureRegistrationState = "Registered"
	SubscriptionFeatureRegistrationStateRegistering   SubscriptionFeatureRegistrationState = "Registering"
	SubscriptionFeatureRegistrationStateUnregistered  SubscriptionFeatureRegistrationState = "Unregistered"
	SubscriptionFeatureRegistrationStateUnregistering SubscriptionFeatureRegistrationState = "Unregistering"
)

func PossibleValuesForSubscriptionFeatureRegistrationState() []string {
	return []string{
		string(SubscriptionFeatureRegistrationStateNotRegistered),
		string(SubscriptionFeatureRegistrationStateNotSpecified),
		string(SubscriptionFeatureRegistrationStatePending),
		string(SubscriptionFeatureRegistrationStateRegistered),
		string(SubscriptionFeatureRegistrationStateRegistering),
		string(SubscriptionFeatureRegistrationStateUnregistered),
		string(SubscriptionFeatureRegistrationStateUnregistering),
	}
}

func (s *SubscriptionFeatureRegistrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionFeatureRegistrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionFeatureRegistrationState(input string) (*SubscriptionFeatureRegistrationState, error) {
	vals := map[string]SubscriptionFeatureRegistrationState{
		"notregistered": SubscriptionFeatureRegistrationStateNotRegistered,
		"notspecified":  SubscriptionFeatureRegistrationStateNotSpecified,
		"pending":       SubscriptionFeatureRegistrationStatePending,
		"registered":    SubscriptionFeatureRegistrationStateRegistered,
		"registering":   SubscriptionFeatureRegistrationStateRegistering,
		"unregistered":  SubscriptionFeatureRegistrationStateUnregistered,
		"unregistering": SubscriptionFeatureRegistrationStateUnregistering,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionFeatureRegistrationState(input)
	return &out, nil
}
