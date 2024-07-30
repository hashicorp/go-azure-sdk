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

type ListOnlineMeetingTranscriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingTranscriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingTranscripts ...
func (c OnlineMeetingTranscriptClient) ListOnlineMeetingTranscripts(ctx context.Context, id MeOnlineMeetingId) (result ListOnlineMeetingTranscriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOnlineMeetingTranscriptsCustomPager{},
		Path:       fmt.Sprintf("%s/transcripts", id.ID()),
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
func (c OnlineMeetingTranscriptClient) ListOnlineMeetingTranscriptsComplete(ctx context.Context, id MeOnlineMeetingId) (ListOnlineMeetingTranscriptsCompleteResult, error) {
	return c.ListOnlineMeetingTranscriptsCompleteMatchingPredicate(ctx, id, CallTranscriptOperationPredicate{})
}

// ListOnlineMeetingTranscriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingTranscriptClient) ListOnlineMeetingTranscriptsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate CallTranscriptOperationPredicate) (result ListOnlineMeetingTranscriptsCompleteResult, err error) {
	items := make([]stable.CallTranscript, 0)

	resp, err := c.ListOnlineMeetingTranscripts(ctx, id)
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
