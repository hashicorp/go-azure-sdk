package groupsitecontenttypecolumn

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

type ListGroupByIdSiteByIdContentTypeByIdColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ColumnDefinitionCollectionResponse
}

type ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteResult struct {
	Items []models.ColumnDefinitionCollectionResponse
}

// ListGroupByIdSiteByIdContentTypeByIdColumns ...
func (c GroupSiteContentTypeColumnClient) ListGroupByIdSiteByIdContentTypeByIdColumns(ctx context.Context, id GroupSiteContentTypeId) (result ListGroupByIdSiteByIdContentTypeByIdColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/columns", id.ID()),
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

// ListGroupByIdSiteByIdContentTypeByIdColumnsComplete retrieves all the results into a single object
func (c GroupSiteContentTypeColumnClient) ListGroupByIdSiteByIdContentTypeByIdColumnsComplete(ctx context.Context, id GroupSiteContentTypeId) (ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteMatchingPredicate(ctx, id, models.ColumnDefinitionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteContentTypeColumnClient) ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteMatchingPredicate(ctx context.Context, id GroupSiteContentTypeId, predicate models.ColumnDefinitionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteResult, err error) {
	items := make([]models.ColumnDefinitionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdContentTypeByIdColumns(ctx, id)
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

	result = ListGroupByIdSiteByIdContentTypeByIdColumnsCompleteResult{
		Items: items,
	}
	return
}
