package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyRule struct {
	ODataType      *string                                       `json:"@odata.type,omitempty"`
	Operation      *ExternalConnectorsPropertyRuleOperation      `json:"operation,omitempty"`
	Property       *string                                       `json:"property,omitempty"`
	Values         *[]string                                     `json:"values,omitempty"`
	ValuesJoinedBy *ExternalConnectorsPropertyRuleValuesJoinedBy `json:"valuesJoinedBy,omitempty"`
}
