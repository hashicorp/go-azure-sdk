package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlScore struct {
	ControlCategory *string  `json:"controlCategory,omitempty"`
	ControlName     *string  `json:"controlName,omitempty"`
	Description     *string  `json:"description,omitempty"`
	ODataType       *string  `json:"@odata.type,omitempty"`
	Score           *float64 `json:"score,omitempty"`
}
