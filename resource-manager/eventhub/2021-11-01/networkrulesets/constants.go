package networkrulesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

func PossibleValuesForNetworkRuleIPAction() []string {
	return []string{
		string(NetworkRuleIPActionAllow),
	}
}

type PublicNetworkAccessFlag string

const (
	PublicNetworkAccessFlagDisabled PublicNetworkAccessFlag = "Disabled"
	PublicNetworkAccessFlagEnabled  PublicNetworkAccessFlag = "Enabled"
)

func PossibleValuesForPublicNetworkAccessFlag() []string {
	return []string{
		string(PublicNetworkAccessFlagDisabled),
		string(PublicNetworkAccessFlagEnabled),
	}
}
