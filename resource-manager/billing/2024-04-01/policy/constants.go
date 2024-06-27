package policy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentAccountOwnerViewCharges string

const (
	EnrollmentAccountOwnerViewChargesAllowed    EnrollmentAccountOwnerViewCharges = "Allowed"
	EnrollmentAccountOwnerViewChargesDisabled   EnrollmentAccountOwnerViewCharges = "Disabled"
	EnrollmentAccountOwnerViewChargesNotAllowed EnrollmentAccountOwnerViewCharges = "NotAllowed"
	EnrollmentAccountOwnerViewChargesOther      EnrollmentAccountOwnerViewCharges = "Other"
)

func PossibleValuesForEnrollmentAccountOwnerViewCharges() []string {
	return []string{
		string(EnrollmentAccountOwnerViewChargesAllowed),
		string(EnrollmentAccountOwnerViewChargesDisabled),
		string(EnrollmentAccountOwnerViewChargesNotAllowed),
		string(EnrollmentAccountOwnerViewChargesOther),
	}
}

func (s *EnrollmentAccountOwnerViewCharges) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentAccountOwnerViewCharges(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentAccountOwnerViewCharges(input string) (*EnrollmentAccountOwnerViewCharges, error) {
	vals := map[string]EnrollmentAccountOwnerViewCharges{
		"allowed":    EnrollmentAccountOwnerViewChargesAllowed,
		"disabled":   EnrollmentAccountOwnerViewChargesDisabled,
		"notallowed": EnrollmentAccountOwnerViewChargesNotAllowed,
		"other":      EnrollmentAccountOwnerViewChargesOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentAccountOwnerViewCharges(input)
	return &out, nil
}

type EnrollmentAuthLevelState string

const (
	EnrollmentAuthLevelStateMicrosoftAccountOnly             EnrollmentAuthLevelState = "MicrosoftAccountOnly"
	EnrollmentAuthLevelStateMixedAccount                     EnrollmentAuthLevelState = "MixedAccount"
	EnrollmentAuthLevelStateOrganizationalAccountCrossTenant EnrollmentAuthLevelState = "OrganizationalAccountCrossTenant"
	EnrollmentAuthLevelStateOrganizationalAccountOnly        EnrollmentAuthLevelState = "OrganizationalAccountOnly"
	EnrollmentAuthLevelStateOther                            EnrollmentAuthLevelState = "Other"
)

func PossibleValuesForEnrollmentAuthLevelState() []string {
	return []string{
		string(EnrollmentAuthLevelStateMicrosoftAccountOnly),
		string(EnrollmentAuthLevelStateMixedAccount),
		string(EnrollmentAuthLevelStateOrganizationalAccountCrossTenant),
		string(EnrollmentAuthLevelStateOrganizationalAccountOnly),
		string(EnrollmentAuthLevelStateOther),
	}
}

func (s *EnrollmentAuthLevelState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentAuthLevelState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentAuthLevelState(input string) (*EnrollmentAuthLevelState, error) {
	vals := map[string]EnrollmentAuthLevelState{
		"microsoftaccountonly":             EnrollmentAuthLevelStateMicrosoftAccountOnly,
		"mixedaccount":                     EnrollmentAuthLevelStateMixedAccount,
		"organizationalaccountcrosstenant": EnrollmentAuthLevelStateOrganizationalAccountCrossTenant,
		"organizationalaccountonly":        EnrollmentAuthLevelStateOrganizationalAccountOnly,
		"other":                            EnrollmentAuthLevelStateOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentAuthLevelState(input)
	return &out, nil
}

type EnrollmentDepartmentAdminViewCharges string

const (
	EnrollmentDepartmentAdminViewChargesAllowed    EnrollmentDepartmentAdminViewCharges = "Allowed"
	EnrollmentDepartmentAdminViewChargesDisabled   EnrollmentDepartmentAdminViewCharges = "Disabled"
	EnrollmentDepartmentAdminViewChargesNotAllowed EnrollmentDepartmentAdminViewCharges = "NotAllowed"
	EnrollmentDepartmentAdminViewChargesOther      EnrollmentDepartmentAdminViewCharges = "Other"
)

func PossibleValuesForEnrollmentDepartmentAdminViewCharges() []string {
	return []string{
		string(EnrollmentDepartmentAdminViewChargesAllowed),
		string(EnrollmentDepartmentAdminViewChargesDisabled),
		string(EnrollmentDepartmentAdminViewChargesNotAllowed),
		string(EnrollmentDepartmentAdminViewChargesOther),
	}
}

func (s *EnrollmentDepartmentAdminViewCharges) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentDepartmentAdminViewCharges(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentDepartmentAdminViewCharges(input string) (*EnrollmentDepartmentAdminViewCharges, error) {
	vals := map[string]EnrollmentDepartmentAdminViewCharges{
		"allowed":    EnrollmentDepartmentAdminViewChargesAllowed,
		"disabled":   EnrollmentDepartmentAdminViewChargesDisabled,
		"notallowed": EnrollmentDepartmentAdminViewChargesNotAllowed,
		"other":      EnrollmentDepartmentAdminViewChargesOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentDepartmentAdminViewCharges(input)
	return &out, nil
}

type InvoiceSectionLabelManagementPolicy string

const (
	InvoiceSectionLabelManagementPolicyAllowed    InvoiceSectionLabelManagementPolicy = "Allowed"
	InvoiceSectionLabelManagementPolicyNotAllowed InvoiceSectionLabelManagementPolicy = "NotAllowed"
	InvoiceSectionLabelManagementPolicyOther      InvoiceSectionLabelManagementPolicy = "Other"
)

func PossibleValuesForInvoiceSectionLabelManagementPolicy() []string {
	return []string{
		string(InvoiceSectionLabelManagementPolicyAllowed),
		string(InvoiceSectionLabelManagementPolicyNotAllowed),
		string(InvoiceSectionLabelManagementPolicyOther),
	}
}

func (s *InvoiceSectionLabelManagementPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceSectionLabelManagementPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceSectionLabelManagementPolicy(input string) (*InvoiceSectionLabelManagementPolicy, error) {
	vals := map[string]InvoiceSectionLabelManagementPolicy{
		"allowed":    InvoiceSectionLabelManagementPolicyAllowed,
		"notallowed": InvoiceSectionLabelManagementPolicyNotAllowed,
		"other":      InvoiceSectionLabelManagementPolicyOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceSectionLabelManagementPolicy(input)
	return &out, nil
}

type MarketplacePurchasesPolicy string

const (
	MarketplacePurchasesPolicyAllAllowed      MarketplacePurchasesPolicy = "AllAllowed"
	MarketplacePurchasesPolicyDisabled        MarketplacePurchasesPolicy = "Disabled"
	MarketplacePurchasesPolicyNotAllowed      MarketplacePurchasesPolicy = "NotAllowed"
	MarketplacePurchasesPolicyOnlyFreeAllowed MarketplacePurchasesPolicy = "OnlyFreeAllowed"
	MarketplacePurchasesPolicyOther           MarketplacePurchasesPolicy = "Other"
)

func PossibleValuesForMarketplacePurchasesPolicy() []string {
	return []string{
		string(MarketplacePurchasesPolicyAllAllowed),
		string(MarketplacePurchasesPolicyDisabled),
		string(MarketplacePurchasesPolicyNotAllowed),
		string(MarketplacePurchasesPolicyOnlyFreeAllowed),
		string(MarketplacePurchasesPolicyOther),
	}
}

func (s *MarketplacePurchasesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMarketplacePurchasesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMarketplacePurchasesPolicy(input string) (*MarketplacePurchasesPolicy, error) {
	vals := map[string]MarketplacePurchasesPolicy{
		"allallowed":      MarketplacePurchasesPolicyAllAllowed,
		"disabled":        MarketplacePurchasesPolicyDisabled,
		"notallowed":      MarketplacePurchasesPolicyNotAllowed,
		"onlyfreeallowed": MarketplacePurchasesPolicyOnlyFreeAllowed,
		"other":           MarketplacePurchasesPolicyOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MarketplacePurchasesPolicy(input)
	return &out, nil
}

type PolicyType string

const (
	PolicyTypeOther            PolicyType = "Other"
	PolicyTypeSystemControlled PolicyType = "SystemControlled"
	PolicyTypeUserControlled   PolicyType = "UserControlled"
)

func PossibleValuesForPolicyType() []string {
	return []string{
		string(PolicyTypeOther),
		string(PolicyTypeSystemControlled),
		string(PolicyTypeUserControlled),
	}
}

func (s *PolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyType(input string) (*PolicyType, error) {
	vals := map[string]PolicyType{
		"other":            PolicyTypeOther,
		"systemcontrolled": PolicyTypeSystemControlled,
		"usercontrolled":   PolicyTypeUserControlled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNew          ProvisioningState = "New"
	ProvisioningStatePending      ProvisioningState = "Pending"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNew),
		string(ProvisioningStatePending),
		string(ProvisioningStateProvisioning),
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
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"new":          ProvisioningStateNew,
		"pending":      ProvisioningStatePending,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type ReservationPurchasesPolicy string

const (
	ReservationPurchasesPolicyAllowed    ReservationPurchasesPolicy = "Allowed"
	ReservationPurchasesPolicyDisabled   ReservationPurchasesPolicy = "Disabled"
	ReservationPurchasesPolicyNotAllowed ReservationPurchasesPolicy = "NotAllowed"
	ReservationPurchasesPolicyOther      ReservationPurchasesPolicy = "Other"
)

func PossibleValuesForReservationPurchasesPolicy() []string {
	return []string{
		string(ReservationPurchasesPolicyAllowed),
		string(ReservationPurchasesPolicyDisabled),
		string(ReservationPurchasesPolicyNotAllowed),
		string(ReservationPurchasesPolicyOther),
	}
}

func (s *ReservationPurchasesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationPurchasesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationPurchasesPolicy(input string) (*ReservationPurchasesPolicy, error) {
	vals := map[string]ReservationPurchasesPolicy{
		"allowed":    ReservationPurchasesPolicyAllowed,
		"disabled":   ReservationPurchasesPolicyDisabled,
		"notallowed": ReservationPurchasesPolicyNotAllowed,
		"other":      ReservationPurchasesPolicyOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationPurchasesPolicy(input)
	return &out, nil
}

type SavingsPlanPurchasesPolicy string

const (
	SavingsPlanPurchasesPolicyAllowed    SavingsPlanPurchasesPolicy = "Allowed"
	SavingsPlanPurchasesPolicyDisabled   SavingsPlanPurchasesPolicy = "Disabled"
	SavingsPlanPurchasesPolicyNotAllowed SavingsPlanPurchasesPolicy = "NotAllowed"
	SavingsPlanPurchasesPolicyOther      SavingsPlanPurchasesPolicy = "Other"
)

func PossibleValuesForSavingsPlanPurchasesPolicy() []string {
	return []string{
		string(SavingsPlanPurchasesPolicyAllowed),
		string(SavingsPlanPurchasesPolicyDisabled),
		string(SavingsPlanPurchasesPolicyNotAllowed),
		string(SavingsPlanPurchasesPolicyOther),
	}
}

func (s *SavingsPlanPurchasesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSavingsPlanPurchasesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSavingsPlanPurchasesPolicy(input string) (*SavingsPlanPurchasesPolicy, error) {
	vals := map[string]SavingsPlanPurchasesPolicy{
		"allowed":    SavingsPlanPurchasesPolicyAllowed,
		"disabled":   SavingsPlanPurchasesPolicyDisabled,
		"notallowed": SavingsPlanPurchasesPolicyNotAllowed,
		"other":      SavingsPlanPurchasesPolicyOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SavingsPlanPurchasesPolicy(input)
	return &out, nil
}

type ViewChargesPolicy string

const (
	ViewChargesPolicyAllowed    ViewChargesPolicy = "Allowed"
	ViewChargesPolicyNotAllowed ViewChargesPolicy = "NotAllowed"
	ViewChargesPolicyOther      ViewChargesPolicy = "Other"
)

func PossibleValuesForViewChargesPolicy() []string {
	return []string{
		string(ViewChargesPolicyAllowed),
		string(ViewChargesPolicyNotAllowed),
		string(ViewChargesPolicyOther),
	}
}

func (s *ViewChargesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseViewChargesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseViewChargesPolicy(input string) (*ViewChargesPolicy, error) {
	vals := map[string]ViewChargesPolicy{
		"allowed":    ViewChargesPolicyAllowed,
		"notallowed": ViewChargesPolicyNotAllowed,
		"other":      ViewChargesPolicyOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ViewChargesPolicy(input)
	return &out, nil
}
