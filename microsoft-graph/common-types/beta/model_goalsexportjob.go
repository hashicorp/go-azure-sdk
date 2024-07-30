package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GoalsExportJob struct {
	Content             *string               `json:"content,omitempty"`
	CreatedDateTime     *string               `json:"createdDateTime,omitempty"`
	ExpirationDateTime  *string               `json:"expirationDateTime,omitempty"`
	ExplorerViewId      *string               `json:"explorerViewId,omitempty"`
	GoalsOrganizationId *string               `json:"goalsOrganizationId,omitempty"`
	Id                  *string               `json:"id,omitempty"`
	LastActionDateTime  *string               `json:"lastActionDateTime,omitempty"`
	ODataType           *string               `json:"@odata.type,omitempty"`
	ResourceLocation    *string               `json:"resourceLocation,omitempty"`
	Status              *GoalsExportJobStatus `json:"status,omitempty"`
	StatusDetail        *string               `json:"statusDetail,omitempty"`
}
