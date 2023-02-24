package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedServiceDebugResource struct {
	Name       *string       `json:"name,omitempty"`
	Properties LinkedService `json:"properties"`
}
