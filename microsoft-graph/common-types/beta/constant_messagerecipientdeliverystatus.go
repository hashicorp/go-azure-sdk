package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRecipientDeliveryStatus string

const (
	MessageRecipientDeliveryStatus_Delivered      MessageRecipientDeliveryStatus = "delivered"
	MessageRecipientDeliveryStatus_Expanded       MessageRecipientDeliveryStatus = "expanded"
	MessageRecipientDeliveryStatus_Failed         MessageRecipientDeliveryStatus = "failed"
	MessageRecipientDeliveryStatus_FilteredAsSpam MessageRecipientDeliveryStatus = "filteredAsSpam"
	MessageRecipientDeliveryStatus_GettingStatus  MessageRecipientDeliveryStatus = "gettingStatus"
	MessageRecipientDeliveryStatus_Pending        MessageRecipientDeliveryStatus = "pending"
	MessageRecipientDeliveryStatus_Quarantined    MessageRecipientDeliveryStatus = "quarantined"
)

func PossibleValuesForMessageRecipientDeliveryStatus() []string {
	return []string{
		string(MessageRecipientDeliveryStatus_Delivered),
		string(MessageRecipientDeliveryStatus_Expanded),
		string(MessageRecipientDeliveryStatus_Failed),
		string(MessageRecipientDeliveryStatus_FilteredAsSpam),
		string(MessageRecipientDeliveryStatus_GettingStatus),
		string(MessageRecipientDeliveryStatus_Pending),
		string(MessageRecipientDeliveryStatus_Quarantined),
	}
}

func (s *MessageRecipientDeliveryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRecipientDeliveryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRecipientDeliveryStatus(input string) (*MessageRecipientDeliveryStatus, error) {
	vals := map[string]MessageRecipientDeliveryStatus{
		"delivered":      MessageRecipientDeliveryStatus_Delivered,
		"expanded":       MessageRecipientDeliveryStatus_Expanded,
		"failed":         MessageRecipientDeliveryStatus_Failed,
		"filteredasspam": MessageRecipientDeliveryStatus_FilteredAsSpam,
		"gettingstatus":  MessageRecipientDeliveryStatus_GettingStatus,
		"pending":        MessageRecipientDeliveryStatus_Pending,
		"quarantined":    MessageRecipientDeliveryStatus_Quarantined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRecipientDeliveryStatus(input)
	return &out, nil
}
