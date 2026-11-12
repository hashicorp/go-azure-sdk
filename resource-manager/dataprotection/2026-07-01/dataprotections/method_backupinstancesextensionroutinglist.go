package dataprotections

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

type BackupInstancesExtensionRoutingListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BackupInstanceResource
}

type BackupInstancesExtensionRoutingListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BackupInstanceResource
}

type BackupInstancesExtensionRoutingListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BackupInstancesExtensionRoutingListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BackupInstancesExtensionRoutingList ...
func (c DataprotectionsClient) BackupInstancesExtensionRoutingList(ctx context.Context, id commonids.ScopeId) (result BackupInstancesExtensionRoutingListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &BackupInstancesExtensionRoutingListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.DataProtection/backupInstances", id.ID()),
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
		Values *[]BackupInstanceResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BackupInstancesExtensionRoutingListComplete retrieves all the results into a single object
func (c DataprotectionsClient) BackupInstancesExtensionRoutingListComplete(ctx context.Context, id commonids.ScopeId) (BackupInstancesExtensionRoutingListCompleteResult, error) {
	return c.BackupInstancesExtensionRoutingListCompleteMatchingPredicate(ctx, id, BackupInstanceResourceOperationPredicate{})
}

// BackupInstancesExtensionRoutingListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataprotectionsClient) BackupInstancesExtensionRoutingListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate BackupInstanceResourceOperationPredicate) (result BackupInstancesExtensionRoutingListCompleteResult, err error) {
	items := make([]BackupInstanceResource, 0)

	resp, err := c.BackupInstancesExtensionRoutingList(ctx, id)
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

	result = BackupInstancesExtensionRoutingListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
