package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskProfile struct {
	HumanCount    *int64  `json:"humanCount,omitempty"`
	NonHumanCount *int64  `json:"nonHumanCount,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
