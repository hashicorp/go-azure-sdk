package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeOrgData struct {
	CostCenter *string `json:"costCenter,omitempty"`
	Division   *string `json:"division,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
}
