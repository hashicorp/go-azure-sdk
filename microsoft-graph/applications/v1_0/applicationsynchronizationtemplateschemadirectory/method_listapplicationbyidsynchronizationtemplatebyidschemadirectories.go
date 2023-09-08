package applicationsynchronizationtemplateschemadirectory

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

type ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryDefinitionCollectionResponse
}

type ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteResult struct {
	Items []models.DirectoryDefinitionCollectionResponse
}

// ListApplicationByIdSynchronizationTemplateByIdSchemaDirectories ...
func (c ApplicationSynchronizationTemplateSchemaDirectoryClient) ListApplicationByIdSynchronizationTemplateByIdSchemaDirectories(ctx context.Context, id ApplicationSynchronizationTemplateId) (result ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schema/directories", id.ID()),
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
		Values *[]models.DirectoryDefinitionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesComplete retrieves all the results into a single object
func (c ApplicationSynchronizationTemplateSchemaDirectoryClient) ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesComplete(ctx context.Context, id ApplicationSynchronizationTemplateId) (ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteResult, error) {
	return c.ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteMatchingPredicate(ctx, id, models.DirectoryDefinitionCollectionResponseOperationPredicate{})
}

// ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationSynchronizationTemplateSchemaDirectoryClient) ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteMatchingPredicate(ctx context.Context, id ApplicationSynchronizationTemplateId, predicate models.DirectoryDefinitionCollectionResponseOperationPredicate) (result ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteResult, err error) {
	items := make([]models.DirectoryDefinitionCollectionResponse, 0)

	resp, err := c.ListApplicationByIdSynchronizationTemplateByIdSchemaDirectories(ctx, id)
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

	result = ListApplicationByIdSynchronizationTemplateByIdSchemaDirectoriesCompleteResult{
		Items: items,
	}
	return
}
