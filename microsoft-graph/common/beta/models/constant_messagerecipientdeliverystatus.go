package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRecipientDeliveryStatus string

const (
	MessageRecipientDeliveryStatusdelivered      MessageRecipientDeliveryStatus = "Delivered"
	MessageRecipientDeliveryStatusexpanded       MessageRecipientDeliveryStatus = "Expanded"
	MessageRecipientDeliveryStatusfailed         MessageRecipientDeliveryStatus = "Failed"
	MessageRecipientDeliveryStatusfilteredAsSpam MessageRecipientDeliveryStatus = "FilteredAsSpam"
	MessageRecipientDeliveryStatusgettingStatus  MessageRecipientDeliveryStatus = "GettingStatus"
	MessageRecipientDeliveryStatuspending        MessageRecipientDeliveryStatus = "Pending"
	MessageRecipientDeliveryStatusquarantined    MessageRecipientDeliveryStatus = "Quarantined"
)

func PossibleValuesForMessageRecipientDeliveryStatus() []string {
	return []string{
		string(MessageRecipientDeliveryStatusdelivered),
		string(MessageRecipientDeliveryStatusexpanded),
		string(MessageRecipientDeliveryStatusfailed),
		string(MessageRecipientDeliveryStatusfilteredAsSpam),
		string(MessageRecipientDeliveryStatusgettingStatus),
		string(MessageRecipientDeliveryStatuspending),
		string(MessageRecipientDeliveryStatusquarantined),
	}
}

func parseMessageRecipientDeliveryStatus(input string) (*MessageRecipientDeliveryStatus, error) {
	vals := map[string]MessageRecipientDeliveryStatus{
		"delivered":      MessageRecipientDeliveryStatusdelivered,
		"expanded":       MessageRecipientDeliveryStatusexpanded,
		"failed":         MessageRecipientDeliveryStatusfailed,
		"filteredasspam": MessageRecipientDeliveryStatusfilteredAsSpam,
		"gettingstatus":  MessageRecipientDeliveryStatusgettingStatus,
		"pending":        MessageRecipientDeliveryStatuspending,
		"quarantined":    MessageRecipientDeliveryStatusquarantined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRecipientDeliveryStatus(input)
	return &out, nil
}
