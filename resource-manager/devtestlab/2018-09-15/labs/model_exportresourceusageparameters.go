package labs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportResourceUsageParameters struct {
	BlobStorageAbsoluteSasUri *string `json:"blobStorageAbsoluteSasUri,omitempty"`
	UsageStartDate            *string `json:"usageStartDate,omitempty"`
}
