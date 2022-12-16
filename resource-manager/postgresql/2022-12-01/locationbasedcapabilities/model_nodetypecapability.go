package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeTypeCapability struct {
	Name     *string `json:"name,omitempty"`
	NodeType *string `json:"nodeType,omitempty"`
	Status   *string `json:"status,omitempty"`
}
