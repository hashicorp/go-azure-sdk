package homerealmdiscoverypolicyappliesto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListHomeRealmDiscoveryPolicyAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions struct {
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

func DefaultListHomeRealmDiscoveryPolicyAppliesTosOperationOptions() ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions {
	return ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions{}
}

func (o ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions) ToOData() *odata.Query {
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

func (o ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHomeRealmDiscoveryPolicyAppliesTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHomeRealmDiscoveryPolicyAppliesTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHomeRealmDiscoveryPolicyAppliesTos - List appliesTo. Get a list of directoryObject objects that a
// homeRealmDiscoveryPolicy object has been applied to.
func (c HomeRealmDiscoveryPolicyAppliesToClient) ListHomeRealmDiscoveryPolicyAppliesTos(ctx context.Context, id stable.PolicyHomeRealmDiscoveryPolicyId, options ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions) (result ListHomeRealmDiscoveryPolicyAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHomeRealmDiscoveryPolicyAppliesTosCustomPager{},
		Path:          fmt.Sprintf("%s/appliesTo", id.ID()),
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListHomeRealmDiscoveryPolicyAppliesTosComplete retrieves all the results into a single object
func (c HomeRealmDiscoveryPolicyAppliesToClient) ListHomeRealmDiscoveryPolicyAppliesTosComplete(ctx context.Context, id stable.PolicyHomeRealmDiscoveryPolicyId, options ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions) (ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult, error) {
	return c.ListHomeRealmDiscoveryPolicyAppliesTosCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListHomeRealmDiscoveryPolicyAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HomeRealmDiscoveryPolicyAppliesToClient) ListHomeRealmDiscoveryPolicyAppliesTosCompleteMatchingPredicate(ctx context.Context, id stable.PolicyHomeRealmDiscoveryPolicyId, options ListHomeRealmDiscoveryPolicyAppliesTosOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListHomeRealmDiscoveryPolicyAppliesTos(ctx, id, options)
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

	result = ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
