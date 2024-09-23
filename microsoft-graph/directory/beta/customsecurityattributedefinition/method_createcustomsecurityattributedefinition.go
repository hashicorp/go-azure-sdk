package customsecurityattributedefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCustomSecurityAttributeDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CustomSecurityAttributeDefinition
}

type CreateCustomSecurityAttributeDefinitionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCustomSecurityAttributeDefinitionOperationOptions() CreateCustomSecurityAttributeDefinitionOperationOptions {
	return CreateCustomSecurityAttributeDefinitionOperationOptions{}
}

func (o CreateCustomSecurityAttributeDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCustomSecurityAttributeDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCustomSecurityAttributeDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCustomSecurityAttributeDefinition - Create customSecurityAttributeDefinition. Create a new
// customSecurityAttributeDefinition object.
func (c CustomSecurityAttributeDefinitionClient) CreateCustomSecurityAttributeDefinition(ctx context.Context, input beta.CustomSecurityAttributeDefinition, options CreateCustomSecurityAttributeDefinitionOperationOptions) (result CreateCustomSecurityAttributeDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/directory/customSecurityAttributeDefinitions",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.CustomSecurityAttributeDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
