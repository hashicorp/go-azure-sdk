package vmhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMHostUpdateRequest struct {
	State         *VMHostUpdateStates `json:"state,omitempty"`
	VMResourceIds *[]VMResources      `json:"vmResourceIds,omitempty"`
}
