package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealth struct {
	Id        *string               `json:"id,omitempty"`
	Issues    *[]ServiceHealthIssue `json:"issues,omitempty"`
	ODataType *string               `json:"@odata.type,omitempty"`
	Service   *string               `json:"service,omitempty"`
	Status    *ServiceHealthStatus  `json:"status,omitempty"`
}
