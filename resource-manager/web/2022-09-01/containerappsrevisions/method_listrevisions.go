package containerappsrevisions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRevisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Revision
}

type ListRevisionsCompleteResult struct {
	Items []Revision
}

// ListRevisions ...
func (c ContainerAppsRevisionsClient) ListRevisions(ctx context.Context, id ProviderContainerAppId) (result ListRevisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/revisions", id.ID()),
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
		Values *[]Revision `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRevisionsComplete retrieves all the results into a single object
func (c ContainerAppsRevisionsClient) ListRevisionsComplete(ctx context.Context, id ProviderContainerAppId) (ListRevisionsCompleteResult, error) {
	return c.ListRevisionsCompleteMatchingPredicate(ctx, id, RevisionOperationPredicate{})
}

// ListRevisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainerAppsRevisionsClient) ListRevisionsCompleteMatchingPredicate(ctx context.Context, id ProviderContainerAppId, predicate RevisionOperationPredicate) (result ListRevisionsCompleteResult, err error) {
	items := make([]Revision, 0)

	resp, err := c.ListRevisions(ctx, id)
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

	result = ListRevisionsCompleteResult{
		Items: items,
	}
	return
}
