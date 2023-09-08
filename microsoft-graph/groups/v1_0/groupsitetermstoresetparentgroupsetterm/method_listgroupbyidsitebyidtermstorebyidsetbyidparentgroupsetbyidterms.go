package groupsitetermstoresetparentgroupsetterm

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

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTerms ...
func (c GroupSiteTermStoreSetParentGroupSetTermClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTerms(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetId) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetParentGroupSetTermClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsComplete(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetId) (ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetParentGroupSetTermClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTerms(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdTermsCompleteResult{
		Items: items,
	}
	return
}
