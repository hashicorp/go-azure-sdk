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

type AppServiceEnvironmentsListWebAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Site
}

type AppServiceEnvironmentsListWebAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Site
}

type AppServiceEnvironmentsListWebAppsOperationOptions struct {
	PropertiesToInclude *string
}

func DefaultAppServiceEnvironmentsListWebAppsOperationOptions() AppServiceEnvironmentsListWebAppsOperationOptions {
	return AppServiceEnvironmentsListWebAppsOperationOptions{}
}

func (o AppServiceEnvironmentsListWebAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AppServiceEnvironmentsListWebAppsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AppServiceEnvironmentsListWebAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.PropertiesToInclude != nil {
		out.Append("propertiesToInclude", fmt.Sprintf("%v", *o.PropertiesToInclude))
	}
	return &out
}

type AppServiceEnvironmentsListWebAppsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListWebAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListWebApps ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListWebApps(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsListWebAppsOperationOptions) (result AppServiceEnvironmentsListWebAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AppServiceEnvironmentsListWebAppsCustomPager{},
		Path:          fmt.Sprintf("%s/sites", id.ID()),
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
		Values *[]Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListWebAppsComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListWebAppsComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsListWebAppsOperationOptions) (AppServiceEnvironmentsListWebAppsCompleteResult, error) {
	return c.AppServiceEnvironmentsListWebAppsCompleteMatchingPredicate(ctx, id, options, SiteOperationPredicate{})
}

// AppServiceEnvironmentsListWebAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListWebAppsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsListWebAppsOperationOptions, predicate SiteOperationPredicate) (result AppServiceEnvironmentsListWebAppsCompleteResult, err error) {
	items := make([]Site, 0)

	resp, err := c.AppServiceEnvironmentsListWebApps(ctx, id, options)
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

	result = AppServiceEnvironmentsListWebAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
