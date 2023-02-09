package operationstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncInnerErrorDetails struct {
	CallStack               *string `json:"callStack,omitempty"`
	InnerException          *string `json:"innerException,omitempty"`
	InnerExceptionCallStack *string `json:"innerExceptionCallStack,omitempty"`
	Message                 *string `json:"message,omitempty"`
}
