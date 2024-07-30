package securityinformationprotectionsensitivitylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDowngradeJustification struct {
	IsDowngradeJustified *bool   `json:"isDowngradeJustified,omitempty"`
	JustificationMessage *string `json:"justificationMessage,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
