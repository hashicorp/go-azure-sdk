package userprofileproject

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

type ListUserByIdProfileProjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ProjectParticipationCollectionResponse
}

type ListUserByIdProfileProjectsCompleteResult struct {
	Items []models.ProjectParticipationCollectionResponse
}

// ListUserByIdProfileProjects ...
func (c UserProfileProjectClient) ListUserByIdProfileProjects(ctx context.Context, id UserId) (result ListUserByIdProfileProjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/projects", id.ID()),
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
		Values *[]models.ProjectParticipationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileProjectsComplete retrieves all the results into a single object
func (c UserProfileProjectClient) ListUserByIdProfileProjectsComplete(ctx context.Context, id UserId) (ListUserByIdProfileProjectsCompleteResult, error) {
	return c.ListUserByIdProfileProjectsCompleteMatchingPredicate(ctx, id, models.ProjectParticipationCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileProjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileProjectClient) ListUserByIdProfileProjectsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ProjectParticipationCollectionResponseOperationPredicate) (result ListUserByIdProfileProjectsCompleteResult, err error) {
	items := make([]models.ProjectParticipationCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileProjects(ctx, id)
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

	result = ListUserByIdProfileProjectsCompleteResult{
		Items: items,
	}
	return
}
