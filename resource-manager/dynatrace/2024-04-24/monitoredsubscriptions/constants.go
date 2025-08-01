package monitoredsubscriptions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"creating":     ProvisioningStateCreating,
		"deleted":      ProvisioningStateDeleted,
		"deleting":     ProvisioningStateDeleting,
		"failed":       ProvisioningStateFailed,
		"notspecified": ProvisioningStateNotSpecified,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type SendAadLogsStatus string

const (
	SendAadLogsStatusDisabled SendAadLogsStatus = "Disabled"
	SendAadLogsStatusEnabled  SendAadLogsStatus = "Enabled"
)

func PossibleValuesForSendAadLogsStatus() []string {
	return []string{
		string(SendAadLogsStatusDisabled),
		string(SendAadLogsStatusEnabled),
	}
}

func (s *SendAadLogsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendAadLogsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendAadLogsStatus(input string) (*SendAadLogsStatus, error) {
	vals := map[string]SendAadLogsStatus{
		"disabled": SendAadLogsStatusDisabled,
		"enabled":  SendAadLogsStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendAadLogsStatus(input)
	return &out, nil
}

type SendActivityLogsStatus string

const (
	SendActivityLogsStatusDisabled SendActivityLogsStatus = "Disabled"
	SendActivityLogsStatusEnabled  SendActivityLogsStatus = "Enabled"
)

func PossibleValuesForSendActivityLogsStatus() []string {
	return []string{
		string(SendActivityLogsStatusDisabled),
		string(SendActivityLogsStatusEnabled),
	}
}

func (s *SendActivityLogsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendActivityLogsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendActivityLogsStatus(input string) (*SendActivityLogsStatus, error) {
	vals := map[string]SendActivityLogsStatus{
		"disabled": SendActivityLogsStatusDisabled,
		"enabled":  SendActivityLogsStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendActivityLogsStatus(input)
	return &out, nil
}

type SendSubscriptionLogsStatus string

const (
	SendSubscriptionLogsStatusDisabled SendSubscriptionLogsStatus = "Disabled"
	SendSubscriptionLogsStatusEnabled  SendSubscriptionLogsStatus = "Enabled"
)

func PossibleValuesForSendSubscriptionLogsStatus() []string {
	return []string{
		string(SendSubscriptionLogsStatusDisabled),
		string(SendSubscriptionLogsStatusEnabled),
	}
}

func (s *SendSubscriptionLogsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendSubscriptionLogsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendSubscriptionLogsStatus(input string) (*SendSubscriptionLogsStatus, error) {
	vals := map[string]SendSubscriptionLogsStatus{
		"disabled": SendSubscriptionLogsStatusDisabled,
		"enabled":  SendSubscriptionLogsStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendSubscriptionLogsStatus(input)
	return &out, nil
}

type SendingMetricsStatus string

const (
	SendingMetricsStatusDisabled SendingMetricsStatus = "Disabled"
	SendingMetricsStatusEnabled  SendingMetricsStatus = "Enabled"
)

func PossibleValuesForSendingMetricsStatus() []string {
	return []string{
		string(SendingMetricsStatusDisabled),
		string(SendingMetricsStatusEnabled),
	}
}

func (s *SendingMetricsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendingMetricsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendingMetricsStatus(input string) (*SendingMetricsStatus, error) {
	vals := map[string]SendingMetricsStatus{
		"disabled": SendingMetricsStatusDisabled,
		"enabled":  SendingMetricsStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendingMetricsStatus(input)
	return &out, nil
}

type Status string

const (
	StatusActive     Status = "Active"
	StatusDeleting   Status = "Deleting"
	StatusFailed     Status = "Failed"
	StatusInProgress Status = "InProgress"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusActive),
		string(StatusDeleting),
		string(StatusFailed),
		string(StatusInProgress),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"active":     StatusActive,
		"deleting":   StatusDeleting,
		"failed":     StatusFailed,
		"inprogress": StatusInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}

type SubscriptionListOperation string

const (
	SubscriptionListOperationActive         SubscriptionListOperation = "Active"
	SubscriptionListOperationAddBegin       SubscriptionListOperation = "AddBegin"
	SubscriptionListOperationAddComplete    SubscriptionListOperation = "AddComplete"
	SubscriptionListOperationDeleteBegin    SubscriptionListOperation = "DeleteBegin"
	SubscriptionListOperationDeleteComplete SubscriptionListOperation = "DeleteComplete"
)

func PossibleValuesForSubscriptionListOperation() []string {
	return []string{
		string(SubscriptionListOperationActive),
		string(SubscriptionListOperationAddBegin),
		string(SubscriptionListOperationAddComplete),
		string(SubscriptionListOperationDeleteBegin),
		string(SubscriptionListOperationDeleteComplete),
	}
}

func (s *SubscriptionListOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionListOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionListOperation(input string) (*SubscriptionListOperation, error) {
	vals := map[string]SubscriptionListOperation{
		"active":         SubscriptionListOperationActive,
		"addbegin":       SubscriptionListOperationAddBegin,
		"addcomplete":    SubscriptionListOperationAddComplete,
		"deletebegin":    SubscriptionListOperationDeleteBegin,
		"deletecomplete": SubscriptionListOperationDeleteComplete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionListOperation(input)
	return &out, nil
}

type TagAction string

const (
	TagActionExclude TagAction = "Exclude"
	TagActionInclude TagAction = "Include"
)

func PossibleValuesForTagAction() []string {
	return []string{
		string(TagActionExclude),
		string(TagActionInclude),
	}
}

func (s *TagAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTagAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTagAction(input string) (*TagAction, error) {
	vals := map[string]TagAction{
		"exclude": TagActionExclude,
		"include": TagActionInclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TagAction(input)
	return &out, nil
}
