package paymentmethods

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentMethodFamily string

const (
	PaymentMethodFamilyCheckWire  PaymentMethodFamily = "CheckWire"
	PaymentMethodFamilyCreditCard PaymentMethodFamily = "CreditCard"
)

func PossibleValuesForPaymentMethodFamily() []string {
	return []string{
		string(PaymentMethodFamilyCheckWire),
		string(PaymentMethodFamilyCreditCard),
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
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentMethodFamily(input)
	return &out, nil
}

type PaymentMethodStatus string

const (
	PaymentMethodStatusActive   PaymentMethodStatus = "active"
	PaymentMethodStatusInactive PaymentMethodStatus = "inactive"
)

func PossibleValuesForPaymentMethodStatus() []string {
	return []string{
		string(PaymentMethodStatusActive),
		string(PaymentMethodStatusInactive),
	}
}

func (s *PaymentMethodStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePaymentMethodStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePaymentMethodStatus(input string) (*PaymentMethodStatus, error) {
	vals := map[string]PaymentMethodStatus{
		"active":   PaymentMethodStatusActive,
		"inactive": PaymentMethodStatusInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentMethodStatus(input)
	return &out, nil
}
