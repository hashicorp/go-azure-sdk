package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestDestinationRoutingReason string

const (
	MailAssessmentRequestDestinationRoutingReason_AdvancedSpamFiltering MailAssessmentRequestDestinationRoutingReason = "advancedSpamFiltering"
	MailAssessmentRequestDestinationRoutingReason_AutoPurgeToDeleted    MailAssessmentRequestDestinationRoutingReason = "autoPurgeToDeleted"
	MailAssessmentRequestDestinationRoutingReason_AutoPurgeToInbox      MailAssessmentRequestDestinationRoutingReason = "autoPurgeToInbox"
	MailAssessmentRequestDestinationRoutingReason_AutoPurgeToJunk       MailAssessmentRequestDestinationRoutingReason = "autoPurgeToJunk"
	MailAssessmentRequestDestinationRoutingReason_BlockedSender         MailAssessmentRequestDestinationRoutingReason = "blockedSender"
	MailAssessmentRequestDestinationRoutingReason_DomainAllowList       MailAssessmentRequestDestinationRoutingReason = "domainAllowList"
	MailAssessmentRequestDestinationRoutingReason_DomainBlockList       MailAssessmentRequestDestinationRoutingReason = "domainBlockList"
	MailAssessmentRequestDestinationRoutingReason_FirstTimeSender       MailAssessmentRequestDestinationRoutingReason = "firstTimeSender"
	MailAssessmentRequestDestinationRoutingReason_Junk                  MailAssessmentRequestDestinationRoutingReason = "junk"
	MailAssessmentRequestDestinationRoutingReason_MailFlowRule          MailAssessmentRequestDestinationRoutingReason = "mailFlowRule"
	MailAssessmentRequestDestinationRoutingReason_None                  MailAssessmentRequestDestinationRoutingReason = "none"
	MailAssessmentRequestDestinationRoutingReason_NotInAddressBook      MailAssessmentRequestDestinationRoutingReason = "notInAddressBook"
	MailAssessmentRequestDestinationRoutingReason_NotJunk               MailAssessmentRequestDestinationRoutingReason = "notJunk"
	MailAssessmentRequestDestinationRoutingReason_Outbound              MailAssessmentRequestDestinationRoutingReason = "outbound"
	MailAssessmentRequestDestinationRoutingReason_SafeSender            MailAssessmentRequestDestinationRoutingReason = "safeSender"
)

func PossibleValuesForMailAssessmentRequestDestinationRoutingReason() []string {
	return []string{
		string(MailAssessmentRequestDestinationRoutingReason_AdvancedSpamFiltering),
		string(MailAssessmentRequestDestinationRoutingReason_AutoPurgeToDeleted),
		string(MailAssessmentRequestDestinationRoutingReason_AutoPurgeToInbox),
		string(MailAssessmentRequestDestinationRoutingReason_AutoPurgeToJunk),
		string(MailAssessmentRequestDestinationRoutingReason_BlockedSender),
		string(MailAssessmentRequestDestinationRoutingReason_DomainAllowList),
		string(MailAssessmentRequestDestinationRoutingReason_DomainBlockList),
		string(MailAssessmentRequestDestinationRoutingReason_FirstTimeSender),
		string(MailAssessmentRequestDestinationRoutingReason_Junk),
		string(MailAssessmentRequestDestinationRoutingReason_MailFlowRule),
		string(MailAssessmentRequestDestinationRoutingReason_None),
		string(MailAssessmentRequestDestinationRoutingReason_NotInAddressBook),
		string(MailAssessmentRequestDestinationRoutingReason_NotJunk),
		string(MailAssessmentRequestDestinationRoutingReason_Outbound),
		string(MailAssessmentRequestDestinationRoutingReason_SafeSender),
	}
}

func (s *MailAssessmentRequestDestinationRoutingReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailAssessmentRequestDestinationRoutingReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailAssessmentRequestDestinationRoutingReason(input string) (*MailAssessmentRequestDestinationRoutingReason, error) {
	vals := map[string]MailAssessmentRequestDestinationRoutingReason{
		"advancedspamfiltering": MailAssessmentRequestDestinationRoutingReason_AdvancedSpamFiltering,
		"autopurgetodeleted":    MailAssessmentRequestDestinationRoutingReason_AutoPurgeToDeleted,
		"autopurgetoinbox":      MailAssessmentRequestDestinationRoutingReason_AutoPurgeToInbox,
		"autopurgetojunk":       MailAssessmentRequestDestinationRoutingReason_AutoPurgeToJunk,
		"blockedsender":         MailAssessmentRequestDestinationRoutingReason_BlockedSender,
		"domainallowlist":       MailAssessmentRequestDestinationRoutingReason_DomainAllowList,
		"domainblocklist":       MailAssessmentRequestDestinationRoutingReason_DomainBlockList,
		"firsttimesender":       MailAssessmentRequestDestinationRoutingReason_FirstTimeSender,
		"junk":                  MailAssessmentRequestDestinationRoutingReason_Junk,
		"mailflowrule":          MailAssessmentRequestDestinationRoutingReason_MailFlowRule,
		"none":                  MailAssessmentRequestDestinationRoutingReason_None,
		"notinaddressbook":      MailAssessmentRequestDestinationRoutingReason_NotInAddressBook,
		"notjunk":               MailAssessmentRequestDestinationRoutingReason_NotJunk,
		"outbound":              MailAssessmentRequestDestinationRoutingReason_Outbound,
		"safesender":            MailAssessmentRequestDestinationRoutingReason_SafeSender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestDestinationRoutingReason(input)
	return &out, nil
}
