package billingaccount

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountStatus string

const (
	AccountStatusActive      AccountStatus = "Active"
	AccountStatusDeleted     AccountStatus = "Deleted"
	AccountStatusDisabled    AccountStatus = "Disabled"
	AccountStatusExpired     AccountStatus = "Expired"
	AccountStatusExtended    AccountStatus = "Extended"
	AccountStatusNew         AccountStatus = "New"
	AccountStatusOther       AccountStatus = "Other"
	AccountStatusPending     AccountStatus = "Pending"
	AccountStatusTerminated  AccountStatus = "Terminated"
	AccountStatusTransferred AccountStatus = "Transferred"
	AccountStatusUnderReview AccountStatus = "UnderReview"
)

func PossibleValuesForAccountStatus() []string {
	return []string{
		string(AccountStatusActive),
		string(AccountStatusDeleted),
		string(AccountStatusDisabled),
		string(AccountStatusExpired),
		string(AccountStatusExtended),
		string(AccountStatusNew),
		string(AccountStatusOther),
		string(AccountStatusPending),
		string(AccountStatusTerminated),
		string(AccountStatusTransferred),
		string(AccountStatusUnderReview),
	}
}

func (s *AccountStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccountStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccountStatus(input string) (*AccountStatus, error) {
	vals := map[string]AccountStatus{
		"active":      AccountStatusActive,
		"deleted":     AccountStatusDeleted,
		"disabled":    AccountStatusDisabled,
		"expired":     AccountStatusExpired,
		"extended":    AccountStatusExtended,
		"new":         AccountStatusNew,
		"other":       AccountStatusOther,
		"pending":     AccountStatusPending,
		"terminated":  AccountStatusTerminated,
		"transferred": AccountStatusTransferred,
		"underreview": AccountStatusUnderReview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountStatus(input)
	return &out, nil
}

type AccountSubType string

const (
	AccountSubTypeEnterprise   AccountSubType = "Enterprise"
	AccountSubTypeIndividual   AccountSubType = "Individual"
	AccountSubTypeNone         AccountSubType = "None"
	AccountSubTypeOther        AccountSubType = "Other"
	AccountSubTypeProfessional AccountSubType = "Professional"
)

func PossibleValuesForAccountSubType() []string {
	return []string{
		string(AccountSubTypeEnterprise),
		string(AccountSubTypeIndividual),
		string(AccountSubTypeNone),
		string(AccountSubTypeOther),
		string(AccountSubTypeProfessional),
	}
}

func (s *AccountSubType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccountSubType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccountSubType(input string) (*AccountSubType, error) {
	vals := map[string]AccountSubType{
		"enterprise":   AccountSubTypeEnterprise,
		"individual":   AccountSubTypeIndividual,
		"none":         AccountSubTypeNone,
		"other":        AccountSubTypeOther,
		"professional": AccountSubTypeProfessional,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountSubType(input)
	return &out, nil
}

type AccountType string

const (
	AccountTypeBusiness       AccountType = "Business"
	AccountTypeClassicPartner AccountType = "ClassicPartner"
	AccountTypeEnterprise     AccountType = "Enterprise"
	AccountTypeIndividual     AccountType = "Individual"
	AccountTypeInternal       AccountType = "Internal"
	AccountTypeOther          AccountType = "Other"
	AccountTypePartner        AccountType = "Partner"
	AccountTypeReseller       AccountType = "Reseller"
	AccountTypeTenant         AccountType = "Tenant"
)

func PossibleValuesForAccountType() []string {
	return []string{
		string(AccountTypeBusiness),
		string(AccountTypeClassicPartner),
		string(AccountTypeEnterprise),
		string(AccountTypeIndividual),
		string(AccountTypeInternal),
		string(AccountTypeOther),
		string(AccountTypePartner),
		string(AccountTypeReseller),
		string(AccountTypeTenant),
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
		"business":       AccountTypeBusiness,
		"classicpartner": AccountTypeClassicPartner,
		"enterprise":     AccountTypeEnterprise,
		"individual":     AccountTypeIndividual,
		"internal":       AccountTypeInternal,
		"other":          AccountTypeOther,
		"partner":        AccountTypePartner,
		"reseller":       AccountTypeReseller,
		"tenant":         AccountTypeTenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountType(input)
	return &out, nil
}

type AddressValidationStatus string

const (
	AddressValidationStatusInvalid AddressValidationStatus = "Invalid"
	AddressValidationStatusOther   AddressValidationStatus = "Other"
	AddressValidationStatusValid   AddressValidationStatus = "Valid"
)

func PossibleValuesForAddressValidationStatus() []string {
	return []string{
		string(AddressValidationStatusInvalid),
		string(AddressValidationStatusOther),
		string(AddressValidationStatusValid),
	}
}

func (s *AddressValidationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddressValidationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddressValidationStatus(input string) (*AddressValidationStatus, error) {
	vals := map[string]AddressValidationStatus{
		"invalid": AddressValidationStatusInvalid,
		"other":   AddressValidationStatusOther,
		"valid":   AddressValidationStatusValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddressValidationStatus(input)
	return &out, nil
}

type AgreementType string

const (
	AgreementTypeEnterpriseAgreement            AgreementType = "EnterpriseAgreement"
	AgreementTypeMicrosoftCustomerAgreement     AgreementType = "MicrosoftCustomerAgreement"
	AgreementTypeMicrosoftOnlineServicesProgram AgreementType = "MicrosoftOnlineServicesProgram"
	AgreementTypeMicrosoftPartnerAgreement      AgreementType = "MicrosoftPartnerAgreement"
	AgreementTypeOther                          AgreementType = "Other"
)

func PossibleValuesForAgreementType() []string {
	return []string{
		string(AgreementTypeEnterpriseAgreement),
		string(AgreementTypeMicrosoftCustomerAgreement),
		string(AgreementTypeMicrosoftOnlineServicesProgram),
		string(AgreementTypeMicrosoftPartnerAgreement),
		string(AgreementTypeOther),
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
		"other":                          AgreementTypeOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgreementType(input)
	return &out, nil
}

type BillingAccountStatusReasonCode string

const (
	BillingAccountStatusReasonCodeExpired             BillingAccountStatusReasonCode = "Expired"
	BillingAccountStatusReasonCodeManuallyTerminated  BillingAccountStatusReasonCode = "ManuallyTerminated"
	BillingAccountStatusReasonCodeOther               BillingAccountStatusReasonCode = "Other"
	BillingAccountStatusReasonCodeTerminateProcessing BillingAccountStatusReasonCode = "TerminateProcessing"
	BillingAccountStatusReasonCodeTransferred         BillingAccountStatusReasonCode = "Transferred"
	BillingAccountStatusReasonCodeUnusualActivity     BillingAccountStatusReasonCode = "UnusualActivity"
)

func PossibleValuesForBillingAccountStatusReasonCode() []string {
	return []string{
		string(BillingAccountStatusReasonCodeExpired),
		string(BillingAccountStatusReasonCodeManuallyTerminated),
		string(BillingAccountStatusReasonCodeOther),
		string(BillingAccountStatusReasonCodeTerminateProcessing),
		string(BillingAccountStatusReasonCodeTransferred),
		string(BillingAccountStatusReasonCodeUnusualActivity),
	}
}

func (s *BillingAccountStatusReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingAccountStatusReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingAccountStatusReasonCode(input string) (*BillingAccountStatusReasonCode, error) {
	vals := map[string]BillingAccountStatusReasonCode{
		"expired":             BillingAccountStatusReasonCodeExpired,
		"manuallyterminated":  BillingAccountStatusReasonCodeManuallyTerminated,
		"other":               BillingAccountStatusReasonCodeOther,
		"terminateprocessing": BillingAccountStatusReasonCodeTerminateProcessing,
		"transferred":         BillingAccountStatusReasonCodeTransferred,
		"unusualactivity":     BillingAccountStatusReasonCodeUnusualActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingAccountStatusReasonCode(input)
	return &out, nil
}

type BillingProfileStatus string

const (
	BillingProfileStatusActive      BillingProfileStatus = "Active"
	BillingProfileStatusDeleted     BillingProfileStatus = "Deleted"
	BillingProfileStatusDisabled    BillingProfileStatus = "Disabled"
	BillingProfileStatusOther       BillingProfileStatus = "Other"
	BillingProfileStatusUnderReview BillingProfileStatus = "UnderReview"
	BillingProfileStatusWarned      BillingProfileStatus = "Warned"
)

func PossibleValuesForBillingProfileStatus() []string {
	return []string{
		string(BillingProfileStatusActive),
		string(BillingProfileStatusDeleted),
		string(BillingProfileStatusDisabled),
		string(BillingProfileStatusOther),
		string(BillingProfileStatusUnderReview),
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
		"active":      BillingProfileStatusActive,
		"deleted":     BillingProfileStatusDeleted,
		"disabled":    BillingProfileStatusDisabled,
		"other":       BillingProfileStatusOther,
		"underreview": BillingProfileStatusUnderReview,
		"warned":      BillingProfileStatusWarned,
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
	BillingProfileStatusReasonCodeOther                BillingProfileStatusReasonCode = "Other"
	BillingProfileStatusReasonCodePastDue              BillingProfileStatusReasonCode = "PastDue"
	BillingProfileStatusReasonCodeSpendingLimitExpired BillingProfileStatusReasonCode = "SpendingLimitExpired"
	BillingProfileStatusReasonCodeSpendingLimitReached BillingProfileStatusReasonCode = "SpendingLimitReached"
	BillingProfileStatusReasonCodeUnusualActivity      BillingProfileStatusReasonCode = "UnusualActivity"
)

func PossibleValuesForBillingProfileStatusReasonCode() []string {
	return []string{
		string(BillingProfileStatusReasonCodeOther),
		string(BillingProfileStatusReasonCodePastDue),
		string(BillingProfileStatusReasonCodeSpendingLimitExpired),
		string(BillingProfileStatusReasonCodeSpendingLimitReached),
		string(BillingProfileStatusReasonCodeUnusualActivity),
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
		"other":                BillingProfileStatusReasonCodeOther,
		"pastdue":              BillingProfileStatusReasonCodePastDue,
		"spendinglimitexpired": BillingProfileStatusReasonCodeSpendingLimitExpired,
		"spendinglimitreached": BillingProfileStatusReasonCodeSpendingLimitReached,
		"unusualactivity":      BillingProfileStatusReasonCodeUnusualActivity,
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
	BillingRelationshipTypeCSPCustomer      BillingRelationshipType = "CSPCustomer"
	BillingRelationshipTypeCSPPartner       BillingRelationshipType = "CSPPartner"
	BillingRelationshipTypeDirect           BillingRelationshipType = "Direct"
	BillingRelationshipTypeIndirectCustomer BillingRelationshipType = "IndirectCustomer"
	BillingRelationshipTypeIndirectPartner  BillingRelationshipType = "IndirectPartner"
	BillingRelationshipTypeOther            BillingRelationshipType = "Other"
)

func PossibleValuesForBillingRelationshipType() []string {
	return []string{
		string(BillingRelationshipTypeCSPCustomer),
		string(BillingRelationshipTypeCSPPartner),
		string(BillingRelationshipTypeDirect),
		string(BillingRelationshipTypeIndirectCustomer),
		string(BillingRelationshipTypeIndirectPartner),
		string(BillingRelationshipTypeOther),
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
		"cspcustomer":      BillingRelationshipTypeCSPCustomer,
		"csppartner":       BillingRelationshipTypeCSPPartner,
		"direct":           BillingRelationshipTypeDirect,
		"indirectcustomer": BillingRelationshipTypeIndirectCustomer,
		"indirectpartner":  BillingRelationshipTypeIndirectPartner,
		"other":            BillingRelationshipTypeOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingRelationshipType(input)
	return &out, nil
}

type ExtendedTermOption string

const (
	ExtendedTermOptionOptedNegativeIn  ExtendedTermOption = "Opted-In"
	ExtendedTermOptionOptedNegativeOut ExtendedTermOption = "Opted-Out"
	ExtendedTermOptionOther            ExtendedTermOption = "Other"
)

func PossibleValuesForExtendedTermOption() []string {
	return []string{
		string(ExtendedTermOptionOptedNegativeIn),
		string(ExtendedTermOptionOptedNegativeOut),
		string(ExtendedTermOptionOther),
	}
}

func (s *ExtendedTermOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedTermOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedTermOption(input string) (*ExtendedTermOption, error) {
	vals := map[string]ExtendedTermOption{
		"opted-in":  ExtendedTermOptionOptedNegativeIn,
		"opted-out": ExtendedTermOptionOptedNegativeOut,
		"other":     ExtendedTermOptionOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedTermOption(input)
	return &out, nil
}

type MarkupStatus string

const (
	MarkupStatusDisabled  MarkupStatus = "Disabled"
	MarkupStatusLocked    MarkupStatus = "Locked"
	MarkupStatusOther     MarkupStatus = "Other"
	MarkupStatusPreview   MarkupStatus = "Preview"
	MarkupStatusPublished MarkupStatus = "Published"
)

func PossibleValuesForMarkupStatus() []string {
	return []string{
		string(MarkupStatusDisabled),
		string(MarkupStatusLocked),
		string(MarkupStatusOther),
		string(MarkupStatusPreview),
		string(MarkupStatusPublished),
	}
}

func (s *MarkupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMarkupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMarkupStatus(input string) (*MarkupStatus, error) {
	vals := map[string]MarkupStatus{
		"disabled":  MarkupStatusDisabled,
		"locked":    MarkupStatusLocked,
		"other":     MarkupStatusOther,
		"preview":   MarkupStatusPreview,
		"published": MarkupStatusPublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MarkupStatus(input)
	return &out, nil
}

type PaymentTermsEligibilityCode string

const (
	PaymentTermsEligibilityCodeBillingAccountNotFound         PaymentTermsEligibilityCode = "BillingAccountNotFound"
	PaymentTermsEligibilityCodeInactiveBillingAccount         PaymentTermsEligibilityCode = "InactiveBillingAccount"
	PaymentTermsEligibilityCodeIneligibleBillingAccountStatus PaymentTermsEligibilityCode = "IneligibleBillingAccountStatus"
	PaymentTermsEligibilityCodeInvalidBillingAccountType      PaymentTermsEligibilityCode = "InvalidBillingAccountType"
	PaymentTermsEligibilityCodeInvalidDateFormat              PaymentTermsEligibilityCode = "InvalidDateFormat"
	PaymentTermsEligibilityCodeInvalidDateRange               PaymentTermsEligibilityCode = "InvalidDateRange"
	PaymentTermsEligibilityCodeInvalidTerms                   PaymentTermsEligibilityCode = "InvalidTerms"
	PaymentTermsEligibilityCodeNullOrEmptyPaymentTerms        PaymentTermsEligibilityCode = "NullOrEmptyPaymentTerms"
	PaymentTermsEligibilityCodeOther                          PaymentTermsEligibilityCode = "Other"
	PaymentTermsEligibilityCodeOverlappingPaymentTerms        PaymentTermsEligibilityCode = "OverlappingPaymentTerms"
)

func PossibleValuesForPaymentTermsEligibilityCode() []string {
	return []string{
		string(PaymentTermsEligibilityCodeBillingAccountNotFound),
		string(PaymentTermsEligibilityCodeInactiveBillingAccount),
		string(PaymentTermsEligibilityCodeIneligibleBillingAccountStatus),
		string(PaymentTermsEligibilityCodeInvalidBillingAccountType),
		string(PaymentTermsEligibilityCodeInvalidDateFormat),
		string(PaymentTermsEligibilityCodeInvalidDateRange),
		string(PaymentTermsEligibilityCodeInvalidTerms),
		string(PaymentTermsEligibilityCodeNullOrEmptyPaymentTerms),
		string(PaymentTermsEligibilityCodeOther),
		string(PaymentTermsEligibilityCodeOverlappingPaymentTerms),
	}
}

func (s *PaymentTermsEligibilityCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePaymentTermsEligibilityCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePaymentTermsEligibilityCode(input string) (*PaymentTermsEligibilityCode, error) {
	vals := map[string]PaymentTermsEligibilityCode{
		"billingaccountnotfound":         PaymentTermsEligibilityCodeBillingAccountNotFound,
		"inactivebillingaccount":         PaymentTermsEligibilityCodeInactiveBillingAccount,
		"ineligiblebillingaccountstatus": PaymentTermsEligibilityCodeIneligibleBillingAccountStatus,
		"invalidbillingaccounttype":      PaymentTermsEligibilityCodeInvalidBillingAccountType,
		"invaliddateformat":              PaymentTermsEligibilityCodeInvalidDateFormat,
		"invaliddaterange":               PaymentTermsEligibilityCodeInvalidDateRange,
		"invalidterms":                   PaymentTermsEligibilityCodeInvalidTerms,
		"nulloremptypaymentterms":        PaymentTermsEligibilityCodeNullOrEmptyPaymentTerms,
		"other":                          PaymentTermsEligibilityCodeOther,
		"overlappingpaymentterms":        PaymentTermsEligibilityCodeOverlappingPaymentTerms,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentTermsEligibilityCode(input)
	return &out, nil
}

type PaymentTermsEligibilityStatus string

const (
	PaymentTermsEligibilityStatusInvalid PaymentTermsEligibilityStatus = "Invalid"
	PaymentTermsEligibilityStatusOther   PaymentTermsEligibilityStatus = "Other"
	PaymentTermsEligibilityStatusValid   PaymentTermsEligibilityStatus = "Valid"
)

func PossibleValuesForPaymentTermsEligibilityStatus() []string {
	return []string{
		string(PaymentTermsEligibilityStatusInvalid),
		string(PaymentTermsEligibilityStatusOther),
		string(PaymentTermsEligibilityStatusValid),
	}
}

func (s *PaymentTermsEligibilityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePaymentTermsEligibilityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePaymentTermsEligibilityStatus(input string) (*PaymentTermsEligibilityStatus, error) {
	vals := map[string]PaymentTermsEligibilityStatus{
		"invalid": PaymentTermsEligibilityStatusInvalid,
		"other":   PaymentTermsEligibilityStatusOther,
		"valid":   PaymentTermsEligibilityStatusValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentTermsEligibilityStatus(input)
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

type SupportLevel string

const (
	SupportLevelDeveloper         SupportLevel = "Developer"
	SupportLevelOther             SupportLevel = "Other"
	SupportLevelProNegativeDirect SupportLevel = "Pro-Direct"
	SupportLevelStandard          SupportLevel = "Standard"
)

func PossibleValuesForSupportLevel() []string {
	return []string{
		string(SupportLevelDeveloper),
		string(SupportLevelOther),
		string(SupportLevelProNegativeDirect),
		string(SupportLevelStandard),
	}
}

func (s *SupportLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSupportLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSupportLevel(input string) (*SupportLevel, error) {
	vals := map[string]SupportLevel{
		"developer":  SupportLevelDeveloper,
		"other":      SupportLevelOther,
		"pro-direct": SupportLevelProNegativeDirect,
		"standard":   SupportLevelStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportLevel(input)
	return &out, nil
}

type TaxIdentifierStatus string

const (
	TaxIdentifierStatusInvalid TaxIdentifierStatus = "Invalid"
	TaxIdentifierStatusOther   TaxIdentifierStatus = "Other"
	TaxIdentifierStatusValid   TaxIdentifierStatus = "Valid"
)

func PossibleValuesForTaxIdentifierStatus() []string {
	return []string{
		string(TaxIdentifierStatusInvalid),
		string(TaxIdentifierStatusOther),
		string(TaxIdentifierStatusValid),
	}
}

func (s *TaxIdentifierStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaxIdentifierStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaxIdentifierStatus(input string) (*TaxIdentifierStatus, error) {
	vals := map[string]TaxIdentifierStatus{
		"invalid": TaxIdentifierStatusInvalid,
		"other":   TaxIdentifierStatusOther,
		"valid":   TaxIdentifierStatusValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaxIdentifierStatus(input)
	return &out, nil
}

type TaxIdentifierType string

const (
	TaxIdentifierTypeBrazilCcmId                  TaxIdentifierType = "BrazilCcmId"
	TaxIdentifierTypeBrazilCnpjId                 TaxIdentifierType = "BrazilCnpjId"
	TaxIdentifierTypeBrazilCpfId                  TaxIdentifierType = "BrazilCpfId"
	TaxIdentifierTypeCanadianFederalExempt        TaxIdentifierType = "CanadianFederalExempt"
	TaxIdentifierTypeCanadianProvinceExempt       TaxIdentifierType = "CanadianProvinceExempt"
	TaxIdentifierTypeExternalTaxation             TaxIdentifierType = "ExternalTaxation"
	TaxIdentifierTypeIndiaFederalServiceTaxId     TaxIdentifierType = "IndiaFederalServiceTaxId"
	TaxIdentifierTypeIndiaFederalTanId            TaxIdentifierType = "IndiaFederalTanId"
	TaxIdentifierTypeIndiaPanId                   TaxIdentifierType = "IndiaPanId"
	TaxIdentifierTypeIndiaStateCstId              TaxIdentifierType = "IndiaStateCstId"
	TaxIdentifierTypeIndiaStateGstINId            TaxIdentifierType = "IndiaStateGstINId"
	TaxIdentifierTypeIndiaStateVatId              TaxIdentifierType = "IndiaStateVatId"
	TaxIdentifierTypeIntlExempt                   TaxIdentifierType = "IntlExempt"
	TaxIdentifierTypeLoveCode                     TaxIdentifierType = "LoveCode"
	TaxIdentifierTypeMobileBarCode                TaxIdentifierType = "MobileBarCode"
	TaxIdentifierTypeNationalIdentificationNumber TaxIdentifierType = "NationalIdentificationNumber"
	TaxIdentifierTypeOther                        TaxIdentifierType = "Other"
	TaxIdentifierTypePublicSectorId               TaxIdentifierType = "PublicSectorId"
	TaxIdentifierTypeUSExempt                     TaxIdentifierType = "USExempt"
	TaxIdentifierTypeVatId                        TaxIdentifierType = "VatId"
)

func PossibleValuesForTaxIdentifierType() []string {
	return []string{
		string(TaxIdentifierTypeBrazilCcmId),
		string(TaxIdentifierTypeBrazilCnpjId),
		string(TaxIdentifierTypeBrazilCpfId),
		string(TaxIdentifierTypeCanadianFederalExempt),
		string(TaxIdentifierTypeCanadianProvinceExempt),
		string(TaxIdentifierTypeExternalTaxation),
		string(TaxIdentifierTypeIndiaFederalServiceTaxId),
		string(TaxIdentifierTypeIndiaFederalTanId),
		string(TaxIdentifierTypeIndiaPanId),
		string(TaxIdentifierTypeIndiaStateCstId),
		string(TaxIdentifierTypeIndiaStateGstINId),
		string(TaxIdentifierTypeIndiaStateVatId),
		string(TaxIdentifierTypeIntlExempt),
		string(TaxIdentifierTypeLoveCode),
		string(TaxIdentifierTypeMobileBarCode),
		string(TaxIdentifierTypeNationalIdentificationNumber),
		string(TaxIdentifierTypeOther),
		string(TaxIdentifierTypePublicSectorId),
		string(TaxIdentifierTypeUSExempt),
		string(TaxIdentifierTypeVatId),
	}
}

func (s *TaxIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaxIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaxIdentifierType(input string) (*TaxIdentifierType, error) {
	vals := map[string]TaxIdentifierType{
		"brazilccmid":                  TaxIdentifierTypeBrazilCcmId,
		"brazilcnpjid":                 TaxIdentifierTypeBrazilCnpjId,
		"brazilcpfid":                  TaxIdentifierTypeBrazilCpfId,
		"canadianfederalexempt":        TaxIdentifierTypeCanadianFederalExempt,
		"canadianprovinceexempt":       TaxIdentifierTypeCanadianProvinceExempt,
		"externaltaxation":             TaxIdentifierTypeExternalTaxation,
		"indiafederalservicetaxid":     TaxIdentifierTypeIndiaFederalServiceTaxId,
		"indiafederaltanid":            TaxIdentifierTypeIndiaFederalTanId,
		"indiapanid":                   TaxIdentifierTypeIndiaPanId,
		"indiastatecstid":              TaxIdentifierTypeIndiaStateCstId,
		"indiastategstinid":            TaxIdentifierTypeIndiaStateGstINId,
		"indiastatevatid":              TaxIdentifierTypeIndiaStateVatId,
		"intlexempt":                   TaxIdentifierTypeIntlExempt,
		"lovecode":                     TaxIdentifierTypeLoveCode,
		"mobilebarcode":                TaxIdentifierTypeMobileBarCode,
		"nationalidentificationnumber": TaxIdentifierTypeNationalIdentificationNumber,
		"other":                        TaxIdentifierTypeOther,
		"publicsectorid":               TaxIdentifierTypePublicSectorId,
		"usexempt":                     TaxIdentifierTypeUSExempt,
		"vatid":                        TaxIdentifierTypeVatId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaxIdentifierType(input)
	return &out, nil
}
