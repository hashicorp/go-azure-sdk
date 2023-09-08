package groupsitelistoperation

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

type ListGroupByIdSiteByIdListByIdOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.RichLongRunningOperationCollectionResponse
}

type ListGroupByIdSiteByIdListByIdOperationsCompleteResult struct {
	Items []models.RichLongRunningOperationCollectionResponse
}

// ListGroupByIdSiteByIdListByIdOperations ...
func (c GroupSiteListOperationClient) ListGroupByIdSiteByIdListByIdOperations(ctx context.Context, id GroupSiteListId) (result ListGroupByIdSiteByIdListByIdOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]models.RichLongRunningOperationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdOperationsComplete retrieves all the results into a single object
func (c GroupSiteListOperationClient) ListGroupByIdSiteByIdListByIdOperationsComplete(ctx context.Context, id GroupSiteListId) (ListGroupByIdSiteByIdListByIdOperationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdOperationsCompleteMatchingPredicate(ctx, id, models.RichLongRunningOperationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListOperationClient) ListGroupByIdSiteByIdListByIdOperationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteListId, predicate models.RichLongRunningOperationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdOperationsCompleteResult, err error) {
	items := make([]models.RichLongRunningOperationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdOperations(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdOperationsCompleteResult{
		Items: items,
	}
	return
}
