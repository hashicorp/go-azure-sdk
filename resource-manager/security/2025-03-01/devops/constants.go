package devops

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionableRemediationState string

const (
	ActionableRemediationStateDisabled ActionableRemediationState = "Disabled"
	ActionableRemediationStateEnabled  ActionableRemediationState = "Enabled"
	ActionableRemediationStateNone     ActionableRemediationState = "None"
)

func PossibleValuesForActionableRemediationState() []string {
	return []string{
		string(ActionableRemediationStateDisabled),
		string(ActionableRemediationStateEnabled),
		string(ActionableRemediationStateNone),
	}
}

func (s *ActionableRemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActionableRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActionableRemediationState(input string) (*ActionableRemediationState, error) {
	vals := map[string]ActionableRemediationState{
		"disabled": ActionableRemediationStateDisabled,
		"enabled":  ActionableRemediationStateEnabled,
		"none":     ActionableRemediationStateNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionableRemediationState(input)
	return &out, nil
}

type AgentlessEnablement string

const (
	AgentlessEnablementDisabled      AgentlessEnablement = "Disabled"
	AgentlessEnablementEnabled       AgentlessEnablement = "Enabled"
	AgentlessEnablementNotApplicable AgentlessEnablement = "NotApplicable"
)

func PossibleValuesForAgentlessEnablement() []string {
	return []string{
		string(AgentlessEnablementDisabled),
		string(AgentlessEnablementEnabled),
		string(AgentlessEnablementNotApplicable),
	}
}

func (s *AgentlessEnablement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentlessEnablement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentlessEnablement(input string) (*AgentlessEnablement, error) {
	vals := map[string]AgentlessEnablement{
		"disabled":      AgentlessEnablementDisabled,
		"enabled":       AgentlessEnablementEnabled,
		"notapplicable": AgentlessEnablementNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentlessEnablement(input)
	return &out, nil
}

type AnnotateDefaultBranchState string

const (
	AnnotateDefaultBranchStateDisabled AnnotateDefaultBranchState = "Disabled"
	AnnotateDefaultBranchStateEnabled  AnnotateDefaultBranchState = "Enabled"
)

func PossibleValuesForAnnotateDefaultBranchState() []string {
	return []string{
		string(AnnotateDefaultBranchStateDisabled),
		string(AnnotateDefaultBranchStateEnabled),
	}
}

func (s *AnnotateDefaultBranchState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAnnotateDefaultBranchState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAnnotateDefaultBranchState(input string) (*AnnotateDefaultBranchState, error) {
	vals := map[string]AnnotateDefaultBranchState{
		"disabled": AnnotateDefaultBranchStateDisabled,
		"enabled":  AnnotateDefaultBranchStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AnnotateDefaultBranchState(input)
	return &out, nil
}

type AutoDiscovery string

const (
	AutoDiscoveryDisabled      AutoDiscovery = "Disabled"
	AutoDiscoveryEnabled       AutoDiscovery = "Enabled"
	AutoDiscoveryNotApplicable AutoDiscovery = "NotApplicable"
)

func PossibleValuesForAutoDiscovery() []string {
	return []string{
		string(AutoDiscoveryDisabled),
		string(AutoDiscoveryEnabled),
		string(AutoDiscoveryNotApplicable),
	}
}

func (s *AutoDiscovery) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoDiscovery(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoDiscovery(input string) (*AutoDiscovery, error) {
	vals := map[string]AutoDiscovery{
		"disabled":      AutoDiscoveryDisabled,
		"enabled":       AutoDiscoveryEnabled,
		"notapplicable": AutoDiscoveryNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoDiscovery(input)
	return &out, nil
}

type DevOpsProvisioningState string

const (
	DevOpsProvisioningStateCanceled        DevOpsProvisioningState = "Canceled"
	DevOpsProvisioningStateDeletionFailure DevOpsProvisioningState = "DeletionFailure"
	DevOpsProvisioningStateDeletionSuccess DevOpsProvisioningState = "DeletionSuccess"
	DevOpsProvisioningStateFailed          DevOpsProvisioningState = "Failed"
	DevOpsProvisioningStatePending         DevOpsProvisioningState = "Pending"
	DevOpsProvisioningStatePendingDeletion DevOpsProvisioningState = "PendingDeletion"
	DevOpsProvisioningStateSucceeded       DevOpsProvisioningState = "Succeeded"
)

func PossibleValuesForDevOpsProvisioningState() []string {
	return []string{
		string(DevOpsProvisioningStateCanceled),
		string(DevOpsProvisioningStateDeletionFailure),
		string(DevOpsProvisioningStateDeletionSuccess),
		string(DevOpsProvisioningStateFailed),
		string(DevOpsProvisioningStatePending),
		string(DevOpsProvisioningStatePendingDeletion),
		string(DevOpsProvisioningStateSucceeded),
	}
}

func (s *DevOpsProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDevOpsProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDevOpsProvisioningState(input string) (*DevOpsProvisioningState, error) {
	vals := map[string]DevOpsProvisioningState{
		"canceled":        DevOpsProvisioningStateCanceled,
		"deletionfailure": DevOpsProvisioningStateDeletionFailure,
		"deletionsuccess": DevOpsProvisioningStateDeletionSuccess,
		"failed":          DevOpsProvisioningStateFailed,
		"pending":         DevOpsProvisioningStatePending,
		"pendingdeletion": DevOpsProvisioningStatePendingDeletion,
		"succeeded":       DevOpsProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DevOpsProvisioningState(input)
	return &out, nil
}

type InheritFromParentState string

const (
	InheritFromParentStateDisabled InheritFromParentState = "Disabled"
	InheritFromParentStateEnabled  InheritFromParentState = "Enabled"
)

func PossibleValuesForInheritFromParentState() []string {
	return []string{
		string(InheritFromParentStateDisabled),
		string(InheritFromParentStateEnabled),
	}
}

func (s *InheritFromParentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInheritFromParentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInheritFromParentState(input string) (*InheritFromParentState, error) {
	vals := map[string]InheritFromParentState{
		"disabled": InheritFromParentStateDisabled,
		"enabled":  InheritFromParentStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InheritFromParentState(input)
	return &out, nil
}

type InventoryKind string

const (
	InventoryKindAzureDevOpsOrganization InventoryKind = "AzureDevOpsOrganization"
	InventoryKindAzureDevOpsProject      InventoryKind = "AzureDevOpsProject"
	InventoryKindAzureDevOpsRepository   InventoryKind = "AzureDevOpsRepository"
	InventoryKindGitHubOwner             InventoryKind = "GitHubOwner"
	InventoryKindGitHubRepository        InventoryKind = "GitHubRepository"
)

func PossibleValuesForInventoryKind() []string {
	return []string{
		string(InventoryKindAzureDevOpsOrganization),
		string(InventoryKindAzureDevOpsProject),
		string(InventoryKindAzureDevOpsRepository),
		string(InventoryKindGitHubOwner),
		string(InventoryKindGitHubRepository),
	}
}

func (s *InventoryKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInventoryKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInventoryKind(input string) (*InventoryKind, error) {
	vals := map[string]InventoryKind{
		"azuredevopsorganization": InventoryKindAzureDevOpsOrganization,
		"azuredevopsproject":      InventoryKindAzureDevOpsProject,
		"azuredevopsrepository":   InventoryKindAzureDevOpsRepository,
		"githubowner":             InventoryKindGitHubOwner,
		"githubrepository":        InventoryKindGitHubRepository,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InventoryKind(input)
	return &out, nil
}

type InventoryListKind string

const (
	InventoryListKindExclusion InventoryListKind = "Exclusion"
	InventoryListKindInclusion InventoryListKind = "Inclusion"
)

func PossibleValuesForInventoryListKind() []string {
	return []string{
		string(InventoryListKindExclusion),
		string(InventoryListKindInclusion),
	}
}

func (s *InventoryListKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInventoryListKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInventoryListKind(input string) (*InventoryListKind, error) {
	vals := map[string]InventoryListKind{
		"exclusion": InventoryListKindExclusion,
		"inclusion": InventoryListKindInclusion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InventoryListKind(input)
	return &out, nil
}

type OnboardingState string

const (
	OnboardingStateNotApplicable             OnboardingState = "NotApplicable"
	OnboardingStateNotOnboarded              OnboardingState = "NotOnboarded"
	OnboardingStateOnboarded                 OnboardingState = "Onboarded"
	OnboardingStateOnboardedByOtherConnector OnboardingState = "OnboardedByOtherConnector"
)

func PossibleValuesForOnboardingState() []string {
	return []string{
		string(OnboardingStateNotApplicable),
		string(OnboardingStateNotOnboarded),
		string(OnboardingStateOnboarded),
		string(OnboardingStateOnboardedByOtherConnector),
	}
}

func (s *OnboardingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnboardingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnboardingState(input string) (*OnboardingState, error) {
	vals := map[string]OnboardingState{
		"notapplicable":             OnboardingStateNotApplicable,
		"notonboarded":              OnboardingStateNotOnboarded,
		"onboarded":                 OnboardingStateOnboarded,
		"onboardedbyotherconnector": OnboardingStateOnboardedByOtherConnector,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnboardingState(input)
	return &out, nil
}

type RuleCategory string

const (
	RuleCategoryArtifacts    RuleCategory = "Artifacts"
	RuleCategoryCode         RuleCategory = "Code"
	RuleCategoryContainers   RuleCategory = "Containers"
	RuleCategoryDependencies RuleCategory = "Dependencies"
	RuleCategoryIaC          RuleCategory = "IaC"
	RuleCategorySecrets      RuleCategory = "Secrets"
)

func PossibleValuesForRuleCategory() []string {
	return []string{
		string(RuleCategoryArtifacts),
		string(RuleCategoryCode),
		string(RuleCategoryContainers),
		string(RuleCategoryDependencies),
		string(RuleCategoryIaC),
		string(RuleCategorySecrets),
	}
}

func (s *RuleCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleCategory(input string) (*RuleCategory, error) {
	vals := map[string]RuleCategory{
		"artifacts":    RuleCategoryArtifacts,
		"code":         RuleCategoryCode,
		"containers":   RuleCategoryContainers,
		"dependencies": RuleCategoryDependencies,
		"iac":          RuleCategoryIaC,
		"secrets":      RuleCategorySecrets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleCategory(input)
	return &out, nil
}
