package groupsitesite

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

type ListGroupByIdSiteByIdSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SiteCollectionResponse
}

type ListGroupByIdSiteByIdSitesCompleteResult struct {
	Items []models.SiteCollectionResponse
}

// ListGroupByIdSiteByIdSites ...
func (c GroupSiteSiteClient) ListGroupByIdSiteByIdSites(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdSitesOperationResponse, err error) {
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

// ListGroupByIdSiteByIdSitesComplete retrieves all the results into a single object
func (c GroupSiteSiteClient) ListGroupByIdSiteByIdSitesComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdSitesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdSitesCompleteMatchingPredicate(ctx, id, models.SiteCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteSiteClient) ListGroupByIdSiteByIdSitesCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.SiteCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdSitesCompleteResult, err error) {
	items := make([]models.SiteCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdSites(ctx, id)
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

	result = ListGroupByIdSiteByIdSitesCompleteResult{
		Items: items,
	}
	return
}
