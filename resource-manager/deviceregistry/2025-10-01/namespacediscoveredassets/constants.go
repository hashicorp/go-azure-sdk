package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetDestinationTarget string

const (
	DatasetDestinationTargetBrokerStateStore DatasetDestinationTarget = "BrokerStateStore"
	DatasetDestinationTargetMqtt             DatasetDestinationTarget = "Mqtt"
	DatasetDestinationTargetStorage          DatasetDestinationTarget = "Storage"
)

func PossibleValuesForDatasetDestinationTarget() []string {
	return []string{
		string(DatasetDestinationTargetBrokerStateStore),
		string(DatasetDestinationTargetMqtt),
		string(DatasetDestinationTargetStorage),
	}
}

func (s *DatasetDestinationTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatasetDestinationTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatasetDestinationTarget(input string) (*DatasetDestinationTarget, error) {
	vals := map[string]DatasetDestinationTarget{
		"brokerstatestore": DatasetDestinationTargetBrokerStateStore,
		"mqtt":             DatasetDestinationTargetMqtt,
		"storage":          DatasetDestinationTargetStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatasetDestinationTarget(input)
	return &out, nil
}

type EventDestinationTarget string

const (
	EventDestinationTargetMqtt    EventDestinationTarget = "Mqtt"
	EventDestinationTargetStorage EventDestinationTarget = "Storage"
)

func PossibleValuesForEventDestinationTarget() []string {
	return []string{
		string(EventDestinationTargetMqtt),
		string(EventDestinationTargetStorage),
	}
}

func (s *EventDestinationTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventDestinationTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventDestinationTarget(input string) (*EventDestinationTarget, error) {
	vals := map[string]EventDestinationTarget{
		"mqtt":    EventDestinationTargetMqtt,
		"storage": EventDestinationTargetStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventDestinationTarget(input)
	return &out, nil
}

type MqttDestinationQos string

const (
	MqttDestinationQosQosOne  MqttDestinationQos = "Qos1"
	MqttDestinationQosQosZero MqttDestinationQos = "Qos0"
)

func PossibleValuesForMqttDestinationQos() []string {
	return []string{
		string(MqttDestinationQosQosOne),
		string(MqttDestinationQosQosZero),
	}
}

func (s *MqttDestinationQos) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMqttDestinationQos(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMqttDestinationQos(input string) (*MqttDestinationQos, error) {
	vals := map[string]MqttDestinationQos{
		"qos1": MqttDestinationQosQosOne,
		"qos0": MqttDestinationQosQosZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MqttDestinationQos(input)
	return &out, nil
}

type NamespaceDiscoveredManagementActionType string

const (
	NamespaceDiscoveredManagementActionTypeCall  NamespaceDiscoveredManagementActionType = "Call"
	NamespaceDiscoveredManagementActionTypeRead  NamespaceDiscoveredManagementActionType = "Read"
	NamespaceDiscoveredManagementActionTypeWrite NamespaceDiscoveredManagementActionType = "Write"
)

func PossibleValuesForNamespaceDiscoveredManagementActionType() []string {
	return []string{
		string(NamespaceDiscoveredManagementActionTypeCall),
		string(NamespaceDiscoveredManagementActionTypeRead),
		string(NamespaceDiscoveredManagementActionTypeWrite),
	}
}

func (s *NamespaceDiscoveredManagementActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNamespaceDiscoveredManagementActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNamespaceDiscoveredManagementActionType(input string) (*NamespaceDiscoveredManagementActionType, error) {
	vals := map[string]NamespaceDiscoveredManagementActionType{
		"call":  NamespaceDiscoveredManagementActionTypeCall,
		"read":  NamespaceDiscoveredManagementActionTypeRead,
		"write": NamespaceDiscoveredManagementActionTypeWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NamespaceDiscoveredManagementActionType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
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
		"accepted":  ProvisioningStateAccepted,
		"canceled":  ProvisioningStateCanceled,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type StreamDestinationTarget string

const (
	StreamDestinationTargetMqtt    StreamDestinationTarget = "Mqtt"
	StreamDestinationTargetStorage StreamDestinationTarget = "Storage"
)

func PossibleValuesForStreamDestinationTarget() []string {
	return []string{
		string(StreamDestinationTargetMqtt),
		string(StreamDestinationTargetStorage),
	}
}

func (s *StreamDestinationTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStreamDestinationTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStreamDestinationTarget(input string) (*StreamDestinationTarget, error) {
	vals := map[string]StreamDestinationTarget{
		"mqtt":    StreamDestinationTargetMqtt,
		"storage": StreamDestinationTargetStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StreamDestinationTarget(input)
	return &out, nil
}

type TopicRetainType string

const (
	TopicRetainTypeKeep  TopicRetainType = "Keep"
	TopicRetainTypeNever TopicRetainType = "Never"
)

func PossibleValuesForTopicRetainType() []string {
	return []string{
		string(TopicRetainTypeKeep),
		string(TopicRetainTypeNever),
	}
}

func (s *TopicRetainType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTopicRetainType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTopicRetainType(input string) (*TopicRetainType, error) {
	vals := map[string]TopicRetainType{
		"keep":  TopicRetainTypeKeep,
		"never": TopicRetainTypeNever,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TopicRetainType(input)
	return &out, nil
}
