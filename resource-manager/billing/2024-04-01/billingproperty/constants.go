package billingproperty

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

type BillingSubscriptionStatus string

const (
	BillingSubscriptionStatusActive    BillingSubscriptionStatus = "Active"
	BillingSubscriptionStatusAutoRenew BillingSubscriptionStatus = "AutoRenew"
	BillingSubscriptionStatusCancelled BillingSubscriptionStatus = "Cancelled"
	BillingSubscriptionStatusDeleted   BillingSubscriptionStatus = "Deleted"
	BillingSubscriptionStatusDisabled  BillingSubscriptionStatus = "Disabled"
	BillingSubscriptionStatusExpired   BillingSubscriptionStatus = "Expired"
	BillingSubscriptionStatusExpiring  BillingSubscriptionStatus = "Expiring"
	BillingSubscriptionStatusFailed    BillingSubscriptionStatus = "Failed"
	BillingSubscriptionStatusOther     BillingSubscriptionStatus = "Other"
	BillingSubscriptionStatusSuspended BillingSubscriptionStatus = "Suspended"
	BillingSubscriptionStatusUnknown   BillingSubscriptionStatus = "Unknown"
	BillingSubscriptionStatusWarned    BillingSubscriptionStatus = "Warned"
)

func PossibleValuesForBillingSubscriptionStatus() []string {
	return []string{
		string(BillingSubscriptionStatusActive),
		string(BillingSubscriptionStatusAutoRenew),
		string(BillingSubscriptionStatusCancelled),
		string(BillingSubscriptionStatusDeleted),
		string(BillingSubscriptionStatusDisabled),
		string(BillingSubscriptionStatusExpired),
		string(BillingSubscriptionStatusExpiring),
		string(BillingSubscriptionStatusFailed),
		string(BillingSubscriptionStatusOther),
		string(BillingSubscriptionStatusSuspended),
		string(BillingSubscriptionStatusUnknown),
		string(BillingSubscriptionStatusWarned),
	}
}

