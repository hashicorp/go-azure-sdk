package profilewebsite

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

type ListProfileWebsitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonWebsite
}

type ListProfileWebsitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonWebsite
}

type ListProfileWebsitesOperationOptions struct {
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

func DefaultListProfileWebsitesOperationOptions() ListProfileWebsitesOperationOptions {
	return ListProfileWebsitesOperationOptions{}
}

func (o ListProfileWebsitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileWebsitesOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileWebsitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileWebsitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileWebsitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileWebsites - Get websites from users. Represents detailed information about websites associated with a user
// in various services.
func (c ProfileWebsiteClient) ListProfileWebsites(ctx context.Context, id beta.UserId, options ListProfileWebsitesOperationOptions) (result ListProfileWebsitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileWebsitesCustomPager{},
		Path:          fmt.Sprintf("%s/profile/websites", id.ID()),
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
		Values *[]beta.PersonWebsite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileWebsitesComplete retrieves all the results into a single object
func (c ProfileWebsiteClient) ListProfileWebsitesComplete(ctx context.Context, id beta.UserId, options ListProfileWebsitesOperationOptions) (ListProfileWebsitesCompleteResult, error) {
	return c.ListProfileWebsitesCompleteMatchingPredicate(ctx, id, options, PersonWebsiteOperationPredicate{})
}

// ListProfileWebsitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileWebsiteClient) ListProfileWebsitesCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfileWebsitesOperationOptions, predicate PersonWebsiteOperationPredicate) (result ListProfileWebsitesCompleteResult, err error) {
	items := make([]beta.PersonWebsite, 0)

	resp, err := c.ListProfileWebsites(ctx, id, options)
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

	result = ListProfileWebsitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
