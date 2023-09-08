package groupsite

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

type AddGroupByIdSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.Site
}

type AddGroupByIdSitesCompleteResult struct {
	Items []models.Site
}

// AddGroupByIdSites ...
func (c GroupSiteClient) AddGroupByIdSites(ctx context.Context, id GroupId) (result AddGroupByIdSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sites/add", id.ID()),
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

// AddGroupByIdSitesComplete retrieves all the results into a single object
func (c GroupSiteClient) AddGroupByIdSitesComplete(ctx context.Context, id GroupId) (AddGroupByIdSitesCompleteResult, error) {
	return c.AddGroupByIdSitesCompleteMatchingPredicate(ctx, id, models.SiteOperationPredicate{})
}

// AddGroupByIdSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteClient) AddGroupByIdSitesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.SiteOperationPredicate) (result AddGroupByIdSitesCompleteResult, err error) {
	items := make([]models.Site, 0)

	resp, err := c.AddGroupByIdSites(ctx, id)
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

	result = AddGroupByIdSitesCompleteResult{
		Items: items,
	}
	return
}
