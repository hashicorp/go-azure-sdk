package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableSqlDatabasePropertiesResourceDatabase struct {
	Colls             *string                `json:"_colls,omitempty"`
	CreateMode        *CreateMode            `json:"createMode,omitempty"`
	Etag              *string                `json:"_etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	RestoreParameters *RestoreParametersBase `json:"restoreParameters,omitempty"`
	Rid               *string                `json:"_rid,omitempty"`
	Self              *string                `json:"_self,omitempty"`
	Ts                *float64               `json:"_ts,omitempty"`
	Users             *string                `json:"_users,omitempty"`
}
