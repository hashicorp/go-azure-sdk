package entitlementmanagementsubject

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementSubjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageSubject
}

type ListEntitlementManagementSubjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageSubject
}

type ListEntitlementManagementSubjectsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementSubjectsOperationOptions() ListEntitlementManagementSubjectsOperationOptions {
	return ListEntitlementManagementSubjectsOperationOptions{}
}

func (o ListEntitlementManagementSubjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementSubjectsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementSubjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementSubjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementSubjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementSubjects - Get accessPackageSubject. Get the properties of an existing accessPackageSubject
// object.
func (c EntitlementManagementSubjectClient) ListEntitlementManagementSubjects(ctx context.Context, options ListEntitlementManagementSubjectsOperationOptions) (result ListEntitlementManagementSubjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementSubjectsCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/subjects",
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
		Values *[]beta.AccessPackageSubject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementSubjectsComplete retrieves all the results into a single object
func (c EntitlementManagementSubjectClient) ListEntitlementManagementSubjectsComplete(ctx context.Context, options ListEntitlementManagementSubjectsOperationOptions) (ListEntitlementManagementSubjectsCompleteResult, error) {
	return c.ListEntitlementManagementSubjectsCompleteMatchingPredicate(ctx, options, AccessPackageSubjectOperationPredicate{})
}

// ListEntitlementManagementSubjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementSubjectClient) ListEntitlementManagementSubjectsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementSubjectsOperationOptions, predicate AccessPackageSubjectOperationPredicate) (result ListEntitlementManagementSubjectsCompleteResult, err error) {
	items := make([]beta.AccessPackageSubject, 0)

	resp, err := c.ListEntitlementManagementSubjects(ctx, options)
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

	result = ListEntitlementManagementSubjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
