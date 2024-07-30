package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterClause struct {
	ODataType         *string        `json:"@odata.type,omitempty"`
	OperatorName      *string        `json:"operatorName,omitempty"`
	SourceOperandName *string        `json:"sourceOperandName,omitempty"`
	TargetOperand     *FilterOperand `json:"targetOperand,omitempty"`
}
