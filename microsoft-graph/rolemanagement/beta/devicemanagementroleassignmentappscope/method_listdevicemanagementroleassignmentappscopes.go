package devicemanagementroleassignmentappscope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeviceManagementRoleAssignmentAppScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppScope
}

type ListDeviceManagementRoleAssignmentAppScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppScope
}

type ListDeviceManagementRoleAssignmentAppScopesOperationOptions struct {
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

func DefaultListDeviceManagementRoleAssignmentAppScopesOperationOptions() ListDeviceManagementRoleAssignmentAppScopesOperationOptions {
	return ListDeviceManagementRoleAssignmentAppScopesOperationOptions{}
}

func (o ListDeviceManagementRoleAssignmentAppScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementRoleAssignmentAppScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementRoleAssignmentAppScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementRoleAssignmentAppScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentAppScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignmentAppScopes - Get appScopes from roleManagement. Read-only collection with details of
// the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
func (c DeviceManagementRoleAssignmentAppScopeClient) ListDeviceManagementRoleAssignmentAppScopes(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentAppScopesOperationOptions) (result ListDeviceManagementRoleAssignmentAppScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementRoleAssignmentAppScopesCustomPager{},
		Path:          fmt.Sprintf("%s/appScopes", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.AppScope, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalAppScopeImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.AppScope (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDeviceManagementRoleAssignmentAppScopesComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentAppScopeClient) ListDeviceManagementRoleAssignmentAppScopesComplete(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentAppScopesOperationOptions) (ListDeviceManagementRoleAssignmentAppScopesCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentAppScopesCompleteMatchingPredicate(ctx, id, options, AppScopeOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentAppScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentAppScopeClient) ListDeviceManagementRoleAssignmentAppScopesCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentAppScopesOperationOptions, predicate AppScopeOperationPredicate) (result ListDeviceManagementRoleAssignmentAppScopesCompleteResult, err error) {
	items := make([]beta.AppScope, 0)

	resp, err := c.ListDeviceManagementRoleAssignmentAppScopes(ctx, id, options)
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

	result = ListDeviceManagementRoleAssignmentAppScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
