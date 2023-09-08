package groupsitetermstoregroupsetterm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTerms ...
func (c GroupSiteTermStoreGroupSetTermClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTerms(ctx context.Context, id GroupSiteTermStoreGroupSetId) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetTermClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsComplete(ctx context.Context, id GroupSiteTermStoreGroupSetId) (ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetTermClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupSetId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTerms(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermsCompleteResult{
		Items: items,
	}
	return
}
