package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentItemAttachmentType string

const (
	AttachmentItemAttachmentTypefile      AttachmentItemAttachmentType = "File"
	AttachmentItemAttachmentTypeitem      AttachmentItemAttachmentType = "Item"
	AttachmentItemAttachmentTypereference AttachmentItemAttachmentType = "Reference"
)

func PossibleValuesForAttachmentItemAttachmentType() []string {
	return []string{
		string(AttachmentItemAttachmentTypefile),
		string(AttachmentItemAttachmentTypeitem),
		string(AttachmentItemAttachmentTypereference),
	}
}

func parseAttachmentItemAttachmentType(input string) (*AttachmentItemAttachmentType, error) {
	vals := map[string]AttachmentItemAttachmentType{
		"file":      AttachmentItemAttachmentTypefile,
		"item":      AttachmentItemAttachmentTypeitem,
		"reference": AttachmentItemAttachmentTypereference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttachmentItemAttachmentType(input)
	return &out, nil
}
