package meprofileproject

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

type ListMeProfileProjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ProjectParticipationCollectionResponse
}

type ListMeProfileProjectsCompleteResult struct {
	Items []models.ProjectParticipationCollectionResponse
}

// ListMeProfileProjects ...
func (c MeProfileProjectClient) ListMeProfileProjects(ctx context.Context) (result ListMeProfileProjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/projects",
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

// ListMeProfileProjectsComplete retrieves all the results into a single object
func (c MeProfileProjectClient) ListMeProfileProjectsComplete(ctx context.Context) (ListMeProfileProjectsCompleteResult, error) {
	return c.ListMeProfileProjectsCompleteMatchingPredicate(ctx, models.ProjectParticipationCollectionResponseOperationPredicate{})
}

// ListMeProfileProjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileProjectClient) ListMeProfileProjectsCompleteMatchingPredicate(ctx context.Context, predicate models.ProjectParticipationCollectionResponseOperationPredicate) (result ListMeProfileProjectsCompleteResult, err error) {
	items := make([]models.ProjectParticipationCollectionResponse, 0)

	resp, err := c.ListMeProfileProjects(ctx)
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

	result = ListMeProfileProjectsCompleteResult{
		Items: items,
	}
	return
}
