package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnectionWithSystemData
}

type PrivateEndpointConnectionsListByWorkspaceCompleteResult struct {
	Items []PrivateEndpointConnectionWithSystemData
}

// PrivateEndpointConnectionsListByWorkspace ...
func (c PrivateLinkClient) PrivateEndpointConnectionsListByWorkspace(ctx context.Context, id WorkspaceId) (result PrivateEndpointConnectionsListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
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
		Values *[]PrivateEndpointConnectionWithSystemData `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateEndpointConnectionsListByWorkspaceComplete retrieves all the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (PrivateEndpointConnectionsListByWorkspaceCompleteResult, error) {
	return c.PrivateEndpointConnectionsListByWorkspaceCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionWithSystemDataOperationPredicate{})
}

// PrivateEndpointConnectionsListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate PrivateEndpointConnectionWithSystemDataOperationPredicate) (result PrivateEndpointConnectionsListByWorkspaceCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionWithSystemData, 0)

	resp, err := c.PrivateEndpointConnectionsListByWorkspace(ctx, id)
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

	result = PrivateEndpointConnectionsListByWorkspaceCompleteResult{
		Items: items,
	}
	return
}
