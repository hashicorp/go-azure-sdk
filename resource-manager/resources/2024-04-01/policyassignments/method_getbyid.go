package policyassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PolicyAssignment
}

type GetByIdOperationOptions struct {
	Expand *string
}

func DefaultGetByIdOperationOptions() GetByIdOperationOptions {
	return GetByIdOperationOptions{}
}

func (o GetByIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetByIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetByIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

// GetById ...
func (c PolicyAssignmentsClient) GetById(ctx context.Context, id PolicyAssignmentIdId, options GetByIdOperationOptions) (result GetByIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model PolicyAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
