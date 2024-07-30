package synchronizationtemplateschemadirectory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSynchronizationTemplateSchemaDirectoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryDefinition
}

type ListSynchronizationTemplateSchemaDirectoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryDefinition
}

type ListSynchronizationTemplateSchemaDirectoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationTemplateSchemaDirectoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationTemplateSchemaDirectories ...
func (c SynchronizationTemplateSchemaDirectoryClient) ListSynchronizationTemplateSchemaDirectories(ctx context.Context, id ApplicationIdSynchronizationTemplateId) (result ListSynchronizationTemplateSchemaDirectoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSynchronizationTemplateSchemaDirectoriesCustomPager{},
		Path:       fmt.Sprintf("%s/schema/directories", id.ID()),
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
		Values *[]stable.DirectoryDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationTemplateSchemaDirectoriesComplete retrieves all the results into a single object
func (c SynchronizationTemplateSchemaDirectoryClient) ListSynchronizationTemplateSchemaDirectoriesComplete(ctx context.Context, id ApplicationIdSynchronizationTemplateId) (ListSynchronizationTemplateSchemaDirectoriesCompleteResult, error) {
	return c.ListSynchronizationTemplateSchemaDirectoriesCompleteMatchingPredicate(ctx, id, DirectoryDefinitionOperationPredicate{})
}

// ListSynchronizationTemplateSchemaDirectoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationTemplateSchemaDirectoryClient) ListSynchronizationTemplateSchemaDirectoriesCompleteMatchingPredicate(ctx context.Context, id ApplicationIdSynchronizationTemplateId, predicate DirectoryDefinitionOperationPredicate) (result ListSynchronizationTemplateSchemaDirectoriesCompleteResult, err error) {
	items := make([]stable.DirectoryDefinition, 0)

	resp, err := c.ListSynchronizationTemplateSchemaDirectories(ctx, id)
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

	result = ListSynchronizationTemplateSchemaDirectoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
