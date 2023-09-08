package serviceprincipalsynchronizationjobschemadirectory

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

type ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryDefinitionCollectionResponse
}

type ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult struct {
	Items []models.DirectoryDefinitionCollectionResponse
}

// ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectories ...
func (c ServicePrincipalSynchronizationJobSchemaDirectoryClient) ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectories(ctx context.Context, id ServicePrincipalSynchronizationJobId) (result ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesOperationResponse, err error) {
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

// ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesComplete retrieves all the results into a single object
func (c ServicePrincipalSynchronizationJobSchemaDirectoryClient) ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesComplete(ctx context.Context, id ServicePrincipalSynchronizationJobId) (ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult, error) {
	return c.ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteMatchingPredicate(ctx, id, models.DirectoryDefinitionCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalSynchronizationJobSchemaDirectoryClient) ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalSynchronizationJobId, predicate models.DirectoryDefinitionCollectionResponseOperationPredicate) (result ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult, err error) {
	items := make([]models.DirectoryDefinitionCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectories(ctx, id)
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

	result = ListServicePrincipalByIdSynchronizationJobByIdSchemaDirectoriesCompleteResult{
		Items: items,
	}
	return
}
