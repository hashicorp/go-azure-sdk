package cloudpcroleassignmentprincipal

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

type ListCloudPCRoleAssignmentPrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListCloudPCRoleAssignmentPrincipalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListCloudPCRoleAssignmentPrincipalsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCloudPCRoleAssignmentPrincipalsOperationOptions() ListCloudPCRoleAssignmentPrincipalsOperationOptions {
	return ListCloudPCRoleAssignmentPrincipalsOperationOptions{}
}

func (o ListCloudPCRoleAssignmentPrincipalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudPCRoleAssignmentPrincipalsOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudPCRoleAssignmentPrincipalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudPCRoleAssignmentPrincipalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCRoleAssignmentPrincipalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCRoleAssignmentPrincipals - Get principals from roleManagement. Read-only collection that references the
// assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the
// role assignment. Read-only. Supports $expand.
func (c CloudPCRoleAssignmentPrincipalClient) ListCloudPCRoleAssignmentPrincipals(ctx context.Context, id beta.RoleManagementCloudPCRoleAssignmentId, options ListCloudPCRoleAssignmentPrincipalsOperationOptions) (result ListCloudPCRoleAssignmentPrincipalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudPCRoleAssignmentPrincipalsCustomPager{},
		Path:          fmt.Sprintf("%s/principals", id.ID()),
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

// ListCloudPCRoleAssignmentPrincipalsComplete retrieves all the results into a single object
func (c CloudPCRoleAssignmentPrincipalClient) ListCloudPCRoleAssignmentPrincipalsComplete(ctx context.Context, id beta.RoleManagementCloudPCRoleAssignmentId, options ListCloudPCRoleAssignmentPrincipalsOperationOptions) (ListCloudPCRoleAssignmentPrincipalsCompleteResult, error) {
	return c.ListCloudPCRoleAssignmentPrincipalsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListCloudPCRoleAssignmentPrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCRoleAssignmentPrincipalClient) ListCloudPCRoleAssignmentPrincipalsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementCloudPCRoleAssignmentId, options ListCloudPCRoleAssignmentPrincipalsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListCloudPCRoleAssignmentPrincipalsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListCloudPCRoleAssignmentPrincipals(ctx, id, options)
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

	result = ListCloudPCRoleAssignmentPrincipalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
