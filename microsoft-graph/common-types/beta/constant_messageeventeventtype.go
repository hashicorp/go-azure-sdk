package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageEventEventType string

const (
	MessageEventEventType_Delayed                     MessageEventEventType = "delayed"
	MessageEventEventType_Delivered                   MessageEventEventType = "delivered"
	MessageEventEventType_DistributionGroupExpanded   MessageEventEventType = "distributionGroupExpanded"
	MessageEventEventType_DlpRuleTriggered            MessageEventEventType = "dlpRuleTriggered"
	MessageEventEventType_Dropped                     MessageEventEventType = "dropped"
	MessageEventEventType_Failed                      MessageEventEventType = "failed"
	MessageEventEventType_Journaled                   MessageEventEventType = "journaled"
	MessageEventEventType_MalwareDetected             MessageEventEventType = "malwareDetected"
	MessageEventEventType_MalwareDetectedInAttachment MessageEventEventType = "malwareDetectedInAttachment"
	MessageEventEventType_MalwareDetectedInMessage    MessageEventEventType = "malwareDetectedInMessage"
	MessageEventEventType_ProcessingFailed            MessageEventEventType = "processingFailed"
	MessageEventEventType_Received                    MessageEventEventType = "received"
	MessageEventEventType_RecipientsAdded             MessageEventEventType = "recipientsAdded"
	MessageEventEventType_Redirected                  MessageEventEventType = "redirected"
	MessageEventEventType_Resolved                    MessageEventEventType = "resolved"
	MessageEventEventType_Sent                        MessageEventEventType = "sent"
	MessageEventEventType_SpamDetected                MessageEventEventType = "spamDetected"
	MessageEventEventType_Submitted                   MessageEventEventType = "submitted"
	MessageEventEventType_TransportRuleTriggered      MessageEventEventType = "transportRuleTriggered"
	MessageEventEventType_TtDelivered                 MessageEventEventType = "ttDelivered"
	MessageEventEventType_TtZapped                    MessageEventEventType = "ttZapped"
)

func PossibleValuesForMessageEventEventType() []string {
	return []string{
		string(MessageEventEventType_Delayed),
		string(MessageEventEventType_Delivered),
		string(MessageEventEventType_DistributionGroupExpanded),
		string(MessageEventEventType_DlpRuleTriggered),
		string(MessageEventEventType_Dropped),
		string(MessageEventEventType_Failed),
		string(MessageEventEventType_Journaled),
		string(MessageEventEventType_MalwareDetected),
		string(MessageEventEventType_MalwareDetectedInAttachment),
		string(MessageEventEventType_MalwareDetectedInMessage),
		string(MessageEventEventType_ProcessingFailed),
		string(MessageEventEventType_Received),
		string(MessageEventEventType_RecipientsAdded),
		string(MessageEventEventType_Redirected),
		string(MessageEventEventType_Resolved),
		string(MessageEventEventType_Sent),
		string(MessageEventEventType_SpamDetected),
		string(MessageEventEventType_Submitted),
		string(MessageEventEventType_TransportRuleTriggered),
		string(MessageEventEventType_TtDelivered),
		string(MessageEventEventType_TtZapped),
	}
}

func (s *MessageEventEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageEventEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageEventEventType(input string) (*MessageEventEventType, error) {
	vals := map[string]MessageEventEventType{
		"delayed":                     MessageEventEventType_Delayed,
		"delivered":                   MessageEventEventType_Delivered,
		"distributiongroupexpanded":   MessageEventEventType_DistributionGroupExpanded,
		"dlpruletriggered":            MessageEventEventType_DlpRuleTriggered,
		"dropped":                     MessageEventEventType_Dropped,
		"failed":                      MessageEventEventType_Failed,
		"journaled":                   MessageEventEventType_Journaled,
		"malwaredetected":             MessageEventEventType_MalwareDetected,
		"malwaredetectedinattachment": MessageEventEventType_MalwareDetectedInAttachment,
		"malwaredetectedinmessage":    MessageEventEventType_MalwareDetectedInMessage,
		"processingfailed":            MessageEventEventType_ProcessingFailed,
		"received":                    MessageEventEventType_Received,
		"recipientsadded":             MessageEventEventType_RecipientsAdded,
		"redirected":                  MessageEventEventType_Redirected,
		"resolved":                    MessageEventEventType_Resolved,
		"sent":                        MessageEventEventType_Sent,
		"spamdetected":                MessageEventEventType_SpamDetected,
		"submitted":                   MessageEventEventType_Submitted,
		"transportruletriggered":      MessageEventEventType_TransportRuleTriggered,
		"ttdelivered":                 MessageEventEventType_TtDelivered,
		"ttzapped":                    MessageEventEventType_TtZapped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageEventEventType(input)
	return &out, nil
}
