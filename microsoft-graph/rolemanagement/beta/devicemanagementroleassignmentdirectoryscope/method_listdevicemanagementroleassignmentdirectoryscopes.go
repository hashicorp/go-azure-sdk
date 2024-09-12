package devicemanagementroleassignmentdirectoryscope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeviceManagementRoleAssignmentDirectoryScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions() ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions {
	return ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions{}
}

func (o ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementRoleAssignmentDirectoryScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentDirectoryScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignmentDirectoryScopes - Get directoryScopes from roleManagement. Read-only collection
// that references the directory objects that are scope of the assignment. Provided so that callers can get the
// directory objects using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (c DeviceManagementRoleAssignmentDirectoryScopeClient) ListDeviceManagementRoleAssignmentDirectoryScopes(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions) (result ListDeviceManagementRoleAssignmentDirectoryScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementRoleAssignmentDirectoryScopesCustomPager{},
		Path:          fmt.Sprintf("%s/directoryScopes", id.ID()),
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

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDeviceManagementRoleAssignmentDirectoryScopesComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentDirectoryScopeClient) ListDeviceManagementRoleAssignmentDirectoryScopesComplete(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions) (ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentDirectoryScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentDirectoryScopeClient) ListDeviceManagementRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentDirectoryScopesOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceManagementRoleAssignmentDirectoryScopes(ctx, id, options)
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

	result = ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
