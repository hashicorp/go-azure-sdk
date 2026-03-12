package appserviceenvironmentresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentsListUsagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CsmUsageQuota
}

type AppServiceEnvironmentsListUsagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CsmUsageQuota
}

type AppServiceEnvironmentsListUsagesOperationOptions struct {
	Filter *string
}

func DefaultAppServiceEnvironmentsListUsagesOperationOptions() AppServiceEnvironmentsListUsagesOperationOptions {
	return AppServiceEnvironmentsListUsagesOperationOptions{}
}

func (o AppServiceEnvironmentsListUsagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AppServiceEnvironmentsListUsagesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AppServiceEnvironmentsListUsagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type AppServiceEnvironmentsListUsagesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListUsagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListUsages ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListUsages(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsListUsagesOperationOptions) (result AppServiceEnvironmentsListUsagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AppServiceEnvironmentsListUsagesCustomPager{},
		Path:          fmt.Sprintf("%s/usages", id.ID()),
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
		Values *[]CsmUsageQuota `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListUsagesComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListUsagesComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsListUsagesOperationOptions) (AppServiceEnvironmentsListUsagesCompleteResult, error) {
	return c.AppServiceEnvironmentsListUsagesCompleteMatchingPredicate(ctx, id, options, CsmUsageQuotaOperationPredicate{})
}

// AppServiceEnvironmentsListUsagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListUsagesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsListUsagesOperationOptions, predicate CsmUsageQuotaOperationPredicate) (result AppServiceEnvironmentsListUsagesCompleteResult, err error) {
	items := make([]CsmUsageQuota, 0)

	resp, err := c.AppServiceEnvironmentsListUsages(ctx, id, options)
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

	result = AppServiceEnvironmentsListUsagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
