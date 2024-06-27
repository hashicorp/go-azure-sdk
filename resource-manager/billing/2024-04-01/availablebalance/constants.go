package availablebalance

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
