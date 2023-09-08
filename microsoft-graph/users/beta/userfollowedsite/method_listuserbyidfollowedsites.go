package userfollowedsite

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

type ListUserByIdFollowedSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SiteCollectionResponse
}

type ListUserByIdFollowedSitesCompleteResult struct {
	Items []models.SiteCollectionResponse
}

// ListUserByIdFollowedSites ...
func (c UserFollowedSiteClient) ListUserByIdFollowedSites(ctx context.Context, id UserId) (result ListUserByIdFollowedSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/followedSites", id.ID()),
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

// ListUserByIdFollowedSitesComplete retrieves all the results into a single object
func (c UserFollowedSiteClient) ListUserByIdFollowedSitesComplete(ctx context.Context, id UserId) (ListUserByIdFollowedSitesCompleteResult, error) {
	return c.ListUserByIdFollowedSitesCompleteMatchingPredicate(ctx, id, models.SiteCollectionResponseOperationPredicate{})
}

// ListUserByIdFollowedSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserFollowedSiteClient) ListUserByIdFollowedSitesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.SiteCollectionResponseOperationPredicate) (result ListUserByIdFollowedSitesCompleteResult, err error) {
	items := make([]models.SiteCollectionResponse, 0)

	resp, err := c.ListUserByIdFollowedSites(ctx, id)
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

	result = ListUserByIdFollowedSitesCompleteResult{
		Items: items,
	}
	return
}
