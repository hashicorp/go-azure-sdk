package memailfolderchildfoldermessagerule

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

type ListMeMailFolderByIdChildFolderByIdMessageRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MessageRuleCollectionResponse
}

type ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteResult struct {
	Items []models.MessageRuleCollectionResponse
}

// ListMeMailFolderByIdChildFolderByIdMessageRules ...
func (c MeMailFolderChildFolderMessageRuleClient) ListMeMailFolderByIdChildFolderByIdMessageRules(ctx context.Context, id MeMailFolderChildFolderId) (result ListMeMailFolderByIdChildFolderByIdMessageRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/messageRules", id.ID()),
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
		Values *[]models.MessageRuleCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeMailFolderByIdChildFolderByIdMessageRulesComplete retrieves all the results into a single object
func (c MeMailFolderChildFolderMessageRuleClient) ListMeMailFolderByIdChildFolderByIdMessageRulesComplete(ctx context.Context, id MeMailFolderChildFolderId) (ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteResult, error) {
	return c.ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteMatchingPredicate(ctx, id, models.MessageRuleCollectionResponseOperationPredicate{})
}

// ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderChildFolderMessageRuleClient) ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteMatchingPredicate(ctx context.Context, id MeMailFolderChildFolderId, predicate models.MessageRuleCollectionResponseOperationPredicate) (result ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteResult, err error) {
	items := make([]models.MessageRuleCollectionResponse, 0)

	resp, err := c.ListMeMailFolderByIdChildFolderByIdMessageRules(ctx, id)
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

	result = ListMeMailFolderByIdChildFolderByIdMessageRulesCompleteResult{
		Items: items,
	}
	return
}
