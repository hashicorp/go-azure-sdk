package jobtargetexecutions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByStepOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JobExecution
}

type ListByStepCompleteResult struct {
	Items []JobExecution
}

type ListByStepOperationOptions struct {
	CreateTimeMax *string
	CreateTimeMin *string
	EndTimeMax    *string
	EndTimeMin    *string
	IsActive      *bool
	Skip          *int64
	Top           *int64
}

func DefaultListByStepOperationOptions() ListByStepOperationOptions {
	return ListByStepOperationOptions{}
}

func (o ListByStepOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByStepOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByStepOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.CreateTimeMax != nil {
		out.Append("createTimeMax", fmt.Sprintf("%v", *o.CreateTimeMax))
	}
	if o.CreateTimeMin != nil {
		out.Append("createTimeMin", fmt.Sprintf("%v", *o.CreateTimeMin))
	}
	if o.EndTimeMax != nil {
		out.Append("endTimeMax", fmt.Sprintf("%v", *o.EndTimeMax))
	}
	if o.EndTimeMin != nil {
		out.Append("endTimeMin", fmt.Sprintf("%v", *o.EndTimeMin))
	}
	if o.IsActive != nil {
		out.Append("isActive", fmt.Sprintf("%v", *o.IsActive))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByStep ...
func (c JobTargetExecutionsClient) ListByStep(ctx context.Context, id ExecutionStepId, options ListByStepOperationOptions) (result ListByStepOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/targets", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]JobExecution `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByStepComplete retrieves all the results into a single object
func (c JobTargetExecutionsClient) ListByStepComplete(ctx context.Context, id ExecutionStepId, options ListByStepOperationOptions) (ListByStepCompleteResult, error) {
	return c.ListByStepCompleteMatchingPredicate(ctx, id, options, JobExecutionOperationPredicate{})
}

// ListByStepCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JobTargetExecutionsClient) ListByStepCompleteMatchingPredicate(ctx context.Context, id ExecutionStepId, options ListByStepOperationOptions, predicate JobExecutionOperationPredicate) (result ListByStepCompleteResult, err error) {
	items := make([]JobExecution, 0)

	resp, err := c.ListByStep(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByStepCompleteResult{
		Items: items,
	}
	return
}
