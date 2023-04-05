package registries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Action string

const (
	ActionAllow Action = "Allow"
)

func PossibleValuesForAction() []string {
	return []string{
		string(ActionAllow),
	}
}

func (s *Action) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAction() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Action(decoded)
	return nil
}

type ActionsRequired string

const (
	ActionsRequiredNone     ActionsRequired = "None"
	ActionsRequiredRecreate ActionsRequired = "Recreate"
)

func PossibleValuesForActionsRequired() []string {
	return []string{
		string(ActionsRequiredNone),
		string(ActionsRequiredRecreate),
	}
}

func (s *ActionsRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForActionsRequired() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ActionsRequired(decoded)
	return nil
}

type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

func PossibleValuesForConnectionStatus() []string {
	return []string{
		string(ConnectionStatusApproved),
		string(ConnectionStatusDisconnected),
		string(ConnectionStatusPending),
		string(ConnectionStatusRejected),
	}
}

func (s *ConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForConnectionStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ConnectionStatus(decoded)
	return nil
}

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

func (s *DefaultAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForDefaultAction() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = DefaultAction(decoded)
	return nil
}

type EncryptionStatus string

const (
	EncryptionStatusDisabled EncryptionStatus = "disabled"
	EncryptionStatusEnabled  EncryptionStatus = "enabled"
)

func PossibleValuesForEncryptionStatus() []string {
	return []string{
		string(EncryptionStatusDisabled),
		string(EncryptionStatusEnabled),
	}
}

func (s *EncryptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEncryptionStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EncryptionStatus(decoded)
	return nil
}

type ExportPolicyStatus string

const (
	ExportPolicyStatusDisabled ExportPolicyStatus = "disabled"
	ExportPolicyStatusEnabled  ExportPolicyStatus = "enabled"
)

func PossibleValuesForExportPolicyStatus() []string {
	return []string{
		string(ExportPolicyStatusDisabled),
		string(ExportPolicyStatusEnabled),
	}
}

func (s *ExportPolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForExportPolicyStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ExportPolicyStatus(decoded)
	return nil
}

type ImportMode string

const (
	ImportModeForce   ImportMode = "Force"
	ImportModeNoForce ImportMode = "NoForce"
)

func PossibleValuesForImportMode() []string {
	return []string{
		string(ImportModeForce),
		string(ImportModeNoForce),
	}
}

func (s *ImportMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForImportMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ImportMode(decoded)
	return nil
}

type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

func PossibleValuesForNetworkRuleBypassOptions() []string {
	return []string{
		string(NetworkRuleBypassOptionsAzureServices),
		string(NetworkRuleBypassOptionsNone),
	}
}

func (s *NetworkRuleBypassOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForNetworkRuleBypassOptions() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = NetworkRuleBypassOptions(decoded)
	return nil
}

type PasswordName string

const (
	PasswordNamePassword    PasswordName = "password"
	PasswordNamePasswordTwo PasswordName = "password2"
)

func PossibleValuesForPasswordName() []string {
	return []string{
		string(PasswordNamePassword),
		string(PasswordNamePasswordTwo),
	}
}

func (s *PasswordName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPasswordName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PasswordName(decoded)
	return nil
}

type PolicyStatus string

const (
	PolicyStatusDisabled PolicyStatus = "disabled"
	PolicyStatusEnabled  PolicyStatus = "enabled"
)

func PossibleValuesForPolicyStatus() []string {
	return []string{
		string(PolicyStatusDisabled),
		string(PolicyStatusEnabled),
	}
}

func (s *PolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPolicyStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PolicyStatus(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

func (s *PublicNetworkAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPublicNetworkAccess() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PublicNetworkAccess(decoded)
	return nil
}

type RegistryUsageUnit string

const (
	RegistryUsageUnitBytes RegistryUsageUnit = "Bytes"
	RegistryUsageUnitCount RegistryUsageUnit = "Count"
)

func PossibleValuesForRegistryUsageUnit() []string {
	return []string{
		string(RegistryUsageUnitBytes),
		string(RegistryUsageUnitCount),
	}
}

func (s *RegistryUsageUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRegistryUsageUnit() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = RegistryUsageUnit(decoded)
	return nil
}

type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNameClassic  SkuName = "Classic"
	SkuNamePremium  SkuName = "Premium"
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNameClassic),
		string(SkuNamePremium),
		string(SkuNameStandard),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSkuName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SkuName(decoded)
	return nil
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierClassic  SkuTier = "Classic"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierClassic),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSkuTier() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SkuTier(decoded)
	return nil
}

type TokenPasswordName string

const (
	TokenPasswordNamePasswordOne TokenPasswordName = "password1"
	TokenPasswordNamePasswordTwo TokenPasswordName = "password2"
)

func PossibleValuesForTokenPasswordName() []string {
	return []string{
		string(TokenPasswordNamePasswordOne),
		string(TokenPasswordNamePasswordTwo),
	}
}

func (s *TokenPasswordName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTokenPasswordName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TokenPasswordName(decoded)
	return nil
}

type TrustPolicyType string

const (
	TrustPolicyTypeNotary TrustPolicyType = "Notary"
)

func PossibleValuesForTrustPolicyType() []string {
	return []string{
		string(TrustPolicyTypeNotary),
	}
}

func (s *TrustPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTrustPolicyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TrustPolicyType(decoded)
	return nil
}

type ZoneRedundancy string

const (
	ZoneRedundancyDisabled ZoneRedundancy = "Disabled"
	ZoneRedundancyEnabled  ZoneRedundancy = "Enabled"
)

func PossibleValuesForZoneRedundancy() []string {
	return []string{
		string(ZoneRedundancyDisabled),
		string(ZoneRedundancyEnabled),
	}
}

func (s *ZoneRedundancy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForZoneRedundancy() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ZoneRedundancy(decoded)
	return nil
}
