package groupsite

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

type ListGroupByIdSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SiteCollectionResponse
}

type ListGroupByIdSitesCompleteResult struct {
	Items []models.SiteCollectionResponse
}

// ListGroupByIdSites ...
func (c GroupSiteClient) ListGroupByIdSites(ctx context.Context, id GroupId) (result ListGroupByIdSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sites", id.ID()),
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
		Values *[]models.SiteCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSitesComplete retrieves all the results into a single object
func (c GroupSiteClient) ListGroupByIdSitesComplete(ctx context.Context, id GroupId) (ListGroupByIdSitesCompleteResult, error) {
	return c.ListGroupByIdSitesCompleteMatchingPredicate(ctx, id, models.SiteCollectionResponseOperationPredicate{})
}

// ListGroupByIdSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteClient) ListGroupByIdSitesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.SiteCollectionResponseOperationPredicate) (result ListGroupByIdSitesCompleteResult, err error) {
	items := make([]models.SiteCollectionResponse, 0)

	resp, err := c.ListGroupByIdSites(ctx, id)
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

	result = ListGroupByIdSitesCompleteResult{
		Items: items,
	}
	return
}
