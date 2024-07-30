package mailfolderchildfoldermessagemention

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

type ListMailFolderChildFolderMessageMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Mention
}

type ListMailFolderChildFolderMessageMentionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Mention
}

type ListMailFolderChildFolderMessageMentionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFolderMessageMentionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolderMessageMentions ...
func (c MailFolderChildFolderMessageMentionClient) ListMailFolderChildFolderMessageMentions(ctx context.Context, id UserIdMailFolderIdChildFolderIdMessageId) (result ListMailFolderChildFolderMessageMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderChildFolderMessageMentionsCustomPager{},
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

// ListMailFolderChildFolderMessageMentionsComplete retrieves all the results into a single object
func (c MailFolderChildFolderMessageMentionClient) ListMailFolderChildFolderMessageMentionsComplete(ctx context.Context, id UserIdMailFolderIdChildFolderIdMessageId) (ListMailFolderChildFolderMessageMentionsCompleteResult, error) {
	return c.ListMailFolderChildFolderMessageMentionsCompleteMatchingPredicate(ctx, id, MentionOperationPredicate{})
}

// ListMailFolderChildFolderMessageMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderMessageMentionClient) ListMailFolderChildFolderMessageMentionsCompleteMatchingPredicate(ctx context.Context, id UserIdMailFolderIdChildFolderIdMessageId, predicate MentionOperationPredicate) (result ListMailFolderChildFolderMessageMentionsCompleteResult, err error) {
	items := make([]beta.Mention, 0)

	resp, err := c.ListMailFolderChildFolderMessageMentions(ctx, id)
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

	result = ListMailFolderChildFolderMessageMentionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
