package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowBlockListEntryResultEntryType string

const (
	SecurityTenantAllowBlockListEntryResultEntryTypefileHash  SecurityTenantAllowBlockListEntryResultEntryType = "FileHash"
	SecurityTenantAllowBlockListEntryResultEntryTyperecipient SecurityTenantAllowBlockListEntryResultEntryType = "Recipient"
	SecurityTenantAllowBlockListEntryResultEntryTypesender    SecurityTenantAllowBlockListEntryResultEntryType = "Sender"
	SecurityTenantAllowBlockListEntryResultEntryTypeurl       SecurityTenantAllowBlockListEntryResultEntryType = "Url"
)

func PossibleValuesForSecurityTenantAllowBlockListEntryResultEntryType() []string {
	return []string{
		string(SecurityTenantAllowBlockListEntryResultEntryTypefileHash),
		string(SecurityTenantAllowBlockListEntryResultEntryTyperecipient),
		string(SecurityTenantAllowBlockListEntryResultEntryTypesender),
		string(SecurityTenantAllowBlockListEntryResultEntryTypeurl),
	}
}

func parseSecurityTenantAllowBlockListEntryResultEntryType(input string) (*SecurityTenantAllowBlockListEntryResultEntryType, error) {
	vals := map[string]SecurityTenantAllowBlockListEntryResultEntryType{
		"filehash":  SecurityTenantAllowBlockListEntryResultEntryTypefileHash,
		"recipient": SecurityTenantAllowBlockListEntryResultEntryTyperecipient,
		"sender":    SecurityTenantAllowBlockListEntryResultEntryTypesender,
		"url":       SecurityTenantAllowBlockListEntryResultEntryTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowBlockListEntryResultEntryType(input)
	return &out, nil
}
