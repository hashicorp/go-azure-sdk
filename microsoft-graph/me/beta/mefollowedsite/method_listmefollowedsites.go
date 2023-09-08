package mefollowedsite

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

type ListMeFollowedSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SiteCollectionResponse
}

type ListMeFollowedSitesCompleteResult struct {
	Items []models.SiteCollectionResponse
}

// ListMeFollowedSites ...
func (c MeFollowedSiteClient) ListMeFollowedSites(ctx context.Context) (result ListMeFollowedSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/followedSites",
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

// ListMeFollowedSitesComplete retrieves all the results into a single object
func (c MeFollowedSiteClient) ListMeFollowedSitesComplete(ctx context.Context) (ListMeFollowedSitesCompleteResult, error) {
	return c.ListMeFollowedSitesCompleteMatchingPredicate(ctx, models.SiteCollectionResponseOperationPredicate{})
}

// ListMeFollowedSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeFollowedSiteClient) ListMeFollowedSitesCompleteMatchingPredicate(ctx context.Context, predicate models.SiteCollectionResponseOperationPredicate) (result ListMeFollowedSitesCompleteResult, err error) {
	items := make([]models.SiteCollectionResponse, 0)

	resp, err := c.ListMeFollowedSites(ctx)
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

	result = ListMeFollowedSitesCompleteResult{
		Items: items,
	}
	return
}
