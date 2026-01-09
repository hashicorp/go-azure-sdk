package mobileappmanagementpolicy

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

type ListMobileAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MobilityManagementPolicy
}

type ListMobileAppManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MobilityManagementPolicy
}

type ListMobileAppManagementPoliciesOperationOptions struct {
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

func DefaultListMobileAppManagementPoliciesOperationOptions() ListMobileAppManagementPoliciesOperationOptions {
	return ListMobileAppManagementPoliciesOperationOptions{}
}

func (o ListMobileAppManagementPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileAppManagementPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListMobileAppManagementPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileAppManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppManagementPolicies - List mobileAppManagementPolicies. Get a list of the mobilityManagementPolicy
// objects and their properties.
func (c MobileAppManagementPolicyClient) ListMobileAppManagementPolicies(ctx context.Context, options ListMobileAppManagementPoliciesOperationOptions) (result ListMobileAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileAppManagementPoliciesCustomPager{},
		Path:          "/policies/mobileAppManagementPolicies",
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
		Values *[]beta.MobilityManagementPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileAppManagementPoliciesComplete retrieves all the results into a single object
func (c MobileAppManagementPolicyClient) ListMobileAppManagementPoliciesComplete(ctx context.Context, options ListMobileAppManagementPoliciesOperationOptions) (ListMobileAppManagementPoliciesCompleteResult, error) {
	return c.ListMobileAppManagementPoliciesCompleteMatchingPredicate(ctx, options, MobilityManagementPolicyOperationPredicate{})
}

// ListMobileAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppManagementPolicyClient) ListMobileAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, options ListMobileAppManagementPoliciesOperationOptions, predicate MobilityManagementPolicyOperationPredicate) (result ListMobileAppManagementPoliciesCompleteResult, err error) {
	items := make([]beta.MobilityManagementPolicy, 0)

	resp, err := c.ListMobileAppManagementPolicies(ctx, options)
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

	result = ListMobileAppManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
