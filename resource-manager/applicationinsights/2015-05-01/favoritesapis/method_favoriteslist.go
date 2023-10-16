package favoritesapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FavoritesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationInsightsComponentFavorite
}

type FavoritesListOperationOptions struct {
	CanFetchContent *bool
	FavoriteType    *FavoriteType
	SourceType      *FavoriteSourceType
	Tags            *string
}

func DefaultFavoritesListOperationOptions() FavoritesListOperationOptions {
	return FavoritesListOperationOptions{}
}

func (o FavoritesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o FavoritesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o FavoritesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.CanFetchContent != nil {
		out.Append("canFetchContent", fmt.Sprintf("%v", *o.CanFetchContent))
	}
	if o.FavoriteType != nil {
		out.Append("favoriteType", fmt.Sprintf("%v", *o.FavoriteType))
	}
	if o.SourceType != nil {
		out.Append("sourceType", fmt.Sprintf("%v", *o.SourceType))
	}
	if o.Tags != nil {
		out.Append("tags", fmt.Sprintf("%v", *o.Tags))
	}
	return &out
}

// FavoritesList ...
func (c FavoritesAPIsClient) FavoritesList(ctx context.Context, id ComponentId, options FavoritesListOperationOptions) (result FavoritesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/favorites", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
