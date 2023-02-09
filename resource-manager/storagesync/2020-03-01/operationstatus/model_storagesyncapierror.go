package operationstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncApiError struct {
	Code       *string                       `json:"code,omitempty"`
	Details    *StorageSyncErrorDetails      `json:"details,omitempty"`
	Innererror *StorageSyncInnerErrorDetails `json:"innererror,omitempty"`
	Message    *string                       `json:"message,omitempty"`
	Target     *string                       `json:"target,omitempty"`
}
