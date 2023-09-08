package userprofileskill

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

type ListUserByIdProfileSkillsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SkillProficiencyCollectionResponse
}

type ListUserByIdProfileSkillsCompleteResult struct {
	Items []models.SkillProficiencyCollectionResponse
}

// ListUserByIdProfileSkills ...
func (c UserProfileSkillClient) ListUserByIdProfileSkills(ctx context.Context, id UserId) (result ListUserByIdProfileSkillsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/skills", id.ID()),
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
		Values *[]models.SkillProficiencyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileSkillsComplete retrieves all the results into a single object
func (c UserProfileSkillClient) ListUserByIdProfileSkillsComplete(ctx context.Context, id UserId) (ListUserByIdProfileSkillsCompleteResult, error) {
	return c.ListUserByIdProfileSkillsCompleteMatchingPredicate(ctx, id, models.SkillProficiencyCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileSkillsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileSkillClient) ListUserByIdProfileSkillsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.SkillProficiencyCollectionResponseOperationPredicate) (result ListUserByIdProfileSkillsCompleteResult, err error) {
	items := make([]models.SkillProficiencyCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileSkills(ctx, id)
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

	result = ListUserByIdProfileSkillsCompleteResult{
		Items: items,
	}
	return
}
