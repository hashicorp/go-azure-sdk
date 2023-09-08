package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageEventEventType string

const (
	MessageEventEventTypedelayed                     MessageEventEventType = "Delayed"
	MessageEventEventTypedelivered                   MessageEventEventType = "Delivered"
	MessageEventEventTypedistributionGroupExpanded   MessageEventEventType = "DistributionGroupExpanded"
	MessageEventEventTypedlpRuleTriggered            MessageEventEventType = "DlpRuleTriggered"
	MessageEventEventTypedropped                     MessageEventEventType = "Dropped"
	MessageEventEventTypefailed                      MessageEventEventType = "Failed"
	MessageEventEventTypejournaled                   MessageEventEventType = "Journaled"
	MessageEventEventTypemalwareDetected             MessageEventEventType = "MalwareDetected"
	MessageEventEventTypemalwareDetectedInAttachment MessageEventEventType = "MalwareDetectedInAttachment"
	MessageEventEventTypemalwareDetectedInMessage    MessageEventEventType = "MalwareDetectedInMessage"
	MessageEventEventTypeprocessingFailed            MessageEventEventType = "ProcessingFailed"
	MessageEventEventTypereceived                    MessageEventEventType = "Received"
	MessageEventEventTyperecipientsAdded             MessageEventEventType = "RecipientsAdded"
	MessageEventEventTyperedirected                  MessageEventEventType = "Redirected"
	MessageEventEventTyperesolved                    MessageEventEventType = "Resolved"
	MessageEventEventTypesent                        MessageEventEventType = "Sent"
	MessageEventEventTypespamDetected                MessageEventEventType = "SpamDetected"
	MessageEventEventTypesubmitted                   MessageEventEventType = "Submitted"
	MessageEventEventTypetransportRuleTriggered      MessageEventEventType = "TransportRuleTriggered"
	MessageEventEventTypettDelivered                 MessageEventEventType = "TtDelivered"
	MessageEventEventTypettZapped                    MessageEventEventType = "TtZapped"
)

func PossibleValuesForMessageEventEventType() []string {
	return []string{
		string(MessageEventEventTypedelayed),
		string(MessageEventEventTypedelivered),
		string(MessageEventEventTypedistributionGroupExpanded),
		string(MessageEventEventTypedlpRuleTriggered),
		string(MessageEventEventTypedropped),
		string(MessageEventEventTypefailed),
		string(MessageEventEventTypejournaled),
		string(MessageEventEventTypemalwareDetected),
		string(MessageEventEventTypemalwareDetectedInAttachment),
		string(MessageEventEventTypemalwareDetectedInMessage),
		string(MessageEventEventTypeprocessingFailed),
		string(MessageEventEventTypereceived),
		string(MessageEventEventTyperecipientsAdded),
		string(MessageEventEventTyperedirected),
		string(MessageEventEventTyperesolved),
		string(MessageEventEventTypesent),
		string(MessageEventEventTypespamDetected),
		string(MessageEventEventTypesubmitted),
		string(MessageEventEventTypetransportRuleTriggered),
		string(MessageEventEventTypettDelivered),
		string(MessageEventEventTypettZapped),
	}
}

func parseMessageEventEventType(input string) (*MessageEventEventType, error) {
	vals := map[string]MessageEventEventType{
		"delayed":                     MessageEventEventTypedelayed,
		"delivered":                   MessageEventEventTypedelivered,
		"distributiongroupexpanded":   MessageEventEventTypedistributionGroupExpanded,
		"dlpruletriggered":            MessageEventEventTypedlpRuleTriggered,
		"dropped":                     MessageEventEventTypedropped,
		"failed":                      MessageEventEventTypefailed,
		"journaled":                   MessageEventEventTypejournaled,
		"malwaredetected":             MessageEventEventTypemalwareDetected,
		"malwaredetectedinattachment": MessageEventEventTypemalwareDetectedInAttachment,
		"malwaredetectedinmessage":    MessageEventEventTypemalwareDetectedInMessage,
		"processingfailed":            MessageEventEventTypeprocessingFailed,
		"received":                    MessageEventEventTypereceived,
		"recipientsadded":             MessageEventEventTyperecipientsAdded,
		"redirected":                  MessageEventEventTyperedirected,
		"resolved":                    MessageEventEventTyperesolved,
		"sent":                        MessageEventEventTypesent,
		"spamdetected":                MessageEventEventTypespamDetected,
		"submitted":                   MessageEventEventTypesubmitted,
		"transportruletriggered":      MessageEventEventTypetransportRuleTriggered,
		"ttdelivered":                 MessageEventEventTypettDelivered,
		"ttzapped":                    MessageEventEventTypettZapped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageEventEventType(input)
	return &out, nil
}
