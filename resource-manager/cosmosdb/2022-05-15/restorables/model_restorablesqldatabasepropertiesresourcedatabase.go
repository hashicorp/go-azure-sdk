package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableSqlDatabasePropertiesResourceDatabase struct {
	Colls *string `json:"_colls,omitempty"`
	Self  *string `json:"_self,omitempty"`
	Users *string `json:"_users,omitempty"`
}
