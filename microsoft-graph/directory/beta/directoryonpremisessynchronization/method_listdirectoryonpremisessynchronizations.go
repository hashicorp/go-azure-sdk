package directoryonpremisessynchronization

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

type ListDirectoryOnPremisesSynchronizationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnPremisesDirectorySynchronizationCollectionResponse
}

type ListDirectoryOnPremisesSynchronizationsCompleteResult struct {
	Items []models.OnPremisesDirectorySynchronizationCollectionResponse
}

// ListDirectoryOnPremisesSynchronizations ...
func (c DirectoryOnPremisesSynchronizationClient) ListDirectoryOnPremisesSynchronizations(ctx context.Context) (result ListDirectoryOnPremisesSynchronizationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/onPremisesSynchronization",
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
		Values *[]models.OnPremisesDirectorySynchronizationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryOnPremisesSynchronizationsComplete retrieves all the results into a single object
func (c DirectoryOnPremisesSynchronizationClient) ListDirectoryOnPremisesSynchronizationsComplete(ctx context.Context) (ListDirectoryOnPremisesSynchronizationsCompleteResult, error) {
	return c.ListDirectoryOnPremisesSynchronizationsCompleteMatchingPredicate(ctx, models.OnPremisesDirectorySynchronizationCollectionResponseOperationPredicate{})
}

// ListDirectoryOnPremisesSynchronizationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryOnPremisesSynchronizationClient) ListDirectoryOnPremisesSynchronizationsCompleteMatchingPredicate(ctx context.Context, predicate models.OnPremisesDirectorySynchronizationCollectionResponseOperationPredicate) (result ListDirectoryOnPremisesSynchronizationsCompleteResult, err error) {
	items := make([]models.OnPremisesDirectorySynchronizationCollectionResponse, 0)

	resp, err := c.ListDirectoryOnPremisesSynchronizations(ctx)
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

	result = ListDirectoryOnPremisesSynchronizationsCompleteResult{
		Items: items,
	}
	return
}
