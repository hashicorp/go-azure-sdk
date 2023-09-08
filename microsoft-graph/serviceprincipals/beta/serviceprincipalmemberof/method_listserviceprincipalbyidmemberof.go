package serviceprincipalmemberof

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

type ListServicePrincipalByIdMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListServicePrincipalByIdMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListServicePrincipalByIdMemberOf ...
func (c ServicePrincipalMemberOfClient) ListServicePrincipalByIdMemberOf(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/memberOf", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalByIdMemberOfComplete retrieves all the results into a single object
func (c ServicePrincipalMemberOfClient) ListServicePrincipalByIdMemberOfComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdMemberOfCompleteResult, error) {
	return c.ListServicePrincipalByIdMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalMemberOfClient) ListServicePrincipalByIdMemberOfCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListServicePrincipalByIdMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdMemberOf(ctx, id)
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

	result = ListServicePrincipalByIdMemberOfCompleteResult{
		Items: items,
	}
	return
}