func (s *BillingSubscriptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingSubscriptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingSubscriptionStatus(input string) (*BillingSubscriptionStatus, error) {
	vals := map[string]BillingSubscriptionStatus{
		"active":    BillingSubscriptionStatusActive,
		"autorenew": BillingSubscriptionStatusAutoRenew,
		"cancelled": BillingSubscriptionStatusCancelled,
		"deleted":   BillingSubscriptionStatusDeleted,
		"disabled":  BillingSubscriptionStatusDisabled,
		"expired":   BillingSubscriptionStatusExpired,
		"expiring":  BillingSubscriptionStatusExpiring,
		"failed":    BillingSubscriptionStatusFailed,
		"other":     BillingSubscriptionStatusOther,
		"suspended": BillingSubscriptionStatusSuspended,
		"unknown":   BillingSubscriptionStatusUnknown,
		"warned":    BillingSubscriptionStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingSubscriptionStatus(input)
	return &out, nil
}

type CustomerStatus string

const (
	CustomerStatusActive      CustomerStatus = "Active"
	CustomerStatusDeleted     CustomerStatus = "Deleted"
	CustomerStatusDisabled    CustomerStatus = "Disabled"
	CustomerStatusOther       CustomerStatus = "Other"
	CustomerStatusPending     CustomerStatus = "Pending"
	CustomerStatusUnderReview CustomerStatus = "UnderReview"
	CustomerStatusWarned      CustomerStatus = "Warned"
)

func PossibleValuesForCustomerStatus() []string {
	return []string{
		string(CustomerStatusActive),
		string(CustomerStatusDeleted),
		string(CustomerStatusDisabled),
		string(CustomerStatusOther),
		string(CustomerStatusPending),
		string(CustomerStatusUnderReview),
		string(CustomerStatusWarned),
	}
}

func (s *CustomerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomerStatus(input string) (*CustomerStatus, error) {
	vals := map[string]CustomerStatus{
		"active":      CustomerStatusActive,
		"deleted":     CustomerStatusDeleted,
		"disabled":    CustomerStatusDisabled,
		"other":       CustomerStatusOther,
		"pending":     CustomerStatusPending,
		"underreview": CustomerStatusUnderReview,
		"warned":      CustomerStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomerStatus(input)
	return &out, nil
}

type InvoiceSectionState string

const (
	InvoiceSectionStateActive      InvoiceSectionState = "Active"
	InvoiceSectionStateDeleted     InvoiceSectionState = "Deleted"
	InvoiceSectionStateDisabled    InvoiceSectionState = "Disabled"
	InvoiceSectionStateOther       InvoiceSectionState = "Other"
	InvoiceSectionStateRestricted  InvoiceSectionState = "Restricted"
	InvoiceSectionStateUnderReview InvoiceSectionState = "UnderReview"
	InvoiceSectionStateWarned      InvoiceSectionState = "Warned"
)

func PossibleValuesForInvoiceSectionState() []string {
	return []string{
		string(InvoiceSectionStateActive),
		string(InvoiceSectionStateDeleted),
		string(InvoiceSectionStateDisabled),
		string(InvoiceSectionStateOther),
		string(InvoiceSectionStateRestricted),
		string(InvoiceSectionStateUnderReview),
		string(InvoiceSectionStateWarned),
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
		"active":      InvoiceSectionStateActive,
		"deleted":     InvoiceSectionStateDeleted,
		"disabled":    InvoiceSectionStateDisabled,
		"other":       InvoiceSectionStateOther,
		"restricted":  InvoiceSectionStateRestricted,
		"underreview": InvoiceSectionStateUnderReview,
		"warned":      InvoiceSectionStateWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceSectionState(input)
	return &out, nil
}

type InvoiceSectionStateReasonCode string

const (
	InvoiceSectionStateReasonCodeOther                InvoiceSectionStateReasonCode = "Other"
	InvoiceSectionStateReasonCodePastDue              InvoiceSectionStateReasonCode = "PastDue"
	InvoiceSectionStateReasonCodeSpendingLimitExpired InvoiceSectionStateReasonCode = "SpendingLimitExpired"
	InvoiceSectionStateReasonCodeSpendingLimitReached InvoiceSectionStateReasonCode = "SpendingLimitReached"
	InvoiceSectionStateReasonCodeUnusualActivity      InvoiceSectionStateReasonCode = "UnusualActivity"
)

func PossibleValuesForInvoiceSectionStateReasonCode() []string {
	return []string{
		string(InvoiceSectionStateReasonCodeOther),
		string(InvoiceSectionStateReasonCodePastDue),
		string(InvoiceSectionStateReasonCodeSpendingLimitExpired),
		string(InvoiceSectionStateReasonCodeSpendingLimitReached),
		string(InvoiceSectionStateReasonCodeUnusualActivity),
	}
}

func (s *InvoiceSectionStateReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceSectionStateReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceSectionStateReasonCode(input string) (*InvoiceSectionStateReasonCode, error) {
	vals := map[string]InvoiceSectionStateReasonCode{
		"other":                InvoiceSectionStateReasonCodeOther,
		"pastdue":              InvoiceSectionStateReasonCodePastDue,
		"spendinglimitexpired": InvoiceSectionStateReasonCodeSpendingLimitExpired,
		"spendinglimitreached": InvoiceSectionStateReasonCodeSpendingLimitReached,
		"unusualactivity":      InvoiceSectionStateReasonCodeUnusualActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceSectionStateReasonCode(input)
	return &out, nil
}

type PaymentMethodFamily string

const (
	PaymentMethodFamilyCheckWire   PaymentMethodFamily = "CheckWire"
	PaymentMethodFamilyCreditCard  PaymentMethodFamily = "CreditCard"
	PaymentMethodFamilyCredits     PaymentMethodFamily = "Credits"
	PaymentMethodFamilyDirectDebit PaymentMethodFamily = "DirectDebit"
	PaymentMethodFamilyEWallet     PaymentMethodFamily = "EWallet"
	PaymentMethodFamilyNone        PaymentMethodFamily = "None"
	PaymentMethodFamilyOther       PaymentMethodFamily = "Other"
	PaymentMethodFamilyTaskOrder   PaymentMethodFamily = "TaskOrder"
)

func PossibleValuesForPaymentMethodFamily() []string {
	return []string{
		string(PaymentMethodFamilyCheckWire),
		string(PaymentMethodFamilyCreditCard),
		string(PaymentMethodFamilyCredits),
		string(PaymentMethodFamilyDirectDebit),
		string(PaymentMethodFamilyEWallet),
		string(PaymentMethodFamilyNone),
		string(PaymentMethodFamilyOther),
		string(PaymentMethodFamilyTaskOrder),
	}
}

func (s *PaymentMethodFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePaymentMethodFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePaymentMethodFamily(input string) (*PaymentMethodFamily, error) {
	vals := map[string]PaymentMethodFamily{
		"checkwire":   PaymentMethodFamilyCheckWire,
		"creditcard":  PaymentMethodFamilyCreditCard,
		"credits":     PaymentMethodFamilyCredits,
		"directdebit": PaymentMethodFamilyDirectDebit,
		"ewallet":     PaymentMethodFamilyEWallet,
		"none":        PaymentMethodFamilyNone,
		"other":       PaymentMethodFamilyOther,
		"taskorder":   PaymentMethodFamilyTaskOrder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentMethodFamily(input)
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

type SpendingLimitStatus string

const (
	SpendingLimitStatusActive       SpendingLimitStatus = "Active"
	SpendingLimitStatusExpired      SpendingLimitStatus = "Expired"
	SpendingLimitStatusLimitReached SpendingLimitStatus = "LimitReached"
	SpendingLimitStatusLimitRemoved SpendingLimitStatus = "LimitRemoved"
	SpendingLimitStatusNone         SpendingLimitStatus = "None"
	SpendingLimitStatusOther        SpendingLimitStatus = "Other"
)

func PossibleValuesForSpendingLimitStatus() []string {
	return []string{
		string(SpendingLimitStatusActive),
		string(SpendingLimitStatusExpired),
		string(SpendingLimitStatusLimitReached),
		string(SpendingLimitStatusLimitRemoved),
		string(SpendingLimitStatusNone),
		string(SpendingLimitStatusOther),
	}
}

func (s *SpendingLimitStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimitStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimitStatus(input string) (*SpendingLimitStatus, error) {
	vals := map[string]SpendingLimitStatus{
		"active":       SpendingLimitStatusActive,
		"expired":      SpendingLimitStatusExpired,
		"limitreached": SpendingLimitStatusLimitReached,
		"limitremoved": SpendingLimitStatusLimitRemoved,
		"none":         SpendingLimitStatusNone,
		"other":        SpendingLimitStatusOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimitStatus(input)
	return &out, nil
}

type SpendingLimitType string

const (
	SpendingLimitTypeAcademicSponsorship     SpendingLimitType = "AcademicSponsorship"
	SpendingLimitTypeAzureConsumptionCredit  SpendingLimitType = "AzureConsumptionCredit"
	SpendingLimitTypeAzureForStudents        SpendingLimitType = "AzureForStudents"
	SpendingLimitTypeAzureForStudentsStarter SpendingLimitType = "AzureForStudentsStarter"
	SpendingLimitTypeAzurePassSponsorship    SpendingLimitType = "AzurePassSponsorship"
	SpendingLimitTypeFreeAccount             SpendingLimitType = "FreeAccount"
	SpendingLimitTypeMSDN                    SpendingLimitType = "MSDN"
	SpendingLimitTypeMpnSponsorship          SpendingLimitType = "MpnSponsorship"
	SpendingLimitTypeNonProfitSponsorship    SpendingLimitType = "NonProfitSponsorship"
	SpendingLimitTypeNone                    SpendingLimitType = "None"
	SpendingLimitTypeOther                   SpendingLimitType = "Other"
	SpendingLimitTypeSandbox                 SpendingLimitType = "Sandbox"
	SpendingLimitTypeSponsorship             SpendingLimitType = "Sponsorship"
	SpendingLimitTypeStartupSponsorship      SpendingLimitType = "StartupSponsorship"
	SpendingLimitTypeVisualStudio            SpendingLimitType = "VisualStudio"
)

func PossibleValuesForSpendingLimitType() []string {
	return []string{
		string(SpendingLimitTypeAcademicSponsorship),
		string(SpendingLimitTypeAzureConsumptionCredit),
		string(SpendingLimitTypeAzureForStudents),
		string(SpendingLimitTypeAzureForStudentsStarter),
		string(SpendingLimitTypeAzurePassSponsorship),
		string(SpendingLimitTypeFreeAccount),
		string(SpendingLimitTypeMSDN),
		string(SpendingLimitTypeMpnSponsorship),
		string(SpendingLimitTypeNonProfitSponsorship),
		string(SpendingLimitTypeNone),
		string(SpendingLimitTypeOther),
		string(SpendingLimitTypeSandbox),
		string(SpendingLimitTypeSponsorship),
		string(SpendingLimitTypeStartupSponsorship),
		string(SpendingLimitTypeVisualStudio),
	}
}

func (s *SpendingLimitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimitType(input string) (*SpendingLimitType, error) {
	vals := map[string]SpendingLimitType{
		"academicsponsorship":     SpendingLimitTypeAcademicSponsorship,
		"azureconsumptioncredit":  SpendingLimitTypeAzureConsumptionCredit,
		"azureforstudents":        SpendingLimitTypeAzureForStudents,
		"azureforstudentsstarter": SpendingLimitTypeAzureForStudentsStarter,
		"azurepasssponsorship":    SpendingLimitTypeAzurePassSponsorship,
		"freeaccount":             SpendingLimitTypeFreeAccount,
		"msdn":                    SpendingLimitTypeMSDN,
		"mpnsponsorship":          SpendingLimitTypeMpnSponsorship,
		"nonprofitsponsorship":    SpendingLimitTypeNonProfitSponsorship,
		"none":                    SpendingLimitTypeNone,
		"other":                   SpendingLimitTypeOther,
		"sandbox":                 SpendingLimitTypeSandbox,
		"sponsorship":             SpendingLimitTypeSponsorship,
		"startupsponsorship":      SpendingLimitTypeStartupSponsorship,
		"visualstudio":            SpendingLimitTypeVisualStudio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimitType(input)
	return &out, nil
}

type SubscriptionBillingType string

const (
	SubscriptionBillingTypeBenefit SubscriptionBillingType = "Benefit"
	SubscriptionBillingTypeFree    SubscriptionBillingType = "Free"
	SubscriptionBillingTypeNone    SubscriptionBillingType = "None"
	SubscriptionBillingTypePaid    SubscriptionBillingType = "Paid"
	SubscriptionBillingTypePrePaid SubscriptionBillingType = "PrePaid"
)

func PossibleValuesForSubscriptionBillingType() []string {
	return []string{
		string(SubscriptionBillingTypeBenefit),
		string(SubscriptionBillingTypeFree),
		string(SubscriptionBillingTypeNone),
		string(SubscriptionBillingTypePaid),
		string(SubscriptionBillingTypePrePaid),
	}
}

func (s *SubscriptionBillingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionBillingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionBillingType(input string) (*SubscriptionBillingType, error) {
	vals := map[string]SubscriptionBillingType{
		"benefit": SubscriptionBillingTypeBenefit,
		"free":    SubscriptionBillingTypeFree,
		"none":    SubscriptionBillingTypeNone,
		"paid":    SubscriptionBillingTypePaid,
		"prepaid": SubscriptionBillingTypePrePaid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionBillingType(input)
	return &out, nil
}

type SubscriptionStatusReason string

const (
	SubscriptionStatusReasonCancelled            SubscriptionStatusReason = "Cancelled"
	SubscriptionStatusReasonExpired              SubscriptionStatusReason = "Expired"
	SubscriptionStatusReasonNone                 SubscriptionStatusReason = "None"
	SubscriptionStatusReasonOther                SubscriptionStatusReason = "Other"
	SubscriptionStatusReasonPastDue              SubscriptionStatusReason = "PastDue"
	SubscriptionStatusReasonPolicyViolation      SubscriptionStatusReason = "PolicyViolation"
	SubscriptionStatusReasonSpendingLimitReached SubscriptionStatusReason = "SpendingLimitReached"
	SubscriptionStatusReasonSuspiciousActivity   SubscriptionStatusReason = "SuspiciousActivity"
	SubscriptionStatusReasonTransferred          SubscriptionStatusReason = "Transferred"
)

func PossibleValuesForSubscriptionStatusReason() []string {
	return []string{
		string(SubscriptionStatusReasonCancelled),
		string(SubscriptionStatusReasonExpired),
		string(SubscriptionStatusReasonNone),
		string(SubscriptionStatusReasonOther),
		string(SubscriptionStatusReasonPastDue),
		string(SubscriptionStatusReasonPolicyViolation),
		string(SubscriptionStatusReasonSpendingLimitReached),
		string(SubscriptionStatusReasonSuspiciousActivity),
		string(SubscriptionStatusReasonTransferred),
	}
}

func (s *SubscriptionStatusReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionStatusReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionStatusReason(input string) (*SubscriptionStatusReason, error) {
	vals := map[string]SubscriptionStatusReason{
		"cancelled":            SubscriptionStatusReasonCancelled,
		"expired":              SubscriptionStatusReasonExpired,
		"none":                 SubscriptionStatusReasonNone,
		"other":                SubscriptionStatusReasonOther,
		"pastdue":              SubscriptionStatusReasonPastDue,
		"policyviolation":      SubscriptionStatusReasonPolicyViolation,
		"spendinglimitreached": SubscriptionStatusReasonSpendingLimitReached,
		"suspiciousactivity":   SubscriptionStatusReasonSuspiciousActivity,
		"transferred":          SubscriptionStatusReasonTransferred,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionStatusReason(input)
	return &out, nil
}

type SubscriptionWorkloadType string

const (
	SubscriptionWorkloadTypeDevTest    SubscriptionWorkloadType = "DevTest"
	SubscriptionWorkloadTypeInternal   SubscriptionWorkloadType = "Internal"
	SubscriptionWorkloadTypeNone       SubscriptionWorkloadType = "None"
	SubscriptionWorkloadTypeProduction SubscriptionWorkloadType = "Production"
)

func PossibleValuesForSubscriptionWorkloadType() []string {
	return []string{
		string(SubscriptionWorkloadTypeDevTest),
		string(SubscriptionWorkloadTypeInternal),
		string(SubscriptionWorkloadTypeNone),
		string(SubscriptionWorkloadTypeProduction),
	}
}

func (s *SubscriptionWorkloadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionWorkloadType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionWorkloadType(input string) (*SubscriptionWorkloadType, error) {
	vals := map[string]SubscriptionWorkloadType{
		"devtest":    SubscriptionWorkloadTypeDevTest,
		"internal":   SubscriptionWorkloadTypeInternal,
		"none":       SubscriptionWorkloadTypeNone,
		"production": SubscriptionWorkloadTypeProduction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionWorkloadType(input)
	return &out, nil
}
