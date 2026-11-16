package capabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeSizeCapability struct {
	MemoryPerVCoreInMb int64 `json:"memoryPerVCoreInMb"`
	VCores             int64 `json:"vCores"`
}
