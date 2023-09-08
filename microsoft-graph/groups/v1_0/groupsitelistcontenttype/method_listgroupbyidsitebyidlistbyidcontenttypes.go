package groupsitelistcontenttype

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

type ListGroupByIdSiteByIdListByIdContentTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContentTypeCollectionResponse
}

type ListGroupByIdSiteByIdListByIdContentTypesCompleteResult struct {
	Items []models.ContentTypeCollectionResponse
}

// ListGroupByIdSiteByIdListByIdContentTypes ...
func (c GroupSiteListContentTypeClient) ListGroupByIdSiteByIdListByIdContentTypes(ctx context.Context, id GroupSiteListId) (result ListGroupByIdSiteByIdListByIdContentTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/contentTypes", id.ID()),
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

// ListGroupByIdSiteByIdListByIdContentTypesComplete retrieves all the results into a single object
func (c GroupSiteListContentTypeClient) ListGroupByIdSiteByIdListByIdContentTypesComplete(ctx context.Context, id GroupSiteListId) (ListGroupByIdSiteByIdListByIdContentTypesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdContentTypesCompleteMatchingPredicate(ctx, id, models.ContentTypeCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdContentTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListContentTypeClient) ListGroupByIdSiteByIdListByIdContentTypesCompleteMatchingPredicate(ctx context.Context, id GroupSiteListId, predicate models.ContentTypeCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdContentTypesCompleteResult, err error) {
	items := make([]models.ContentTypeCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdContentTypes(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdContentTypesCompleteResult{
		Items: items,
	}
	return
}
