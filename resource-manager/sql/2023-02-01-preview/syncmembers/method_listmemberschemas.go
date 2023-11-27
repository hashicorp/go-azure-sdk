package syncmembers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMemberSchemasOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SyncFullSchemaProperties
}

type ListMemberSchemasCompleteResult struct {
	Items []SyncFullSchemaProperties
}

// ListMemberSchemas ...
func (c SyncMembersClient) ListMemberSchemas(ctx context.Context, id SyncMemberId) (result ListMemberSchemasOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schemas", id.ID()),
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
		Values *[]SyncFullSchemaProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMemberSchemasComplete retrieves all the results into a single object
func (c SyncMembersClient) ListMemberSchemasComplete(ctx context.Context, id SyncMemberId) (ListMemberSchemasCompleteResult, error) {
	return c.ListMemberSchemasCompleteMatchingPredicate(ctx, id, SyncFullSchemaPropertiesOperationPredicate{})
}

// ListMemberSchemasCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SyncMembersClient) ListMemberSchemasCompleteMatchingPredicate(ctx context.Context, id SyncMemberId, predicate SyncFullSchemaPropertiesOperationPredicate) (result ListMemberSchemasCompleteResult, err error) {
	items := make([]SyncFullSchemaProperties, 0)

	resp, err := c.ListMemberSchemas(ctx, id)
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

	result = ListMemberSchemasCompleteResult{
		Items: items,
	}
	return
}
