package integrationaccountschemas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountSchemaProperties struct {
	ChangedTime     *string      `json:"changedTime,omitempty"`
	Content         *string      `json:"content,omitempty"`
	ContentLink     *ContentLink `json:"contentLink,omitempty"`
	ContentType     *string      `json:"contentType,omitempty"`
	CreatedTime     *string      `json:"createdTime,omitempty"`
	DocumentName    *string      `json:"documentName,omitempty"`
	FileName        *string      `json:"fileName,omitempty"`
	Metadata        *interface{} `json:"metadata,omitempty"`
	SchemaType      SchemaType   `json:"schemaType"`
	TargetNamespace *string      `json:"targetNamespace,omitempty"`
}
