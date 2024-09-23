package serviceprincipalcreationpolicyexclude

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

type ListServicePrincipalCreationPolicyExcludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipalCreationConditionSet
}

type ListServicePrincipalCreationPolicyExcludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipalCreationConditionSet
}

type ListServicePrincipalCreationPolicyExcludesOperationOptions struct {
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

func DefaultListServicePrincipalCreationPolicyExcludesOperationOptions() ListServicePrincipalCreationPolicyExcludesOperationOptions {
	return ListServicePrincipalCreationPolicyExcludesOperationOptions{}
}

func (o ListServicePrincipalCreationPolicyExcludesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListServicePrincipalCreationPolicyExcludesOperationOptions) ToOData() *odata.Query {
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

func (o ListServicePrincipalCreationPolicyExcludesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListServicePrincipalCreationPolicyExcludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServicePrincipalCreationPolicyExcludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServicePrincipalCreationPolicyExcludes - Get excludes from policies
func (c ServicePrincipalCreationPolicyExcludeClient) ListServicePrincipalCreationPolicyExcludes(ctx context.Context, id beta.PolicyServicePrincipalCreationPolicyId, options ListServicePrincipalCreationPolicyExcludesOperationOptions) (result ListServicePrincipalCreationPolicyExcludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListServicePrincipalCreationPolicyExcludesCustomPager{},
		Path:          fmt.Sprintf("%s/excludes", id.ID()),
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
		Values *[]beta.ServicePrincipalCreationConditionSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalCreationPolicyExcludesComplete retrieves all the results into a single object
func (c ServicePrincipalCreationPolicyExcludeClient) ListServicePrincipalCreationPolicyExcludesComplete(ctx context.Context, id beta.PolicyServicePrincipalCreationPolicyId, options ListServicePrincipalCreationPolicyExcludesOperationOptions) (ListServicePrincipalCreationPolicyExcludesCompleteResult, error) {
	return c.ListServicePrincipalCreationPolicyExcludesCompleteMatchingPredicate(ctx, id, options, ServicePrincipalCreationConditionSetOperationPredicate{})
}

// ListServicePrincipalCreationPolicyExcludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalCreationPolicyExcludeClient) ListServicePrincipalCreationPolicyExcludesCompleteMatchingPredicate(ctx context.Context, id beta.PolicyServicePrincipalCreationPolicyId, options ListServicePrincipalCreationPolicyExcludesOperationOptions, predicate ServicePrincipalCreationConditionSetOperationPredicate) (result ListServicePrincipalCreationPolicyExcludesCompleteResult, err error) {
	items := make([]beta.ServicePrincipalCreationConditionSet, 0)

	resp, err := c.ListServicePrincipalCreationPolicyExcludes(ctx, id, options)
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

	result = ListServicePrincipalCreationPolicyExcludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
