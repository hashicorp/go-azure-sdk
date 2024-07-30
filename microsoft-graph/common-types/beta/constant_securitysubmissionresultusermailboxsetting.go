package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResultUserMailboxSetting string

const (
	SecuritySubmissionResultUserMailboxSetting_CustomRule                             SecuritySubmissionResultUserMailboxSetting = "customRule"
	SecuritySubmissionResultUserMailboxSetting_Exclusive                              SecuritySubmissionResultUserMailboxSetting = "exclusive"
	SecuritySubmissionResultUserMailboxSetting_FromFirstTimeSender                    SecuritySubmissionResultUserMailboxSetting = "fromFirstTimeSender"
	SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressBlockList        SecuritySubmissionResultUserMailboxSetting = "isFromAddressInAddressBlockList"
	SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressBook             SecuritySubmissionResultUserMailboxSetting = "isFromAddressInAddressBook"
	SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressImplicitJunkList SecuritySubmissionResultUserMailboxSetting = "isFromAddressInAddressImplicitJunkList"
	SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressImplicitSafeList SecuritySubmissionResultUserMailboxSetting = "isFromAddressInAddressImplicitSafeList"
	SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressSafeList         SecuritySubmissionResultUserMailboxSetting = "isFromAddressInAddressSafeList"
	SecuritySubmissionResultUserMailboxSetting_IsFromDomainInDomainBlockList          SecuritySubmissionResultUserMailboxSetting = "isFromDomainInDomainBlockList"
	SecuritySubmissionResultUserMailboxSetting_IsFromDomainInDomainSafeList           SecuritySubmissionResultUserMailboxSetting = "isFromDomainInDomainSafeList"
	SecuritySubmissionResultUserMailboxSetting_IsJunkMailRuleEnabled                  SecuritySubmissionResultUserMailboxSetting = "isJunkMailRuleEnabled"
	SecuritySubmissionResultUserMailboxSetting_IsRecipientInRecipientSafeList         SecuritySubmissionResultUserMailboxSetting = "isRecipientInRecipientSafeList"
	SecuritySubmissionResultUserMailboxSetting_JunkMailDeletion                       SecuritySubmissionResultUserMailboxSetting = "junkMailDeletion"
	SecuritySubmissionResultUserMailboxSetting_JunkMailRule                           SecuritySubmissionResultUserMailboxSetting = "junkMailRule"
	SecuritySubmissionResultUserMailboxSetting_None                                   SecuritySubmissionResultUserMailboxSetting = "none"
	SecuritySubmissionResultUserMailboxSetting_PriorSeenPass                          SecuritySubmissionResultUserMailboxSetting = "priorSeenPass"
	SecuritySubmissionResultUserMailboxSetting_SenderAuthenticationSucceeded          SecuritySubmissionResultUserMailboxSetting = "senderAuthenticationSucceeded"
	SecuritySubmissionResultUserMailboxSetting_SenderPraPresent                       SecuritySubmissionResultUserMailboxSetting = "senderPraPresent"
)

func PossibleValuesForSecuritySubmissionResultUserMailboxSetting() []string {
	return []string{
		string(SecuritySubmissionResultUserMailboxSetting_CustomRule),
		string(SecuritySubmissionResultUserMailboxSetting_Exclusive),
		string(SecuritySubmissionResultUserMailboxSetting_FromFirstTimeSender),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressBlockList),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressBook),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressImplicitJunkList),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressImplicitSafeList),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressSafeList),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromDomainInDomainBlockList),
		string(SecuritySubmissionResultUserMailboxSetting_IsFromDomainInDomainSafeList),
		string(SecuritySubmissionResultUserMailboxSetting_IsJunkMailRuleEnabled),
		string(SecuritySubmissionResultUserMailboxSetting_IsRecipientInRecipientSafeList),
		string(SecuritySubmissionResultUserMailboxSetting_JunkMailDeletion),
		string(SecuritySubmissionResultUserMailboxSetting_JunkMailRule),
		string(SecuritySubmissionResultUserMailboxSetting_None),
		string(SecuritySubmissionResultUserMailboxSetting_PriorSeenPass),
		string(SecuritySubmissionResultUserMailboxSetting_SenderAuthenticationSucceeded),
		string(SecuritySubmissionResultUserMailboxSetting_SenderPraPresent),
	}
}

func (s *SecuritySubmissionResultUserMailboxSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionResultUserMailboxSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionResultUserMailboxSetting(input string) (*SecuritySubmissionResultUserMailboxSetting, error) {
	vals := map[string]SecuritySubmissionResultUserMailboxSetting{
		"customrule":                             SecuritySubmissionResultUserMailboxSetting_CustomRule,
		"exclusive":                              SecuritySubmissionResultUserMailboxSetting_Exclusive,
		"fromfirsttimesender":                    SecuritySubmissionResultUserMailboxSetting_FromFirstTimeSender,
		"isfromaddressinaddressblocklist":        SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressBlockList,
		"isfromaddressinaddressbook":             SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressBook,
		"isfromaddressinaddressimplicitjunklist": SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressImplicitJunkList,
		"isfromaddressinaddressimplicitsafelist": SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressImplicitSafeList,
		"isfromaddressinaddresssafelist":         SecuritySubmissionResultUserMailboxSetting_IsFromAddressInAddressSafeList,
		"isfromdomainindomainblocklist":          SecuritySubmissionResultUserMailboxSetting_IsFromDomainInDomainBlockList,
		"isfromdomainindomainsafelist":           SecuritySubmissionResultUserMailboxSetting_IsFromDomainInDomainSafeList,
		"isjunkmailruleenabled":                  SecuritySubmissionResultUserMailboxSetting_IsJunkMailRuleEnabled,
		"isrecipientinrecipientsafelist":         SecuritySubmissionResultUserMailboxSetting_IsRecipientInRecipientSafeList,
		"junkmaildeletion":                       SecuritySubmissionResultUserMailboxSetting_JunkMailDeletion,
		"junkmailrule":                           SecuritySubmissionResultUserMailboxSetting_JunkMailRule,
		"none":                                   SecuritySubmissionResultUserMailboxSetting_None,
		"priorseenpass":                          SecuritySubmissionResultUserMailboxSetting_PriorSeenPass,
		"senderauthenticationsucceeded":          SecuritySubmissionResultUserMailboxSetting_SenderAuthenticationSucceeded,
		"senderprapresent":                       SecuritySubmissionResultUserMailboxSetting_SenderPraPresent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionResultUserMailboxSetting(input)
	return &out, nil
}
