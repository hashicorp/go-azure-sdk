package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SolutionsRoot struct {
	BusinessScenarios *[]BusinessScenario `json:"businessScenarios,omitempty"`
	ODataType         *string             `json:"@odata.type,omitempty"`
	VirtualEvents     *VirtualEventsRoot  `json:"virtualEvents,omitempty"`
}
