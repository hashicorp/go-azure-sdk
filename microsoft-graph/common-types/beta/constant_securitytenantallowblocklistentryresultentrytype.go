package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowBlockListEntryResultEntryType string

const (
	SecurityTenantAllowBlockListEntryResultEntryType_FileHash  SecurityTenantAllowBlockListEntryResultEntryType = "fileHash"
	SecurityTenantAllowBlockListEntryResultEntryType_Recipient SecurityTenantAllowBlockListEntryResultEntryType = "recipient"
	SecurityTenantAllowBlockListEntryResultEntryType_Sender    SecurityTenantAllowBlockListEntryResultEntryType = "sender"
	SecurityTenantAllowBlockListEntryResultEntryType_Url       SecurityTenantAllowBlockListEntryResultEntryType = "url"
)

func PossibleValuesForSecurityTenantAllowBlockListEntryResultEntryType() []string {
	return []string{
		string(SecurityTenantAllowBlockListEntryResultEntryType_FileHash),
		string(SecurityTenantAllowBlockListEntryResultEntryType_Recipient),
		string(SecurityTenantAllowBlockListEntryResultEntryType_Sender),
		string(SecurityTenantAllowBlockListEntryResultEntryType_Url),
	}
}

func (s *SecurityTenantAllowBlockListEntryResultEntryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTenantAllowBlockListEntryResultEntryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTenantAllowBlockListEntryResultEntryType(input string) (*SecurityTenantAllowBlockListEntryResultEntryType, error) {
	vals := map[string]SecurityTenantAllowBlockListEntryResultEntryType{
		"filehash":  SecurityTenantAllowBlockListEntryResultEntryType_FileHash,
		"recipient": SecurityTenantAllowBlockListEntryResultEntryType_Recipient,
		"sender":    SecurityTenantAllowBlockListEntryResultEntryType_Sender,
		"url":       SecurityTenantAllowBlockListEntryResultEntryType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowBlockListEntryResultEntryType(input)
	return &out, nil
}
