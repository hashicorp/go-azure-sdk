package componentpolicystates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ComponentPolicyStatesQueryResults
}

type ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions struct {
	Apply   *string
	Filter  *string
	From    *string
	Orderby *string
	Select  *string
	To      *string
	Top     *int64
}

func DefaultListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions() ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions {
	return ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions{}
}

func (o ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Apply != nil {
		out.Append("$apply", fmt.Sprintf("%v", *o.Apply))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.From != nil {
		out.Append("$from", fmt.Sprintf("%v", *o.From))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.To != nil {
		out.Append("$to", fmt.Sprintf("%v", *o.To))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListQueryResultsForResourceGroupLevelPolicyAssignment ...
func (c ComponentPolicyStatesClient) ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx context.Context, id AuthorizationNamespacePolicyAssignmentId, options ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions) (result ListQueryResultsForResourceGroupLevelPolicyAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/componentPolicyStates/latest/queryResults", id.ID()),
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

	var model ComponentPolicyStatesQueryResults
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
