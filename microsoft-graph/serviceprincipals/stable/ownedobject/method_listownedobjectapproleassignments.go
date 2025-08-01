package ownedobject

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOwnedObjectAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AppRoleAssignment
}

type ListOwnedObjectAppRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AppRoleAssignment
}

type ListOwnedObjectAppRoleAssignmentsOperationOptions struct {
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

func DefaultListOwnedObjectAppRoleAssignmentsOperationOptions() ListOwnedObjectAppRoleAssignmentsOperationOptions {
	return ListOwnedObjectAppRoleAssignmentsOperationOptions{}
}

func (o ListOwnedObjectAppRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnedObjectAppRoleAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListOwnedObjectAppRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnedObjectAppRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnedObjectAppRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnedObjectAppRoleAssignments - Get the items of type microsoft.graph.appRoleAssignment in the
// microsoft.graph.directoryObject collection
func (c OwnedObjectClient) ListOwnedObjectAppRoleAssignments(ctx context.Context, id stable.ServicePrincipalId, options ListOwnedObjectAppRoleAssignmentsOperationOptions) (result ListOwnedObjectAppRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnedObjectAppRoleAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/ownedObjects/appRoleAssignment", id.ID()),
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
		Values *[]stable.AppRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnedObjectAppRoleAssignmentsComplete retrieves all the results into a single object
func (c OwnedObjectClient) ListOwnedObjectAppRoleAssignmentsComplete(ctx context.Context, id stable.ServicePrincipalId, options ListOwnedObjectAppRoleAssignmentsOperationOptions) (ListOwnedObjectAppRoleAssignmentsCompleteResult, error) {
	return c.ListOwnedObjectAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, options, AppRoleAssignmentOperationPredicate{})
}

// ListOwnedObjectAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnedObjectClient) ListOwnedObjectAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id stable.ServicePrincipalId, options ListOwnedObjectAppRoleAssignmentsOperationOptions, predicate AppRoleAssignmentOperationPredicate) (result ListOwnedObjectAppRoleAssignmentsCompleteResult, err error) {
	items := make([]stable.AppRoleAssignment, 0)

	resp, err := c.ListOwnedObjectAppRoleAssignments(ctx, id, options)
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

	result = ListOwnedObjectAppRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
