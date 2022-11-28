package reservedinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatus struct {
	Properties *ReportURL           `json:"properties"`
	Status     *OperationStatusType `json:"status,omitempty"`
}
