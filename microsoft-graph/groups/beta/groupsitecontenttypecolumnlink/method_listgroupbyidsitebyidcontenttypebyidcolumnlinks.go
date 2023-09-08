package groupsitecontenttypecolumnlink

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

type ListGroupByIdSiteByIdContentTypeByIdColumnLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ColumnLinkCollectionResponse
}

type ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteResult struct {
	Items []models.ColumnLinkCollectionResponse
}

// ListGroupByIdSiteByIdContentTypeByIdColumnLinks ...
func (c GroupSiteContentTypeColumnLinkClient) ListGroupByIdSiteByIdContentTypeByIdColumnLinks(ctx context.Context, id GroupSiteContentTypeId) (result ListGroupByIdSiteByIdContentTypeByIdColumnLinksOperationResponse, err error) {
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

// ListGroupByIdSiteByIdContentTypeByIdColumnLinksComplete retrieves all the results into a single object
func (c GroupSiteContentTypeColumnLinkClient) ListGroupByIdSiteByIdContentTypeByIdColumnLinksComplete(ctx context.Context, id GroupSiteContentTypeId) (ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteResult, error) {
	return c.ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteMatchingPredicate(ctx, id, models.ColumnLinkCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteContentTypeColumnLinkClient) ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteMatchingPredicate(ctx context.Context, id GroupSiteContentTypeId, predicate models.ColumnLinkCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteResult, err error) {
	items := make([]models.ColumnLinkCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdContentTypeByIdColumnLinks(ctx, id)
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

	result = ListGroupByIdSiteByIdContentTypeByIdColumnLinksCompleteResult{
		Items: items,
	}
	return
}
