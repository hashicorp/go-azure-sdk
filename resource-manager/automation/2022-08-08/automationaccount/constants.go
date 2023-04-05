package automationaccount

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationAccountState string

const (
	AutomationAccountStateOk          AutomationAccountState = "Ok"
	AutomationAccountStateSuspended   AutomationAccountState = "Suspended"
	AutomationAccountStateUnavailable AutomationAccountState = "Unavailable"
)

func PossibleValuesForAutomationAccountState() []string {
	return []string{
		string(AutomationAccountStateOk),
		string(AutomationAccountStateSuspended),
		string(AutomationAccountStateUnavailable),
	}
}

func (s *AutomationAccountState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAutomationAccountState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AutomationAccountState(decoded)
	return nil
}

type EncryptionKeySourceType string

const (
	EncryptionKeySourceTypeMicrosoftPointAutomation EncryptionKeySourceType = "Microsoft.Automation"
	EncryptionKeySourceTypeMicrosoftPointKeyvault   EncryptionKeySourceType = "Microsoft.Keyvault"
)

func PossibleValuesForEncryptionKeySourceType() []string {
	return []string{
		string(EncryptionKeySourceTypeMicrosoftPointAutomation),
		string(EncryptionKeySourceTypeMicrosoftPointKeyvault),
	}
}

func (s *EncryptionKeySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEncryptionKeySourceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EncryptionKeySourceType(decoded)
	return nil
}

type SkuNameEnum string

const (
	SkuNameEnumBasic SkuNameEnum = "Basic"
	SkuNameEnumFree  SkuNameEnum = "Free"
)

func PossibleValuesForSkuNameEnum() []string {
	return []string{
		string(SkuNameEnumBasic),
		string(SkuNameEnumFree),
	}
}

func (s *SkuNameEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSkuNameEnum() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SkuNameEnum(decoded)
	return nil
}
