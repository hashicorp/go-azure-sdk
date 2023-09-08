package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAlterationOptions struct {
	EnableModification *bool   `json:"enableModification,omitempty"`
	EnableSuggestion   *bool   `json:"enableSuggestion,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
