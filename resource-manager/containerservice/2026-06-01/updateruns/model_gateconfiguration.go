package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GateConfiguration struct {
	DisplayName *string  `json:"displayName,omitempty"`
	Type        GateType `json:"type"`
}
