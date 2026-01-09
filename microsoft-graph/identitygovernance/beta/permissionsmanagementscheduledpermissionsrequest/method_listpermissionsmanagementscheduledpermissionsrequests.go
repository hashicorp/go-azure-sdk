package permissionsmanagementscheduledpermissionsrequest

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

type ListPermissionsManagementScheduledPermissionsRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ScheduledPermissionsRequest
}

type ListPermissionsManagementScheduledPermissionsRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ScheduledPermissionsRequest
}

type ListPermissionsManagementScheduledPermissionsRequestsOperationOptions struct {
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

func DefaultListPermissionsManagementScheduledPermissionsRequestsOperationOptions() ListPermissionsManagementScheduledPermissionsRequestsOperationOptions {
	return ListPermissionsManagementScheduledPermissionsRequestsOperationOptions{}
}

func (o ListPermissionsManagementScheduledPermissionsRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionsManagementScheduledPermissionsRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionsManagementScheduledPermissionsRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionsManagementScheduledPermissionsRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionsManagementScheduledPermissionsRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionsManagementScheduledPermissionsRequests - Get scheduledPermissionsRequests from identityGovernance.
// Represents a permissions request that Permissions Management uses to manage permissions for an identity on resources
// in the authorization system. This request can be granted, rejected or canceled by identities in Permissions
// Management.
func (c PermissionsManagementScheduledPermissionsRequestClient) ListPermissionsManagementScheduledPermissionsRequests(ctx context.Context, options ListPermissionsManagementScheduledPermissionsRequestsOperationOptions) (result ListPermissionsManagementScheduledPermissionsRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionsManagementScheduledPermissionsRequestsCustomPager{},
		Path:          "/identityGovernance/permissionsManagement/scheduledPermissionsRequests",
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
		Values *[]beta.ScheduledPermissionsRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionsManagementScheduledPermissionsRequestsComplete retrieves all the results into a single object
func (c PermissionsManagementScheduledPermissionsRequestClient) ListPermissionsManagementScheduledPermissionsRequestsComplete(ctx context.Context, options ListPermissionsManagementScheduledPermissionsRequestsOperationOptions) (ListPermissionsManagementScheduledPermissionsRequestsCompleteResult, error) {
	return c.ListPermissionsManagementScheduledPermissionsRequestsCompleteMatchingPredicate(ctx, options, ScheduledPermissionsRequestOperationPredicate{})
}

// ListPermissionsManagementScheduledPermissionsRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionsManagementScheduledPermissionsRequestClient) ListPermissionsManagementScheduledPermissionsRequestsCompleteMatchingPredicate(ctx context.Context, options ListPermissionsManagementScheduledPermissionsRequestsOperationOptions, predicate ScheduledPermissionsRequestOperationPredicate) (result ListPermissionsManagementScheduledPermissionsRequestsCompleteResult, err error) {
	items := make([]beta.ScheduledPermissionsRequest, 0)

	resp, err := c.ListPermissionsManagementScheduledPermissionsRequests(ctx, options)
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

	result = ListPermissionsManagementScheduledPermissionsRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
