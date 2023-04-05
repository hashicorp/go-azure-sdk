package webapplicationfirewallpolicies

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

type ManagedRuleExclusionMatchVariable string

const (
	ManagedRuleExclusionMatchVariableQueryStringArgNames     ManagedRuleExclusionMatchVariable = "QueryStringArgNames"
	ManagedRuleExclusionMatchVariableRequestBodyPostArgNames ManagedRuleExclusionMatchVariable = "RequestBodyPostArgNames"
	ManagedRuleExclusionMatchVariableRequestCookieNames      ManagedRuleExclusionMatchVariable = "RequestCookieNames"
	ManagedRuleExclusionMatchVariableRequestHeaderNames      ManagedRuleExclusionMatchVariable = "RequestHeaderNames"
)

func PossibleValuesForManagedRuleExclusionMatchVariable() []string {
	return []string{
		string(ManagedRuleExclusionMatchVariableQueryStringArgNames),
		string(ManagedRuleExclusionMatchVariableRequestBodyPostArgNames),
		string(ManagedRuleExclusionMatchVariableRequestCookieNames),
		string(ManagedRuleExclusionMatchVariableRequestHeaderNames),
	}
}

type ManagedRuleExclusionSelectorMatchOperator string

const (
	ManagedRuleExclusionSelectorMatchOperatorContains   ManagedRuleExclusionSelectorMatchOperator = "Contains"
	ManagedRuleExclusionSelectorMatchOperatorEndsWith   ManagedRuleExclusionSelectorMatchOperator = "EndsWith"
	ManagedRuleExclusionSelectorMatchOperatorEquals     ManagedRuleExclusionSelectorMatchOperator = "Equals"
	ManagedRuleExclusionSelectorMatchOperatorEqualsAny  ManagedRuleExclusionSelectorMatchOperator = "EqualsAny"
	ManagedRuleExclusionSelectorMatchOperatorStartsWith ManagedRuleExclusionSelectorMatchOperator = "StartsWith"
)

func PossibleValuesForManagedRuleExclusionSelectorMatchOperator() []string {
	return []string{
		string(ManagedRuleExclusionSelectorMatchOperatorContains),
		string(ManagedRuleExclusionSelectorMatchOperatorEndsWith),
		string(ManagedRuleExclusionSelectorMatchOperatorEquals),
		string(ManagedRuleExclusionSelectorMatchOperatorEqualsAny),
		string(ManagedRuleExclusionSelectorMatchOperatorStartsWith),
	}
}

type MatchVariable string

const (
	MatchVariableCookies       MatchVariable = "Cookies"
	MatchVariablePostArgs      MatchVariable = "PostArgs"
	MatchVariableQueryString   MatchVariable = "QueryString"
	MatchVariableRemoteAddr    MatchVariable = "RemoteAddr"
	MatchVariableRequestBody   MatchVariable = "RequestBody"
	MatchVariableRequestHeader MatchVariable = "RequestHeader"
	MatchVariableRequestMethod MatchVariable = "RequestMethod"
	MatchVariableRequestUri    MatchVariable = "RequestUri"
	MatchVariableSocketAddr    MatchVariable = "SocketAddr"
)

func PossibleValuesForMatchVariable() []string {
	return []string{
		string(MatchVariableCookies),
		string(MatchVariablePostArgs),
		string(MatchVariableQueryString),
		string(MatchVariableRemoteAddr),
		string(MatchVariableRequestBody),
		string(MatchVariableRequestHeader),
		string(MatchVariableRequestMethod),
		string(MatchVariableRequestUri),
		string(MatchVariableSocketAddr),
	}
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

type RuleType string

const (
	RuleTypeMatchRule     RuleType = "MatchRule"
	RuleTypeRateLimitRule RuleType = "RateLimitRule"
)

func PossibleValuesForRuleType() []string {
	return []string{
		string(RuleTypeMatchRule),
		string(RuleTypeRateLimitRule),
	}
}

type TransformType string

const (
	TransformTypeLowercase   TransformType = "Lowercase"
	TransformTypeRemoveNulls TransformType = "RemoveNulls"
	TransformTypeTrim        TransformType = "Trim"
	TransformTypeUppercase   TransformType = "Uppercase"
	TransformTypeUrlDecode   TransformType = "UrlDecode"
	TransformTypeUrlEncode   TransformType = "UrlEncode"
)

func PossibleValuesForTransformType() []string {
	return []string{
		string(TransformTypeLowercase),
		string(TransformTypeRemoveNulls),
		string(TransformTypeTrim),
		string(TransformTypeUppercase),
		string(TransformTypeUrlDecode),
		string(TransformTypeUrlEncode),
	}
}
