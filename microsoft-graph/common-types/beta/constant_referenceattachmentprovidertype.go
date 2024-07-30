package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachmentProviderType string

const (
	ReferenceAttachmentProviderType_Dropbox          ReferenceAttachmentProviderType = "dropbox"
	ReferenceAttachmentProviderType_OneDriveBusiness ReferenceAttachmentProviderType = "oneDriveBusiness"
	ReferenceAttachmentProviderType_OneDriveConsumer ReferenceAttachmentProviderType = "oneDriveConsumer"
	ReferenceAttachmentProviderType_Other            ReferenceAttachmentProviderType = "other"
)

func PossibleValuesForReferenceAttachmentProviderType() []string {
	return []string{
		string(ReferenceAttachmentProviderType_Dropbox),
		string(ReferenceAttachmentProviderType_OneDriveBusiness),
		string(ReferenceAttachmentProviderType_OneDriveConsumer),
		string(ReferenceAttachmentProviderType_Other),
	}
}

func (s *ReferenceAttachmentProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReferenceAttachmentProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReferenceAttachmentProviderType(input string) (*ReferenceAttachmentProviderType, error) {
	vals := map[string]ReferenceAttachmentProviderType{
		"dropbox":          ReferenceAttachmentProviderType_Dropbox,
		"onedrivebusiness": ReferenceAttachmentProviderType_OneDriveBusiness,
		"onedriveconsumer": ReferenceAttachmentProviderType_OneDriveConsumer,
		"other":            ReferenceAttachmentProviderType_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceAttachmentProviderType(input)
	return &out, nil
}
