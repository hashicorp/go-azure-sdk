package onlinemeetingrecording

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOnlineMeetingRecordingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CallRecording
}

type ListOnlineMeetingRecordingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CallRecording
}

type ListOnlineMeetingRecordingsOperationOptions struct {
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

func DefaultListOnlineMeetingRecordingsOperationOptions() ListOnlineMeetingRecordingsOperationOptions {
	return ListOnlineMeetingRecordingsOperationOptions{}
}

func (o ListOnlineMeetingRecordingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnlineMeetingRecordingsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnlineMeetingRecordingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnlineMeetingRecordingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingRecordingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingRecordings - Get callRecording. Get a callRecording object associated with a scheduled
// onlineMeeting. This API doesn't support getting call recordings from channel meetings. For a recording, this API
// returns the metadata of the single recording associated with the online meeting. For the content of a recording, this
// API returns the stream of bytes associated with the recording.
func (c OnlineMeetingRecordingClient) ListOnlineMeetingRecordings(ctx context.Context, id stable.UserIdOnlineMeetingId, options ListOnlineMeetingRecordingsOperationOptions) (result ListOnlineMeetingRecordingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnlineMeetingRecordingsCustomPager{},
		Path:          fmt.Sprintf("%s/recordings", id.ID()),
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
		Values *[]stable.CallRecording `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingRecordingsComplete retrieves all the results into a single object
func (c OnlineMeetingRecordingClient) ListOnlineMeetingRecordingsComplete(ctx context.Context, id stable.UserIdOnlineMeetingId, options ListOnlineMeetingRecordingsOperationOptions) (ListOnlineMeetingRecordingsCompleteResult, error) {
	return c.ListOnlineMeetingRecordingsCompleteMatchingPredicate(ctx, id, options, CallRecordingOperationPredicate{})
}

// ListOnlineMeetingRecordingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingRecordingClient) ListOnlineMeetingRecordingsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdOnlineMeetingId, options ListOnlineMeetingRecordingsOperationOptions, predicate CallRecordingOperationPredicate) (result ListOnlineMeetingRecordingsCompleteResult, err error) {
	items := make([]stable.CallRecording, 0)

	resp, err := c.ListOnlineMeetingRecordings(ctx, id, options)
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

	result = ListOnlineMeetingRecordingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
