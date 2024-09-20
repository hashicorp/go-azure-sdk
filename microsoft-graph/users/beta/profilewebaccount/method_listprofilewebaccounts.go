package profilewebaccount

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

type ListProfileWebAccountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WebAccount
}

type ListProfileWebAccountsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WebAccount
}

type ListProfileWebAccountsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListProfileWebAccountsOperationOptions() ListProfileWebAccountsOperationOptions {
	return ListProfileWebAccountsOperationOptions{}
}

func (o ListProfileWebAccountsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileWebAccountsOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileWebAccountsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileWebAccountsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileWebAccountsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileWebAccounts - Get webAccounts from users. Represents web accounts the user has indicated they use or has
// added to their user profile.
func (c ProfileWebAccountClient) ListProfileWebAccounts(ctx context.Context, id beta.UserId, options ListProfileWebAccountsOperationOptions) (result ListProfileWebAccountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileWebAccountsCustomPager{},
		Path:          fmt.Sprintf("%s/profile/webAccounts", id.ID()),
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
		Values *[]beta.WebAccount `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileWebAccountsComplete retrieves all the results into a single object
func (c ProfileWebAccountClient) ListProfileWebAccountsComplete(ctx context.Context, id beta.UserId, options ListProfileWebAccountsOperationOptions) (ListProfileWebAccountsCompleteResult, error) {
	return c.ListProfileWebAccountsCompleteMatchingPredicate(ctx, id, options, WebAccountOperationPredicate{})
}

// ListProfileWebAccountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileWebAccountClient) ListProfileWebAccountsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfileWebAccountsOperationOptions, predicate WebAccountOperationPredicate) (result ListProfileWebAccountsCompleteResult, err error) {
	items := make([]beta.WebAccount, 0)

	resp, err := c.ListProfileWebAccounts(ctx, id, options)
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

	result = ListProfileWebAccountsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
