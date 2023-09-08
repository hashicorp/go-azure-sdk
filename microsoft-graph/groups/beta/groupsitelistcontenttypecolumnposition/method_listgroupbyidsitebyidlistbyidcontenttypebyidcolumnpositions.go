package groupsitelistcontenttypecolumnposition

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

type ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ColumnDefinitionCollectionResponse
}

type ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteResult struct {
	Items []models.ColumnDefinitionCollectionResponse
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositions ...
func (c GroupSiteListContentTypeColumnPositionClient) ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositions(ctx context.Context, id GroupSiteListContentTypeId) (result ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/columnPositions", id.ID()),
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
		Values *[]models.ColumnDefinitionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsComplete retrieves all the results into a single object
func (c GroupSiteListContentTypeColumnPositionClient) ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsComplete(ctx context.Context, id GroupSiteListContentTypeId) (ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteMatchingPredicate(ctx, id, models.ColumnDefinitionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListContentTypeColumnPositionClient) ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteMatchingPredicate(ctx context.Context, id GroupSiteListContentTypeId, predicate models.ColumnDefinitionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteResult, err error) {
	items := make([]models.ColumnDefinitionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositions(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdContentTypeByIdColumnPositionsCompleteResult{
		Items: items,
	}
	return
}
