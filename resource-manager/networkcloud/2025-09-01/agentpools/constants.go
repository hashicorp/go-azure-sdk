package agentpools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolDetailedStatus string

const (
	AgentPoolDetailedStatusAvailable    AgentPoolDetailedStatus = "Available"
	AgentPoolDetailedStatusError        AgentPoolDetailedStatus = "Error"
	AgentPoolDetailedStatusProvisioning AgentPoolDetailedStatus = "Provisioning"
)

func PossibleValuesForAgentPoolDetailedStatus() []string {
	return []string{
		string(AgentPoolDetailedStatusAvailable),
		string(AgentPoolDetailedStatusError),
		string(AgentPoolDetailedStatusProvisioning),
	}
}

func (s *AgentPoolDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentPoolDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentPoolDetailedStatus(input string) (*AgentPoolDetailedStatus, error) {
	vals := map[string]AgentPoolDetailedStatus{
		"available":    AgentPoolDetailedStatusAvailable,
		"error":        AgentPoolDetailedStatusError,
		"provisioning": AgentPoolDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolDetailedStatus(input)
	return &out, nil
}

type AgentPoolMode string

const (
	AgentPoolModeNotApplicable AgentPoolMode = "NotApplicable"
	AgentPoolModeSystem        AgentPoolMode = "System"
	AgentPoolModeUser          AgentPoolMode = "User"
)

func PossibleValuesForAgentPoolMode() []string {
	return []string{
		string(AgentPoolModeNotApplicable),
		string(AgentPoolModeSystem),
		string(AgentPoolModeUser),
	}
}

func (s *AgentPoolMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentPoolMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentPoolMode(input string) (*AgentPoolMode, error) {
	vals := map[string]AgentPoolMode{
		"notapplicable": AgentPoolModeNotApplicable,
		"system":        AgentPoolModeSystem,
		"user":          AgentPoolModeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolMode(input)
	return &out, nil
}

type AgentPoolProvisioningState string

const (
	AgentPoolProvisioningStateAccepted   AgentPoolProvisioningState = "Accepted"
	AgentPoolProvisioningStateCanceled   AgentPoolProvisioningState = "Canceled"
	AgentPoolProvisioningStateDeleting   AgentPoolProvisioningState = "Deleting"
	AgentPoolProvisioningStateFailed     AgentPoolProvisioningState = "Failed"
	AgentPoolProvisioningStateInProgress AgentPoolProvisioningState = "InProgress"
	AgentPoolProvisioningStateSucceeded  AgentPoolProvisioningState = "Succeeded"
	AgentPoolProvisioningStateUpdating   AgentPoolProvisioningState = "Updating"
)

func PossibleValuesForAgentPoolProvisioningState() []string {
	return []string{
		string(AgentPoolProvisioningStateAccepted),
		string(AgentPoolProvisioningStateCanceled),
		string(AgentPoolProvisioningStateDeleting),
		string(AgentPoolProvisioningStateFailed),
		string(AgentPoolProvisioningStateInProgress),
		string(AgentPoolProvisioningStateSucceeded),
		string(AgentPoolProvisioningStateUpdating),
	}
}

func (s *AgentPoolProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentPoolProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentPoolProvisioningState(input string) (*AgentPoolProvisioningState, error) {
	vals := map[string]AgentPoolProvisioningState{
		"accepted":   AgentPoolProvisioningStateAccepted,
		"canceled":   AgentPoolProvisioningStateCanceled,
		"deleting":   AgentPoolProvisioningStateDeleting,
		"failed":     AgentPoolProvisioningStateFailed,
		"inprogress": AgentPoolProvisioningStateInProgress,
		"succeeded":  AgentPoolProvisioningStateSucceeded,
		"updating":   AgentPoolProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolProvisioningState(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
	ExtendedLocationTypeEdgeZone       ExtendedLocationType = "EdgeZone"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
		string(ExtendedLocationTypeEdgeZone),
	}
}

func (s *ExtendedLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationType(input string) (*ExtendedLocationType, error) {
	vals := map[string]ExtendedLocationType{
		"customlocation": ExtendedLocationTypeCustomLocation,
		"edgezone":       ExtendedLocationTypeEdgeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}

type HugepagesSize string

const (
	HugepagesSizeOneG HugepagesSize = "1G"
	HugepagesSizeTwoM HugepagesSize = "2M"
)

func PossibleValuesForHugepagesSize() []string {
	return []string{
		string(HugepagesSizeOneG),
		string(HugepagesSizeTwoM),
	}
}

func (s *HugepagesSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHugepagesSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHugepagesSize(input string) (*HugepagesSize, error) {
	vals := map[string]HugepagesSize{
		"1g": HugepagesSizeOneG,
		"2m": HugepagesSizeTwoM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HugepagesSize(input)
	return &out, nil
}

type KubernetesPluginType string

const (
	KubernetesPluginTypeDPDK     KubernetesPluginType = "DPDK"
	KubernetesPluginTypeIPVLAN   KubernetesPluginType = "IPVLAN"
	KubernetesPluginTypeMACVLAN  KubernetesPluginType = "MACVLAN"
	KubernetesPluginTypeOSDevice KubernetesPluginType = "OSDevice"
	KubernetesPluginTypeSRIOV    KubernetesPluginType = "SRIOV"
)

func PossibleValuesForKubernetesPluginType() []string {
	return []string{
		string(KubernetesPluginTypeDPDK),
		string(KubernetesPluginTypeIPVLAN),
		string(KubernetesPluginTypeMACVLAN),
		string(KubernetesPluginTypeOSDevice),
		string(KubernetesPluginTypeSRIOV),
	}
}

func (s *KubernetesPluginType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesPluginType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesPluginType(input string) (*KubernetesPluginType, error) {
	vals := map[string]KubernetesPluginType{
		"dpdk":     KubernetesPluginTypeDPDK,
		"ipvlan":   KubernetesPluginTypeIPVLAN,
		"macvlan":  KubernetesPluginTypeMACVLAN,
		"osdevice": KubernetesPluginTypeOSDevice,
		"sriov":    KubernetesPluginTypeSRIOV,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesPluginType(input)
	return &out, nil
}

type L3NetworkConfigurationIPamEnabled string

const (
	L3NetworkConfigurationIPamEnabledFalse L3NetworkConfigurationIPamEnabled = "False"
	L3NetworkConfigurationIPamEnabledTrue  L3NetworkConfigurationIPamEnabled = "True"
)

func PossibleValuesForL3NetworkConfigurationIPamEnabled() []string {
	return []string{
		string(L3NetworkConfigurationIPamEnabledFalse),
		string(L3NetworkConfigurationIPamEnabledTrue),
	}
}

func (s *L3NetworkConfigurationIPamEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL3NetworkConfigurationIPamEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL3NetworkConfigurationIPamEnabled(input string) (*L3NetworkConfigurationIPamEnabled, error) {
	vals := map[string]L3NetworkConfigurationIPamEnabled{
		"false": L3NetworkConfigurationIPamEnabledFalse,
		"true":  L3NetworkConfigurationIPamEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L3NetworkConfigurationIPamEnabled(input)
	return &out, nil
}
