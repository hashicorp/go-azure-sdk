package databasemigrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoToCosmosDbRUMongoGetForScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseMigrationCosmosDbMongo
}

type MongoToCosmosDbRUMongoGetForScopeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatabaseMigrationCosmosDbMongo
}

type MongoToCosmosDbRUMongoGetForScopeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *MongoToCosmosDbRUMongoGetForScopeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// MongoToCosmosDbRUMongoGetForScope ...
func (c DatabaseMigrationsClient) MongoToCosmosDbRUMongoGetForScope(ctx context.Context, id DatabaseAccountId) (result MongoToCosmosDbRUMongoGetForScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &MongoToCosmosDbRUMongoGetForScopeCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.DataMigration/databaseMigrations", id.ID()),
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
		Values *[]DatabaseMigrationCosmosDbMongo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MongoToCosmosDbRUMongoGetForScopeComplete retrieves all the results into a single object
func (c DatabaseMigrationsClient) MongoToCosmosDbRUMongoGetForScopeComplete(ctx context.Context, id DatabaseAccountId) (MongoToCosmosDbRUMongoGetForScopeCompleteResult, error) {
	return c.MongoToCosmosDbRUMongoGetForScopeCompleteMatchingPredicate(ctx, id, DatabaseMigrationCosmosDbMongoOperationPredicate{})
}

// MongoToCosmosDbRUMongoGetForScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabaseMigrationsClient) MongoToCosmosDbRUMongoGetForScopeCompleteMatchingPredicate(ctx context.Context, id DatabaseAccountId, predicate DatabaseMigrationCosmosDbMongoOperationPredicate) (result MongoToCosmosDbRUMongoGetForScopeCompleteResult, err error) {
	items := make([]DatabaseMigrationCosmosDbMongo, 0)

	resp, err := c.MongoToCosmosDbRUMongoGetForScope(ctx, id)
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

	result = MongoToCosmosDbRUMongoGetForScopeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
