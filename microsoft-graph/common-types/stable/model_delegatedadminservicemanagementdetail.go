package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminServiceManagementDetail struct {
	Id                   *string `json:"id,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	ServiceManagementUrl *string `json:"serviceManagementUrl,omitempty"`
	ServiceName          *string `json:"serviceName,omitempty"`
}
