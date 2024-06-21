package keyvalues

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValueProperties struct {
	ContentType  *string            `json:"contentType,omitempty"`
	ETag         *string            `json:"eTag,omitempty"`
	Key          *string            `json:"key,omitempty"`
	Label        *string            `json:"label,omitempty"`
	LastModified *string            `json:"lastModified,omitempty"`
	Locked       *bool              `json:"locked,omitempty"`
	Tags         *map[string]string `json:"tags,omitempty"`
	Value        *string            `json:"value,omitempty"`
}
