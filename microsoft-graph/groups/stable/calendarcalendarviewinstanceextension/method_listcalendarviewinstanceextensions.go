package calendarcalendarviewinstanceextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCalendarViewInstanceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Extension
}

type ListCalendarViewInstanceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Extension
}

type ListCalendarViewInstanceExtensionsOperationOptions struct {
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

func DefaultListCalendarViewInstanceExtensionsOperationOptions() ListCalendarViewInstanceExtensionsOperationOptions {
	return ListCalendarViewInstanceExtensionsOperationOptions{}
}

func (o ListCalendarViewInstanceExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarViewInstanceExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarViewInstanceExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarViewInstanceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewInstanceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewInstanceExtensions - Get extensions from groups. The collection of open extensions defined for the
// event. Nullable.
func (c CalendarCalendarViewInstanceExtensionClient) ListCalendarViewInstanceExtensions(ctx context.Context, id stable.GroupIdCalendarCalendarViewIdInstanceId, options ListCalendarViewInstanceExtensionsOperationOptions) (result ListCalendarViewInstanceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarViewInstanceExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.Extension, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalExtensionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.Extension (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListCalendarViewInstanceExtensionsComplete retrieves all the results into a single object
func (c CalendarCalendarViewInstanceExtensionClient) ListCalendarViewInstanceExtensionsComplete(ctx context.Context, id stable.GroupIdCalendarCalendarViewIdInstanceId, options ListCalendarViewInstanceExtensionsOperationOptions) (ListCalendarViewInstanceExtensionsCompleteResult, error) {
	return c.ListCalendarViewInstanceExtensionsCompleteMatchingPredicate(ctx, id, options, ExtensionOperationPredicate{})
}

// ListCalendarViewInstanceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewInstanceExtensionClient) ListCalendarViewInstanceExtensionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdCalendarCalendarViewIdInstanceId, options ListCalendarViewInstanceExtensionsOperationOptions, predicate ExtensionOperationPredicate) (result ListCalendarViewInstanceExtensionsCompleteResult, err error) {
	items := make([]stable.Extension, 0)

	resp, err := c.ListCalendarViewInstanceExtensions(ctx, id, options)
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

	result = ListCalendarViewInstanceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
