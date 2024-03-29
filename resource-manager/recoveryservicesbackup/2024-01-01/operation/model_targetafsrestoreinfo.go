package operation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetAFSRestoreInfo struct {
	Name             *string `json:"name,omitempty"`
	TargetResourceId *string `json:"targetResourceId,omitempty"`
}
