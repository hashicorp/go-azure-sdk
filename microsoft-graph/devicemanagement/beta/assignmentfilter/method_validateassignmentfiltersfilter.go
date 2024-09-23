package assignmentfilter

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateAssignmentFiltersFilterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AssignmentFilterValidationResult
}

type ValidateAssignmentFiltersFilterOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultValidateAssignmentFiltersFilterOperationOptions() ValidateAssignmentFiltersFilterOperationOptions {
	return ValidateAssignmentFiltersFilterOperationOptions{}
}

func (o ValidateAssignmentFiltersFilterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateAssignmentFiltersFilterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ValidateAssignmentFiltersFilterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ValidateAssignmentFiltersFilter - Invoke action validateFilter
func (c AssignmentFilterClient) ValidateAssignmentFiltersFilter(ctx context.Context, input ValidateAssignmentFiltersFilterRequest, options ValidateAssignmentFiltersFilterOperationOptions) (result ValidateAssignmentFiltersFilterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/assignmentFilters/validateFilter",
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

	var model beta.AssignmentFilterValidationResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
