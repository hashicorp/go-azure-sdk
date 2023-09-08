package groupsitetermstoresetterm

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

type ListGroupByIdSiteByIdTermStoreSetByIdTermsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreSetByIdTerms ...
func (c GroupSiteTermStoreSetTermClient) ListGroupByIdSiteByIdTermStoreSetByIdTerms(ctx context.Context, id GroupSiteTermStoreSetId) (result ListGroupByIdSiteByIdTermStoreSetByIdTermsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/terms", id.ID()),
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
		Values *[]models.TermStoreTermCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdTermStoreSetByIdTermsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetTermClient) ListGroupByIdSiteByIdTermStoreSetByIdTermsComplete(ctx context.Context, id GroupSiteTermStoreSetId) (ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetTermClient) ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreSetByIdTerms(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreSetByIdTermsCompleteResult{
		Items: items,
	}
	return
}
