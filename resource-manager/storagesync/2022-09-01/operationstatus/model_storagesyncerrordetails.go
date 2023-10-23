package operationstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncErrorDetails struct {
	Code          *string `json:"code,omitempty"`
	ExceptionType *string `json:"exceptionType,omitempty"`
	HTTPErrorCode *string `json:"httpErrorCode,omitempty"`
	HTTPMethod    *string `json:"httpMethod,omitempty"`
	HashedMessage *string `json:"hashedMessage,omitempty"`
	Message       *string `json:"message,omitempty"`
	RequestUri    *string `json:"requestUri,omitempty"`
	Target        *string `json:"target,omitempty"`
}
