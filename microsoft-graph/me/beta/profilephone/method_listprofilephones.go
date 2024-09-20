package profilephone

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

type ListProfilePhonesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemPhone
}

type ListProfilePhonesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemPhone
}

type ListProfilePhonesOperationOptions struct {
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

func DefaultListProfilePhonesOperationOptions() ListProfilePhonesOperationOptions {
	return ListProfilePhonesOperationOptions{}
}

func (o ListProfilePhonesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfilePhonesOperationOptions) ToOData() *odata.Query {
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

func (o ListProfilePhonesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfilePhonesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfilePhonesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfilePhones - List phones. Retrieve a list of itemPhone objects from a user's profile.
func (c ProfilePhoneClient) ListProfilePhones(ctx context.Context, options ListProfilePhonesOperationOptions) (result ListProfilePhonesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfilePhonesCustomPager{},
		Path:          "/me/profile/phones",
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
		Values *[]beta.ItemPhone `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfilePhonesComplete retrieves all the results into a single object
func (c ProfilePhoneClient) ListProfilePhonesComplete(ctx context.Context, options ListProfilePhonesOperationOptions) (ListProfilePhonesCompleteResult, error) {
	return c.ListProfilePhonesCompleteMatchingPredicate(ctx, options, ItemPhoneOperationPredicate{})
}

// ListProfilePhonesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfilePhoneClient) ListProfilePhonesCompleteMatchingPredicate(ctx context.Context, options ListProfilePhonesOperationOptions, predicate ItemPhoneOperationPredicate) (result ListProfilePhonesCompleteResult, err error) {
	items := make([]beta.ItemPhone, 0)

	resp, err := c.ListProfilePhones(ctx, options)
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

	result = ListProfilePhonesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
