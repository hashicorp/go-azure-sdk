package threadpostinreplytomention

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

type ListThreadPostInReplyToMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Mention
}

type ListThreadPostInReplyToMentionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Mention
}

type ListThreadPostInReplyToMentionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListThreadPostInReplyToMentionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListThreadPostInReplyToMentions ...
func (c ThreadPostInReplyToMentionClient) ListThreadPostInReplyToMentions(ctx context.Context, id GroupIdThreadIdPostId) (result ListThreadPostInReplyToMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListThreadPostInReplyToMentionsCustomPager{},
		Path:       fmt.Sprintf("%s/inReplyTo/mentions", id.ID()),
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
		Values *[]beta.Mention `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListThreadPostInReplyToMentionsComplete retrieves all the results into a single object
func (c ThreadPostInReplyToMentionClient) ListThreadPostInReplyToMentionsComplete(ctx context.Context, id GroupIdThreadIdPostId) (ListThreadPostInReplyToMentionsCompleteResult, error) {
	return c.ListThreadPostInReplyToMentionsCompleteMatchingPredicate(ctx, id, MentionOperationPredicate{})
}

// ListThreadPostInReplyToMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadPostInReplyToMentionClient) ListThreadPostInReplyToMentionsCompleteMatchingPredicate(ctx context.Context, id GroupIdThreadIdPostId, predicate MentionOperationPredicate) (result ListThreadPostInReplyToMentionsCompleteResult, err error) {
	items := make([]beta.Mention, 0)

	resp, err := c.ListThreadPostInReplyToMentions(ctx, id)
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

	result = ListThreadPostInReplyToMentionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
