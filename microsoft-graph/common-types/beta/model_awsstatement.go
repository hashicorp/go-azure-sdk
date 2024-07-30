package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsStatement struct {
	Actions      *[]string           `json:"actions,omitempty"`
	Condition    *AwsCondition       `json:"condition,omitempty"`
	Effect       *AwsStatementEffect `json:"effect,omitempty"`
	NotActions   *[]string           `json:"notActions,omitempty"`
	NotResources *[]string           `json:"notResources,omitempty"`
	ODataType    *string             `json:"@odata.type,omitempty"`
	Resources    *[]string           `json:"resources,omitempty"`
	StatementId  *string             `json:"statementId,omitempty"`
}
