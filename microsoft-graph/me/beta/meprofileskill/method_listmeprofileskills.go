package meprofileskill

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

type ListMeProfileSkillsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SkillProficiencyCollectionResponse
}

type ListMeProfileSkillsCompleteResult struct {
	Items []models.SkillProficiencyCollectionResponse
}

// ListMeProfileSkills ...
func (c MeProfileSkillClient) ListMeProfileSkills(ctx context.Context) (result ListMeProfileSkillsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/skills",
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

// ListMeProfileSkillsComplete retrieves all the results into a single object
func (c MeProfileSkillClient) ListMeProfileSkillsComplete(ctx context.Context) (ListMeProfileSkillsCompleteResult, error) {
	return c.ListMeProfileSkillsCompleteMatchingPredicate(ctx, models.SkillProficiencyCollectionResponseOperationPredicate{})
}

// ListMeProfileSkillsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileSkillClient) ListMeProfileSkillsCompleteMatchingPredicate(ctx context.Context, predicate models.SkillProficiencyCollectionResponseOperationPredicate) (result ListMeProfileSkillsCompleteResult, err error) {
	items := make([]models.SkillProficiencyCollectionResponse, 0)

	resp, err := c.ListMeProfileSkills(ctx)
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

	result = ListMeProfileSkillsCompleteResult{
		Items: items,
	}
	return
}
