package integrationaccountmaps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountMapProperties struct {
	ChangedTime      *string                                          `json:"changedTime,omitempty"`
	Content          *string                                          `json:"content,omitempty"`
	ContentLink      *ContentLink                                     `json:"contentLink,omitempty"`
	ContentType      *string                                          `json:"contentType,omitempty"`
	CreatedTime      *string                                          `json:"createdTime,omitempty"`
	MapType          MapType                                          `json:"mapType"`
	Metadata         *interface{}                                     `json:"metadata,omitempty"`
	ParametersSchema *IntegrationAccountMapPropertiesParametersSchema `json:"parametersSchema,omitempty"`
}
