package todolisttaskattachment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentInfoAttachmentType string

const (
	AttachmentInfoAttachmentTypeFile      AttachmentInfoAttachmentType = "file"
	AttachmentInfoAttachmentTypeItem      AttachmentInfoAttachmentType = "item"
	AttachmentInfoAttachmentTypeReference AttachmentInfoAttachmentType = "reference"
)

func PossibleValuesForAttachmentInfoAttachmentType() []string {
	return []string{
		string(AttachmentInfoAttachmentTypeFile),
		string(AttachmentInfoAttachmentTypeItem),
		string(AttachmentInfoAttachmentTypeReference),
	}
}

func (s *AttachmentInfoAttachmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttachmentInfoAttachmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttachmentInfoAttachmentType(input string) (*AttachmentInfoAttachmentType, error) {
	vals := map[string]AttachmentInfoAttachmentType{
		"file":      AttachmentInfoAttachmentTypeFile,
		"item":      AttachmentInfoAttachmentTypeItem,
		"reference": AttachmentInfoAttachmentTypeReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttachmentInfoAttachmentType(input)
	return &out, nil
}
