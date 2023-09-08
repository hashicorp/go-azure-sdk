package applicationsynchronizationjobschemadirectory

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

type ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryDefinitionCollectionResponse
}

type ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult struct {
	Items []models.DirectoryDefinitionCollectionResponse
}

// ListApplicationByIdSynchronizationJobByIdSchemaDirectories ...
func (c ApplicationSynchronizationJobSchemaDirectoryClient) ListApplicationByIdSynchronizationJobByIdSchemaDirectories(ctx context.Context, id ApplicationSynchronizationJobId) (result ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesOperationResponse, err error) {
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

// ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesComplete retrieves all the results into a single object
func (c ApplicationSynchronizationJobSchemaDirectoryClient) ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesComplete(ctx context.Context, id ApplicationSynchronizationJobId) (ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult, error) {
	return c.ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteMatchingPredicate(ctx, id, models.DirectoryDefinitionCollectionResponseOperationPredicate{})
}

// ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationSynchronizationJobSchemaDirectoryClient) ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteMatchingPredicate(ctx context.Context, id ApplicationSynchronizationJobId, predicate models.DirectoryDefinitionCollectionResponseOperationPredicate) (result ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult, err error) {
	items := make([]models.DirectoryDefinitionCollectionResponse, 0)

	resp, err := c.ListApplicationByIdSynchronizationJobByIdSchemaDirectories(ctx, id)
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

	result = ListApplicationByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult{
		Items: items,
	}
	return
}
