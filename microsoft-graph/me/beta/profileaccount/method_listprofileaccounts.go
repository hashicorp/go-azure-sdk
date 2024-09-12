package profileaccount

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

type ListProfileAccountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserAccountInformation
}

type ListProfileAccountsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserAccountInformation
}

type ListProfileAccountsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListProfileAccountsOperationOptions() ListProfileAccountsOperationOptions {
	return ListProfileAccountsOperationOptions{}
}

func (o ListProfileAccountsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileAccountsOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileAccountsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileAccountsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileAccountsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileAccounts - List accounts. Retrieves properties related to the user's accounts from the profile.
func (c ProfileAccountClient) ListProfileAccounts(ctx context.Context, options ListProfileAccountsOperationOptions) (result ListProfileAccountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileAccountsCustomPager{},
		Path:          "/me/profile/account",
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
		Values *[]beta.UserAccountInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileAccountsComplete retrieves all the results into a single object
func (c ProfileAccountClient) ListProfileAccountsComplete(ctx context.Context, options ListProfileAccountsOperationOptions) (ListProfileAccountsCompleteResult, error) {
	return c.ListProfileAccountsCompleteMatchingPredicate(ctx, options, UserAccountInformationOperationPredicate{})
}

// ListProfileAccountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileAccountClient) ListProfileAccountsCompleteMatchingPredicate(ctx context.Context, options ListProfileAccountsOperationOptions, predicate UserAccountInformationOperationPredicate) (result ListProfileAccountsCompleteResult, err error) {
	items := make([]beta.UserAccountInformation, 0)

	resp, err := c.ListProfileAccounts(ctx, options)
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

	result = ListProfileAccountsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
