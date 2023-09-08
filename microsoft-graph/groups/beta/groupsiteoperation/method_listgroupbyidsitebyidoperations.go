package groupsiteoperation

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

type ListGroupByIdSiteByIdOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.RichLongRunningOperationCollectionResponse
}

type ListGroupByIdSiteByIdOperationsCompleteResult struct {
	Items []models.RichLongRunningOperationCollectionResponse
}

// ListGroupByIdSiteByIdOperations ...
func (c GroupSiteOperationClient) ListGroupByIdSiteByIdOperations(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdOperationsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdOperationsComplete retrieves all the results into a single object
func (c GroupSiteOperationClient) ListGroupByIdSiteByIdOperationsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdOperationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOperationsCompleteMatchingPredicate(ctx, id, models.RichLongRunningOperationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOperationClient) ListGroupByIdSiteByIdOperationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.RichLongRunningOperationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOperationsCompleteResult, err error) {
	items := make([]models.RichLongRunningOperationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOperations(ctx, id)
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

	result = ListGroupByIdSiteByIdOperationsCompleteResult{
		Items: items,
	}
	return
}
