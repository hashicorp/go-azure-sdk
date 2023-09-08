package serviceprincipalsynchronizationjob

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

type ListServicePrincipalByIdSynchronizationJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SynchronizationJobCollectionResponse
}

type ListServicePrincipalByIdSynchronizationJobsCompleteResult struct {
	Items []models.SynchronizationJobCollectionResponse
}

// ListServicePrincipalByIdSynchronizationJobs ...
func (c ServicePrincipalSynchronizationJobClient) ListServicePrincipalByIdSynchronizationJobs(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdSynchronizationJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/synchronization/jobs", id.ID()),
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
		Values *[]models.SynchronizationJobCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalByIdSynchronizationJobsComplete retrieves all the results into a single object
func (c ServicePrincipalSynchronizationJobClient) ListServicePrincipalByIdSynchronizationJobsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdSynchronizationJobsCompleteResult, error) {
	return c.ListServicePrincipalByIdSynchronizationJobsCompleteMatchingPredicate(ctx, id, models.SynchronizationJobCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdSynchronizationJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalSynchronizationJobClient) ListServicePrincipalByIdSynchronizationJobsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.SynchronizationJobCollectionResponseOperationPredicate) (result ListServicePrincipalByIdSynchronizationJobsCompleteResult, err error) {
	items := make([]models.SynchronizationJobCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdSynchronizationJobs(ctx, id)
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

	result = ListServicePrincipalByIdSynchronizationJobsCompleteResult{
		Items: items,
	}
	return
}
