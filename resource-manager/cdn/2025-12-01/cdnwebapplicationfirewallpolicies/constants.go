package cdnwebapplicationfirewallpolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionType string

const (
	ActionTypeAllow    ActionType = "Allow"
	ActionTypeBlock    ActionType = "Block"
	ActionTypeLog      ActionType = "Log"
	ActionTypeRedirect ActionType = "Redirect"
)

func PossibleValuesForActionType() []string {
	return []string{
		string(ActionTypeAllow),
		string(ActionTypeBlock),
		string(ActionTypeLog),
		string(ActionTypeRedirect),
	}
}

func (s *ActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActionType(input string) (*ActionType, error) {
	vals := map[string]ActionType{
		"allow":    ActionTypeAllow,
		"block":    ActionTypeBlock,
		"log":      ActionTypeLog,
		"redirect": ActionTypeRedirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionType(input)
	return &out, nil
}

type CustomRuleEnabledState string

const (
	CustomRuleEnabledStateDisabled CustomRuleEnabledState = "Disabled"
	CustomRuleEnabledStateEnabled  CustomRuleEnabledState = "Enabled"
)

func PossibleValuesForCustomRuleEnabledState() []string {
	return []string{
		string(CustomRuleEnabledStateDisabled),
		string(CustomRuleEnabledStateEnabled),
	}
}

func (s *CustomRuleEnabledState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomRuleEnabledState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomRuleEnabledState(input string) (*CustomRuleEnabledState, error) {
	vals := map[string]CustomRuleEnabledState{
		"disabled": CustomRuleEnabledStateDisabled,
		"enabled":  CustomRuleEnabledStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomRuleEnabledState(input)
	return &out, nil
}

type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledStateDisabled ManagedRuleEnabledState = "Disabled"
	ManagedRuleEnabledStateEnabled  ManagedRuleEnabledState = "Enabled"
)

func PossibleValuesForManagedRuleEnabledState() []string {
	return []string{
		string(ManagedRuleEnabledStateDisabled),
		string(ManagedRuleEnabledStateEnabled),
	}
}

func (s *ManagedRuleEnabledState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedRuleEnabledState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedRuleEnabledState(input string) (*ManagedRuleEnabledState, error) {
	vals := map[string]ManagedRuleEnabledState{
		"disabled": ManagedRuleEnabledStateDisabled,
		"enabled":  ManagedRuleEnabledStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedRuleEnabledState(input)
	return &out, nil
}

type Operator string

const (
	OperatorAny                Operator = "Any"
	OperatorBeginsWith         Operator = "BeginsWith"
	OperatorContains           Operator = "Contains"
	OperatorEndsWith           Operator = "EndsWith"
	OperatorEqual              Operator = "Equal"
	OperatorGeoMatch           Operator = "GeoMatch"
	OperatorGreaterThan        Operator = "GreaterThan"
	OperatorGreaterThanOrEqual Operator = "GreaterThanOrEqual"
	OperatorIPMatch            Operator = "IPMatch"
	OperatorLessThan           Operator = "LessThan"
	OperatorLessThanOrEqual    Operator = "LessThanOrEqual"
	OperatorRegEx              Operator = "RegEx"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(OperatorAny),
		string(OperatorBeginsWith),
		string(OperatorContains),
		string(OperatorEndsWith),
		string(OperatorEqual),
		string(OperatorGeoMatch),
		string(OperatorGreaterThan),
		string(OperatorGreaterThanOrEqual),
		string(OperatorIPMatch),
		string(OperatorLessThan),
		string(OperatorLessThanOrEqual),
		string(OperatorRegEx),
	}
}

func (s *Operator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperator(input string) (*Operator, error) {
	vals := map[string]Operator{
		"any":                OperatorAny,
		"beginswith":         OperatorBeginsWith,
		"contains":           OperatorContains,
		"endswith":           OperatorEndsWith,
		"equal":              OperatorEqual,
		"geomatch":           OperatorGeoMatch,
		"greaterthan":        OperatorGreaterThan,
		"greaterthanorequal": OperatorGreaterThanOrEqual,
		"ipmatch":            OperatorIPMatch,
		"lessthan":           OperatorLessThan,
		"lessthanorequal":    OperatorLessThanOrEqual,
		"regex":              OperatorRegEx,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Operator(input)
	return &out, nil
}

type PolicyEnabledState string

const (
	PolicyEnabledStateDisabled PolicyEnabledState = "Disabled"
	PolicyEnabledStateEnabled  PolicyEnabledState = "Enabled"
)

func PossibleValuesForPolicyEnabledState() []string {
	return []string{
		string(PolicyEnabledStateDisabled),
		string(PolicyEnabledStateEnabled),
	}
}

func (s *PolicyEnabledState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyEnabledState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyEnabledState(input string) (*PolicyEnabledState, error) {
	vals := map[string]PolicyEnabledState{
		"disabled": PolicyEnabledStateDisabled,
		"enabled":  PolicyEnabledStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyEnabledState(input)
	return &out, nil
}

type PolicyMode string

const (
	PolicyModeDetection  PolicyMode = "Detection"
	PolicyModePrevention PolicyMode = "Prevention"
)

func PossibleValuesForPolicyMode() []string {
	return []string{
		string(PolicyModeDetection),
		string(PolicyModePrevention),
	}
}

func (s *PolicyMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyMode(input string) (*PolicyMode, error) {
	vals := map[string]PolicyMode{
		"detection":  PolicyModeDetection,
		"prevention": PolicyModePrevention,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyMode(input)
	return &out, nil
}

type PolicyResourceState string

const (
	PolicyResourceStateCreating  PolicyResourceState = "Creating"
	PolicyResourceStateDeleting  PolicyResourceState = "Deleting"
	PolicyResourceStateDisabled  PolicyResourceState = "Disabled"
	PolicyResourceStateDisabling PolicyResourceState = "Disabling"
	PolicyResourceStateEnabled   PolicyResourceState = "Enabled"
	PolicyResourceStateEnabling  PolicyResourceState = "Enabling"
)

func PossibleValuesForPolicyResourceState() []string {
	return []string{
		string(PolicyResourceStateCreating),
		string(PolicyResourceStateDeleting),
		string(PolicyResourceStateDisabled),
		string(PolicyResourceStateDisabling),
		string(PolicyResourceStateEnabled),
		string(PolicyResourceStateEnabling),
	}
}

func (s *PolicyResourceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyResourceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyResourceState(input string) (*PolicyResourceState, error) {
	vals := map[string]PolicyResourceState{
		"creating":  PolicyResourceStateCreating,
		"deleting":  PolicyResourceStateDeleting,
		"disabled":  PolicyResourceStateDisabled,
		"disabling": PolicyResourceStateDisabling,
		"enabled":   PolicyResourceStateEnabled,
		"enabling":  PolicyResourceStateEnabling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyResourceState(input)
	return &out, nil
}

type PolicySettingsDefaultCustomBlockResponseStatusCode float64

const (
	PolicySettingsDefaultCustomBlockResponseStatusCodeFourTwoNine   PolicySettingsDefaultCustomBlockResponseStatusCode = 429
	PolicySettingsDefaultCustomBlockResponseStatusCodeFourZeroFive  PolicySettingsDefaultCustomBlockResponseStatusCode = 405
	PolicySettingsDefaultCustomBlockResponseStatusCodeFourZeroSix   PolicySettingsDefaultCustomBlockResponseStatusCode = 406
	PolicySettingsDefaultCustomBlockResponseStatusCodeFourZeroThree PolicySettingsDefaultCustomBlockResponseStatusCode = 403
	PolicySettingsDefaultCustomBlockResponseStatusCodeTwoHundred    PolicySettingsDefaultCustomBlockResponseStatusCode = 200
)

func PossibleValuesForPolicySettingsDefaultCustomBlockResponseStatusCode() []float64 {
	return []float64{
		float64(PolicySettingsDefaultCustomBlockResponseStatusCodeFourTwoNine),
		float64(PolicySettingsDefaultCustomBlockResponseStatusCodeFourZeroFive),
		float64(PolicySettingsDefaultCustomBlockResponseStatusCodeFourZeroSix),
		float64(PolicySettingsDefaultCustomBlockResponseStatusCodeFourZeroThree),
		float64(PolicySettingsDefaultCustomBlockResponseStatusCodeTwoHundred),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameCustomVerizon                             SkuName = "Custom_Verizon"
	SkuNamePremiumAzureFrontDoor                     SkuName = "Premium_AzureFrontDoor"
	SkuNamePremiumVerizon                            SkuName = "Premium_Verizon"
	SkuNameStandardAkamai                            SkuName = "Standard_Akamai"
	SkuNameStandardAvgBandWidthChinaCdn              SkuName = "Standard_AvgBandWidth_ChinaCdn"
	SkuNameStandardAzureFrontDoor                    SkuName = "Standard_AzureFrontDoor"
	SkuNameStandardChinaCdn                          SkuName = "Standard_ChinaCdn"
	SkuNameStandardMicrosoft                         SkuName = "Standard_Microsoft"
	SkuNameStandardNineFiveFiveBandWidthChinaCdn     SkuName = "Standard_955BandWidth_ChinaCdn"
	SkuNameStandardPlusAvgBandWidthChinaCdn          SkuName = "StandardPlus_AvgBandWidth_ChinaCdn"
	SkuNameStandardPlusChinaCdn                      SkuName = "StandardPlus_ChinaCdn"
	SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn SkuName = "StandardPlus_955BandWidth_ChinaCdn"
	SkuNameStandardVerizon                           SkuName = "Standard_Verizon"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameCustomVerizon),
		string(SkuNamePremiumAzureFrontDoor),
		string(SkuNamePremiumVerizon),
		string(SkuNameStandardAkamai),
		string(SkuNameStandardAvgBandWidthChinaCdn),
		string(SkuNameStandardAzureFrontDoor),
		string(SkuNameStandardChinaCdn),
		string(SkuNameStandardMicrosoft),
		string(SkuNameStandardNineFiveFiveBandWidthChinaCdn),
		string(SkuNameStandardPlusAvgBandWidthChinaCdn),
		string(SkuNameStandardPlusChinaCdn),
		string(SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn),
		string(SkuNameStandardVerizon),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"custom_verizon":                     SkuNameCustomVerizon,
		"premium_azurefrontdoor":             SkuNamePremiumAzureFrontDoor,
		"premium_verizon":                    SkuNamePremiumVerizon,
		"standard_akamai":                    SkuNameStandardAkamai,
		"standard_avgbandwidth_chinacdn":     SkuNameStandardAvgBandWidthChinaCdn,
		"standard_azurefrontdoor":            SkuNameStandardAzureFrontDoor,
		"standard_chinacdn":                  SkuNameStandardChinaCdn,
		"standard_microsoft":                 SkuNameStandardMicrosoft,
		"standard_955bandwidth_chinacdn":     SkuNameStandardNineFiveFiveBandWidthChinaCdn,
		"standardplus_avgbandwidth_chinacdn": SkuNameStandardPlusAvgBandWidthChinaCdn,
		"standardplus_chinacdn":              SkuNameStandardPlusChinaCdn,
		"standardplus_955bandwidth_chinacdn": SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn,
		"standard_verizon":                   SkuNameStandardVerizon,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type TransformType string

const (
	TransformTypeLowercase   TransformType = "Lowercase"
	TransformTypeRemoveNulls TransformType = "RemoveNulls"
	TransformTypeTrim        TransformType = "Trim"
	TransformTypeURLDecode   TransformType = "UrlDecode"
	TransformTypeURLEncode   TransformType = "UrlEncode"
	TransformTypeUppercase   TransformType = "Uppercase"
)

func PossibleValuesForTransformType() []string {
	return []string{
		string(TransformTypeLowercase),
		string(TransformTypeRemoveNulls),
		string(TransformTypeTrim),
		string(TransformTypeURLDecode),
		string(TransformTypeURLEncode),
		string(TransformTypeUppercase),
	}
}

func (s *TransformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransformType(input string) (*TransformType, error) {
	vals := map[string]TransformType{
		"lowercase":   TransformTypeLowercase,
		"removenulls": TransformTypeRemoveNulls,
		"trim":        TransformTypeTrim,
		"urldecode":   TransformTypeURLDecode,
		"urlencode":   TransformTypeURLEncode,
		"uppercase":   TransformTypeUppercase,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransformType(input)
	return &out, nil
}

type WafMatchVariable string

const (
	WafMatchVariableCookies       WafMatchVariable = "Cookies"
	WafMatchVariablePostArgs      WafMatchVariable = "PostArgs"
	WafMatchVariableQueryString   WafMatchVariable = "QueryString"
	WafMatchVariableRemoteAddr    WafMatchVariable = "RemoteAddr"
	WafMatchVariableRequestBody   WafMatchVariable = "RequestBody"
	WafMatchVariableRequestHeader WafMatchVariable = "RequestHeader"
	WafMatchVariableRequestMethod WafMatchVariable = "RequestMethod"
	WafMatchVariableRequestUri    WafMatchVariable = "RequestUri"
	WafMatchVariableSocketAddr    WafMatchVariable = "SocketAddr"
)

func PossibleValuesForWafMatchVariable() []string {
	return []string{
		string(WafMatchVariableCookies),
		string(WafMatchVariablePostArgs),
		string(WafMatchVariableQueryString),
		string(WafMatchVariableRemoteAddr),
		string(WafMatchVariableRequestBody),
		string(WafMatchVariableRequestHeader),
		string(WafMatchVariableRequestMethod),
		string(WafMatchVariableRequestUri),
		string(WafMatchVariableSocketAddr),
	}
}

func (s *WafMatchVariable) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWafMatchVariable(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWafMatchVariable(input string) (*WafMatchVariable, error) {
	vals := map[string]WafMatchVariable{
		"cookies":       WafMatchVariableCookies,
		"postargs":      WafMatchVariablePostArgs,
		"querystring":   WafMatchVariableQueryString,
		"remoteaddr":    WafMatchVariableRemoteAddr,
		"requestbody":   WafMatchVariableRequestBody,
		"requestheader": WafMatchVariableRequestHeader,
		"requestmethod": WafMatchVariableRequestMethod,
		"requesturi":    WafMatchVariableRequestUri,
		"socketaddr":    WafMatchVariableSocketAddr,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WafMatchVariable(input)
	return &out, nil
}
