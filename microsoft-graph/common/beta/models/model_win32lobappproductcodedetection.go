package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppProductCodeDetection struct {
	ODataType              *string                                                `json:"@odata.type,omitempty"`
	ProductCode            *string                                                `json:"productCode,omitempty"`
	ProductVersion         *string                                                `json:"productVersion,omitempty"`
	ProductVersionOperator *Win32LobAppProductCodeDetectionProductVersionOperator `json:"productVersionOperator,omitempty"`
}
