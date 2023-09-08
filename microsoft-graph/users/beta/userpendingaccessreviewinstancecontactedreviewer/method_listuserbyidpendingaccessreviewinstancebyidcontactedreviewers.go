package userpendingaccessreviewinstancecontactedreviewer

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

type ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewReviewerCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteResult struct {
	Items []models.AccessReviewReviewerCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdContactedReviewers ...
func (c UserPendingAccessReviewInstanceContactedReviewerClient) ListUserByIdPendingAccessReviewInstanceByIdContactedReviewers(ctx context.Context, id UserPendingAccessReviewInstanceId) (result ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceContactedReviewerClient) ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersComplete(ctx context.Context, id UserPendingAccessReviewInstanceId) (ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteMatchingPredicate(ctx, id, models.AccessReviewReviewerCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceContactedReviewerClient) ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceId, predicate models.AccessReviewReviewerCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteResult, err error) {
	items := make([]models.AccessReviewReviewerCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdContactedReviewers(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdContactedReviewersCompleteResult{
		Items: items,
	}
	return
}
