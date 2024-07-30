package b2xuserflowuserattributeassignment

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentOrder struct {
	ODataType *string   `json:"@odata.type,omitempty"`
	Order     *[]string `json:"order,omitempty"`
}
