package devices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceMoveDetails struct {
	OperationInProgress                 *ResourceMoveStatus `json:"operationInProgress,omitempty"`
	OperationInProgressLockTimeoutInUTC *string             `json:"operationInProgressLockTimeoutInUTC,omitempty"`
}
