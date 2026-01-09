package b2xuserflowlanguagedefaultpage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListB2xUserFlowLanguageDefaultPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageDefaultPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageDefaultPagesOperationOptions struct {
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

func DefaultListB2xUserFlowLanguageDefaultPagesOperationOptions() ListB2xUserFlowLanguageDefaultPagesOperationOptions {
	return ListB2xUserFlowLanguageDefaultPagesOperationOptions{}
}

func (o ListB2xUserFlowLanguageDefaultPagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListB2xUserFlowLanguageDefaultPagesOperationOptions) ToOData() *odata.Query {
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

func (o ListB2xUserFlowLanguageDefaultPagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListB2xUserFlowLanguageDefaultPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowLanguageDefaultPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowLanguageDefaultPages - Get userFlowLanguagePage. Read the values in a userFlowLanguagePage object for
// a language in a user flow. These values are shown to a user during a user journey defined by a user flow.
func (c B2xUserFlowLanguageDefaultPageClient) ListB2xUserFlowLanguageDefaultPages(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options ListB2xUserFlowLanguageDefaultPagesOperationOptions) (result ListB2xUserFlowLanguageDefaultPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListB2xUserFlowLanguageDefaultPagesCustomPager{},
		Path:          fmt.Sprintf("%s/defaultPages", id.ID()),
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
		Values *[]stable.UserFlowLanguagePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowLanguageDefaultPagesComplete retrieves all the results into a single object
func (c B2xUserFlowLanguageDefaultPageClient) ListB2xUserFlowLanguageDefaultPagesComplete(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options ListB2xUserFlowLanguageDefaultPagesOperationOptions) (ListB2xUserFlowLanguageDefaultPagesCompleteResult, error) {
	return c.ListB2xUserFlowLanguageDefaultPagesCompleteMatchingPredicate(ctx, id, options, UserFlowLanguagePageOperationPredicate{})
}

// ListB2xUserFlowLanguageDefaultPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowLanguageDefaultPageClient) ListB2xUserFlowLanguageDefaultPagesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options ListB2xUserFlowLanguageDefaultPagesOperationOptions, predicate UserFlowLanguagePageOperationPredicate) (result ListB2xUserFlowLanguageDefaultPagesCompleteResult, err error) {
	items := make([]stable.UserFlowLanguagePage, 0)

	resp, err := c.ListB2xUserFlowLanguageDefaultPages(ctx, id, options)
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

	result = ListB2xUserFlowLanguageDefaultPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
