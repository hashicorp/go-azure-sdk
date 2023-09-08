package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCCrossCloudGovernmentOrganizationMapping struct {
	Id                          *string   `json:"id,omitempty"`
	ODataType                   *string   `json:"@odata.type,omitempty"`
	OrganizationIdsInUSGovCloud *[]string `json:"organizationIdsInUSGovCloud,omitempty"`
}
