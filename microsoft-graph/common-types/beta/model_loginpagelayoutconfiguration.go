package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageLayoutConfiguration struct {
	IsFooterShown      *bool                                           `json:"isFooterShown,omitempty"`
	IsHeaderShown      *bool                                           `json:"isHeaderShown,omitempty"`
	LayoutTemplateType *LoginPageLayoutConfigurationLayoutTemplateType `json:"layoutTemplateType,omitempty"`
	ODataType          *string                                         `json:"@odata.type,omitempty"`
}
