package protectionintent

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreValidateEnableBackupResponse struct {
	ContainerName     *string           `json:"containerName,omitempty"`
	ErrorCode         *string           `json:"errorCode,omitempty"`
	ErrorMessage      *string           `json:"errorMessage,omitempty"`
	ProtectedItemName *string           `json:"protectedItemName,omitempty"`
	Recommendation    *string           `json:"recommendation,omitempty"`
	Status            *ValidationStatus `json:"status,omitempty"`
}
