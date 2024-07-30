package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsCreepIndex struct {
	ODataType *string `json:"@odata.type,omitempty"`
	Score     *int64  `json:"score,omitempty"`
}
