package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardWebPart struct {
	ContainerTextWebPartId *string      `json:"containerTextWebPartId,omitempty"`
	Data                   *WebPartData `json:"data,omitempty"`
	Id                     *string      `json:"id,omitempty"`
	ODataType              *string      `json:"@odata.type,omitempty"`
	WebPartType            *string      `json:"webPartType,omitempty"`
}
