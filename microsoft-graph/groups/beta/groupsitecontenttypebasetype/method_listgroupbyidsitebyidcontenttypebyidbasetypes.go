package groupsitecontenttypebasetype

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

type ListGroupByIdSiteByIdContentTypeByIdBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContentTypeCollectionResponse
}

type ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteResult struct {
	Items []models.ContentTypeCollectionResponse
}

// ListGroupByIdSiteByIdContentTypeByIdBaseTypes ...
func (c GroupSiteContentTypeBaseTypeClient) ListGroupByIdSiteByIdContentTypeByIdBaseTypes(ctx context.Context, id GroupSiteContentTypeId) (result ListGroupByIdSiteByIdContentTypeByIdBaseTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/baseTypes", id.ID()),
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
		Values *[]models.ContentTypeCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdContentTypeByIdBaseTypesComplete retrieves all the results into a single object
func (c GroupSiteContentTypeBaseTypeClient) ListGroupByIdSiteByIdContentTypeByIdBaseTypesComplete(ctx context.Context, id GroupSiteContentTypeId) (ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteMatchingPredicate(ctx, id, models.ContentTypeCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteContentTypeBaseTypeClient) ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteMatchingPredicate(ctx context.Context, id GroupSiteContentTypeId, predicate models.ContentTypeCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteResult, err error) {
	items := make([]models.ContentTypeCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdContentTypeByIdBaseTypes(ctx, id)
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

	result = ListGroupByIdSiteByIdContentTypeByIdBaseTypesCompleteResult{
		Items: items,
	}
	return
}
