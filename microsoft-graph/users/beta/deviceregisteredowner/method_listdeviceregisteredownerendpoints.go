package deviceregisteredowner

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

type ListDeviceRegisteredOwnerEndpointsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Endpoint
}

type ListDeviceRegisteredOwnerEndpointsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Endpoint
}

type ListDeviceRegisteredOwnerEndpointsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListDeviceRegisteredOwnerEndpointsOperationOptions() ListDeviceRegisteredOwnerEndpointsOperationOptions {
	return ListDeviceRegisteredOwnerEndpointsOperationOptions{}
}

func (o ListDeviceRegisteredOwnerEndpointsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceRegisteredOwnerEndpointsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o ListDeviceRegisteredOwnerEndpointsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceRegisteredOwnerEndpointsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceRegisteredOwnerEndpointsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceRegisteredOwnerEndpoints - Get the items of type microsoft.graph.endpoint in the
// microsoft.graph.directoryObject collection
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnerEndpoints(ctx context.Context, id beta.UserIdDeviceId, options ListDeviceRegisteredOwnerEndpointsOperationOptions) (result ListDeviceRegisteredOwnerEndpointsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceRegisteredOwnerEndpointsCustomPager{},
		Path:          fmt.Sprintf("%s/registeredOwners/endpoint", id.ID()),
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
		Values *[]beta.Endpoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceRegisteredOwnerEndpointsComplete retrieves all the results into a single object
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnerEndpointsComplete(ctx context.Context, id beta.UserIdDeviceId, options ListDeviceRegisteredOwnerEndpointsOperationOptions) (ListDeviceRegisteredOwnerEndpointsCompleteResult, error) {
	return c.ListDeviceRegisteredOwnerEndpointsCompleteMatchingPredicate(ctx, id, options, EndpointOperationPredicate{})
}

// ListDeviceRegisteredOwnerEndpointsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnerEndpointsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDeviceId, options ListDeviceRegisteredOwnerEndpointsOperationOptions, predicate EndpointOperationPredicate) (result ListDeviceRegisteredOwnerEndpointsCompleteResult, err error) {
	items := make([]beta.Endpoint, 0)

	resp, err := c.ListDeviceRegisteredOwnerEndpoints(ctx, id, options)
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

	result = ListDeviceRegisteredOwnerEndpointsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
