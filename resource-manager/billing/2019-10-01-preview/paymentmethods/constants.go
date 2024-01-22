package paymentmethods

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentMethodType string

const (
	PaymentMethodTypeChequeWire PaymentMethodType = "ChequeWire"
	PaymentMethodTypeCredits    PaymentMethodType = "Credits"
)

func PossibleValuesForPaymentMethodType() []string {
	return []string{
		string(PaymentMethodTypeChequeWire),
		string(PaymentMethodTypeCredits),
	}
}

func (s *PaymentMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePaymentMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePaymentMethodType(input string) (*PaymentMethodType, error) {
	vals := map[string]PaymentMethodType{
		"chequewire": PaymentMethodTypeChequeWire,
		"credits":    PaymentMethodTypeCredits,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentMethodType(input)
	return &out, nil
}
