package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSecurityState struct {
	FileHash  *FileHash `json:"fileHash,omitempty"`
	Name      *string   `json:"name,omitempty"`
	ODataType *string   `json:"@odata.type,omitempty"`
	Path      *string   `json:"path,omitempty"`
	RiskScore *string   `json:"riskScore,omitempty"`
}
