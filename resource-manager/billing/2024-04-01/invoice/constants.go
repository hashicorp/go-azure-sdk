package invoice

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSource string

const (
	DocumentSourceDRS   DocumentSource = "DRS"
	DocumentSourceENF   DocumentSource = "ENF"
	DocumentSourceOther DocumentSource = "Other"
)

func PossibleValuesForDocumentSource() []string {
	return []string{
		string(DocumentSourceDRS),
		string(DocumentSourceENF),
		string(DocumentSourceOther),
	}
}

func (s *DocumentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDocumentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDocumentSource(input string) (*DocumentSource, error) {
	vals := map[string]DocumentSource{
		"drs":   DocumentSourceDRS,
		"enf":   DocumentSourceENF,
		"other": DocumentSourceOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DocumentSource(input)
	return &out, nil
}

type FailedPaymentReason string

const (
	FailedPaymentReasonBankDeclined         FailedPaymentReason = "BankDeclined"
	FailedPaymentReasonCardExpired          FailedPaymentReason = "CardExpired"
	FailedPaymentReasonIncorrectCardDetails FailedPaymentReason = "IncorrectCardDetails"
	FailedPaymentReasonOther                FailedPaymentReason = "Other"
)

func PossibleValuesForFailedPaymentReason() []string {
	return []string{
		string(FailedPaymentReasonBankDeclined),
		string(FailedPaymentReasonCardExpired),
		string(FailedPaymentReasonIncorrectCardDetails),
		string(FailedPaymentReasonOther),
	}
}

func (s *FailedPaymentReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFailedPaymentReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFailedPaymentReason(input string) (*FailedPaymentReason, error) {
	vals := map[string]FailedPaymentReason{
		"bankdeclined":         FailedPaymentReasonBankDeclined,
		"cardexpired":          FailedPaymentReasonCardExpired,
		"incorrectcarddetails": FailedPaymentReasonIncorrectCardDetails,
		"other":                FailedPaymentReasonOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FailedPaymentReason(input)
	return &out, nil
}

type InvoiceDocumentType string

const (
	InvoiceDocumentTypeCreditNote   InvoiceDocumentType = "CreditNote"
	InvoiceDocumentTypeInvoice      InvoiceDocumentType = "Invoice"
	InvoiceDocumentTypeOther        InvoiceDocumentType = "Other"
	InvoiceDocumentTypeSummary      InvoiceDocumentType = "Summary"
	InvoiceDocumentTypeTaxReceipt   InvoiceDocumentType = "TaxReceipt"
	InvoiceDocumentTypeTransactions InvoiceDocumentType = "Transactions"
	InvoiceDocumentTypeVoidNote     InvoiceDocumentType = "VoidNote"
)

func PossibleValuesForInvoiceDocumentType() []string {
	return []string{
		string(InvoiceDocumentTypeCreditNote),
		string(InvoiceDocumentTypeInvoice),
		string(InvoiceDocumentTypeOther),
		string(InvoiceDocumentTypeSummary),
		string(InvoiceDocumentTypeTaxReceipt),
		string(InvoiceDocumentTypeTransactions),
		string(InvoiceDocumentTypeVoidNote),
	}
}

func (s *InvoiceDocumentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceDocumentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceDocumentType(input string) (*InvoiceDocumentType, error) {
	vals := map[string]InvoiceDocumentType{
		"creditnote":   InvoiceDocumentTypeCreditNote,
		"invoice":      InvoiceDocumentTypeInvoice,
		"other":        InvoiceDocumentTypeOther,
		"summary":      InvoiceDocumentTypeSummary,
		"taxreceipt":   InvoiceDocumentTypeTaxReceipt,
		"transactions": InvoiceDocumentTypeTransactions,
		"voidnote":     InvoiceDocumentTypeVoidNote,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceDocumentType(input)
	return &out, nil
}

type InvoiceStatus string

const (
	InvoiceStatusDue     InvoiceStatus = "Due"
	InvoiceStatusLocked  InvoiceStatus = "Locked"
	InvoiceStatusOther   InvoiceStatus = "Other"
	InvoiceStatusOverDue InvoiceStatus = "OverDue"
	InvoiceStatusPaid    InvoiceStatus = "Paid"
	InvoiceStatusVoid    InvoiceStatus = "Void"
)

func PossibleValuesForInvoiceStatus() []string {
	return []string{
		string(InvoiceStatusDue),
		string(InvoiceStatusLocked),
		string(InvoiceStatusOther),
		string(InvoiceStatusOverDue),
		string(InvoiceStatusPaid),
		string(InvoiceStatusVoid),
	}
}

func (s *InvoiceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceStatus(input string) (*InvoiceStatus, error) {
	vals := map[string]InvoiceStatus{
		"due":     InvoiceStatusDue,
		"locked":  InvoiceStatusLocked,
		"other":   InvoiceStatusOther,
		"overdue": InvoiceStatusOverDue,
		"paid":    InvoiceStatusPaid,
		"void":    InvoiceStatusVoid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceStatus(input)
	return &out, nil
}

type InvoiceType string

const (
	InvoiceTypeAzureMarketplace InvoiceType = "AzureMarketplace"
	InvoiceTypeAzureServices    InvoiceType = "AzureServices"
	InvoiceTypeAzureSupport     InvoiceType = "AzureSupport"
	InvoiceTypeOther            InvoiceType = "Other"
)

func PossibleValuesForInvoiceType() []string {
	return []string{
		string(InvoiceTypeAzureMarketplace),
		string(InvoiceTypeAzureServices),
		string(InvoiceTypeAzureSupport),
		string(InvoiceTypeOther),
	}
}

func (s *InvoiceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceType(input string) (*InvoiceType, error) {
	vals := map[string]InvoiceType{
		"azuremarketplace": InvoiceTypeAzureMarketplace,
		"azureservices":    InvoiceTypeAzureServices,
		"azuresupport":     InvoiceTypeAzureSupport,
		"other":            InvoiceTypeOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceType(input)
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

type RefundReasonCode string

const (
	RefundReasonCodeAccidentalConversion RefundReasonCode = "AccidentalConversion"
	RefundReasonCodeAccidentalPurchase   RefundReasonCode = "AccidentalPurchase"
	RefundReasonCodeForgotToCancel       RefundReasonCode = "ForgotToCancel"
	RefundReasonCodeOther                RefundReasonCode = "Other"
	RefundReasonCodeUnclearDocumentation RefundReasonCode = "UnclearDocumentation"
	RefundReasonCodeUnclearPricing       RefundReasonCode = "UnclearPricing"
)

func PossibleValuesForRefundReasonCode() []string {
	return []string{
		string(RefundReasonCodeAccidentalConversion),
		string(RefundReasonCodeAccidentalPurchase),
		string(RefundReasonCodeForgotToCancel),
		string(RefundReasonCodeOther),
		string(RefundReasonCodeUnclearDocumentation),
		string(RefundReasonCodeUnclearPricing),
	}
}

func (s *RefundReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRefundReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRefundReasonCode(input string) (*RefundReasonCode, error) {
	vals := map[string]RefundReasonCode{
		"accidentalconversion": RefundReasonCodeAccidentalConversion,
		"accidentalpurchase":   RefundReasonCodeAccidentalPurchase,
		"forgottocancel":       RefundReasonCodeForgotToCancel,
		"other":                RefundReasonCodeOther,
		"uncleardocumentation": RefundReasonCodeUnclearDocumentation,
		"unclearpricing":       RefundReasonCodeUnclearPricing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RefundReasonCode(input)
	return &out, nil
}

type RefundStatus string

const (
	RefundStatusApproved  RefundStatus = "Approved"
	RefundStatusCancelled RefundStatus = "Cancelled"
	RefundStatusCompleted RefundStatus = "Completed"
	RefundStatusDeclined  RefundStatus = "Declined"
	RefundStatusExpired   RefundStatus = "Expired"
	RefundStatusOther     RefundStatus = "Other"
	RefundStatusPending   RefundStatus = "Pending"
)

func PossibleValuesForRefundStatus() []string {
	return []string{
		string(RefundStatusApproved),
		string(RefundStatusCancelled),
		string(RefundStatusCompleted),
		string(RefundStatusDeclined),
		string(RefundStatusExpired),
		string(RefundStatusOther),
		string(RefundStatusPending),
	}
}

func (s *RefundStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRefundStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRefundStatus(input string) (*RefundStatus, error) {
	vals := map[string]RefundStatus{
		"approved":  RefundStatusApproved,
		"cancelled": RefundStatusCancelled,
		"completed": RefundStatusCompleted,
		"declined":  RefundStatusDeclined,
		"expired":   RefundStatusExpired,
		"other":     RefundStatusOther,
		"pending":   RefundStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RefundStatus(input)
	return &out, nil
}

type SpecialTaxationType string

const (
	SpecialTaxationTypeInvoiceLevel  SpecialTaxationType = "InvoiceLevel"
	SpecialTaxationTypeSubtotalLevel SpecialTaxationType = "SubtotalLevel"
)

func PossibleValuesForSpecialTaxationType() []string {
	return []string{
		string(SpecialTaxationTypeInvoiceLevel),
		string(SpecialTaxationTypeSubtotalLevel),
	}
}

func (s *SpecialTaxationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpecialTaxationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpecialTaxationType(input string) (*SpecialTaxationType, error) {
	vals := map[string]SpecialTaxationType{
		"invoicelevel":  SpecialTaxationTypeInvoiceLevel,
		"subtotallevel": SpecialTaxationTypeSubtotalLevel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpecialTaxationType(input)
	return &out, nil
}
