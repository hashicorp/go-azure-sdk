package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartHoldMusicOperation struct {
	ClientContext *string                        `json:"clientContext,omitempty"`
	Id            *string                        `json:"id,omitempty"`
	ODataType     *string                        `json:"@odata.type,omitempty"`
	ResultInfo    *ResultInfo                    `json:"resultInfo,omitempty"`
	Status        *StartHoldMusicOperationStatus `json:"status,omitempty"`
}
