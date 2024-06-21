package databases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseKey struct {
	CreationDate *string          `json:"creationDate,omitempty"`
	Subregion    *string          `json:"subregion,omitempty"`
	Thumbprint   *string          `json:"thumbprint,omitempty"`
	Type         *DatabaseKeyType `json:"type,omitempty"`
}
