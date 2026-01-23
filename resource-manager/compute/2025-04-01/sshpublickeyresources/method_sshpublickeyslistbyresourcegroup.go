package sshpublickeyresources

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

type SshPublicKeysListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SshPublicKeyResource
}

type SshPublicKeysListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SshPublicKeyResource
}

type SshPublicKeysListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SshPublicKeysListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SshPublicKeysListByResourceGroup ...
func (c SshPublicKeyResourcesClient) SshPublicKeysListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result SshPublicKeysListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SshPublicKeysListByResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Compute/sshPublicKeys", id.ID()),
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
		Values *[]SshPublicKeyResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SshPublicKeysListByResourceGroupComplete retrieves all the results into a single object
func (c SshPublicKeyResourcesClient) SshPublicKeysListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (SshPublicKeysListByResourceGroupCompleteResult, error) {
	return c.SshPublicKeysListByResourceGroupCompleteMatchingPredicate(ctx, id, SshPublicKeyResourceOperationPredicate{})
}

// SshPublicKeysListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SshPublicKeyResourcesClient) SshPublicKeysListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate SshPublicKeyResourceOperationPredicate) (result SshPublicKeysListByResourceGroupCompleteResult, err error) {
	items := make([]SshPublicKeyResource, 0)

	resp, err := c.SshPublicKeysListByResourceGroup(ctx, id)
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

	result = SshPublicKeysListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
