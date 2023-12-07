package costdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobInfo struct {
	BlobLink  *string `json:"blobLink,omitempty"`
	ByteCount *int64  `json:"byteCount,omitempty"`
}
