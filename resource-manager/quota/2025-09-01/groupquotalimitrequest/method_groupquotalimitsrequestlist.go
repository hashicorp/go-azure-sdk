package groupquotalimitrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaLimitsRequestListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubmittedResourceRequestStatus
}

type GroupQuotaLimitsRequestListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubmittedResourceRequestStatus
}

type GroupQuotaLimitsRequestListOperationOptions struct {
	Filter *string
}

func DefaultGroupQuotaLimitsRequestListOperationOptions() GroupQuotaLimitsRequestListOperationOptions {
	return GroupQuotaLimitsRequestListOperationOptions{}
}

func (o GroupQuotaLimitsRequestListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GroupQuotaLimitsRequestListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GroupQuotaLimitsRequestListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type GroupQuotaLimitsRequestListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GroupQuotaLimitsRequestListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GroupQuotaLimitsRequestList ...
func (c GroupQuotaLimitRequestClient) GroupQuotaLimitsRequestList(ctx context.Context, id ResourceProviderId, options GroupQuotaLimitsRequestListOperationOptions) (result GroupQuotaLimitsRequestListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GroupQuotaLimitsRequestListCustomPager{},
		Path:          fmt.Sprintf("%s/groupQuotaRequests", id.ID()),
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
		Values *[]SubmittedResourceRequestStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GroupQuotaLimitsRequestListComplete retrieves all the results into a single object
func (c GroupQuotaLimitRequestClient) GroupQuotaLimitsRequestListComplete(ctx context.Context, id ResourceProviderId, options GroupQuotaLimitsRequestListOperationOptions) (GroupQuotaLimitsRequestListCompleteResult, error) {
	return c.GroupQuotaLimitsRequestListCompleteMatchingPredicate(ctx, id, options, SubmittedResourceRequestStatusOperationPredicate{})
}

// GroupQuotaLimitsRequestListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupQuotaLimitRequestClient) GroupQuotaLimitsRequestListCompleteMatchingPredicate(ctx context.Context, id ResourceProviderId, options GroupQuotaLimitsRequestListOperationOptions, predicate SubmittedResourceRequestStatusOperationPredicate) (result GroupQuotaLimitsRequestListCompleteResult, err error) {
	items := make([]SubmittedResourceRequestStatus, 0)

	resp, err := c.GroupQuotaLimitsRequestList(ctx, id, options)
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

	result = GroupQuotaLimitsRequestListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
