package memailfolderchildfoldermessagemention

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MentionCollectionResponse
}

type ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteResult struct {
	Items []models.MentionCollectionResponse
}

// ListMeMailFolderByIdChildFolderByIdMessageByIdMentions ...
func (c MeMailFolderChildFolderMessageMentionClient) ListMeMailFolderByIdChildFolderByIdMessageByIdMentions(ctx context.Context, id MeMailFolderChildFolderMessageId) (result ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.MentionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsComplete retrieves all the results into a single object
func (c MeMailFolderChildFolderMessageMentionClient) ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsComplete(ctx context.Context, id MeMailFolderChildFolderMessageId) (ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteResult, error) {
	return c.ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteMatchingPredicate(ctx, id, models.MentionCollectionResponseOperationPredicate{})
}

// ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderChildFolderMessageMentionClient) ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteMatchingPredicate(ctx context.Context, id MeMailFolderChildFolderMessageId, predicate models.MentionCollectionResponseOperationPredicate) (result ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteResult, err error) {
	items := make([]models.MentionCollectionResponse, 0)

	resp, err := c.ListMeMailFolderByIdChildFolderByIdMessageByIdMentions(ctx, id)
	if err != nil {
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

	result = ListMeMailFolderByIdChildFolderByIdMessageByIdMentionsCompleteResult{
		Items: items,
	}
	return
}
