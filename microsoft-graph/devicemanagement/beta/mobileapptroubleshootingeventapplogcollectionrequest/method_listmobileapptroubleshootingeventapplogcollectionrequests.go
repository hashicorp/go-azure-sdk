package mobileapptroubleshootingeventapplogcollectionrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppLogCollectionRequest
}

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppLogCollectionRequest
}

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions() ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions {
	return ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions{}
}

func (o ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppTroubleshootingEventAppLogCollectionRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppTroubleshootingEventAppLogCollectionRequests - Get appLogCollectionRequests from deviceManagement. The
// collection property of AppLogUploadRequest.
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMobileAppTroubleshootingEventAppLogCollectionRequests(ctx context.Context, id beta.DeviceManagementMobileAppTroubleshootingEventId, options ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions) (result ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileAppTroubleshootingEventAppLogCollectionRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/appLogCollectionRequests", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.AppLogCollectionRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileAppTroubleshootingEventAppLogCollectionRequestsComplete retrieves all the results into a single object
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMobileAppTroubleshootingEventAppLogCollectionRequestsComplete(ctx context.Context, id beta.DeviceManagementMobileAppTroubleshootingEventId, options ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions) (ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult, error) {
	return c.ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteMatchingPredicate(ctx, id, options, AppLogCollectionRequestOperationPredicate{})
}

// ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementMobileAppTroubleshootingEventId, options ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationOptions, predicate AppLogCollectionRequestOperationPredicate) (result ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult, err error) {
	items := make([]beta.AppLogCollectionRequest, 0)

	resp, err := c.ListMobileAppTroubleshootingEventAppLogCollectionRequests(ctx, id, options)
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

	result = ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
