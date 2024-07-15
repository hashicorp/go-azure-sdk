package containerappspatches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatchDetailsNewLayer struct {
	FrameworkAndVersion *string `json:"frameworkAndVersion,omitempty"`
	Name                *string `json:"name,omitempty"`
	OsAndVersion        *string `json:"osAndVersion,omitempty"`
}
