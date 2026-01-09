package homerealmdiscoverypolicy

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

type ListHomeRealmDiscoveryPolicyRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListHomeRealmDiscoveryPolicyRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListHomeRealmDiscoveryPolicyRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListHomeRealmDiscoveryPolicyRefsOperationOptions() ListHomeRealmDiscoveryPolicyRefsOperationOptions {
	return ListHomeRealmDiscoveryPolicyRefsOperationOptions{}
}

func (o ListHomeRealmDiscoveryPolicyRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHomeRealmDiscoveryPolicyRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListHomeRealmDiscoveryPolicyRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHomeRealmDiscoveryPolicyRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHomeRealmDiscoveryPolicyRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHomeRealmDiscoveryPolicyRefs - List assigned homeRealmDiscoveryPolicy. List the homeRealmDiscoveryPolicy objects
// that are assigned to a servicePrincipal.
func (c HomeRealmDiscoveryPolicyClient) ListHomeRealmDiscoveryPolicyRefs(ctx context.Context, id beta.ServicePrincipalId, options ListHomeRealmDiscoveryPolicyRefsOperationOptions) (result ListHomeRealmDiscoveryPolicyRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHomeRealmDiscoveryPolicyRefsCustomPager{},
		Path:          fmt.Sprintf("%s/homeRealmDiscoveryPolicies/$ref", id.ID()),
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

// ListHomeRealmDiscoveryPolicyRefsComplete retrieves all the results into a single object
func (c HomeRealmDiscoveryPolicyClient) ListHomeRealmDiscoveryPolicyRefsComplete(ctx context.Context, id beta.ServicePrincipalId, options ListHomeRealmDiscoveryPolicyRefsOperationOptions) (ListHomeRealmDiscoveryPolicyRefsCompleteResult, error) {
	return c.ListHomeRealmDiscoveryPolicyRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListHomeRealmDiscoveryPolicyRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HomeRealmDiscoveryPolicyClient) ListHomeRealmDiscoveryPolicyRefsCompleteMatchingPredicate(ctx context.Context, id beta.ServicePrincipalId, options ListHomeRealmDiscoveryPolicyRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListHomeRealmDiscoveryPolicyRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListHomeRealmDiscoveryPolicyRefs(ctx, id, options)
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

	result = ListHomeRealmDiscoveryPolicyRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
