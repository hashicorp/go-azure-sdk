package transfers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
)

func PossibleValuesForProductType() []string {
	return []string{
		string(ProductTypeAzureReservation),
		string(ProductTypeAzureSubscription),
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
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductType(input)
	return &out, nil
}

type TransferStatus string

const (
	TransferStatusCanceled            TransferStatus = "Canceled"
	TransferStatusCompleted           TransferStatus = "Completed"
	TransferStatusCompletedWithErrors TransferStatus = "CompletedWithErrors"
	TransferStatusDeclined            TransferStatus = "Declined"
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
