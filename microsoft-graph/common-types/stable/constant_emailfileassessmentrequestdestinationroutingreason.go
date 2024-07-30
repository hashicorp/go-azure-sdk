package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestDestinationRoutingReason string

const (
	EmailFileAssessmentRequestDestinationRoutingReason_AdvancedSpamFiltering EmailFileAssessmentRequestDestinationRoutingReason = "advancedSpamFiltering"
	EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToDeleted    EmailFileAssessmentRequestDestinationRoutingReason = "autoPurgeToDeleted"
	EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToInbox      EmailFileAssessmentRequestDestinationRoutingReason = "autoPurgeToInbox"
	EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToJunk       EmailFileAssessmentRequestDestinationRoutingReason = "autoPurgeToJunk"
	EmailFileAssessmentRequestDestinationRoutingReason_BlockedSender         EmailFileAssessmentRequestDestinationRoutingReason = "blockedSender"
	EmailFileAssessmentRequestDestinationRoutingReason_DomainAllowList       EmailFileAssessmentRequestDestinationRoutingReason = "domainAllowList"
	EmailFileAssessmentRequestDestinationRoutingReason_DomainBlockList       EmailFileAssessmentRequestDestinationRoutingReason = "domainBlockList"
	EmailFileAssessmentRequestDestinationRoutingReason_FirstTimeSender       EmailFileAssessmentRequestDestinationRoutingReason = "firstTimeSender"
	EmailFileAssessmentRequestDestinationRoutingReason_Junk                  EmailFileAssessmentRequestDestinationRoutingReason = "junk"
	EmailFileAssessmentRequestDestinationRoutingReason_MailFlowRule          EmailFileAssessmentRequestDestinationRoutingReason = "mailFlowRule"
	EmailFileAssessmentRequestDestinationRoutingReason_None                  EmailFileAssessmentRequestDestinationRoutingReason = "none"
	EmailFileAssessmentRequestDestinationRoutingReason_NotInAddressBook      EmailFileAssessmentRequestDestinationRoutingReason = "notInAddressBook"
	EmailFileAssessmentRequestDestinationRoutingReason_NotJunk               EmailFileAssessmentRequestDestinationRoutingReason = "notJunk"
	EmailFileAssessmentRequestDestinationRoutingReason_Outbound              EmailFileAssessmentRequestDestinationRoutingReason = "outbound"
	EmailFileAssessmentRequestDestinationRoutingReason_SafeSender            EmailFileAssessmentRequestDestinationRoutingReason = "safeSender"
)

func PossibleValuesForEmailFileAssessmentRequestDestinationRoutingReason() []string {
	return []string{
		string(EmailFileAssessmentRequestDestinationRoutingReason_AdvancedSpamFiltering),
		string(EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToDeleted),
		string(EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToInbox),
		string(EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToJunk),
		string(EmailFileAssessmentRequestDestinationRoutingReason_BlockedSender),
		string(EmailFileAssessmentRequestDestinationRoutingReason_DomainAllowList),
		string(EmailFileAssessmentRequestDestinationRoutingReason_DomainBlockList),
		string(EmailFileAssessmentRequestDestinationRoutingReason_FirstTimeSender),
		string(EmailFileAssessmentRequestDestinationRoutingReason_Junk),
		string(EmailFileAssessmentRequestDestinationRoutingReason_MailFlowRule),
		string(EmailFileAssessmentRequestDestinationRoutingReason_None),
		string(EmailFileAssessmentRequestDestinationRoutingReason_NotInAddressBook),
		string(EmailFileAssessmentRequestDestinationRoutingReason_NotJunk),
		string(EmailFileAssessmentRequestDestinationRoutingReason_Outbound),
		string(EmailFileAssessmentRequestDestinationRoutingReason_SafeSender),
	}
}

func (s *EmailFileAssessmentRequestDestinationRoutingReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailFileAssessmentRequestDestinationRoutingReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailFileAssessmentRequestDestinationRoutingReason(input string) (*EmailFileAssessmentRequestDestinationRoutingReason, error) {
	vals := map[string]EmailFileAssessmentRequestDestinationRoutingReason{
		"advancedspamfiltering": EmailFileAssessmentRequestDestinationRoutingReason_AdvancedSpamFiltering,
		"autopurgetodeleted":    EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToDeleted,
		"autopurgetoinbox":      EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToInbox,
		"autopurgetojunk":       EmailFileAssessmentRequestDestinationRoutingReason_AutoPurgeToJunk,
		"blockedsender":         EmailFileAssessmentRequestDestinationRoutingReason_BlockedSender,
		"domainallowlist":       EmailFileAssessmentRequestDestinationRoutingReason_DomainAllowList,
		"domainblocklist":       EmailFileAssessmentRequestDestinationRoutingReason_DomainBlockList,
		"firsttimesender":       EmailFileAssessmentRequestDestinationRoutingReason_FirstTimeSender,
		"junk":                  EmailFileAssessmentRequestDestinationRoutingReason_Junk,
		"mailflowrule":          EmailFileAssessmentRequestDestinationRoutingReason_MailFlowRule,
		"none":                  EmailFileAssessmentRequestDestinationRoutingReason_None,
		"notinaddressbook":      EmailFileAssessmentRequestDestinationRoutingReason_NotInAddressBook,
		"notjunk":               EmailFileAssessmentRequestDestinationRoutingReason_NotJunk,
		"outbound":              EmailFileAssessmentRequestDestinationRoutingReason_Outbound,
		"safesender":            EmailFileAssessmentRequestDestinationRoutingReason_SafeSender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestDestinationRoutingReason(input)
	return &out, nil
}
