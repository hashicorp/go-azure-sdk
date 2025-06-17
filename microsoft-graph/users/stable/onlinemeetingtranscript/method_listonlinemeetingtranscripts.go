package onlinemeetingtranscript

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

type ListOnlineMeetingTranscriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CallTranscript
}

type ListOnlineMeetingTranscriptsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CallTranscript
}

type ListOnlineMeetingTranscriptsOperationOptions struct {
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

func DefaultListOnlineMeetingTranscriptsOperationOptions() ListOnlineMeetingTranscriptsOperationOptions {
	return ListOnlineMeetingTranscriptsOperationOptions{}
}

func (o ListOnlineMeetingTranscriptsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnlineMeetingTranscriptsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnlineMeetingTranscriptsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnlineMeetingTranscriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingTranscriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingTranscripts - List transcripts. Retrieve the list of callTranscript objects associated with a
// scheduled onlineMeeting. This API supports the retrieval of call recordings from private chat meetings and channel
// meetings. However, private channel meetings are not supported at this time.
func (c OnlineMeetingTranscriptClient) ListOnlineMeetingTranscripts(ctx context.Context, id stable.UserIdOnlineMeetingId, options ListOnlineMeetingTranscriptsOperationOptions) (result ListOnlineMeetingTranscriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnlineMeetingTranscriptsCustomPager{},
		Path:          fmt.Sprintf("%s/transcripts", id.ID()),
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
		Values *[]stable.CallTranscript `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingTranscriptsComplete retrieves all the results into a single object
func (c OnlineMeetingTranscriptClient) ListOnlineMeetingTranscriptsComplete(ctx context.Context, id stable.UserIdOnlineMeetingId, options ListOnlineMeetingTranscriptsOperationOptions) (ListOnlineMeetingTranscriptsCompleteResult, error) {
	return c.ListOnlineMeetingTranscriptsCompleteMatchingPredicate(ctx, id, options, CallTranscriptOperationPredicate{})
}

// ListOnlineMeetingTranscriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingTranscriptClient) ListOnlineMeetingTranscriptsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdOnlineMeetingId, options ListOnlineMeetingTranscriptsOperationOptions, predicate CallTranscriptOperationPredicate) (result ListOnlineMeetingTranscriptsCompleteResult, err error) {
	items := make([]stable.CallTranscript, 0)

	resp, err := c.ListOnlineMeetingTranscripts(ctx, id, options)
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

	result = ListOnlineMeetingTranscriptsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
