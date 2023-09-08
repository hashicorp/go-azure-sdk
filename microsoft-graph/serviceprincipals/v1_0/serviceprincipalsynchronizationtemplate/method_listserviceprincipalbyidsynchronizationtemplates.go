package serviceprincipalsynchronizationtemplate

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

type ListServicePrincipalByIdSynchronizationTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SynchronizationTemplateCollectionResponse
}

type ListServicePrincipalByIdSynchronizationTemplatesCompleteResult struct {
	Items []models.SynchronizationTemplateCollectionResponse
}

// ListServicePrincipalByIdSynchronizationTemplates ...
func (c ServicePrincipalSynchronizationTemplateClient) ListServicePrincipalByIdSynchronizationTemplates(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdSynchronizationTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/synchronization/templates", id.ID()),
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
		Values *[]models.SynchronizationTemplateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalByIdSynchronizationTemplatesComplete retrieves all the results into a single object
func (c ServicePrincipalSynchronizationTemplateClient) ListServicePrincipalByIdSynchronizationTemplatesComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdSynchronizationTemplatesCompleteResult, error) {
	return c.ListServicePrincipalByIdSynchronizationTemplatesCompleteMatchingPredicate(ctx, id, models.SynchronizationTemplateCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdSynchronizationTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalSynchronizationTemplateClient) ListServicePrincipalByIdSynchronizationTemplatesCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.SynchronizationTemplateCollectionResponseOperationPredicate) (result ListServicePrincipalByIdSynchronizationTemplatesCompleteResult, err error) {
	items := make([]models.SynchronizationTemplateCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdSynchronizationTemplates(ctx, id)
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

	result = ListServicePrincipalByIdSynchronizationTemplatesCompleteResult{
		Items: items,
	}
	return
}
