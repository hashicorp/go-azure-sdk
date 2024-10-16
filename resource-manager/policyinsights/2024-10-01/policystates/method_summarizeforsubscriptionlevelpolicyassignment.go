package policystates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummarizeForSubscriptionLevelPolicyAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SummarizeResults
}

type SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions struct {
	Filter *string
	From   *string
	To     *string
	Top    *int64
}

func DefaultSummarizeForSubscriptionLevelPolicyAssignmentOperationOptions() SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions {
	return SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions{}
}

func (o SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.From != nil {
		out.Append("$from", fmt.Sprintf("%v", *o.From))
	}
	if o.To != nil {
		out.Append("$to", fmt.Sprintf("%v", *o.To))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// SummarizeForSubscriptionLevelPolicyAssignment ...
func (c PolicyStatesClient) SummarizeForSubscriptionLevelPolicyAssignment(ctx context.Context, id PolicyAssignmentId, options SummarizeForSubscriptionLevelPolicyAssignmentOperationOptions) (result SummarizeForSubscriptionLevelPolicyAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/policyStates/latest/summarize", id.ID()),
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

	var model SummarizeResults
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
