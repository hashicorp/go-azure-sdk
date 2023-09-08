package applicationsynchronizationtemplate

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

type ListApplicationByIdSynchronizationTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SynchronizationTemplateCollectionResponse
}

type ListApplicationByIdSynchronizationTemplatesCompleteResult struct {
	Items []models.SynchronizationTemplateCollectionResponse
}

// ListApplicationByIdSynchronizationTemplates ...
func (c ApplicationSynchronizationTemplateClient) ListApplicationByIdSynchronizationTemplates(ctx context.Context, id ApplicationId) (result ListApplicationByIdSynchronizationTemplatesOperationResponse, err error) {
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

// ListApplicationByIdSynchronizationTemplatesComplete retrieves all the results into a single object
func (c ApplicationSynchronizationTemplateClient) ListApplicationByIdSynchronizationTemplatesComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdSynchronizationTemplatesCompleteResult, error) {
	return c.ListApplicationByIdSynchronizationTemplatesCompleteMatchingPredicate(ctx, id, models.SynchronizationTemplateCollectionResponseOperationPredicate{})
}

// ListApplicationByIdSynchronizationTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationSynchronizationTemplateClient) ListApplicationByIdSynchronizationTemplatesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.SynchronizationTemplateCollectionResponseOperationPredicate) (result ListApplicationByIdSynchronizationTemplatesCompleteResult, err error) {
	items := make([]models.SynchronizationTemplateCollectionResponse, 0)

	resp, err := c.ListApplicationByIdSynchronizationTemplates(ctx, id)
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

	result = ListApplicationByIdSynchronizationTemplatesCompleteResult{
		Items: items,
	}
	return
}
