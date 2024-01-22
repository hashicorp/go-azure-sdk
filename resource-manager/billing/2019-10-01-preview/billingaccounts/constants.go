package billingaccounts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountType string

const (
	AccountTypeEnterprise AccountType = "Enterprise"
	AccountTypeIndividual AccountType = "Individual"
	AccountTypePartner    AccountType = "Partner"
)

func PossibleValuesForAccountType() []string {
	return []string{
		string(AccountTypeEnterprise),
		string(AccountTypeIndividual),
		string(AccountTypePartner),
	}
}

func (s *AccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccountType(input string) (*AccountType, error) {
	vals := map[string]AccountType{
		"enterprise": AccountTypeEnterprise,
		"individual": AccountTypeIndividual,
		"partner":    AccountTypePartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountType(input)
	return &out, nil
}

type AgreementType string

const (
	AgreementTypeEnterpriseAgreement            AgreementType = "EnterpriseAgreement"
	AgreementTypeMicrosoftCustomerAgreement     AgreementType = "MicrosoftCustomerAgreement"
	AgreementTypeMicrosoftOnlineServicesProgram AgreementType = "MicrosoftOnlineServicesProgram"
	AgreementTypeMicrosoftPartnerAgreement      AgreementType = "MicrosoftPartnerAgreement"
)

func PossibleValuesForAgreementType() []string {
	return []string{
		string(AgreementTypeEnterpriseAgreement),
		string(AgreementTypeMicrosoftCustomerAgreement),
		string(AgreementTypeMicrosoftOnlineServicesProgram),
		string(AgreementTypeMicrosoftPartnerAgreement),
	}
}

func (s *AgreementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgreementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgreementType(input string) (*AgreementType, error) {
	vals := map[string]AgreementType{
		"enterpriseagreement":            AgreementTypeEnterpriseAgreement,
		"microsoftcustomeragreement":     AgreementTypeMicrosoftCustomerAgreement,
		"microsoftonlineservicesprogram": AgreementTypeMicrosoftOnlineServicesProgram,
		"microsoftpartneragreement":      AgreementTypeMicrosoftPartnerAgreement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgreementType(input)
	return &out, nil
}

type BillingProfileSpendingLimit string

const (
	BillingProfileSpendingLimitOff BillingProfileSpendingLimit = "Off"
	BillingProfileSpendingLimitOn  BillingProfileSpendingLimit = "On"
)

func PossibleValuesForBillingProfileSpendingLimit() []string {
	return []string{
		string(BillingProfileSpendingLimitOff),
		string(BillingProfileSpendingLimitOn),
	}
}

func (s *BillingProfileSpendingLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileSpendingLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileSpendingLimit(input string) (*BillingProfileSpendingLimit, error) {
	vals := map[string]BillingProfileSpendingLimit{
		"off": BillingProfileSpendingLimitOff,
		"on":  BillingProfileSpendingLimitOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileSpendingLimit(input)
	return &out, nil
}

type BillingProfileStatus string

const (
	BillingProfileStatusActive   BillingProfileStatus = "Active"
	BillingProfileStatusDisabled BillingProfileStatus = "Disabled"
	BillingProfileStatusWarned   BillingProfileStatus = "Warned"
)

func PossibleValuesForBillingProfileStatus() []string {
	return []string{
		string(BillingProfileStatusActive),
		string(BillingProfileStatusDisabled),
		string(BillingProfileStatusWarned),
	}
}

func (s *BillingProfileStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatus(input string) (*BillingProfileStatus, error) {
	vals := map[string]BillingProfileStatus{
		"active":   BillingProfileStatusActive,
		"disabled": BillingProfileStatusDisabled,
		"warned":   BillingProfileStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatus(input)
	return &out, nil
}

type BillingProfileStatusReasonCode string

const (
	BillingProfileStatusReasonCodePastDue              BillingProfileStatusReasonCode = "PastDue"
	BillingProfileStatusReasonCodeSpendingLimitExpired BillingProfileStatusReasonCode = "SpendingLimitExpired"
	BillingProfileStatusReasonCodeSpendingLimitReached BillingProfileStatusReasonCode = "SpendingLimitReached"
)

func PossibleValuesForBillingProfileStatusReasonCode() []string {
	return []string{
		string(BillingProfileStatusReasonCodePastDue),
		string(BillingProfileStatusReasonCodeSpendingLimitExpired),
		string(BillingProfileStatusReasonCodeSpendingLimitReached),
	}
}

func (s *BillingProfileStatusReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatusReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatusReasonCode(input string) (*BillingProfileStatusReasonCode, error) {
	vals := map[string]BillingProfileStatusReasonCode{
		"pastdue":              BillingProfileStatusReasonCodePastDue,
		"spendinglimitexpired": BillingProfileStatusReasonCodeSpendingLimitExpired,
		"spendinglimitreached": BillingProfileStatusReasonCodeSpendingLimitReached,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatusReasonCode(input)
	return &out, nil
}

type BillingRelationshipType string

const (
	BillingRelationshipTypeCSPPartner       BillingRelationshipType = "CSPPartner"
	BillingRelationshipTypeDirect           BillingRelationshipType = "Direct"
	BillingRelationshipTypeIndirectCustomer BillingRelationshipType = "IndirectCustomer"
	BillingRelationshipTypeIndirectPartner  BillingRelationshipType = "IndirectPartner"
)

func PossibleValuesForBillingRelationshipType() []string {
	return []string{
		string(BillingRelationshipTypeCSPPartner),
		string(BillingRelationshipTypeDirect),
		string(BillingRelationshipTypeIndirectCustomer),
		string(BillingRelationshipTypeIndirectPartner),
	}
}

func (s *BillingRelationshipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingRelationshipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingRelationshipType(input string) (*BillingRelationshipType, error) {
	vals := map[string]BillingRelationshipType{
		"csppartner":       BillingRelationshipTypeCSPPartner,
		"direct":           BillingRelationshipTypeDirect,
		"indirectcustomer": BillingRelationshipTypeIndirectCustomer,
		"indirectpartner":  BillingRelationshipTypeIndirectPartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingRelationshipType(input)
	return &out, nil
}

type CustomerType string

const (
	CustomerTypeEnterprise CustomerType = "Enterprise"
	CustomerTypeIndividual CustomerType = "Individual"
	CustomerTypePartner    CustomerType = "Partner"
)

func PossibleValuesForCustomerType() []string {
	return []string{
		string(CustomerTypeEnterprise),
		string(CustomerTypeIndividual),
		string(CustomerTypePartner),
	}
}

func (s *CustomerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomerType(input string) (*CustomerType, error) {
	vals := map[string]CustomerType{
		"enterprise": CustomerTypeEnterprise,
		"individual": CustomerTypeIndividual,
		"partner":    CustomerTypePartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomerType(input)
	return &out, nil
}

type InvoiceSectionState string

const (
	InvoiceSectionStateActive     InvoiceSectionState = "Active"
	InvoiceSectionStateRestricted InvoiceSectionState = "Restricted"
)

func PossibleValuesForInvoiceSectionState() []string {
	return []string{
		string(InvoiceSectionStateActive),
		string(InvoiceSectionStateRestricted),
	}
}

func (s *InvoiceSectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceSectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceSectionState(input string) (*InvoiceSectionState, error) {
	vals := map[string]InvoiceSectionState{
		"active":     InvoiceSectionStateActive,
		"restricted": InvoiceSectionStateRestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceSectionState(input)
	return &out, nil
}

type SpendingLimit string

const (
	SpendingLimitOff SpendingLimit = "Off"
	SpendingLimitOn  SpendingLimit = "On"
)

func PossibleValuesForSpendingLimit() []string {
	return []string{
		string(SpendingLimitOff),
		string(SpendingLimitOn),
	}
}

func (s *SpendingLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimit(input string) (*SpendingLimit, error) {
	vals := map[string]SpendingLimit{
		"off": SpendingLimitOff,
		"on":  SpendingLimitOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimit(input)
	return &out, nil
}

type StatusReasonCode string

const (
	StatusReasonCodePastDue              StatusReasonCode = "PastDue"
	StatusReasonCodeSpendingLimitExpired StatusReasonCode = "SpendingLimitExpired"
	StatusReasonCodeSpendingLimitReached StatusReasonCode = "SpendingLimitReached"
)

func PossibleValuesForStatusReasonCode() []string {
	return []string{
		string(StatusReasonCodePastDue),
		string(StatusReasonCodeSpendingLimitExpired),
		string(StatusReasonCodeSpendingLimitReached),
	}
}

func (s *StatusReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusReasonCode(input string) (*StatusReasonCode, error) {
	vals := map[string]StatusReasonCode{
		"pastdue":              StatusReasonCodePastDue,
		"spendinglimitexpired": StatusReasonCodeSpendingLimitExpired,
		"spendinglimitreached": StatusReasonCodeSpendingLimitReached,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusReasonCode(input)
	return &out, nil
}

type TargetCloud string

const (
	TargetCloudUSGov TargetCloud = "USGov"
	TargetCloudUSNat TargetCloud = "USNat"
	TargetCloudUSSec TargetCloud = "USSec"
)

func PossibleValuesForTargetCloud() []string {
	return []string{
		string(TargetCloudUSGov),
		string(TargetCloudUSNat),
		string(TargetCloudUSSec),
	}
}

func (s *TargetCloud) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetCloud(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetCloud(input string) (*TargetCloud, error) {
	vals := map[string]TargetCloud{
		"usgov": TargetCloudUSGov,
		"usnat": TargetCloudUSNat,
		"ussec": TargetCloudUSSec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetCloud(input)
	return &out, nil
}
