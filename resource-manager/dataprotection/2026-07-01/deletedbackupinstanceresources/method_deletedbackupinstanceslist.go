package deletedbackupinstanceresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedBackupInstancesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeletedBackupInstanceResource
}

type DeletedBackupInstancesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeletedBackupInstanceResource
}

type DeletedBackupInstancesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeletedBackupInstancesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeletedBackupInstancesList ...
func (c DeletedBackupInstanceResourcesClient) DeletedBackupInstancesList(ctx context.Context, id BackupVaultId) (result DeletedBackupInstancesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DeletedBackupInstancesListCustomPager{},
		Path:       fmt.Sprintf("%s/deletedBackupInstances", id.ID()),
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
		Values *[]DeletedBackupInstanceResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DeletedBackupInstancesListComplete retrieves all the results into a single object
func (c DeletedBackupInstanceResourcesClient) DeletedBackupInstancesListComplete(ctx context.Context, id BackupVaultId) (DeletedBackupInstancesListCompleteResult, error) {
	return c.DeletedBackupInstancesListCompleteMatchingPredicate(ctx, id, DeletedBackupInstanceResourceOperationPredicate{})
}

// DeletedBackupInstancesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedBackupInstanceResourcesClient) DeletedBackupInstancesListCompleteMatchingPredicate(ctx context.Context, id BackupVaultId, predicate DeletedBackupInstanceResourceOperationPredicate) (result DeletedBackupInstancesListCompleteResult, err error) {
	items := make([]DeletedBackupInstanceResource, 0)

	resp, err := c.DeletedBackupInstancesList(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = DeletedBackupInstancesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
