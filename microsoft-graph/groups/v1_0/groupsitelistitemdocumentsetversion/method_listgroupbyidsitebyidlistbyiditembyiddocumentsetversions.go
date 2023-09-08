package groupsitelistitemdocumentsetversion

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

type ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DocumentSetVersionCollectionResponse
}

type ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteResult struct {
	Items []models.DocumentSetVersionCollectionResponse
}

// ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersions ...
func (c GroupSiteListItemDocumentSetVersionClient) ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersions(ctx context.Context, id GroupSiteListItemId) (result ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/documentSetVersions", id.ID()),
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
		Values *[]models.DocumentSetVersionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsComplete retrieves all the results into a single object
func (c GroupSiteListItemDocumentSetVersionClient) ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsComplete(ctx context.Context, id GroupSiteListItemId) (ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteMatchingPredicate(ctx, id, models.DocumentSetVersionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListItemDocumentSetVersionClient) ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteMatchingPredicate(ctx context.Context, id GroupSiteListItemId, predicate models.DocumentSetVersionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteResult, err error) {
	items := make([]models.DocumentSetVersionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersions(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdItemByIdDocumentSetVersionsCompleteResult{
		Items: items,
	}
	return
}
