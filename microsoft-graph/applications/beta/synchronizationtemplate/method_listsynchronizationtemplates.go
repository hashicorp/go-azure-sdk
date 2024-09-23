package synchronizationtemplate

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

type ListSynchronizationTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SynchronizationTemplate
}

type ListSynchronizationTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SynchronizationTemplate
}

type ListSynchronizationTemplatesOperationOptions struct {
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

func DefaultListSynchronizationTemplatesOperationOptions() ListSynchronizationTemplatesOperationOptions {
	return ListSynchronizationTemplatesOperationOptions{}
}

func (o ListSynchronizationTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSynchronizationTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListSynchronizationTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSynchronizationTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationTemplates - Get templates from applications. Pre-configured synchronization settings for a
// particular application.
func (c SynchronizationTemplateClient) ListSynchronizationTemplates(ctx context.Context, id beta.ApplicationId, options ListSynchronizationTemplatesOperationOptions) (result ListSynchronizationTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSynchronizationTemplatesCustomPager{},
		Path:          fmt.Sprintf("%s/synchronization/templates", id.ID()),
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
		Values *[]beta.SynchronizationTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationTemplatesComplete retrieves all the results into a single object
func (c SynchronizationTemplateClient) ListSynchronizationTemplatesComplete(ctx context.Context, id beta.ApplicationId, options ListSynchronizationTemplatesOperationOptions) (ListSynchronizationTemplatesCompleteResult, error) {
	return c.ListSynchronizationTemplatesCompleteMatchingPredicate(ctx, id, options, SynchronizationTemplateOperationPredicate{})
}

// ListSynchronizationTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationTemplateClient) ListSynchronizationTemplatesCompleteMatchingPredicate(ctx context.Context, id beta.ApplicationId, options ListSynchronizationTemplatesOperationOptions, predicate SynchronizationTemplateOperationPredicate) (result ListSynchronizationTemplatesCompleteResult, err error) {
	items := make([]beta.SynchronizationTemplate, 0)

	resp, err := c.ListSynchronizationTemplates(ctx, id, options)
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

	result = ListSynchronizationTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
