package mobileappintentandstate

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

type ListMobileAppIntentAndStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MobileAppIntentAndState
}

type ListMobileAppIntentAndStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MobileAppIntentAndState
}

type ListMobileAppIntentAndStatesOperationOptions struct {
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

func DefaultListMobileAppIntentAndStatesOperationOptions() ListMobileAppIntentAndStatesOperationOptions {
	return ListMobileAppIntentAndStatesOperationOptions{}
}

func (o ListMobileAppIntentAndStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileAppIntentAndStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListMobileAppIntentAndStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileAppIntentAndStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppIntentAndStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppIntentAndStates - Get mobileAppIntentAndStates from me. The list of troubleshooting events for this
// user.
func (c MobileAppIntentAndStateClient) ListMobileAppIntentAndStates(ctx context.Context, options ListMobileAppIntentAndStatesOperationOptions) (result ListMobileAppIntentAndStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileAppIntentAndStatesCustomPager{},
		Path:          "/me/mobileAppIntentAndStates",
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
		Values *[]beta.MobileAppIntentAndState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileAppIntentAndStatesComplete retrieves all the results into a single object
func (c MobileAppIntentAndStateClient) ListMobileAppIntentAndStatesComplete(ctx context.Context, options ListMobileAppIntentAndStatesOperationOptions) (ListMobileAppIntentAndStatesCompleteResult, error) {
	return c.ListMobileAppIntentAndStatesCompleteMatchingPredicate(ctx, options, MobileAppIntentAndStateOperationPredicate{})
}

// ListMobileAppIntentAndStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppIntentAndStateClient) ListMobileAppIntentAndStatesCompleteMatchingPredicate(ctx context.Context, options ListMobileAppIntentAndStatesOperationOptions, predicate MobileAppIntentAndStateOperationPredicate) (result ListMobileAppIntentAndStatesCompleteResult, err error) {
	items := make([]beta.MobileAppIntentAndState, 0)

	resp, err := c.ListMobileAppIntentAndStates(ctx, options)
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

	result = ListMobileAppIntentAndStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
