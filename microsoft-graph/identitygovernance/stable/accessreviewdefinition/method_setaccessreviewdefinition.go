package accessreviewdefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetAccessReviewDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetAccessReviewDefinitionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetAccessReviewDefinitionOperationOptions() SetAccessReviewDefinitionOperationOptions {
	return SetAccessReviewDefinitionOperationOptions{}
}

func (o SetAccessReviewDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetAccessReviewDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetAccessReviewDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetAccessReviewDefinition - Update accessReviewScheduleDefinition. Update an existing accessReviewScheduleDefinition
// object to change one or more of its properties.
func (c AccessReviewDefinitionClient) SetAccessReviewDefinition(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionId, input stable.AccessReviewScheduleDefinition, options SetAccessReviewDefinitionOperationOptions) (result SetAccessReviewDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
