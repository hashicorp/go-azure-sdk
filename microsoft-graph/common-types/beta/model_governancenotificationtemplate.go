package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceNotificationTemplate struct {
	Culture   *string `json:"culture,omitempty"`
	Id        *string `json:"id,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Source    *string `json:"source,omitempty"`
	Type      *string `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}
