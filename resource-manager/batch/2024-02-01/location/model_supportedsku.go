package location

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedSku struct {
	BatchSupportEndOfLife *string          `json:"batchSupportEndOfLife,omitempty"`
	Capabilities          *[]SkuCapability `json:"capabilities,omitempty"`
	FamilyName            *string          `json:"familyName,omitempty"`
	Name                  *string          `json:"name,omitempty"`
}
