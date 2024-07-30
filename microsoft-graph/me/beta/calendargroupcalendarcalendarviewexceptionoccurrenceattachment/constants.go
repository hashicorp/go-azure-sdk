package calendargroupcalendarcalendarviewexceptionoccurrenceattachment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentItemAttachmentType string

const (
	AttachmentItemAttachmentTypeFile      AttachmentItemAttachmentType = "file"
	AttachmentItemAttachmentTypeItem      AttachmentItemAttachmentType = "item"
	AttachmentItemAttachmentTypeReference AttachmentItemAttachmentType = "reference"
)

func PossibleValuesForAttachmentItemAttachmentType() []string {
	return []string{
		string(AttachmentItemAttachmentTypeFile),
		string(AttachmentItemAttachmentTypeItem),
		string(AttachmentItemAttachmentTypeReference),
	}
}

func (s *AttachmentItemAttachmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttachmentItemAttachmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttachmentItemAttachmentType(input string) (*AttachmentItemAttachmentType, error) {
	vals := map[string]AttachmentItemAttachmentType{
		"file":      AttachmentItemAttachmentTypeFile,
		"item":      AttachmentItemAttachmentTypeItem,
		"reference": AttachmentItemAttachmentTypeReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttachmentItemAttachmentType(input)
	return &out, nil
}
