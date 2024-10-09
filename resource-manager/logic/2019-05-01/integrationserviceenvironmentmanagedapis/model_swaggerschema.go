package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerSchema struct {
	AdditionalProperties     *interface{}                    `json:"additionalProperties,omitempty"`
	AllOf                    *[]SwaggerSchema                `json:"allOf,omitempty"`
	Discriminator            *string                         `json:"discriminator,omitempty"`
	DynamicListNew           *SwaggerCustomDynamicList       `json:"dynamicListNew,omitempty"`
	DynamicSchemaNew         *SwaggerCustomDynamicProperties `json:"dynamicSchemaNew,omitempty"`
	DynamicSchemaOld         *SwaggerCustomDynamicSchema     `json:"dynamicSchemaOld,omitempty"`
	DynamicTree              *SwaggerCustomDynamicTree       `json:"dynamicTree,omitempty"`
	Example                  *interface{}                    `json:"example,omitempty"`
	ExternalDocs             *SwaggerExternalDocumentation   `json:"externalDocs,omitempty"`
	Items                    *SwaggerSchema                  `json:"items,omitempty"`
	MaxProperties            *int64                          `json:"maxProperties,omitempty"`
	MinProperties            *int64                          `json:"minProperties,omitempty"`
	NotificationURLExtension *bool                           `json:"notificationUrlExtension,omitempty"`
	Properties               *map[string]SwaggerSchema       `json:"properties,omitempty"`
	ReadOnly                 *bool                           `json:"readOnly,omitempty"`
	Ref                      *string                         `json:"ref,omitempty"`
	Required                 *[]string                       `json:"required,omitempty"`
	Title                    *string                         `json:"title,omitempty"`
	Type                     *SwaggerSchemaType              `json:"type,omitempty"`
	Xml                      *SwaggerXml                     `json:"xml,omitempty"`
}
