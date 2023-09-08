package groupsitelistcontenttypebasetype

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

type ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContentTypeCollectionResponse
}

type ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteResult struct {
	Items []models.ContentTypeCollectionResponse
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypes ...
func (c GroupSiteListContentTypeBaseTypeClient) ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypes(ctx context.Context, id GroupSiteListContentTypeId) (result ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesOperationResponse, err error) {
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

// ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesComplete retrieves all the results into a single object
func (c GroupSiteListContentTypeBaseTypeClient) ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesComplete(ctx context.Context, id GroupSiteListContentTypeId) (ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteMatchingPredicate(ctx, id, models.ContentTypeCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListContentTypeBaseTypeClient) ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteMatchingPredicate(ctx context.Context, id GroupSiteListContentTypeId, predicate models.ContentTypeCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteResult, err error) {
	items := make([]models.ContentTypeCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypes(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdContentTypeByIdBaseTypesCompleteResult{
		Items: items,
	}
	return
}
