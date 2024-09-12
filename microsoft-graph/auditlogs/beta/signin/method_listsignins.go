package signin

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

type ListSignInsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SignIn
}

type ListSignInsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SignIn
}

type ListSignInsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListSignInsOperationOptions() ListSignInsOperationOptions {
	return ListSignInsOperationOptions{}
}

func (o ListSignInsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSignInsOperationOptions) ToOData() *odata.Query {
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

func (o ListSignInsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSignInsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSignInsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSignIns - List signIns. Get a list of signIn objects. The list contains the user sign-ins for your Microsoft
// Entra tenant. Sign-ins where a username and password are passed as part of authorization token, and successful
// federated sign-ins are currently included in the sign-in logs. The maximum and default page size is 1,000 objects and
// by default, the most recent sign-ins are returned first. Only sign-in events that occurred within the Microsoft Entra
// ID default retention period are available.
func (c SignInClient) ListSignIns(ctx context.Context, options ListSignInsOperationOptions) (result ListSignInsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSignInsCustomPager{},
		Path:          "/auditLogs/signIns",
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
		Values *[]beta.SignIn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSignInsComplete retrieves all the results into a single object
func (c SignInClient) ListSignInsComplete(ctx context.Context, options ListSignInsOperationOptions) (ListSignInsCompleteResult, error) {
	return c.ListSignInsCompleteMatchingPredicate(ctx, options, SignInOperationPredicate{})
}

// ListSignInsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SignInClient) ListSignInsCompleteMatchingPredicate(ctx context.Context, options ListSignInsOperationOptions, predicate SignInOperationPredicate) (result ListSignInsCompleteResult, err error) {
	items := make([]beta.SignIn, 0)

	resp, err := c.ListSignIns(ctx, options)
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

	result = ListSignInsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
