package userprofileemail

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

type ListUserByIdProfileEmailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemEmailCollectionResponse
}

type ListUserByIdProfileEmailsCompleteResult struct {
	Items []models.ItemEmailCollectionResponse
}

// ListUserByIdProfileEmails ...
func (c UserProfileEmailClient) ListUserByIdProfileEmails(ctx context.Context, id UserId) (result ListUserByIdProfileEmailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/emails", id.ID()),
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
		Values *[]models.ItemEmailCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileEmailsComplete retrieves all the results into a single object
func (c UserProfileEmailClient) ListUserByIdProfileEmailsComplete(ctx context.Context, id UserId) (ListUserByIdProfileEmailsCompleteResult, error) {
	return c.ListUserByIdProfileEmailsCompleteMatchingPredicate(ctx, id, models.ItemEmailCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileEmailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileEmailClient) ListUserByIdProfileEmailsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ItemEmailCollectionResponseOperationPredicate) (result ListUserByIdProfileEmailsCompleteResult, err error) {
	items := make([]models.ItemEmailCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileEmails(ctx, id)
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

	result = ListUserByIdProfileEmailsCompleteResult{
		Items: items,
	}
	return
}
