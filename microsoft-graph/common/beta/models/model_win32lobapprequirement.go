package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRequirement struct {
	DetectionValue *string                         `json:"detectionValue,omitempty"`
	ODataType      *string                         `json:"@odata.type,omitempty"`
	Operator       *Win32LobAppRequirementOperator `json:"operator,omitempty"`
}
