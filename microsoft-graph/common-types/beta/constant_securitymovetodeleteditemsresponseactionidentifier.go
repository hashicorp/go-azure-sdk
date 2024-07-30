package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMoveToDeletedItemsResponseActionIdentifier string

const (
	SecurityMoveToDeletedItemsResponseActionIdentifier_NetworkMessageId      SecurityMoveToDeletedItemsResponseActionIdentifier = "networkMessageId"
	SecurityMoveToDeletedItemsResponseActionIdentifier_RecipientEmailAddress SecurityMoveToDeletedItemsResponseActionIdentifier = "recipientEmailAddress"
)

func PossibleValuesForSecurityMoveToDeletedItemsResponseActionIdentifier() []string {
	return []string{
		string(SecurityMoveToDeletedItemsResponseActionIdentifier_NetworkMessageId),
		string(SecurityMoveToDeletedItemsResponseActionIdentifier_RecipientEmailAddress),
	}
}

func (s *SecurityMoveToDeletedItemsResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMoveToDeletedItemsResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMoveToDeletedItemsResponseActionIdentifier(input string) (*SecurityMoveToDeletedItemsResponseActionIdentifier, error) {
	vals := map[string]SecurityMoveToDeletedItemsResponseActionIdentifier{
		"networkmessageid":      SecurityMoveToDeletedItemsResponseActionIdentifier_NetworkMessageId,
		"recipientemailaddress": SecurityMoveToDeletedItemsResponseActionIdentifier_RecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMoveToDeletedItemsResponseActionIdentifier(input)
	return &out, nil
}
