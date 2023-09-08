package memailfoldermessagerule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeMailFolderByIdMessageRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MessageRuleCollectionResponse
}

type ListMeMailFolderByIdMessageRulesCompleteResult struct {
	Items []models.MessageRuleCollectionResponse
}

// ListMeMailFolderByIdMessageRules ...
func (c MeMailFolderMessageRuleClient) ListMeMailFolderByIdMessageRules(ctx context.Context, id MeMailFolderId) (result ListMeMailFolderByIdMessageRulesOperationResponse, err error) {
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

// ListMeMailFolderByIdMessageRulesComplete retrieves all the results into a single object
func (c MeMailFolderMessageRuleClient) ListMeMailFolderByIdMessageRulesComplete(ctx context.Context, id MeMailFolderId) (ListMeMailFolderByIdMessageRulesCompleteResult, error) {
	return c.ListMeMailFolderByIdMessageRulesCompleteMatchingPredicate(ctx, id, models.MessageRuleCollectionResponseOperationPredicate{})
}

// ListMeMailFolderByIdMessageRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderMessageRuleClient) ListMeMailFolderByIdMessageRulesCompleteMatchingPredicate(ctx context.Context, id MeMailFolderId, predicate models.MessageRuleCollectionResponseOperationPredicate) (result ListMeMailFolderByIdMessageRulesCompleteResult, err error) {
	items := make([]models.MessageRuleCollectionResponse, 0)

	resp, err := c.ListMeMailFolderByIdMessageRules(ctx, id)
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

	result = ListMeMailFolderByIdMessageRulesCompleteResult{
		Items: items,
	}
	return
}
