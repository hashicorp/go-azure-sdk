package devicemanagementroleassignmentprincipal

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

type ListDeviceManagementRoleAssignmentPrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentPrincipalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentPrincipalsOperationOptions struct {
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

func DefaultListDeviceManagementRoleAssignmentPrincipalsOperationOptions() ListDeviceManagementRoleAssignmentPrincipalsOperationOptions {
	return ListDeviceManagementRoleAssignmentPrincipalsOperationOptions{}
}

func (o ListDeviceManagementRoleAssignmentPrincipalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementRoleAssignmentPrincipalsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementRoleAssignmentPrincipalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementRoleAssignmentPrincipalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentPrincipalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignmentPrincipals - Get principals from roleManagement. Read-only collection that
// references the assigned principals. Provided so that callers can get the principals using $expand at the same time as
// getting the role assignment. Read-only. Supports $expand.
func (c DeviceManagementRoleAssignmentPrincipalClient) ListDeviceManagementRoleAssignmentPrincipals(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentPrincipalsOperationOptions) (result ListDeviceManagementRoleAssignmentPrincipalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementRoleAssignmentPrincipalsCustomPager{},
		Path:          fmt.Sprintf("%s/principals", id.ID()),
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

// ListDeviceManagementRoleAssignmentPrincipalsComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentPrincipalClient) ListDeviceManagementRoleAssignmentPrincipalsComplete(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentPrincipalsOperationOptions) (ListDeviceManagementRoleAssignmentPrincipalsCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentPrincipalsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentPrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentPrincipalClient) ListDeviceManagementRoleAssignmentPrincipalsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options ListDeviceManagementRoleAssignmentPrincipalsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDeviceManagementRoleAssignmentPrincipalsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceManagementRoleAssignmentPrincipals(ctx, id, options)
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

	result = ListDeviceManagementRoleAssignmentPrincipalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
