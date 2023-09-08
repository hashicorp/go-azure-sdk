package groupsitecontenttype

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

type ListGroupByIdSiteByIdContentTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContentTypeCollectionResponse
}

type ListGroupByIdSiteByIdContentTypesCompleteResult struct {
	Items []models.ContentTypeCollectionResponse
}

// ListGroupByIdSiteByIdContentTypes ...
func (c GroupSiteContentTypeClient) ListGroupByIdSiteByIdContentTypes(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdContentTypesOperationResponse, err error) {
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

// ListGroupByIdSiteByIdContentTypesComplete retrieves all the results into a single object
func (c GroupSiteContentTypeClient) ListGroupByIdSiteByIdContentTypesComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdContentTypesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdContentTypesCompleteMatchingPredicate(ctx, id, models.ContentTypeCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdContentTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteContentTypeClient) ListGroupByIdSiteByIdContentTypesCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.ContentTypeCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdContentTypesCompleteResult, err error) {
	items := make([]models.ContentTypeCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdContentTypes(ctx, id)
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

	result = ListGroupByIdSiteByIdContentTypesCompleteResult{
		Items: items,
	}
	return
}
