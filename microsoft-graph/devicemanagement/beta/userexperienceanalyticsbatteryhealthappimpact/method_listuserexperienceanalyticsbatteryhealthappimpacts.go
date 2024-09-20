package userexperienceanalyticsbatteryhealthappimpact

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBatteryHealthAppImpact
}

type ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBatteryHealthAppImpact
}

type ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions() ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions {
	return ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions{}
}

func (o ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsBatteryHealthAppImpactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBatteryHealthAppImpactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBatteryHealthAppImpacts - Get userExperienceAnalyticsBatteryHealthAppImpact from
// deviceManagement. User Experience Analytics Battery Health App Impact
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) ListUserExperienceAnalyticsBatteryHealthAppImpacts(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions) (result ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsBatteryHealthAppImpactsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact",
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
		Values *[]beta.UserExperienceAnalyticsBatteryHealthAppImpact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBatteryHealthAppImpactsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) ListUserExperienceAnalyticsBatteryHealthAppImpactsComplete(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions) (ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsBatteryHealthAppImpactOperationPredicate{})
}

// ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationOptions, predicate UserExperienceAnalyticsBatteryHealthAppImpactOperationPredicate) (result ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBatteryHealthAppImpact, 0)

	resp, err := c.ListUserExperienceAnalyticsBatteryHealthAppImpacts(ctx, options)
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

	result = ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
