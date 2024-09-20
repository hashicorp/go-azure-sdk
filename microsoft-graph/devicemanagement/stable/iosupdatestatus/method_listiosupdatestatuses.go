package iosupdatestatus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIosUpdateStatusesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IosUpdateDeviceStatus
}

type ListIosUpdateStatusesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IosUpdateDeviceStatus
}

type ListIosUpdateStatusesOperationOptions struct {
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

func DefaultListIosUpdateStatusesOperationOptions() ListIosUpdateStatusesOperationOptions {
	return ListIosUpdateStatusesOperationOptions{}
}

func (o ListIosUpdateStatusesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListIosUpdateStatusesOperationOptions) ToOData() *odata.Query {
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

func (o ListIosUpdateStatusesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListIosUpdateStatusesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIosUpdateStatusesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIosUpdateStatuses - List iosUpdateDeviceStatuses. List properties and relationships of the iosUpdateDeviceStatus
// objects.
func (c IosUpdateStatusClient) ListIosUpdateStatuses(ctx context.Context, options ListIosUpdateStatusesOperationOptions) (result ListIosUpdateStatusesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListIosUpdateStatusesCustomPager{},
		Path:          "/deviceManagement/iosUpdateStatuses",
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
		Values *[]stable.IosUpdateDeviceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIosUpdateStatusesComplete retrieves all the results into a single object
func (c IosUpdateStatusClient) ListIosUpdateStatusesComplete(ctx context.Context, options ListIosUpdateStatusesOperationOptions) (ListIosUpdateStatusesCompleteResult, error) {
	return c.ListIosUpdateStatusesCompleteMatchingPredicate(ctx, options, IosUpdateDeviceStatusOperationPredicate{})
}

// ListIosUpdateStatusesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IosUpdateStatusClient) ListIosUpdateStatusesCompleteMatchingPredicate(ctx context.Context, options ListIosUpdateStatusesOperationOptions, predicate IosUpdateDeviceStatusOperationPredicate) (result ListIosUpdateStatusesCompleteResult, err error) {
	items := make([]stable.IosUpdateDeviceStatus, 0)

	resp, err := c.ListIosUpdateStatuses(ctx, options)
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

	result = ListIosUpdateStatusesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
