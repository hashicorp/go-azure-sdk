package calendareventinstanceextension

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

type ListCalendarEventInstanceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Extension
}

type ListCalendarEventInstanceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Extension
}

type ListCalendarEventInstanceExtensionsOperationOptions struct {
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

func DefaultListCalendarEventInstanceExtensionsOperationOptions() ListCalendarEventInstanceExtensionsOperationOptions {
	return ListCalendarEventInstanceExtensionsOperationOptions{}
}

func (o ListCalendarEventInstanceExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarEventInstanceExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarEventInstanceExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarEventInstanceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarEventInstanceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarEventInstanceExtensions - Get extensions from groups. The collection of open extensions defined for the
// event. Nullable.
func (c CalendarEventInstanceExtensionClient) ListCalendarEventInstanceExtensions(ctx context.Context, id stable.GroupIdCalendarEventIdInstanceId, options ListCalendarEventInstanceExtensionsOperationOptions) (result ListCalendarEventInstanceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarEventInstanceExtensionsCustomPager{},
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

// ListCalendarEventInstanceExtensionsComplete retrieves all the results into a single object
func (c CalendarEventInstanceExtensionClient) ListCalendarEventInstanceExtensionsComplete(ctx context.Context, id stable.GroupIdCalendarEventIdInstanceId, options ListCalendarEventInstanceExtensionsOperationOptions) (ListCalendarEventInstanceExtensionsCompleteResult, error) {
	return c.ListCalendarEventInstanceExtensionsCompleteMatchingPredicate(ctx, id, options, ExtensionOperationPredicate{})
}

// ListCalendarEventInstanceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarEventInstanceExtensionClient) ListCalendarEventInstanceExtensionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdCalendarEventIdInstanceId, options ListCalendarEventInstanceExtensionsOperationOptions, predicate ExtensionOperationPredicate) (result ListCalendarEventInstanceExtensionsCompleteResult, err error) {
	items := make([]stable.Extension, 0)

	resp, err := c.ListCalendarEventInstanceExtensions(ctx, id, options)
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

	result = ListCalendarEventInstanceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
