package backupvaultoperationresults

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupVaultsGetInResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BackupVaultResource
}

type BackupVaultsGetInResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BackupVaultResource
}

type BackupVaultsGetInResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BackupVaultsGetInResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BackupVaultsGetInResourceGroup ...
func (c BackupVaultOperationResultsClient) BackupVaultsGetInResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result BackupVaultsGetInResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &BackupVaultsGetInResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.DataProtection/backupVaults", id.ID()),
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
		Values *[]BackupVaultResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BackupVaultsGetInResourceGroupComplete retrieves all the results into a single object
func (c BackupVaultOperationResultsClient) BackupVaultsGetInResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (BackupVaultsGetInResourceGroupCompleteResult, error) {
	return c.BackupVaultsGetInResourceGroupCompleteMatchingPredicate(ctx, id, BackupVaultResourceOperationPredicate{})
}

// BackupVaultsGetInResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BackupVaultOperationResultsClient) BackupVaultsGetInResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate BackupVaultResourceOperationPredicate) (result BackupVaultsGetInResourceGroupCompleteResult, err error) {
	items := make([]BackupVaultResource, 0)

	resp, err := c.BackupVaultsGetInResourceGroup(ctx, id)
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

	result = BackupVaultsGetInResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
