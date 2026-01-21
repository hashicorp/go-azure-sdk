package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeRemoveParameter struct {
	NodeDeallocationOption *ComputeNodeDeallocationOption `json:"nodeDeallocationOption,omitempty"`
	NodeList               []string                       `json:"nodeList"`
	ResizeTimeout          *string                        `json:"resizeTimeout,omitempty"`
}
