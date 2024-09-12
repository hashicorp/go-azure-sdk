package synchronizationtemplateschema

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParseSynchronizationTemplateSchemaExpressionRequest struct {
	Expression                nullable.Type[string]       `json:"expression,omitempty"`
	TargetAttributeDefinition *beta.AttributeDefinition   `json:"targetAttributeDefinition,omitempty"`
	TestInputObject           *beta.ExpressionInputObject `json:"testInputObject,omitempty"`
}
