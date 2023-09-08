package meprofilelanguage

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

type ListMeProfileLanguagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LanguageProficiencyCollectionResponse
}

type ListMeProfileLanguagesCompleteResult struct {
	Items []models.LanguageProficiencyCollectionResponse
}

// ListMeProfileLanguages ...
func (c MeProfileLanguageClient) ListMeProfileLanguages(ctx context.Context) (result ListMeProfileLanguagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/languages",
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
		Values *[]models.LanguageProficiencyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeProfileLanguagesComplete retrieves all the results into a single object
func (c MeProfileLanguageClient) ListMeProfileLanguagesComplete(ctx context.Context) (ListMeProfileLanguagesCompleteResult, error) {
	return c.ListMeProfileLanguagesCompleteMatchingPredicate(ctx, models.LanguageProficiencyCollectionResponseOperationPredicate{})
}

// ListMeProfileLanguagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileLanguageClient) ListMeProfileLanguagesCompleteMatchingPredicate(ctx context.Context, predicate models.LanguageProficiencyCollectionResponseOperationPredicate) (result ListMeProfileLanguagesCompleteResult, err error) {
	items := make([]models.LanguageProficiencyCollectionResponse, 0)

	resp, err := c.ListMeProfileLanguages(ctx)
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

	result = ListMeProfileLanguagesCompleteResult{
		Items: items,
	}
	return
}
