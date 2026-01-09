package registereddevice

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

type ListRegisteredDeviceAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppRoleAssignment
}

type ListRegisteredDeviceAppRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppRoleAssignment
}

type ListRegisteredDeviceAppRoleAssignmentsOperationOptions struct {
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

func DefaultListRegisteredDeviceAppRoleAssignmentsOperationOptions() ListRegisteredDeviceAppRoleAssignmentsOperationOptions {
	return ListRegisteredDeviceAppRoleAssignmentsOperationOptions{}
}

func (o ListRegisteredDeviceAppRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRegisteredDeviceAppRoleAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListRegisteredDeviceAppRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRegisteredDeviceAppRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRegisteredDeviceAppRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRegisteredDeviceAppRoleAssignments - Get the items of type microsoft.graph.appRoleAssignment in the
// microsoft.graph.directoryObject collection
func (c RegisteredDeviceClient) ListRegisteredDeviceAppRoleAssignments(ctx context.Context, id beta.UserId, options ListRegisteredDeviceAppRoleAssignmentsOperationOptions) (result ListRegisteredDeviceAppRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRegisteredDeviceAppRoleAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/registeredDevices/appRoleAssignment", id.ID()),
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
		Values *[]beta.AppRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRegisteredDeviceAppRoleAssignmentsComplete retrieves all the results into a single object
func (c RegisteredDeviceClient) ListRegisteredDeviceAppRoleAssignmentsComplete(ctx context.Context, id beta.UserId, options ListRegisteredDeviceAppRoleAssignmentsOperationOptions) (ListRegisteredDeviceAppRoleAssignmentsCompleteResult, error) {
	return c.ListRegisteredDeviceAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, options, AppRoleAssignmentOperationPredicate{})
}

// ListRegisteredDeviceAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegisteredDeviceClient) ListRegisteredDeviceAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListRegisteredDeviceAppRoleAssignmentsOperationOptions, predicate AppRoleAssignmentOperationPredicate) (result ListRegisteredDeviceAppRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.AppRoleAssignment, 0)

	resp, err := c.ListRegisteredDeviceAppRoleAssignments(ctx, id, options)
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

	result = ListRegisteredDeviceAppRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
