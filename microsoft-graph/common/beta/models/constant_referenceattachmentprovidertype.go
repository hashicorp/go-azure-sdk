package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachmentProviderType string

const (
	ReferenceAttachmentProviderTypedropbox          ReferenceAttachmentProviderType = "Dropbox"
	ReferenceAttachmentProviderTypeoneDriveBusiness ReferenceAttachmentProviderType = "OneDriveBusiness"
	ReferenceAttachmentProviderTypeoneDriveConsumer ReferenceAttachmentProviderType = "OneDriveConsumer"
	ReferenceAttachmentProviderTypeother            ReferenceAttachmentProviderType = "Other"
)

func PossibleValuesForReferenceAttachmentProviderType() []string {
	return []string{
		string(ReferenceAttachmentProviderTypedropbox),
		string(ReferenceAttachmentProviderTypeoneDriveBusiness),
		string(ReferenceAttachmentProviderTypeoneDriveConsumer),
		string(ReferenceAttachmentProviderTypeother),
	}
}

func parseReferenceAttachmentProviderType(input string) (*ReferenceAttachmentProviderType, error) {
	vals := map[string]ReferenceAttachmentProviderType{
		"dropbox":          ReferenceAttachmentProviderTypedropbox,
		"onedrivebusiness": ReferenceAttachmentProviderTypeoneDriveBusiness,
		"onedriveconsumer": ReferenceAttachmentProviderTypeoneDriveConsumer,
		"other":            ReferenceAttachmentProviderTypeother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceAttachmentProviderType(input)
	return &out, nil
}
