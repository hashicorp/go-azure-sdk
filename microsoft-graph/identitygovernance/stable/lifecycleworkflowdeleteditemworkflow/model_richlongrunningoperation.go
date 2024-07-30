package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RichLongRunningOperation struct {
	CreatedDateTime    *string                         `json:"createdDateTime,omitempty"`
	Error              *PublicError                    `json:"error,omitempty"`
	Id                 *string                         `json:"id,omitempty"`
	LastActionDateTime *string                         `json:"lastActionDateTime,omitempty"`
	ODataType          *string                         `json:"@odata.type,omitempty"`
	PercentageComplete *int64                          `json:"percentageComplete,omitempty"`
	ResourceId         *string                         `json:"resourceId,omitempty"`
	ResourceLocation   *string                         `json:"resourceLocation,omitempty"`
	Status             *RichLongRunningOperationStatus `json:"status,omitempty"`
	StatusDetail       *string                         `json:"statusDetail,omitempty"`
	Type               *string                         `json:"type,omitempty"`
}
