package invoices

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentType string

const (
	DocumentTypeCreditNote DocumentType = "CreditNote"
	DocumentTypeInvoice    DocumentType = "Invoice"
	DocumentTypeTaxReceipt DocumentType = "TaxReceipt"
	DocumentTypeVoidNote   DocumentType = "VoidNote"
)

func PossibleValuesForDocumentType() []string {
	return []string{
		string(DocumentTypeCreditNote),
		string(DocumentTypeInvoice),
		string(DocumentTypeTaxReceipt),
		string(DocumentTypeVoidNote),
	}
}

func (s *DocumentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDocumentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDocumentType(input string) (*DocumentType, error) {
	vals := map[string]DocumentType{
		"creditnote": DocumentTypeCreditNote,
		"invoice":    DocumentTypeInvoice,
		"taxreceipt": DocumentTypeTaxReceipt,
		"voidnote":   DocumentTypeVoidNote,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DocumentType(input)
	return &out, nil
}

type InvoiceStatus string

const (
	InvoiceStatusDue     InvoiceStatus = "Due"
	InvoiceStatusOverDue InvoiceStatus = "OverDue"
	InvoiceStatusPaid    InvoiceStatus = "Paid"
	InvoiceStatusVoid    InvoiceStatus = "Void"
)

func PossibleValuesForInvoiceStatus() []string {
	return []string{
		string(InvoiceStatusDue),
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
	InvoiceTypeAzureService     InvoiceType = "AzureService"
	InvoiceTypeAzureSupport     InvoiceType = "AzureSupport"
)

func PossibleValuesForInvoiceType() []string {
	return []string{
		string(InvoiceTypeAzureMarketplace),
		string(InvoiceTypeAzureService),
		string(InvoiceTypeAzureSupport),
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
		"azureservice":     InvoiceTypeAzureService,
		"azuresupport":     InvoiceTypeAzureSupport,
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
	PaymentMethodFamilyCheckWire  PaymentMethodFamily = "CheckWire"
	PaymentMethodFamilyCreditCard PaymentMethodFamily = "CreditCard"
	PaymentMethodFamilyCredits    PaymentMethodFamily = "Credits"
	PaymentMethodFamilyNone       PaymentMethodFamily = "None"
)

func PossibleValuesForPaymentMethodFamily() []string {
	return []string{
		string(PaymentMethodFamilyCheckWire),
		string(PaymentMethodFamilyCreditCard),
		string(PaymentMethodFamilyCredits),
		string(PaymentMethodFamilyNone),
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
		"checkwire":  PaymentMethodFamilyCheckWire,
		"creditcard": PaymentMethodFamilyCreditCard,
		"credits":    PaymentMethodFamilyCredits,
		"none":       PaymentMethodFamilyNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentMethodFamily(input)
	return &out, nil
}

type RebillDocumentType string

const (
	RebillDocumentTypeCredit   RebillDocumentType = "Credit"
	RebillDocumentTypeOriginal RebillDocumentType = "Original"
	RebillDocumentTypeRebill   RebillDocumentType = "Rebill"
)

func PossibleValuesForRebillDocumentType() []string {
	return []string{
		string(RebillDocumentTypeCredit),
		string(RebillDocumentTypeOriginal),
		string(RebillDocumentTypeRebill),
	}
}

func (s *RebillDocumentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRebillDocumentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRebillDocumentType(input string) (*RebillDocumentType, error) {
	vals := map[string]RebillDocumentType{
		"credit":   RebillDocumentTypeCredit,
		"original": RebillDocumentTypeOriginal,
		"rebill":   RebillDocumentTypeRebill,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RebillDocumentType(input)
	return &out, nil
}
