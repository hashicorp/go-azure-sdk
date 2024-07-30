package informationprotectiondatalosspreventionpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentLabel struct {
	ApplicationMode *CurrentLabelApplicationMode `json:"applicationMode,omitempty"`
	Id              *string                      `json:"id,omitempty"`
	ODataType       *string                      `json:"@odata.type,omitempty"`
}
