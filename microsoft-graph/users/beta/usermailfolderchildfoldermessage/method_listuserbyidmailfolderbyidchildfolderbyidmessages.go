package usermailfolderchildfoldermessage

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

type ListUserByIdMailFolderByIdChildFolderByIdMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MessageCollectionResponse
}

type ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteResult struct {
	Items []models.MessageCollectionResponse
}

// ListUserByIdMailFolderByIdChildFolderByIdMessages ...
func (c UserMailFolderChildFolderMessageClient) ListUserByIdMailFolderByIdChildFolderByIdMessages(ctx context.Context, id UserMailFolderChildFolderId) (result ListUserByIdMailFolderByIdChildFolderByIdMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/messages", id.ID()),
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
		Values *[]models.MessageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdMailFolderByIdChildFolderByIdMessagesComplete retrieves all the results into a single object
func (c UserMailFolderChildFolderMessageClient) ListUserByIdMailFolderByIdChildFolderByIdMessagesComplete(ctx context.Context, id UserMailFolderChildFolderId) (ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteMatchingPredicate(ctx, id, models.MessageCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderChildFolderMessageClient) ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteMatchingPredicate(ctx context.Context, id UserMailFolderChildFolderId, predicate models.MessageCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteResult, err error) {
	items := make([]models.MessageCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdChildFolderByIdMessages(ctx, id)
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

	result = ListUserByIdMailFolderByIdChildFolderByIdMessagesCompleteResult{
		Items: items,
	}
	return
}
