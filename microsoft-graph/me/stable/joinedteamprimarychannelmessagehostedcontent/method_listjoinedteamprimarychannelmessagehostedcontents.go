package joinedteamprimarychannelmessagehostedcontent

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

type ListJoinedTeamPrimaryChannelMessageHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessageHostedContent
}

type ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessageHostedContent
}

type ListJoinedTeamPrimaryChannelMessageHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelMessageHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelMessageHostedContents ...
func (c JoinedTeamPrimaryChannelMessageHostedContentClient) ListJoinedTeamPrimaryChannelMessageHostedContents(ctx context.Context, id MeJoinedTeamIdPrimaryChannelMessageId) (result ListJoinedTeamPrimaryChannelMessageHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamPrimaryChannelMessageHostedContentsCustomPager{},
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

// ListJoinedTeamPrimaryChannelMessageHostedContentsComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelMessageHostedContentClient) ListJoinedTeamPrimaryChannelMessageHostedContentsComplete(ctx context.Context, id MeJoinedTeamIdPrimaryChannelMessageId) (ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate(ctx, id, ChatMessageHostedContentOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelMessageHostedContentClient) ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamIdPrimaryChannelMessageId, predicate ChatMessageHostedContentOperationPredicate) (result ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteResult, err error) {
	items := make([]stable.ChatMessageHostedContent, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelMessageHostedContents(ctx, id)
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

	result = ListJoinedTeamPrimaryChannelMessageHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
