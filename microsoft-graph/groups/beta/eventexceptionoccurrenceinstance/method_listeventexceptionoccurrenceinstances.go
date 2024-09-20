package eventexceptionoccurrenceinstance

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

type ListEventExceptionOccurrenceInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListEventExceptionOccurrenceInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListEventExceptionOccurrenceInstancesOperationOptions struct {
	Count         *bool
	EndDateTime   *string
	Expand        *odata.Expand
	Filter        *string
	Metadata      *odata.Metadata
	OrderBy       *odata.OrderBy
	Search        *string
	Select        *[]string
	Skip          *int64
	StartDateTime *string
	Top           *int64
}

func DefaultListEventExceptionOccurrenceInstancesOperationOptions() ListEventExceptionOccurrenceInstancesOperationOptions {
	return ListEventExceptionOccurrenceInstancesOperationOptions{}
}

func (o ListEventExceptionOccurrenceInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEventExceptionOccurrenceInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListEventExceptionOccurrenceInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDateTime != nil {
		out.Append("endDateTime", fmt.Sprintf("%v", *o.EndDateTime))
	}
	if o.StartDateTime != nil {
		out.Append("startDateTime", fmt.Sprintf("%v", *o.StartDateTime))
	}
	return &out
}

type ListEventExceptionOccurrenceInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEventExceptionOccurrenceInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEventExceptionOccurrenceInstances - Get instances from groups. The occurrences of a recurring series, if the
// event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions
// that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation
// property. Read-only. Nullable.
func (c EventExceptionOccurrenceInstanceClient) ListEventExceptionOccurrenceInstances(ctx context.Context, id beta.GroupIdEventIdExceptionOccurrenceId, options ListEventExceptionOccurrenceInstancesOperationOptions) (result ListEventExceptionOccurrenceInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEventExceptionOccurrenceInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]beta.Event `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEventExceptionOccurrenceInstancesComplete retrieves all the results into a single object
func (c EventExceptionOccurrenceInstanceClient) ListEventExceptionOccurrenceInstancesComplete(ctx context.Context, id beta.GroupIdEventIdExceptionOccurrenceId, options ListEventExceptionOccurrenceInstancesOperationOptions) (ListEventExceptionOccurrenceInstancesCompleteResult, error) {
	return c.ListEventExceptionOccurrenceInstancesCompleteMatchingPredicate(ctx, id, options, EventOperationPredicate{})
}

// ListEventExceptionOccurrenceInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventExceptionOccurrenceInstanceClient) ListEventExceptionOccurrenceInstancesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdEventIdExceptionOccurrenceId, options ListEventExceptionOccurrenceInstancesOperationOptions, predicate EventOperationPredicate) (result ListEventExceptionOccurrenceInstancesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListEventExceptionOccurrenceInstances(ctx, id, options)
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

	result = ListEventExceptionOccurrenceInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
