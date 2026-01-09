package defenderroleassignmentdirectoryscope

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

type ListDefenderRoleAssignmentDirectoryScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDefenderRoleAssignmentDirectoryScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDefenderRoleAssignmentDirectoryScopesOperationOptions struct {
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

func DefaultListDefenderRoleAssignmentDirectoryScopesOperationOptions() ListDefenderRoleAssignmentDirectoryScopesOperationOptions {
	return ListDefenderRoleAssignmentDirectoryScopesOperationOptions{}
}

func (o ListDefenderRoleAssignmentDirectoryScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDefenderRoleAssignmentDirectoryScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListDefenderRoleAssignmentDirectoryScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDefenderRoleAssignmentDirectoryScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDefenderRoleAssignmentDirectoryScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDefenderRoleAssignmentDirectoryScopes - Get directoryScopes from roleManagement. Read-only collection that
// references the directory objects that are scope of the assignment. Provided so that callers can get the directory
// objects using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (c DefenderRoleAssignmentDirectoryScopeClient) ListDefenderRoleAssignmentDirectoryScopes(ctx context.Context, id beta.RoleManagementDefenderRoleAssignmentId, options ListDefenderRoleAssignmentDirectoryScopesOperationOptions) (result ListDefenderRoleAssignmentDirectoryScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDefenderRoleAssignmentDirectoryScopesCustomPager{},
		Path:          fmt.Sprintf("%s/directoryScopes", id.ID()),
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

// ListDefenderRoleAssignmentDirectoryScopesComplete retrieves all the results into a single object
func (c DefenderRoleAssignmentDirectoryScopeClient) ListDefenderRoleAssignmentDirectoryScopesComplete(ctx context.Context, id beta.RoleManagementDefenderRoleAssignmentId, options ListDefenderRoleAssignmentDirectoryScopesOperationOptions) (ListDefenderRoleAssignmentDirectoryScopesCompleteResult, error) {
	return c.ListDefenderRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDefenderRoleAssignmentDirectoryScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DefenderRoleAssignmentDirectoryScopeClient) ListDefenderRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDefenderRoleAssignmentId, options ListDefenderRoleAssignmentDirectoryScopesOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDefenderRoleAssignmentDirectoryScopesCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDefenderRoleAssignmentDirectoryScopes(ctx, id, options)
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

	result = ListDefenderRoleAssignmentDirectoryScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
