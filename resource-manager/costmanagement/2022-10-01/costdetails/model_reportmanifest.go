package costdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportManifest struct {
	BlobCount       *int64                 `json:"blobCount,omitempty"`
	Blobs           *[]BlobInfo            `json:"blobs,omitempty"`
	ByteCount       *int64                 `json:"byteCount,omitempty"`
	CompressData    *bool                  `json:"compressData,omitempty"`
	DataFormat      *CostDetailsDataFormat `json:"dataFormat,omitempty"`
	ManifestVersion *string                `json:"manifestVersion,omitempty"`
	RequestContext  *RequestContext        `json:"requestContext"`
}
