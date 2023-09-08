package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestDestinationRoutingReason string

const (
	EmailFileAssessmentRequestDestinationRoutingReasonadvancedSpamFiltering EmailFileAssessmentRequestDestinationRoutingReason = "AdvancedSpamFiltering"
	EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToDeleted    EmailFileAssessmentRequestDestinationRoutingReason = "AutoPurgeToDeleted"
	EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToInbox      EmailFileAssessmentRequestDestinationRoutingReason = "AutoPurgeToInbox"
	EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToJunk       EmailFileAssessmentRequestDestinationRoutingReason = "AutoPurgeToJunk"
	EmailFileAssessmentRequestDestinationRoutingReasonblockedSender         EmailFileAssessmentRequestDestinationRoutingReason = "BlockedSender"
	EmailFileAssessmentRequestDestinationRoutingReasondomainAllowList       EmailFileAssessmentRequestDestinationRoutingReason = "DomainAllowList"
	EmailFileAssessmentRequestDestinationRoutingReasondomainBlockList       EmailFileAssessmentRequestDestinationRoutingReason = "DomainBlockList"
	EmailFileAssessmentRequestDestinationRoutingReasonfirstTimeSender       EmailFileAssessmentRequestDestinationRoutingReason = "FirstTimeSender"
	EmailFileAssessmentRequestDestinationRoutingReasonjunk                  EmailFileAssessmentRequestDestinationRoutingReason = "Junk"
	EmailFileAssessmentRequestDestinationRoutingReasonmailFlowRule          EmailFileAssessmentRequestDestinationRoutingReason = "MailFlowRule"
	EmailFileAssessmentRequestDestinationRoutingReasonnone                  EmailFileAssessmentRequestDestinationRoutingReason = "None"
	EmailFileAssessmentRequestDestinationRoutingReasonnotInAddressBook      EmailFileAssessmentRequestDestinationRoutingReason = "NotInAddressBook"
	EmailFileAssessmentRequestDestinationRoutingReasonnotJunk               EmailFileAssessmentRequestDestinationRoutingReason = "NotJunk"
	EmailFileAssessmentRequestDestinationRoutingReasonoutbound              EmailFileAssessmentRequestDestinationRoutingReason = "Outbound"
	EmailFileAssessmentRequestDestinationRoutingReasonsafeSender            EmailFileAssessmentRequestDestinationRoutingReason = "SafeSender"
)

func PossibleValuesForEmailFileAssessmentRequestDestinationRoutingReason() []string {
	return []string{
		string(EmailFileAssessmentRequestDestinationRoutingReasonadvancedSpamFiltering),
		string(EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToDeleted),
		string(EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToInbox),
		string(EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToJunk),
		string(EmailFileAssessmentRequestDestinationRoutingReasonblockedSender),
		string(EmailFileAssessmentRequestDestinationRoutingReasondomainAllowList),
		string(EmailFileAssessmentRequestDestinationRoutingReasondomainBlockList),
		string(EmailFileAssessmentRequestDestinationRoutingReasonfirstTimeSender),
		string(EmailFileAssessmentRequestDestinationRoutingReasonjunk),
		string(EmailFileAssessmentRequestDestinationRoutingReasonmailFlowRule),
		string(EmailFileAssessmentRequestDestinationRoutingReasonnone),
		string(EmailFileAssessmentRequestDestinationRoutingReasonnotInAddressBook),
		string(EmailFileAssessmentRequestDestinationRoutingReasonnotJunk),
		string(EmailFileAssessmentRequestDestinationRoutingReasonoutbound),
		string(EmailFileAssessmentRequestDestinationRoutingReasonsafeSender),
	}
}

func parseEmailFileAssessmentRequestDestinationRoutingReason(input string) (*EmailFileAssessmentRequestDestinationRoutingReason, error) {
	vals := map[string]EmailFileAssessmentRequestDestinationRoutingReason{
		"advancedspamfiltering": EmailFileAssessmentRequestDestinationRoutingReasonadvancedSpamFiltering,
		"autopurgetodeleted":    EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToDeleted,
		"autopurgetoinbox":      EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToInbox,
		"autopurgetojunk":       EmailFileAssessmentRequestDestinationRoutingReasonautoPurgeToJunk,
		"blockedsender":         EmailFileAssessmentRequestDestinationRoutingReasonblockedSender,
		"domainallowlist":       EmailFileAssessmentRequestDestinationRoutingReasondomainAllowList,
		"domainblocklist":       EmailFileAssessmentRequestDestinationRoutingReasondomainBlockList,
		"firsttimesender":       EmailFileAssessmentRequestDestinationRoutingReasonfirstTimeSender,
		"junk":                  EmailFileAssessmentRequestDestinationRoutingReasonjunk,
		"mailflowrule":          EmailFileAssessmentRequestDestinationRoutingReasonmailFlowRule,
		"none":                  EmailFileAssessmentRequestDestinationRoutingReasonnone,
		"notinaddressbook":      EmailFileAssessmentRequestDestinationRoutingReasonnotInAddressBook,
		"notjunk":               EmailFileAssessmentRequestDestinationRoutingReasonnotJunk,
		"outbound":              EmailFileAssessmentRequestDestinationRoutingReasonoutbound,
		"safesender":            EmailFileAssessmentRequestDestinationRoutingReasonsafeSender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestDestinationRoutingReason(input)
	return &out, nil
}
