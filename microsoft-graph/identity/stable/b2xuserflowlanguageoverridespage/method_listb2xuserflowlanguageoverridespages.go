package b2xuserflowlanguageoverridespage

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

type ListB2xUserFlowLanguageOverridesPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageOverridesPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageOverridesPagesOperationOptions struct {
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

func DefaultListB2xUserFlowLanguageOverridesPagesOperationOptions() ListB2xUserFlowLanguageOverridesPagesOperationOptions {
	return ListB2xUserFlowLanguageOverridesPagesOperationOptions{}
}

func (o ListB2xUserFlowLanguageOverridesPagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListB2xUserFlowLanguageOverridesPagesOperationOptions) ToOData() *odata.Query {
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

func (o ListB2xUserFlowLanguageOverridesPagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListB2xUserFlowLanguageOverridesPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowLanguageOverridesPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowLanguageOverridesPages - List overridesPages. Get the userFlowLanguagePage resources from the
// overridesPages navigation property. These pages are used to customize the values shown to the user during a user
// journey in a user flow.
func (c B2xUserFlowLanguageOverridesPageClient) ListB2xUserFlowLanguageOverridesPages(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options ListB2xUserFlowLanguageOverridesPagesOperationOptions) (result ListB2xUserFlowLanguageOverridesPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListB2xUserFlowLanguageOverridesPagesCustomPager{},
		Path:          fmt.Sprintf("%s/overridesPages", id.ID()),
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

// ListB2xUserFlowLanguageOverridesPagesComplete retrieves all the results into a single object
func (c B2xUserFlowLanguageOverridesPageClient) ListB2xUserFlowLanguageOverridesPagesComplete(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options ListB2xUserFlowLanguageOverridesPagesOperationOptions) (ListB2xUserFlowLanguageOverridesPagesCompleteResult, error) {
	return c.ListB2xUserFlowLanguageOverridesPagesCompleteMatchingPredicate(ctx, id, options, UserFlowLanguagePageOperationPredicate{})
}

// ListB2xUserFlowLanguageOverridesPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowLanguageOverridesPageClient) ListB2xUserFlowLanguageOverridesPagesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options ListB2xUserFlowLanguageOverridesPagesOperationOptions, predicate UserFlowLanguagePageOperationPredicate) (result ListB2xUserFlowLanguageOverridesPagesCompleteResult, err error) {
	items := make([]stable.UserFlowLanguagePage, 0)

	resp, err := c.ListB2xUserFlowLanguageOverridesPages(ctx, id, options)
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

	result = ListB2xUserFlowLanguageOverridesPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
