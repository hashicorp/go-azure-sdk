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

type RemoveGroupByIdSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.Site
}

type RemoveGroupByIdSitesCompleteResult struct {
	Items []models.Site
}

// RemoveGroupByIdSites ...
func (c GroupSiteClient) RemoveGroupByIdSites(ctx context.Context, id GroupId) (result RemoveGroupByIdSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sites/remove", id.ID()),
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
		Values *[]models.Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RemoveGroupByIdSitesComplete retrieves all the results into a single object
func (c GroupSiteClient) RemoveGroupByIdSitesComplete(ctx context.Context, id GroupId) (RemoveGroupByIdSitesCompleteResult, error) {
	return c.RemoveGroupByIdSitesCompleteMatchingPredicate(ctx, id, models.SiteOperationPredicate{})
}

// RemoveGroupByIdSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteClient) RemoveGroupByIdSitesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.SiteOperationPredicate) (result RemoveGroupByIdSitesCompleteResult, err error) {
	items := make([]models.Site, 0)

	resp, err := c.RemoveGroupByIdSites(ctx, id)
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

	result = RemoveGroupByIdSitesCompleteResult{
		Items: items,
	}
	return
}
