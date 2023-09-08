package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestDestinationRoutingReason string

const (
	MailAssessmentRequestDestinationRoutingReasonadvancedSpamFiltering MailAssessmentRequestDestinationRoutingReason = "AdvancedSpamFiltering"
	MailAssessmentRequestDestinationRoutingReasonautoPurgeToDeleted    MailAssessmentRequestDestinationRoutingReason = "AutoPurgeToDeleted"
	MailAssessmentRequestDestinationRoutingReasonautoPurgeToInbox      MailAssessmentRequestDestinationRoutingReason = "AutoPurgeToInbox"
	MailAssessmentRequestDestinationRoutingReasonautoPurgeToJunk       MailAssessmentRequestDestinationRoutingReason = "AutoPurgeToJunk"
	MailAssessmentRequestDestinationRoutingReasonblockedSender         MailAssessmentRequestDestinationRoutingReason = "BlockedSender"
	MailAssessmentRequestDestinationRoutingReasondomainAllowList       MailAssessmentRequestDestinationRoutingReason = "DomainAllowList"
	MailAssessmentRequestDestinationRoutingReasondomainBlockList       MailAssessmentRequestDestinationRoutingReason = "DomainBlockList"
	MailAssessmentRequestDestinationRoutingReasonfirstTimeSender       MailAssessmentRequestDestinationRoutingReason = "FirstTimeSender"
	MailAssessmentRequestDestinationRoutingReasonjunk                  MailAssessmentRequestDestinationRoutingReason = "Junk"
	MailAssessmentRequestDestinationRoutingReasonmailFlowRule          MailAssessmentRequestDestinationRoutingReason = "MailFlowRule"
	MailAssessmentRequestDestinationRoutingReasonnone                  MailAssessmentRequestDestinationRoutingReason = "None"
	MailAssessmentRequestDestinationRoutingReasonnotInAddressBook      MailAssessmentRequestDestinationRoutingReason = "NotInAddressBook"
	MailAssessmentRequestDestinationRoutingReasonnotJunk               MailAssessmentRequestDestinationRoutingReason = "NotJunk"
	MailAssessmentRequestDestinationRoutingReasonoutbound              MailAssessmentRequestDestinationRoutingReason = "Outbound"
	MailAssessmentRequestDestinationRoutingReasonsafeSender            MailAssessmentRequestDestinationRoutingReason = "SafeSender"
)

func PossibleValuesForMailAssessmentRequestDestinationRoutingReason() []string {
	return []string{
		string(MailAssessmentRequestDestinationRoutingReasonadvancedSpamFiltering),
		string(MailAssessmentRequestDestinationRoutingReasonautoPurgeToDeleted),
		string(MailAssessmentRequestDestinationRoutingReasonautoPurgeToInbox),
		string(MailAssessmentRequestDestinationRoutingReasonautoPurgeToJunk),
		string(MailAssessmentRequestDestinationRoutingReasonblockedSender),
		string(MailAssessmentRequestDestinationRoutingReasondomainAllowList),
		string(MailAssessmentRequestDestinationRoutingReasondomainBlockList),
		string(MailAssessmentRequestDestinationRoutingReasonfirstTimeSender),
		string(MailAssessmentRequestDestinationRoutingReasonjunk),
		string(MailAssessmentRequestDestinationRoutingReasonmailFlowRule),
		string(MailAssessmentRequestDestinationRoutingReasonnone),
		string(MailAssessmentRequestDestinationRoutingReasonnotInAddressBook),
		string(MailAssessmentRequestDestinationRoutingReasonnotJunk),
		string(MailAssessmentRequestDestinationRoutingReasonoutbound),
		string(MailAssessmentRequestDestinationRoutingReasonsafeSender),
	}
}

func parseMailAssessmentRequestDestinationRoutingReason(input string) (*MailAssessmentRequestDestinationRoutingReason, error) {
	vals := map[string]MailAssessmentRequestDestinationRoutingReason{
		"advancedspamfiltering": MailAssessmentRequestDestinationRoutingReasonadvancedSpamFiltering,
		"autopurgetodeleted":    MailAssessmentRequestDestinationRoutingReasonautoPurgeToDeleted,
		"autopurgetoinbox":      MailAssessmentRequestDestinationRoutingReasonautoPurgeToInbox,
		"autopurgetojunk":       MailAssessmentRequestDestinationRoutingReasonautoPurgeToJunk,
		"blockedsender":         MailAssessmentRequestDestinationRoutingReasonblockedSender,
		"domainallowlist":       MailAssessmentRequestDestinationRoutingReasondomainAllowList,
		"domainblocklist":       MailAssessmentRequestDestinationRoutingReasondomainBlockList,
		"firsttimesender":       MailAssessmentRequestDestinationRoutingReasonfirstTimeSender,
		"junk":                  MailAssessmentRequestDestinationRoutingReasonjunk,
		"mailflowrule":          MailAssessmentRequestDestinationRoutingReasonmailFlowRule,
		"none":                  MailAssessmentRequestDestinationRoutingReasonnone,
		"notinaddressbook":      MailAssessmentRequestDestinationRoutingReasonnotInAddressBook,
		"notjunk":               MailAssessmentRequestDestinationRoutingReasonnotJunk,
		"outbound":              MailAssessmentRequestDestinationRoutingReasonoutbound,
		"safesender":            MailAssessmentRequestDestinationRoutingReasonsafeSender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestDestinationRoutingReason(input)
	return &out, nil
}
