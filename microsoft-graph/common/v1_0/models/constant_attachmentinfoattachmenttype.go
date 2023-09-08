package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentInfoAttachmentType string

const (
	AttachmentInfoAttachmentTypefile      AttachmentInfoAttachmentType = "File"
	AttachmentInfoAttachmentTypeitem      AttachmentInfoAttachmentType = "Item"
	AttachmentInfoAttachmentTypereference AttachmentInfoAttachmentType = "Reference"
)

func PossibleValuesForAttachmentInfoAttachmentType() []string {
	return []string{
		string(AttachmentInfoAttachmentTypefile),
		string(AttachmentInfoAttachmentTypeitem),
		string(AttachmentInfoAttachmentTypereference),
	}
}

func parseAttachmentInfoAttachmentType(input string) (*AttachmentInfoAttachmentType, error) {
	vals := map[string]AttachmentInfoAttachmentType{
		"file":      AttachmentInfoAttachmentTypefile,
		"item":      AttachmentInfoAttachmentTypeitem,
		"reference": AttachmentInfoAttachmentTypereference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttachmentInfoAttachmentType(input)
	return &out, nil
}
