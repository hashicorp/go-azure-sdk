package logfiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogFileProperties struct {
	CreatedTime      *string `json:"createdTime,omitempty"`
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`
	SizeInKb         *int64  `json:"sizeInKb,omitempty"`
	Type             *string `json:"type,omitempty"`
	Url              *string `json:"url,omitempty"`
}
