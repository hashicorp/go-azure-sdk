package plannerall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPlannerAllsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerDelta
}

type ListPlannerAllsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerDelta
}

type ListPlannerAllsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListPlannerAllsOperationOptions() ListPlannerAllsOperationOptions {
	return ListPlannerAllsOperationOptions{}
}

func (o ListPlannerAllsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPlannerAllsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListPlannerAllsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPlannerAllsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerAllsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerAlls - Get all from users
func (c PlannerAllClient) ListPlannerAlls(ctx context.Context, id beta.UserId, options ListPlannerAllsOperationOptions) (result ListPlannerAllsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPlannerAllsCustomPager{},
		Path:          fmt.Sprintf("%s/planner/all", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.PlannerDelta, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalPlannerDeltaImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.PlannerDelta (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListPlannerAllsComplete retrieves all the results into a single object
func (c PlannerAllClient) ListPlannerAllsComplete(ctx context.Context, id beta.UserId, options ListPlannerAllsOperationOptions) (ListPlannerAllsCompleteResult, error) {
	return c.ListPlannerAllsCompleteMatchingPredicate(ctx, id, options, PlannerDeltaOperationPredicate{})
}

// ListPlannerAllsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerAllClient) ListPlannerAllsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListPlannerAllsOperationOptions, predicate PlannerDeltaOperationPredicate) (result ListPlannerAllsCompleteResult, err error) {
	items := make([]beta.PlannerDelta, 0)

	resp, err := c.ListPlannerAlls(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListPlannerAllsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
