package usermailfoldermessagemention

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

type ListUserByIdMailFolderByIdMessageByIdMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MentionCollectionResponse
}

type ListUserByIdMailFolderByIdMessageByIdMentionsCompleteResult struct {
	Items []models.MentionCollectionResponse
}

// ListUserByIdMailFolderByIdMessageByIdMentions ...
func (c UserMailFolderMessageMentionClient) ListUserByIdMailFolderByIdMessageByIdMentions(ctx context.Context, id UserMailFolderMessageId) (result ListUserByIdMailFolderByIdMessageByIdMentionsOperationResponse, err error) {
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

// ListUserByIdMailFolderByIdMessageByIdMentionsComplete retrieves all the results into a single object
func (c UserMailFolderMessageMentionClient) ListUserByIdMailFolderByIdMessageByIdMentionsComplete(ctx context.Context, id UserMailFolderMessageId) (ListUserByIdMailFolderByIdMessageByIdMentionsCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdMessageByIdMentionsCompleteMatchingPredicate(ctx, id, models.MentionCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdMessageByIdMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderMessageMentionClient) ListUserByIdMailFolderByIdMessageByIdMentionsCompleteMatchingPredicate(ctx context.Context, id UserMailFolderMessageId, predicate models.MentionCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdMessageByIdMentionsCompleteResult, err error) {
	items := make([]models.MentionCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdMessageByIdMentions(ctx, id)
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

	result = ListUserByIdMailFolderByIdMessageByIdMentionsCompleteResult{
		Items: items,
	}
	return
}
