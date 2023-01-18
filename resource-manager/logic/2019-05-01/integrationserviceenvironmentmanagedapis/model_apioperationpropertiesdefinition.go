package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationPropertiesDefinition struct {
	Annotation          *ApiOperationAnnotation   `json:"annotation,omitempty"`
	Api                 *ApiReference             `json:"api,omitempty"`
	Description         *string                   `json:"description,omitempty"`
	InputsDefinition    *SwaggerSchema            `json:"inputsDefinition,omitempty"`
	IsNotification      *bool                     `json:"isNotification,omitempty"`
	IsWebhook           *bool                     `json:"isWebhook,omitempty"`
	Pageable            *bool                     `json:"pageable,omitempty"`
	ResponsesDefinition *map[string]SwaggerSchema `json:"responsesDefinition,omitempty"`
	Summary             *string                   `json:"summary,omitempty"`
	Trigger             *string                   `json:"trigger,omitempty"`
	TriggerHint         *string                   `json:"triggerHint,omitempty"`
	Visibility          *string                   `json:"visibility,omitempty"`
}
