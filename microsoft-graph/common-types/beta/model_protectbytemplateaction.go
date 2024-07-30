package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectByTemplateAction struct {
	ODataType  *string `json:"@odata.type,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
}
