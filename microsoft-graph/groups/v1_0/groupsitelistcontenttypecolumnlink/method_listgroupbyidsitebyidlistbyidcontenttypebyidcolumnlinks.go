package groupsitelistcontenttypecolumnlink

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

type ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ColumnLinkCollectionResponse
}

type ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteResult struct {
	Items []models.ColumnLinkCollectionResponse
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinks ...
func (c GroupSiteListContentTypeColumnLinkClient) ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinks(ctx context.Context, id GroupSiteListContentTypeId) (result ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/columnLinks", id.ID()),
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
		Values *[]models.ColumnLinkCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksComplete retrieves all the results into a single object
func (c GroupSiteListContentTypeColumnLinkClient) ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksComplete(ctx context.Context, id GroupSiteListContentTypeId) (ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteMatchingPredicate(ctx, id, models.ColumnLinkCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListContentTypeColumnLinkClient) ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteMatchingPredicate(ctx context.Context, id GroupSiteListContentTypeId, predicate models.ColumnLinkCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteResult, err error) {
	items := make([]models.ColumnLinkCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinks(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdContentTypeByIdColumnLinksCompleteResult{
		Items: items,
	}
	return
}
