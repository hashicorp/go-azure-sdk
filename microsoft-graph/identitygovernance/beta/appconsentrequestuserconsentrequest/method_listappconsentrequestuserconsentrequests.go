package appconsentrequestuserconsentrequest

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

type ListAppConsentRequestUserConsentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserConsentRequest
}

type ListAppConsentRequestUserConsentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserConsentRequest
}

type ListAppConsentRequestUserConsentRequestsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAppConsentRequestUserConsentRequestsOperationOptions() ListAppConsentRequestUserConsentRequestsOperationOptions {
	return ListAppConsentRequestUserConsentRequestsOperationOptions{}
}

func (o ListAppConsentRequestUserConsentRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppConsentRequestUserConsentRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListAppConsentRequestUserConsentRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppConsentRequestUserConsentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestUserConsentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestUserConsentRequests - List userConsentRequests. Retrieve a collection of userConsentRequest
// objects and their properties.
func (c AppConsentRequestUserConsentRequestClient) ListAppConsentRequestUserConsentRequests(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestId, options ListAppConsentRequestUserConsentRequestsOperationOptions) (result ListAppConsentRequestUserConsentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppConsentRequestUserConsentRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/userConsentRequests", id.ID()),
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
		Values *[]beta.UserConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentRequestUserConsentRequestsComplete retrieves all the results into a single object
func (c AppConsentRequestUserConsentRequestClient) ListAppConsentRequestUserConsentRequestsComplete(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestId, options ListAppConsentRequestUserConsentRequestsOperationOptions) (ListAppConsentRequestUserConsentRequestsCompleteResult, error) {
	return c.ListAppConsentRequestUserConsentRequestsCompleteMatchingPredicate(ctx, id, options, UserConsentRequestOperationPredicate{})
}

// ListAppConsentRequestUserConsentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentRequestUserConsentRequestClient) ListAppConsentRequestUserConsentRequestsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestId, options ListAppConsentRequestUserConsentRequestsOperationOptions, predicate UserConsentRequestOperationPredicate) (result ListAppConsentRequestUserConsentRequestsCompleteResult, err error) {
	items := make([]beta.UserConsentRequest, 0)

	resp, err := c.ListAppConsentRequestUserConsentRequests(ctx, id, options)
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

	result = ListAppConsentRequestUserConsentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
