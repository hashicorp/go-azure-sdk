package serviceprincipalsynchronizationjobschema

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParseServicePrincipalByIdSynchronizationJobByIdSchemaExpressionRequest struct {
	Expression                *string                `json:"expression,omitempty"`
	TargetAttributeDefinition *AttributeDefinition   `json:"targetAttributeDefinition,omitempty"`
	TestInputObject           *ExpressionInputObject `json:"testInputObject,omitempty"`
}
