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

type ListOnlineMeetingRecordingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingRecordingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingRecordings ...
func (c OnlineMeetingRecordingClient) ListOnlineMeetingRecordings(ctx context.Context, id MeOnlineMeetingId) (result ListOnlineMeetingRecordingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOnlineMeetingRecordingsCustomPager{},
		Path:       fmt.Sprintf("%s/recordings", id.ID()),
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
func (c OnlineMeetingRecordingClient) ListOnlineMeetingRecordingsComplete(ctx context.Context, id MeOnlineMeetingId) (ListOnlineMeetingRecordingsCompleteResult, error) {
	return c.ListOnlineMeetingRecordingsCompleteMatchingPredicate(ctx, id, CallRecordingOperationPredicate{})
}

// ListOnlineMeetingRecordingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingRecordingClient) ListOnlineMeetingRecordingsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate CallRecordingOperationPredicate) (result ListOnlineMeetingRecordingsCompleteResult, err error) {
	items := make([]stable.CallRecording, 0)

	resp, err := c.ListOnlineMeetingRecordings(ctx, id)
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
