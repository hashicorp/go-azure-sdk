package comanageddevicelogcollectionrequest

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

type ListComanagedDeviceLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceLogCollectionResponse
}

type ListComanagedDeviceLogCollectionRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceLogCollectionResponse
}

type ListComanagedDeviceLogCollectionRequestsOperationOptions struct {
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

func DefaultListComanagedDeviceLogCollectionRequestsOperationOptions() ListComanagedDeviceLogCollectionRequestsOperationOptions {
	return ListComanagedDeviceLogCollectionRequestsOperationOptions{}
}

func (o ListComanagedDeviceLogCollectionRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListComanagedDeviceLogCollectionRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListComanagedDeviceLogCollectionRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListComanagedDeviceLogCollectionRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceLogCollectionRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceLogCollectionRequests - Get logCollectionRequests from deviceManagement. List of log collection
// requests
func (c ComanagedDeviceLogCollectionRequestClient) ListComanagedDeviceLogCollectionRequests(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceLogCollectionRequestsOperationOptions) (result ListComanagedDeviceLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListComanagedDeviceLogCollectionRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/logCollectionRequests", id.ID()),
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
		Values *[]beta.DeviceLogCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceLogCollectionRequestsComplete retrieves all the results into a single object
func (c ComanagedDeviceLogCollectionRequestClient) ListComanagedDeviceLogCollectionRequestsComplete(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceLogCollectionRequestsOperationOptions) (ListComanagedDeviceLogCollectionRequestsCompleteResult, error) {
	return c.ListComanagedDeviceLogCollectionRequestsCompleteMatchingPredicate(ctx, id, options, DeviceLogCollectionResponseOperationPredicate{})
}

// ListComanagedDeviceLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceLogCollectionRequestClient) ListComanagedDeviceLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceLogCollectionRequestsOperationOptions, predicate DeviceLogCollectionResponseOperationPredicate) (result ListComanagedDeviceLogCollectionRequestsCompleteResult, err error) {
	items := make([]beta.DeviceLogCollectionResponse, 0)

	resp, err := c.ListComanagedDeviceLogCollectionRequests(ctx, id, options)
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

	result = ListComanagedDeviceLogCollectionRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
