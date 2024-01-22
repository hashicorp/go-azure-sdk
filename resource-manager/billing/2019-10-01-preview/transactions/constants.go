package transactions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationType string

const (
	ReservationTypePurchase    ReservationType = "Purchase"
	ReservationTypeUsageCharge ReservationType = "Usage Charge"
)

func PossibleValuesForReservationType() []string {
	return []string{
		string(ReservationTypePurchase),
		string(ReservationTypeUsageCharge),
	}
}

func (s *ReservationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationType(input string) (*ReservationType, error) {
	vals := map[string]ReservationType{
		"purchase":     ReservationTypePurchase,
		"usage charge": ReservationTypeUsageCharge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationType(input)
	return &out, nil
}

type TransactionTypeKind string

const (
	TransactionTypeKindAll         TransactionTypeKind = "all"
	TransactionTypeKindReservation TransactionTypeKind = "reservation"
)

func PossibleValuesForTransactionTypeKind() []string {
	return []string{
		string(TransactionTypeKindAll),
		string(TransactionTypeKindReservation),
	}
}

func (s *TransactionTypeKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransactionTypeKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransactionTypeKind(input string) (*TransactionTypeKind, error) {
	vals := map[string]TransactionTypeKind{
		"all":         TransactionTypeKindAll,
		"reservation": TransactionTypeKindReservation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransactionTypeKind(input)
	return &out, nil
}
