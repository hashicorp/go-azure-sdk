package transaction

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditType string

const (
	CreditTypeAzureCreditOffer    CreditType = "AzureCreditOffer"
	CreditTypeAzureFreeCredit     CreditType = "AzureFreeCredit"
	CreditTypeOther               CreditType = "Other"
	CreditTypeRefund              CreditType = "Refund"
	CreditTypeServiceInterruption CreditType = "ServiceInterruption"
)

func PossibleValuesForCreditType() []string {
	return []string{
		string(CreditTypeAzureCreditOffer),
		string(CreditTypeAzureFreeCredit),
		string(CreditTypeOther),
		string(CreditTypeRefund),
		string(CreditTypeServiceInterruption),
	}
}

func (s *CreditType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreditType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreditType(input string) (*CreditType, error) {
	vals := map[string]CreditType{
		"azurecreditoffer":    CreditTypeAzureCreditOffer,
		"azurefreecredit":     CreditTypeAzureFreeCredit,
		"other":               CreditTypeOther,
		"refund":              CreditTypeRefund,
		"serviceinterruption": CreditTypeServiceInterruption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreditType(input)
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

type TransactionKind string

const (
	TransactionKindAll         TransactionKind = "All"
	TransactionKindOther       TransactionKind = "Other"
	TransactionKindReservation TransactionKind = "Reservation"
)

func PossibleValuesForTransactionKind() []string {
	return []string{
		string(TransactionKindAll),
		string(TransactionKindOther),
		string(TransactionKindReservation),
	}
}

func (s *TransactionKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransactionKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransactionKind(input string) (*TransactionKind, error) {
	vals := map[string]TransactionKind{
		"all":         TransactionKindAll,
		"other":       TransactionKindOther,
		"reservation": TransactionKindReservation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransactionKind(input)
	return &out, nil
}

type TransactionType string

const (
	TransactionTypeBilled   TransactionType = "Billed"
	TransactionTypeOther    TransactionType = "Other"
	TransactionTypeUnbilled TransactionType = "Unbilled"
)

func PossibleValuesForTransactionType() []string {
	return []string{
		string(TransactionTypeBilled),
		string(TransactionTypeOther),
		string(TransactionTypeUnbilled),
	}
}

func (s *TransactionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransactionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransactionType(input string) (*TransactionType, error) {
	vals := map[string]TransactionType{
		"billed":   TransactionTypeBilled,
		"other":    TransactionTypeOther,
		"unbilled": TransactionTypeUnbilled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransactionType(input)
	return &out, nil
}
