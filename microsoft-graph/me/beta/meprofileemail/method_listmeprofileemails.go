package meprofileemail

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

type ListMeProfileEmailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemEmailCollectionResponse
}

type ListMeProfileEmailsCompleteResult struct {
	Items []models.ItemEmailCollectionResponse
}

// ListMeProfileEmails ...
func (c MeProfileEmailClient) ListMeProfileEmails(ctx context.Context) (result ListMeProfileEmailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/emails",
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

// ListMeProfileEmailsComplete retrieves all the results into a single object
func (c MeProfileEmailClient) ListMeProfileEmailsComplete(ctx context.Context) (ListMeProfileEmailsCompleteResult, error) {
	return c.ListMeProfileEmailsCompleteMatchingPredicate(ctx, models.ItemEmailCollectionResponseOperationPredicate{})
}

// ListMeProfileEmailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileEmailClient) ListMeProfileEmailsCompleteMatchingPredicate(ctx context.Context, predicate models.ItemEmailCollectionResponseOperationPredicate) (result ListMeProfileEmailsCompleteResult, err error) {
	items := make([]models.ItemEmailCollectionResponse, 0)

	resp, err := c.ListMeProfileEmails(ctx)
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

	result = ListMeProfileEmailsCompleteResult{
		Items: items,
	}
	return
}
