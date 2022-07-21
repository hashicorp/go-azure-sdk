package sshpublickeys

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/client"
	"github.com/hashicorp/go-azure-sdk/odata"
	"net/http"
)

// Copyright (c) TODO, Inc.

type ListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SshPublicKeyResource
	OData        *odata.OData
}

type ListByResourceGroupCompleteResult struct {
	Items []SshPublicKeyResource
}

// ListByResourceGroup ...
func (c SshPublicKeysClient) ListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ListByResourceGroupOperationResponse, err error) {
	req, err := c.Client2.NewGetRequest(ctx, fmt.Sprintf("%s/providers/Microsoft.Compute/sshPublicKeys", id.ID()), defaultApiVersion, odata.Query{})
	if err != nil {
		return
	}

	var resp *client.Response
	resp, result.OData, _, err = req.ExecutePaged()
	result.HttpResponse = resp.Response
	if err != nil {
		return
	}

	values := &struct {
		Values *[]SshPublicKeyResource `json:"value"`
	}{}
	if err = resp.Unmarshal(values); err != nil {
		return
	}
	result.Model = values.Values

	return
}

// ListByResourceGroupComplete retrieves all the results into a single object
func (c SshPublicKeysClient) ListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ListByResourceGroupCompleteResult, error) {
	return c.ListByResourceGroupCompleteMatchingPredicate(ctx, id, SshPublicKeyResourceOperationPredicate{})
}

// ListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SshPublicKeysClient) ListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate SshPublicKeyResourceOperationPredicate) (resp ListByResourceGroupCompleteResult, err error) {
	items := make([]SshPublicKeyResource, 0)

	result, err := c.ListByResourceGroup(ctx, id)
	if err != nil {
		err = fmt.Errorf("retrieving result: %+v", err)
		return
	}

	if result.Model != nil {
		for _, v := range *result.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	out := ListByResourceGroupCompleteResult{
		Items: items,
	}

	return out, nil
}
