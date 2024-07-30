package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Website struct {
	Address     *string      `json:"address,omitempty"`
	DisplayName *string      `json:"displayName,omitempty"`
	ODataType   *string      `json:"@odata.type,omitempty"`
	Type        *WebsiteType `json:"type,omitempty"`
}
