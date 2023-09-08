package groupsitepage

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

type ListGroupByIdSiteByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.BaseSitePageCollectionResponse
}

type ListGroupByIdSiteByIdPagesCompleteResult struct {
	Items []models.BaseSitePageCollectionResponse
}

// ListGroupByIdSiteByIdPages ...
func (c GroupSitePageClient) ListGroupByIdSiteByIdPages(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/pages", id.ID()),
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
		Values *[]models.BaseSitePageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdPagesComplete retrieves all the results into a single object
func (c GroupSitePageClient) ListGroupByIdSiteByIdPagesComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdPagesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdPagesCompleteMatchingPredicate(ctx, id, models.BaseSitePageCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSitePageClient) ListGroupByIdSiteByIdPagesCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.BaseSitePageCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdPagesCompleteResult, err error) {
	items := make([]models.BaseSitePageCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdPages(ctx, id)
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

	result = ListGroupByIdSiteByIdPagesCompleteResult{
		Items: items,
	}
	return
}
