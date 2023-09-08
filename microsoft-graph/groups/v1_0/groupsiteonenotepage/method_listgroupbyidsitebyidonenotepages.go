package groupsiteonenotepage

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

type ListGroupByIdSiteByIdOnenotePagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListGroupByIdSiteByIdOnenotePagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListGroupByIdSiteByIdOnenotePages ...
func (c GroupSiteOnenotePageClient) ListGroupByIdSiteByIdOnenotePages(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdOnenotePagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/pages", id.ID()),
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
		Values *[]models.OnenotePageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdOnenotePagesComplete retrieves all the results into a single object
func (c GroupSiteOnenotePageClient) ListGroupByIdSiteByIdOnenotePagesComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdOnenotePagesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenotePagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenotePagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenotePageClient) ListGroupByIdSiteByIdOnenotePagesCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenotePagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenotePages(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenotePagesCompleteResult{
		Items: items,
	}
	return
}
