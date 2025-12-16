package databoxedgedevices

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeAzureActiveDirectory AuthenticationType = "AzureActiveDirectory"
	AuthenticationTypeInvalid              AuthenticationType = "Invalid"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeAzureActiveDirectory),
		string(AuthenticationTypeInvalid),
	}
}

func (s *AuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationType(input string) (*AuthenticationType, error) {
	vals := map[string]AuthenticationType{
		"azureactivedirectory": AuthenticationTypeAzureActiveDirectory,
		"invalid":              AuthenticationTypeInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationType(input)
	return &out, nil
}

type ClusterWitnessType string

const (
	ClusterWitnessTypeCloud     ClusterWitnessType = "Cloud"
	ClusterWitnessTypeFileShare ClusterWitnessType = "FileShare"
	ClusterWitnessTypeNone      ClusterWitnessType = "None"
)

func PossibleValuesForClusterWitnessType() []string {
	return []string{
		string(ClusterWitnessTypeCloud),
		string(ClusterWitnessTypeFileShare),
		string(ClusterWitnessTypeNone),
	}
}

func (s *ClusterWitnessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterWitnessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterWitnessType(input string) (*ClusterWitnessType, error) {
	vals := map[string]ClusterWitnessType{
		"cloud":     ClusterWitnessTypeCloud,
		"fileshare": ClusterWitnessTypeFileShare,
		"none":      ClusterWitnessTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterWitnessType(input)
	return &out, nil
}

type DataBoxEdgeDeviceKind string

const (
	DataBoxEdgeDeviceKindAzureDataBoxGateway    DataBoxEdgeDeviceKind = "AzureDataBoxGateway"
	DataBoxEdgeDeviceKindAzureModularDataCentre DataBoxEdgeDeviceKind = "AzureModularDataCentre"
	DataBoxEdgeDeviceKindAzureStackEdge         DataBoxEdgeDeviceKind = "AzureStackEdge"
	DataBoxEdgeDeviceKindAzureStackHub          DataBoxEdgeDeviceKind = "AzureStackHub"
)

func PossibleValuesForDataBoxEdgeDeviceKind() []string {
	return []string{
		string(DataBoxEdgeDeviceKindAzureDataBoxGateway),
		string(DataBoxEdgeDeviceKindAzureModularDataCentre),
		string(DataBoxEdgeDeviceKindAzureStackEdge),
		string(DataBoxEdgeDeviceKindAzureStackHub),
	}
}

func (s *DataBoxEdgeDeviceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataBoxEdgeDeviceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataBoxEdgeDeviceKind(input string) (*DataBoxEdgeDeviceKind, error) {
	vals := map[string]DataBoxEdgeDeviceKind{
		"azuredataboxgateway":    DataBoxEdgeDeviceKindAzureDataBoxGateway,
		"azuremodulardatacentre": DataBoxEdgeDeviceKindAzureModularDataCentre,
		"azurestackedge":         DataBoxEdgeDeviceKindAzureStackEdge,
		"azurestackhub":          DataBoxEdgeDeviceKindAzureStackHub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataBoxEdgeDeviceKind(input)
	return &out, nil
}

type DataBoxEdgeDeviceStatus string

const (
	DataBoxEdgeDeviceStatusDisconnected          DataBoxEdgeDeviceStatus = "Disconnected"
	DataBoxEdgeDeviceStatusMaintenance           DataBoxEdgeDeviceStatus = "Maintenance"
	DataBoxEdgeDeviceStatusNeedsAttention        DataBoxEdgeDeviceStatus = "NeedsAttention"
	DataBoxEdgeDeviceStatusOffline               DataBoxEdgeDeviceStatus = "Offline"
	DataBoxEdgeDeviceStatusOnline                DataBoxEdgeDeviceStatus = "Online"
	DataBoxEdgeDeviceStatusPartiallyDisconnected DataBoxEdgeDeviceStatus = "PartiallyDisconnected"
	DataBoxEdgeDeviceStatusReadyToSetup          DataBoxEdgeDeviceStatus = "ReadyToSetup"
)

func PossibleValuesForDataBoxEdgeDeviceStatus() []string {
	return []string{
		string(DataBoxEdgeDeviceStatusDisconnected),
		string(DataBoxEdgeDeviceStatusMaintenance),
		string(DataBoxEdgeDeviceStatusNeedsAttention),
		string(DataBoxEdgeDeviceStatusOffline),
		string(DataBoxEdgeDeviceStatusOnline),
		string(DataBoxEdgeDeviceStatusPartiallyDisconnected),
		string(DataBoxEdgeDeviceStatusReadyToSetup),
	}
}

func (s *DataBoxEdgeDeviceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataBoxEdgeDeviceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataBoxEdgeDeviceStatus(input string) (*DataBoxEdgeDeviceStatus, error) {
	vals := map[string]DataBoxEdgeDeviceStatus{
		"disconnected":          DataBoxEdgeDeviceStatusDisconnected,
		"maintenance":           DataBoxEdgeDeviceStatusMaintenance,
		"needsattention":        DataBoxEdgeDeviceStatusNeedsAttention,
		"offline":               DataBoxEdgeDeviceStatusOffline,
		"online":                DataBoxEdgeDeviceStatusOnline,
		"partiallydisconnected": DataBoxEdgeDeviceStatusPartiallyDisconnected,
		"readytosetup":          DataBoxEdgeDeviceStatusReadyToSetup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataBoxEdgeDeviceStatus(input)
	return &out, nil
}

type DataResidencyType string

const (
	DataResidencyTypeGeoZoneReplication DataResidencyType = "GeoZoneReplication"
	DataResidencyTypeZoneReplication    DataResidencyType = "ZoneReplication"
)

func PossibleValuesForDataResidencyType() []string {
	return []string{
		string(DataResidencyTypeGeoZoneReplication),
		string(DataResidencyTypeZoneReplication),
	}
}

func (s *DataResidencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataResidencyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataResidencyType(input string) (*DataResidencyType, error) {
	vals := map[string]DataResidencyType{
		"geozonereplication": DataResidencyTypeGeoZoneReplication,
		"zonereplication":    DataResidencyTypeZoneReplication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataResidencyType(input)
	return &out, nil
}

type DeviceType string

const (
	DeviceTypeDataBoxEdgeDevice DeviceType = "DataBoxEdgeDevice"
)

func PossibleValuesForDeviceType() []string {
	return []string{
		string(DeviceTypeDataBoxEdgeDevice),
	}
}

func (s *DeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceType(input string) (*DeviceType, error) {
	vals := map[string]DeviceType{
		"databoxedgedevice": DeviceTypeDataBoxEdgeDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceType(input)
	return &out, nil
}

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmAESTwoFiveSix        EncryptionAlgorithm = "AES256"
	EncryptionAlgorithmNone                 EncryptionAlgorithm = "None"
	EncryptionAlgorithmRSAESPKCSOneVOneFive EncryptionAlgorithm = "RSAES_PKCS1_v_1_5"
)

func PossibleValuesForEncryptionAlgorithm() []string {
	return []string{
		string(EncryptionAlgorithmAESTwoFiveSix),
		string(EncryptionAlgorithmNone),
		string(EncryptionAlgorithmRSAESPKCSOneVOneFive),
	}
}

func (s *EncryptionAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionAlgorithm(input string) (*EncryptionAlgorithm, error) {
	vals := map[string]EncryptionAlgorithm{
		"aes256":            EncryptionAlgorithmAESTwoFiveSix,
		"none":              EncryptionAlgorithmNone,
		"rsaes_pkcs1_v_1_5": EncryptionAlgorithmRSAESPKCSOneVOneFive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionAlgorithm(input)
	return &out, nil
}

type KeyVaultSyncStatus string

const (
	KeyVaultSyncStatusKeyVaultNotConfigured KeyVaultSyncStatus = "KeyVaultNotConfigured"
	KeyVaultSyncStatusKeyVaultNotSynced     KeyVaultSyncStatus = "KeyVaultNotSynced"
	KeyVaultSyncStatusKeyVaultSyncFailed    KeyVaultSyncStatus = "KeyVaultSyncFailed"
	KeyVaultSyncStatusKeyVaultSyncPending   KeyVaultSyncStatus = "KeyVaultSyncPending"
	KeyVaultSyncStatusKeyVaultSynced        KeyVaultSyncStatus = "KeyVaultSynced"
	KeyVaultSyncStatusKeyVaultSyncing       KeyVaultSyncStatus = "KeyVaultSyncing"
)

func PossibleValuesForKeyVaultSyncStatus() []string {
	return []string{
		string(KeyVaultSyncStatusKeyVaultNotConfigured),
		string(KeyVaultSyncStatusKeyVaultNotSynced),
		string(KeyVaultSyncStatusKeyVaultSyncFailed),
		string(KeyVaultSyncStatusKeyVaultSyncPending),
		string(KeyVaultSyncStatusKeyVaultSynced),
		string(KeyVaultSyncStatusKeyVaultSyncing),
	}
}

func (s *KeyVaultSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeyVaultSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeyVaultSyncStatus(input string) (*KeyVaultSyncStatus, error) {
	vals := map[string]KeyVaultSyncStatus{
		"keyvaultnotconfigured": KeyVaultSyncStatusKeyVaultNotConfigured,
		"keyvaultnotsynced":     KeyVaultSyncStatusKeyVaultNotSynced,
		"keyvaultsyncfailed":    KeyVaultSyncStatusKeyVaultSyncFailed,
		"keyvaultsyncpending":   KeyVaultSyncStatusKeyVaultSyncPending,
		"keyvaultsynced":        KeyVaultSyncStatusKeyVaultSynced,
		"keyvaultsyncing":       KeyVaultSyncStatusKeyVaultSyncing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyVaultSyncStatus(input)
	return &out, nil
}

type MsiIdentityType string

const (
	MsiIdentityTypeNone           MsiIdentityType = "None"
	MsiIdentityTypeSystemAssigned MsiIdentityType = "SystemAssigned"
	MsiIdentityTypeUserAssigned   MsiIdentityType = "UserAssigned"
)

func PossibleValuesForMsiIdentityType() []string {
	return []string{
		string(MsiIdentityTypeNone),
		string(MsiIdentityTypeSystemAssigned),
		string(MsiIdentityTypeUserAssigned),
	}
}

func (s *MsiIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMsiIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMsiIdentityType(input string) (*MsiIdentityType, error) {
	vals := map[string]MsiIdentityType{
		"none":           MsiIdentityTypeNone,
		"systemassigned": MsiIdentityTypeSystemAssigned,
		"userassigned":   MsiIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MsiIdentityType(input)
	return &out, nil
}

type NodeStatus string

const (
	NodeStatusDown         NodeStatus = "Down"
	NodeStatusRebooting    NodeStatus = "Rebooting"
	NodeStatusShuttingDown NodeStatus = "ShuttingDown"
	NodeStatusUnknown      NodeStatus = "Unknown"
	NodeStatusUp           NodeStatus = "Up"
)

func PossibleValuesForNodeStatus() []string {
	return []string{
		string(NodeStatusDown),
		string(NodeStatusRebooting),
		string(NodeStatusShuttingDown),
		string(NodeStatusUnknown),
		string(NodeStatusUp),
	}
}

func (s *NodeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNodeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNodeStatus(input string) (*NodeStatus, error) {
	vals := map[string]NodeStatus{
		"down":         NodeStatusDown,
		"rebooting":    NodeStatusRebooting,
		"shuttingdown": NodeStatusShuttingDown,
		"unknown":      NodeStatusUnknown,
		"up":           NodeStatusUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeStatus(input)
	return &out, nil
}

type ResourceMoveStatus string

const (
	ResourceMoveStatusNone                   ResourceMoveStatus = "None"
	ResourceMoveStatusResourceMoveFailed     ResourceMoveStatus = "ResourceMoveFailed"
	ResourceMoveStatusResourceMoveInProgress ResourceMoveStatus = "ResourceMoveInProgress"
)

func PossibleValuesForResourceMoveStatus() []string {
	return []string{
		string(ResourceMoveStatusNone),
		string(ResourceMoveStatusResourceMoveFailed),
		string(ResourceMoveStatusResourceMoveInProgress),
	}
}

func (s *ResourceMoveStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceMoveStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceMoveStatus(input string) (*ResourceMoveStatus, error) {
	vals := map[string]ResourceMoveStatus{
		"none":                   ResourceMoveStatusNone,
		"resourcemovefailed":     ResourceMoveStatusResourceMoveFailed,
		"resourcemoveinprogress": ResourceMoveStatusResourceMoveInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceMoveStatus(input)
	return &out, nil
}

type RoleTypes string

const (
	RoleTypesASA                 RoleTypes = "ASA"
	RoleTypesCloudEdgeManagement RoleTypes = "CloudEdgeManagement"
	RoleTypesCognitive           RoleTypes = "Cognitive"
	RoleTypesFunctions           RoleTypes = "Functions"
	RoleTypesIOT                 RoleTypes = "IOT"
	RoleTypesKubernetes          RoleTypes = "Kubernetes"
	RoleTypesMEC                 RoleTypes = "MEC"
)

func PossibleValuesForRoleTypes() []string {
	return []string{
		string(RoleTypesASA),
		string(RoleTypesCloudEdgeManagement),
		string(RoleTypesCognitive),
		string(RoleTypesFunctions),
		string(RoleTypesIOT),
		string(RoleTypesKubernetes),
		string(RoleTypesMEC),
	}
}

func (s *RoleTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleTypes(input string) (*RoleTypes, error) {
	vals := map[string]RoleTypes{
		"asa":                 RoleTypesASA,
		"cloudedgemanagement": RoleTypesCloudEdgeManagement,
		"cognitive":           RoleTypesCognitive,
		"functions":           RoleTypesFunctions,
		"iot":                 RoleTypesIOT,
		"kubernetes":          RoleTypesKubernetes,
		"mec":                 RoleTypesMEC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleTypes(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameEPTwoOneTwoEightGPUOneMxOneW   SkuName = "EP2_128_GPU1_Mx1_W"
	SkuNameEPTwoOneTwoEightOneTFourMxOneW SkuName = "EP2_128_1T4_Mx1_W"
	SkuNameEPTwoSixFourMxOneW             SkuName = "EP2_64_Mx1_W"
	SkuNameEPTwoSixFourOneVPUW            SkuName = "EP2_64_1VPU_W"
	SkuNameEPTwoTwoFiveSixGPUTwoMxOne     SkuName = "EP2_256_GPU2_Mx1"
	SkuNameEPTwoTwoFiveSixTwoTFourW       SkuName = "EP2_256_2T4_W"
	SkuNameEdge                           SkuName = "Edge"
	SkuNameEdgeMRMini                     SkuName = "EdgeMR_Mini"
	SkuNameEdgeMRTCP                      SkuName = "EdgeMR_TCP"
	SkuNameEdgePBase                      SkuName = "EdgeP_Base"
	SkuNameEdgePHigh                      SkuName = "EdgeP_High"
	SkuNameEdgePRBase                     SkuName = "EdgePR_Base"
	SkuNameEdgePRBaseUPS                  SkuName = "EdgePR_Base_UPS"
	SkuNameGPU                            SkuName = "GPU"
	SkuNameGateway                        SkuName = "Gateway"
	SkuNameManagement                     SkuName = "Management"
	SkuNameRCALarge                       SkuName = "RCA_Large"
	SkuNameRCASmall                       SkuName = "RCA_Small"
	SkuNameRDC                            SkuName = "RDC"
	SkuNameTCALarge                       SkuName = "TCA_Large"
	SkuNameTCASmall                       SkuName = "TCA_Small"
	SkuNameTDC                            SkuName = "TDC"
	SkuNameTEAFourNodeHeater              SkuName = "TEA_4Node_Heater"
	SkuNameTEAFourNodeUPSHeater           SkuName = "TEA_4Node_UPS_Heater"
	SkuNameTEAOneNode                     SkuName = "TEA_1Node"
	SkuNameTEAOneNodeHeater               SkuName = "TEA_1Node_Heater"
	SkuNameTEAOneNodeUPS                  SkuName = "TEA_1Node_UPS"
	SkuNameTEAOneNodeUPSHeater            SkuName = "TEA_1Node_UPS_Heater"
	SkuNameTMA                            SkuName = "TMA"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameEPTwoOneTwoEightGPUOneMxOneW),
		string(SkuNameEPTwoOneTwoEightOneTFourMxOneW),
		string(SkuNameEPTwoSixFourMxOneW),
		string(SkuNameEPTwoSixFourOneVPUW),
		string(SkuNameEPTwoTwoFiveSixGPUTwoMxOne),
		string(SkuNameEPTwoTwoFiveSixTwoTFourW),
		string(SkuNameEdge),
		string(SkuNameEdgeMRMini),
		string(SkuNameEdgeMRTCP),
		string(SkuNameEdgePBase),
		string(SkuNameEdgePHigh),
		string(SkuNameEdgePRBase),
		string(SkuNameEdgePRBaseUPS),
		string(SkuNameGPU),
		string(SkuNameGateway),
		string(SkuNameManagement),
		string(SkuNameRCALarge),
		string(SkuNameRCASmall),
		string(SkuNameRDC),
		string(SkuNameTCALarge),
		string(SkuNameTCASmall),
		string(SkuNameTDC),
		string(SkuNameTEAFourNodeHeater),
		string(SkuNameTEAFourNodeUPSHeater),
		string(SkuNameTEAOneNode),
		string(SkuNameTEAOneNodeHeater),
		string(SkuNameTEAOneNodeUPS),
		string(SkuNameTEAOneNodeUPSHeater),
		string(SkuNameTMA),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"ep2_128_gpu1_mx1_w":   SkuNameEPTwoOneTwoEightGPUOneMxOneW,
		"ep2_128_1t4_mx1_w":    SkuNameEPTwoOneTwoEightOneTFourMxOneW,
		"ep2_64_mx1_w":         SkuNameEPTwoSixFourMxOneW,
		"ep2_64_1vpu_w":        SkuNameEPTwoSixFourOneVPUW,
		"ep2_256_gpu2_mx1":     SkuNameEPTwoTwoFiveSixGPUTwoMxOne,
		"ep2_256_2t4_w":        SkuNameEPTwoTwoFiveSixTwoTFourW,
		"edge":                 SkuNameEdge,
		"edgemr_mini":          SkuNameEdgeMRMini,
		"edgemr_tcp":           SkuNameEdgeMRTCP,
		"edgep_base":           SkuNameEdgePBase,
		"edgep_high":           SkuNameEdgePHigh,
		"edgepr_base":          SkuNameEdgePRBase,
		"edgepr_base_ups":      SkuNameEdgePRBaseUPS,
		"gpu":                  SkuNameGPU,
		"gateway":              SkuNameGateway,
		"management":           SkuNameManagement,
		"rca_large":            SkuNameRCALarge,
		"rca_small":            SkuNameRCASmall,
		"rdc":                  SkuNameRDC,
		"tca_large":            SkuNameTCALarge,
		"tca_small":            SkuNameTCASmall,
		"tdc":                  SkuNameTDC,
		"tea_4node_heater":     SkuNameTEAFourNodeHeater,
		"tea_4node_ups_heater": SkuNameTEAFourNodeUPSHeater,
		"tea_1node":            SkuNameTEAOneNode,
		"tea_1node_heater":     SkuNameTEAOneNodeHeater,
		"tea_1node_ups":        SkuNameTEAOneNodeUPS,
		"tea_1node_ups_heater": SkuNameTEAOneNodeUPSHeater,
		"tma":                  SkuNameTMA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierStandard),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}

type SubscriptionState string

const (
	SubscriptionStateDeleted      SubscriptionState = "Deleted"
	SubscriptionStateRegistered   SubscriptionState = "Registered"
	SubscriptionStateSuspended    SubscriptionState = "Suspended"
	SubscriptionStateUnregistered SubscriptionState = "Unregistered"
	SubscriptionStateWarned       SubscriptionState = "Warned"
)

func PossibleValuesForSubscriptionState() []string {
	return []string{
		string(SubscriptionStateDeleted),
		string(SubscriptionStateRegistered),
		string(SubscriptionStateSuspended),
		string(SubscriptionStateUnregistered),
		string(SubscriptionStateWarned),
	}
}

func (s *SubscriptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionState(input string) (*SubscriptionState, error) {
	vals := map[string]SubscriptionState{
		"deleted":      SubscriptionStateDeleted,
		"registered":   SubscriptionStateRegistered,
		"suspended":    SubscriptionStateSuspended,
		"unregistered": SubscriptionStateUnregistered,
		"warned":       SubscriptionStateWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionState(input)
	return &out, nil
}
