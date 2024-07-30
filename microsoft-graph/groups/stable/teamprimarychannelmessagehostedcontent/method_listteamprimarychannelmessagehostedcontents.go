package teamprimarychannelmessagehostedcontent

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

type ListTeamPrimaryChannelMessageHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessageHostedContent
}

type ListTeamPrimaryChannelMessageHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessageHostedContent
}

type ListTeamPrimaryChannelMessageHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelMessageHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelMessageHostedContents ...
func (c TeamPrimaryChannelMessageHostedContentClient) ListTeamPrimaryChannelMessageHostedContents(ctx context.Context, id GroupIdTeamPrimaryChannelMessageId) (result ListTeamPrimaryChannelMessageHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamPrimaryChannelMessageHostedContentsCustomPager{},
		Path:       fmt.Sprintf("%s/hostedContents", id.ID()),
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
		Values *[]stable.ChatMessageHostedContent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPrimaryChannelMessageHostedContentsComplete retrieves all the results into a single object
func (c TeamPrimaryChannelMessageHostedContentClient) ListTeamPrimaryChannelMessageHostedContentsComplete(ctx context.Context, id GroupIdTeamPrimaryChannelMessageId) (ListTeamPrimaryChannelMessageHostedContentsCompleteResult, error) {
	return c.ListTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate(ctx, id, ChatMessageHostedContentOperationPredicate{})
}

// ListTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelMessageHostedContentClient) ListTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate(ctx context.Context, id GroupIdTeamPrimaryChannelMessageId, predicate ChatMessageHostedContentOperationPredicate) (result ListTeamPrimaryChannelMessageHostedContentsCompleteResult, err error) {
	items := make([]stable.ChatMessageHostedContent, 0)

	resp, err := c.ListTeamPrimaryChannelMessageHostedContents(ctx, id)
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

	result = ListTeamPrimaryChannelMessageHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
