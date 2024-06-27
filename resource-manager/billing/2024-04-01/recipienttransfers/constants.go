package recipienttransfers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EligibleProductType string

const (
	EligibleProductTypeAzureReservation          EligibleProductType = "AzureReservation"
	EligibleProductTypeDevTestAzureSubscription  EligibleProductType = "DevTestAzureSubscription"
	EligibleProductTypeStandardAzureSubscription EligibleProductType = "StandardAzureSubscription"
)

func PossibleValuesForEligibleProductType() []string {
	return []string{
		string(EligibleProductTypeAzureReservation),
		string(EligibleProductTypeDevTestAzureSubscription),
		string(EligibleProductTypeStandardAzureSubscription),
	}
}

func (s *EligibleProductType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEligibleProductType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEligibleProductType(input string) (*EligibleProductType, error) {
	vals := map[string]EligibleProductType{
		"azurereservation":          EligibleProductTypeAzureReservation,
		"devtestazuresubscription":  EligibleProductTypeDevTestAzureSubscription,
		"standardazuresubscription": EligibleProductTypeStandardAzureSubscription,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EligibleProductType(input)
	return &out, nil
}

type InitiatorCustomerType string

const (
	InitiatorCustomerTypeEA      InitiatorCustomerType = "EA"
	InitiatorCustomerTypePartner InitiatorCustomerType = "Partner"
)

func PossibleValuesForInitiatorCustomerType() []string {
	return []string{
		string(InitiatorCustomerTypeEA),
		string(InitiatorCustomerTypePartner),
	}
}

func (s *InitiatorCustomerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInitiatorCustomerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInitiatorCustomerType(input string) (*InitiatorCustomerType, error) {
	vals := map[string]InitiatorCustomerType{
		"ea":      InitiatorCustomerTypeEA,
		"partner": InitiatorCustomerTypePartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InitiatorCustomerType(input)
	return &out, nil
}

type ProductTransferStatus string

const (
	ProductTransferStatusCompleted  ProductTransferStatus = "Completed"
	ProductTransferStatusFailed     ProductTransferStatus = "Failed"
	ProductTransferStatusInProgress ProductTransferStatus = "InProgress"
	ProductTransferStatusNotStarted ProductTransferStatus = "NotStarted"
)

func PossibleValuesForProductTransferStatus() []string {
	return []string{
		string(ProductTransferStatusCompleted),
		string(ProductTransferStatusFailed),
		string(ProductTransferStatusInProgress),
		string(ProductTransferStatusNotStarted),
	}
}

func (s *ProductTransferStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProductTransferStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProductTransferStatus(input string) (*ProductTransferStatus, error) {
	vals := map[string]ProductTransferStatus{
		"completed":  ProductTransferStatusCompleted,
		"failed":     ProductTransferStatusFailed,
		"inprogress": ProductTransferStatusInProgress,
		"notstarted": ProductTransferStatusNotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductTransferStatus(input)
	return &out, nil
}

type ProductType string

const (
	ProductTypeAzureReservation  ProductType = "AzureReservation"
	ProductTypeAzureSubscription ProductType = "AzureSubscription"
	ProductTypeDepartment        ProductType = "Department"
	ProductTypeSAAS              ProductType = "SAAS"
	ProductTypeSavingsPlan       ProductType = "SavingsPlan"
)

func PossibleValuesForProductType() []string {
	return []string{
		string(ProductTypeAzureReservation),
		string(ProductTypeAzureSubscription),
		string(ProductTypeDepartment),
		string(ProductTypeSAAS),
		string(ProductTypeSavingsPlan),
	}
}

func (s *ProductType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProductType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProductType(input string) (*ProductType, error) {
	vals := map[string]ProductType{
		"azurereservation":  ProductTypeAzureReservation,
		"azuresubscription": ProductTypeAzureSubscription,
		"department":        ProductTypeDepartment,
		"saas":              ProductTypeSAAS,
		"savingsplan":       ProductTypeSavingsPlan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductType(input)
	return &out, nil
}

type SupportedAccountType string

const (
	SupportedAccountTypeEnterprise SupportedAccountType = "Enterprise"
	SupportedAccountTypeIndividual SupportedAccountType = "Individual"
	SupportedAccountTypeNone       SupportedAccountType = "None"
	SupportedAccountTypePartner    SupportedAccountType = "Partner"
)

func PossibleValuesForSupportedAccountType() []string {
	return []string{
		string(SupportedAccountTypeEnterprise),
		string(SupportedAccountTypeIndividual),
		string(SupportedAccountTypeNone),
		string(SupportedAccountTypePartner),
	}
}

func (s *SupportedAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSupportedAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSupportedAccountType(input string) (*SupportedAccountType, error) {
	vals := map[string]SupportedAccountType{
		"enterprise": SupportedAccountTypeEnterprise,
		"individual": SupportedAccountTypeIndividual,
		"none":       SupportedAccountTypeNone,
		"partner":    SupportedAccountTypePartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportedAccountType(input)
	return &out, nil
}

type TransferStatus string

const (
	TransferStatusCanceled            TransferStatus = "Canceled"
	TransferStatusCompleted           TransferStatus = "Completed"
	TransferStatusCompletedWithErrors TransferStatus = "CompletedWithErrors"
	TransferStatusDeclined            TransferStatus = "Declined"
	TransferStatusExpired             TransferStatus = "Expired"
	TransferStatusFailed              TransferStatus = "Failed"
	TransferStatusInProgress          TransferStatus = "InProgress"
	TransferStatusPending             TransferStatus = "Pending"
)

func PossibleValuesForTransferStatus() []string {
	return []string{
		string(TransferStatusCanceled),
		string(TransferStatusCompleted),
		string(TransferStatusCompletedWithErrors),
		string(TransferStatusDeclined),
		string(TransferStatusExpired),
		string(TransferStatusFailed),
		string(TransferStatusInProgress),
		string(TransferStatusPending),
	}
}

func (s *TransferStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransferStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransferStatus(input string) (*TransferStatus, error) {
	vals := map[string]TransferStatus{
		"canceled":            TransferStatusCanceled,
		"completed":           TransferStatusCompleted,
		"completedwitherrors": TransferStatusCompletedWithErrors,
		"declined":            TransferStatusDeclined,
		"expired":             TransferStatusExpired,
		"failed":              TransferStatusFailed,
		"inprogress":          TransferStatusInProgress,
		"pending":             TransferStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransferStatus(input)
	return &out, nil
}
