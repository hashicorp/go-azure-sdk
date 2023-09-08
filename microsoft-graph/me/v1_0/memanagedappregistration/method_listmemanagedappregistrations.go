package memanagedappregistration

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

type ListMeManagedAppRegistrationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ManagedAppRegistrationCollectionResponse
}

type ListMeManagedAppRegistrationsCompleteResult struct {
	Items []models.ManagedAppRegistrationCollectionResponse
}

// ListMeManagedAppRegistrations ...
func (c MeManagedAppRegistrationClient) ListMeManagedAppRegistrations(ctx context.Context) (result ListMeManagedAppRegistrationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/managedAppRegistrations",
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
		Values *[]models.ManagedAppRegistrationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeManagedAppRegistrationsComplete retrieves all the results into a single object
func (c MeManagedAppRegistrationClient) ListMeManagedAppRegistrationsComplete(ctx context.Context) (ListMeManagedAppRegistrationsCompleteResult, error) {
	return c.ListMeManagedAppRegistrationsCompleteMatchingPredicate(ctx, models.ManagedAppRegistrationCollectionResponseOperationPredicate{})
}

// ListMeManagedAppRegistrationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedAppRegistrationClient) ListMeManagedAppRegistrationsCompleteMatchingPredicate(ctx context.Context, predicate models.ManagedAppRegistrationCollectionResponseOperationPredicate) (result ListMeManagedAppRegistrationsCompleteResult, err error) {
	items := make([]models.ManagedAppRegistrationCollectionResponse, 0)

	resp, err := c.ListMeManagedAppRegistrations(ctx)
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

	result = ListMeManagedAppRegistrationsCompleteResult{
		Items: items,
	}
	return
}
