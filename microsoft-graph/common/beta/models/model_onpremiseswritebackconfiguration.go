package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesWritebackConfiguration struct {
	ODataType             *string `json:"@odata.type,omitempty"`
	UnifiedGroupContainer *string `json:"unifiedGroupContainer,omitempty"`
	UserContainer         *string `json:"userContainer,omitempty"`
}
