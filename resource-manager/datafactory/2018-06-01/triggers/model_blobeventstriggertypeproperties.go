package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobEventsTriggerTypeProperties struct {
	BlobPathBeginsWith *string          `json:"blobPathBeginsWith,omitempty"`
	BlobPathEndsWith   *string          `json:"blobPathEndsWith,omitempty"`
	Events             []BlobEventTypes `json:"events"`
	IgnoreEmptyBlobs   *bool            `json:"ignoreEmptyBlobs,omitempty"`
	Scope              string           `json:"scope"`
}
