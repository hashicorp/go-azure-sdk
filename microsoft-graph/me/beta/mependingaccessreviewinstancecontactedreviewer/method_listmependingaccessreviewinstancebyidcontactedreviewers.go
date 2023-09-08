package mependingaccessreviewinstancecontactedreviewer

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

type ListMePendingAccessReviewInstanceByIdContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewReviewerCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteResult struct {
	Items []models.AccessReviewReviewerCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdContactedReviewers ...
func (c MePendingAccessReviewInstanceContactedReviewerClient) ListMePendingAccessReviewInstanceByIdContactedReviewers(ctx context.Context, id MePendingAccessReviewInstanceId) (result ListMePendingAccessReviewInstanceByIdContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/contactedReviewers", id.ID()),
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
		Values *[]models.AccessReviewReviewerCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMePendingAccessReviewInstanceByIdContactedReviewersComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceContactedReviewerClient) ListMePendingAccessReviewInstanceByIdContactedReviewersComplete(ctx context.Context, id MePendingAccessReviewInstanceId) (ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteMatchingPredicate(ctx, id, models.AccessReviewReviewerCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceContactedReviewerClient) ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceId, predicate models.AccessReviewReviewerCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteResult, err error) {
	items := make([]models.AccessReviewReviewerCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdContactedReviewers(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdContactedReviewersCompleteResult{
		Items: items,
	}
	return
}
