package department

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepartmentProperties struct {
	CostCenter  *string `json:"costCenter,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Id          *string `json:"id,omitempty"`
	Status      *string `json:"status,omitempty"`
}
