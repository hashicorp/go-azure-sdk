package kubernetesclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvertiseToFabric string

const (
	AdvertiseToFabricFalse AdvertiseToFabric = "False"
	AdvertiseToFabricTrue  AdvertiseToFabric = "True"
)

func PossibleValuesForAdvertiseToFabric() []string {
	return []string{
		string(AdvertiseToFabricFalse),
		string(AdvertiseToFabricTrue),
	}
}

func (s *AdvertiseToFabric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvertiseToFabric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvertiseToFabric(input string) (*AdvertiseToFabric, error) {
	vals := map[string]AdvertiseToFabric{
		"false": AdvertiseToFabricFalse,
		"true":  AdvertiseToFabricTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvertiseToFabric(input)
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

type AvailabilityLifecycle string

const (
	AvailabilityLifecycleGenerallyAvailable AvailabilityLifecycle = "GenerallyAvailable"
	AvailabilityLifecyclePreview            AvailabilityLifecycle = "Preview"
)

func PossibleValuesForAvailabilityLifecycle() []string {
	return []string{
		string(AvailabilityLifecycleGenerallyAvailable),
		string(AvailabilityLifecyclePreview),
	}
}

func (s *AvailabilityLifecycle) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAvailabilityLifecycle(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAvailabilityLifecycle(input string) (*AvailabilityLifecycle, error) {
	vals := map[string]AvailabilityLifecycle{
		"generallyavailable": AvailabilityLifecycleGenerallyAvailable,
		"preview":            AvailabilityLifecyclePreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AvailabilityLifecycle(input)
	return &out, nil
}

type BfdEnabled string

const (
	BfdEnabledFalse BfdEnabled = "False"
	BfdEnabledTrue  BfdEnabled = "True"
)

func PossibleValuesForBfdEnabled() []string {
	return []string{
		string(BfdEnabledFalse),
		string(BfdEnabledTrue),
	}
}

func (s *BfdEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBfdEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBfdEnabled(input string) (*BfdEnabled, error) {
	vals := map[string]BfdEnabled{
		"false": BfdEnabledFalse,
		"true":  BfdEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BfdEnabled(input)
	return &out, nil
}

type BgpMultiHop string

const (
	BgpMultiHopFalse BgpMultiHop = "False"
	BgpMultiHopTrue  BgpMultiHop = "True"
)

func PossibleValuesForBgpMultiHop() []string {
	return []string{
		string(BgpMultiHopFalse),
		string(BgpMultiHopTrue),
	}
}

func (s *BgpMultiHop) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBgpMultiHop(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBgpMultiHop(input string) (*BgpMultiHop, error) {
	vals := map[string]BgpMultiHop{
		"false": BgpMultiHopFalse,
		"true":  BgpMultiHopTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BgpMultiHop(input)
	return &out, nil
}

type DefaultGateway string

const (
	DefaultGatewayFalse DefaultGateway = "False"
	DefaultGatewayTrue  DefaultGateway = "True"
)

func PossibleValuesForDefaultGateway() []string {
	return []string{
		string(DefaultGatewayFalse),
		string(DefaultGatewayTrue),
	}
}

func (s *DefaultGateway) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultGateway(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultGateway(input string) (*DefaultGateway, error) {
	vals := map[string]DefaultGateway{
		"false": DefaultGatewayFalse,
		"true":  DefaultGatewayTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultGateway(input)
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

type FabricPeeringEnabled string

const (
	FabricPeeringEnabledFalse FabricPeeringEnabled = "False"
	FabricPeeringEnabledTrue  FabricPeeringEnabled = "True"
)

func PossibleValuesForFabricPeeringEnabled() []string {
	return []string{
		string(FabricPeeringEnabledFalse),
		string(FabricPeeringEnabledTrue),
	}
}

func (s *FabricPeeringEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFabricPeeringEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFabricPeeringEnabled(input string) (*FabricPeeringEnabled, error) {
	vals := map[string]FabricPeeringEnabled{
		"false": FabricPeeringEnabledFalse,
		"true":  FabricPeeringEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FabricPeeringEnabled(input)
	return &out, nil
}

type FeatureDetailedStatus string

const (
	FeatureDetailedStatusFailed  FeatureDetailedStatus = "Failed"
	FeatureDetailedStatusRunning FeatureDetailedStatus = "Running"
	FeatureDetailedStatusUnknown FeatureDetailedStatus = "Unknown"
)

func PossibleValuesForFeatureDetailedStatus() []string {
	return []string{
		string(FeatureDetailedStatusFailed),
		string(FeatureDetailedStatusRunning),
		string(FeatureDetailedStatusUnknown),
	}
}

func (s *FeatureDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureDetailedStatus(input string) (*FeatureDetailedStatus, error) {
	vals := map[string]FeatureDetailedStatus{
		"failed":  FeatureDetailedStatusFailed,
		"running": FeatureDetailedStatusRunning,
		"unknown": FeatureDetailedStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureDetailedStatus(input)
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

type KubernetesClusterDetailedStatus string

const (
	KubernetesClusterDetailedStatusAvailable    KubernetesClusterDetailedStatus = "Available"
	KubernetesClusterDetailedStatusError        KubernetesClusterDetailedStatus = "Error"
	KubernetesClusterDetailedStatusProvisioning KubernetesClusterDetailedStatus = "Provisioning"
)

func PossibleValuesForKubernetesClusterDetailedStatus() []string {
	return []string{
		string(KubernetesClusterDetailedStatusAvailable),
		string(KubernetesClusterDetailedStatusError),
		string(KubernetesClusterDetailedStatusProvisioning),
	}
}

func (s *KubernetesClusterDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterDetailedStatus(input string) (*KubernetesClusterDetailedStatus, error) {
	vals := map[string]KubernetesClusterDetailedStatus{
		"available":    KubernetesClusterDetailedStatusAvailable,
		"error":        KubernetesClusterDetailedStatusError,
		"provisioning": KubernetesClusterDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterDetailedStatus(input)
	return &out, nil
}

type KubernetesClusterNodeDetailedStatus string

const (
	KubernetesClusterNodeDetailedStatusAvailable    KubernetesClusterNodeDetailedStatus = "Available"
	KubernetesClusterNodeDetailedStatusError        KubernetesClusterNodeDetailedStatus = "Error"
	KubernetesClusterNodeDetailedStatusProvisioning KubernetesClusterNodeDetailedStatus = "Provisioning"
	KubernetesClusterNodeDetailedStatusRunning      KubernetesClusterNodeDetailedStatus = "Running"
	KubernetesClusterNodeDetailedStatusScheduling   KubernetesClusterNodeDetailedStatus = "Scheduling"
	KubernetesClusterNodeDetailedStatusStopped      KubernetesClusterNodeDetailedStatus = "Stopped"
	KubernetesClusterNodeDetailedStatusTerminating  KubernetesClusterNodeDetailedStatus = "Terminating"
	KubernetesClusterNodeDetailedStatusUnknown      KubernetesClusterNodeDetailedStatus = "Unknown"
)

func PossibleValuesForKubernetesClusterNodeDetailedStatus() []string {
	return []string{
		string(KubernetesClusterNodeDetailedStatusAvailable),
		string(KubernetesClusterNodeDetailedStatusError),
		string(KubernetesClusterNodeDetailedStatusProvisioning),
		string(KubernetesClusterNodeDetailedStatusRunning),
		string(KubernetesClusterNodeDetailedStatusScheduling),
		string(KubernetesClusterNodeDetailedStatusStopped),
		string(KubernetesClusterNodeDetailedStatusTerminating),
		string(KubernetesClusterNodeDetailedStatusUnknown),
	}
}

func (s *KubernetesClusterNodeDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterNodeDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterNodeDetailedStatus(input string) (*KubernetesClusterNodeDetailedStatus, error) {
	vals := map[string]KubernetesClusterNodeDetailedStatus{
		"available":    KubernetesClusterNodeDetailedStatusAvailable,
		"error":        KubernetesClusterNodeDetailedStatusError,
		"provisioning": KubernetesClusterNodeDetailedStatusProvisioning,
		"running":      KubernetesClusterNodeDetailedStatusRunning,
		"scheduling":   KubernetesClusterNodeDetailedStatusScheduling,
		"stopped":      KubernetesClusterNodeDetailedStatusStopped,
		"terminating":  KubernetesClusterNodeDetailedStatusTerminating,
		"unknown":      KubernetesClusterNodeDetailedStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterNodeDetailedStatus(input)
	return &out, nil
}

type KubernetesClusterProvisioningState string

const (
	KubernetesClusterProvisioningStateAccepted   KubernetesClusterProvisioningState = "Accepted"
	KubernetesClusterProvisioningStateCanceled   KubernetesClusterProvisioningState = "Canceled"
	KubernetesClusterProvisioningStateCreated    KubernetesClusterProvisioningState = "Created"
	KubernetesClusterProvisioningStateDeleting   KubernetesClusterProvisioningState = "Deleting"
	KubernetesClusterProvisioningStateFailed     KubernetesClusterProvisioningState = "Failed"
	KubernetesClusterProvisioningStateInProgress KubernetesClusterProvisioningState = "InProgress"
	KubernetesClusterProvisioningStateSucceeded  KubernetesClusterProvisioningState = "Succeeded"
	KubernetesClusterProvisioningStateUpdating   KubernetesClusterProvisioningState = "Updating"
)

func PossibleValuesForKubernetesClusterProvisioningState() []string {
	return []string{
		string(KubernetesClusterProvisioningStateAccepted),
		string(KubernetesClusterProvisioningStateCanceled),
		string(KubernetesClusterProvisioningStateCreated),
		string(KubernetesClusterProvisioningStateDeleting),
		string(KubernetesClusterProvisioningStateFailed),
		string(KubernetesClusterProvisioningStateInProgress),
		string(KubernetesClusterProvisioningStateSucceeded),
		string(KubernetesClusterProvisioningStateUpdating),
	}
}

func (s *KubernetesClusterProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterProvisioningState(input string) (*KubernetesClusterProvisioningState, error) {
	vals := map[string]KubernetesClusterProvisioningState{
		"accepted":   KubernetesClusterProvisioningStateAccepted,
		"canceled":   KubernetesClusterProvisioningStateCanceled,
		"created":    KubernetesClusterProvisioningStateCreated,
		"deleting":   KubernetesClusterProvisioningStateDeleting,
		"failed":     KubernetesClusterProvisioningStateFailed,
		"inprogress": KubernetesClusterProvisioningStateInProgress,
		"succeeded":  KubernetesClusterProvisioningStateSucceeded,
		"updating":   KubernetesClusterProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterProvisioningState(input)
	return &out, nil
}

type KubernetesNodePowerState string

const (
	KubernetesNodePowerStateOff     KubernetesNodePowerState = "Off"
	KubernetesNodePowerStateOn      KubernetesNodePowerState = "On"
	KubernetesNodePowerStateUnknown KubernetesNodePowerState = "Unknown"
)

func PossibleValuesForKubernetesNodePowerState() []string {
	return []string{
		string(KubernetesNodePowerStateOff),
		string(KubernetesNodePowerStateOn),
		string(KubernetesNodePowerStateUnknown),
	}
}

func (s *KubernetesNodePowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesNodePowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesNodePowerState(input string) (*KubernetesNodePowerState, error) {
	vals := map[string]KubernetesNodePowerState{
		"off":     KubernetesNodePowerStateOff,
		"on":      KubernetesNodePowerStateOn,
		"unknown": KubernetesNodePowerStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesNodePowerState(input)
	return &out, nil
}

type KubernetesNodeRole string

const (
	KubernetesNodeRoleControlPlane KubernetesNodeRole = "ControlPlane"
	KubernetesNodeRoleWorker       KubernetesNodeRole = "Worker"
)

func PossibleValuesForKubernetesNodeRole() []string {
	return []string{
		string(KubernetesNodeRoleControlPlane),
		string(KubernetesNodeRoleWorker),
	}
}

func (s *KubernetesNodeRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesNodeRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesNodeRole(input string) (*KubernetesNodeRole, error) {
	vals := map[string]KubernetesNodeRole{
		"controlplane": KubernetesNodeRoleControlPlane,
		"worker":       KubernetesNodeRoleWorker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesNodeRole(input)
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

type VirtualMachineIPAllocationMethod string

const (
	VirtualMachineIPAllocationMethodDisabled VirtualMachineIPAllocationMethod = "Disabled"
	VirtualMachineIPAllocationMethodDynamic  VirtualMachineIPAllocationMethod = "Dynamic"
	VirtualMachineIPAllocationMethodStatic   VirtualMachineIPAllocationMethod = "Static"
)

func PossibleValuesForVirtualMachineIPAllocationMethod() []string {
	return []string{
		string(VirtualMachineIPAllocationMethodDisabled),
		string(VirtualMachineIPAllocationMethodDynamic),
		string(VirtualMachineIPAllocationMethodStatic),
	}
}

func (s *VirtualMachineIPAllocationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineIPAllocationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineIPAllocationMethod(input string) (*VirtualMachineIPAllocationMethod, error) {
	vals := map[string]VirtualMachineIPAllocationMethod{
		"disabled": VirtualMachineIPAllocationMethodDisabled,
		"dynamic":  VirtualMachineIPAllocationMethodDynamic,
		"static":   VirtualMachineIPAllocationMethodStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineIPAllocationMethod(input)
	return &out, nil
}
