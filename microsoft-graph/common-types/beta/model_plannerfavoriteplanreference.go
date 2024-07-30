package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerFavoritePlanReference struct {
	ODataType *string `json:"@odata.type,omitempty"`
	OrderHint *string `json:"orderHint,omitempty"`
	PlanTitle *string `json:"planTitle,omitempty"`
}
