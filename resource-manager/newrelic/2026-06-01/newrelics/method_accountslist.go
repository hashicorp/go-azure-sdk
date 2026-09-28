package newrelics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AccountResource
}

type AccountsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AccountResource
}

type AccountsListOperationOptions struct {
	Location  *string
	UserEmail *string
}

func DefaultAccountsListOperationOptions() AccountsListOperationOptions {
	return AccountsListOperationOptions{}
}

func (o AccountsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AccountsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AccountsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Location != nil {
		out.Append("location", fmt.Sprintf("%v", *o.Location))
	}
	if o.UserEmail != nil {
		out.Append("userEmail", fmt.Sprintf("%v", *o.UserEmail))
	}
	return &out
}

type AccountsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AccountsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AccountsList ...
func (c NewRelicsClient) AccountsList(ctx context.Context, id commonids.SubscriptionId, options AccountsListOperationOptions) (result AccountsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AccountsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/NewRelic.Observability/accounts", id.ID()),
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
		Values *[]AccountResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AccountsListComplete retrieves all the results into a single object
func (c NewRelicsClient) AccountsListComplete(ctx context.Context, id commonids.SubscriptionId, options AccountsListOperationOptions) (AccountsListCompleteResult, error) {
	return c.AccountsListCompleteMatchingPredicate(ctx, id, options, AccountResourceOperationPredicate{})
}

// AccountsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NewRelicsClient) AccountsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options AccountsListOperationOptions, predicate AccountResourceOperationPredicate) (result AccountsListCompleteResult, err error) {
	items := make([]AccountResource, 0)

	resp, err := c.AccountsList(ctx, id, options)
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

	result = AccountsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
