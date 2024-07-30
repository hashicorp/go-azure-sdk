package mailfoldermessagemention

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

type ListMailFolderMessageMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Mention
}

type ListMailFolderMessageMentionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Mention
}

type ListMailFolderMessageMentionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderMessageMentionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderMessageMentions ...
func (c MailFolderMessageMentionClient) ListMailFolderMessageMentions(ctx context.Context, id UserIdMailFolderIdMessageId) (result ListMailFolderMessageMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderMessageMentionsCustomPager{},
		Path:       fmt.Sprintf("%s/mentions", id.ID()),
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

// ListMailFolderMessageMentionsComplete retrieves all the results into a single object
func (c MailFolderMessageMentionClient) ListMailFolderMessageMentionsComplete(ctx context.Context, id UserIdMailFolderIdMessageId) (ListMailFolderMessageMentionsCompleteResult, error) {
	return c.ListMailFolderMessageMentionsCompleteMatchingPredicate(ctx, id, MentionOperationPredicate{})
}

// ListMailFolderMessageMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderMessageMentionClient) ListMailFolderMessageMentionsCompleteMatchingPredicate(ctx context.Context, id UserIdMailFolderIdMessageId, predicate MentionOperationPredicate) (result ListMailFolderMessageMentionsCompleteResult, err error) {
	items := make([]beta.Mention, 0)

	resp, err := c.ListMailFolderMessageMentions(ctx, id)
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

	result = ListMailFolderMessageMentionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
