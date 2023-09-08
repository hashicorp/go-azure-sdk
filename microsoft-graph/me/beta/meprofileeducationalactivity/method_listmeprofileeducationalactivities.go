package meprofileeducationalactivity

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

type ListMeProfileEducationalActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EducationalActivityCollectionResponse
}

type ListMeProfileEducationalActivitiesCompleteResult struct {
	Items []models.EducationalActivityCollectionResponse
}

// ListMeProfileEducationalActivities ...
func (c MeProfileEducationalActivityClient) ListMeProfileEducationalActivities(ctx context.Context) (result ListMeProfileEducationalActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/educationalActivities",
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
		Values *[]models.EducationalActivityCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeProfileEducationalActivitiesComplete retrieves all the results into a single object
func (c MeProfileEducationalActivityClient) ListMeProfileEducationalActivitiesComplete(ctx context.Context) (ListMeProfileEducationalActivitiesCompleteResult, error) {
	return c.ListMeProfileEducationalActivitiesCompleteMatchingPredicate(ctx, models.EducationalActivityCollectionResponseOperationPredicate{})
}

// ListMeProfileEducationalActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileEducationalActivityClient) ListMeProfileEducationalActivitiesCompleteMatchingPredicate(ctx context.Context, predicate models.EducationalActivityCollectionResponseOperationPredicate) (result ListMeProfileEducationalActivitiesCompleteResult, err error) {
	items := make([]models.EducationalActivityCollectionResponse, 0)

	resp, err := c.ListMeProfileEducationalActivities(ctx)
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

	result = ListMeProfileEducationalActivitiesCompleteResult{
		Items: items,
	}
	return
}
