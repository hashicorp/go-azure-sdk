package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDiagnosticStatus struct {
	MitigationInstruction *string `json:"mitigationInstruction,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	State                 *string `json:"state,omitempty"`
	ValidationName        *string `json:"validationName,omitempty"`
}
