package mongorbacs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoRoleDefinitionType string

const (
	MongoRoleDefinitionTypeBuiltInRole MongoRoleDefinitionType = "BuiltInRole"
	MongoRoleDefinitionTypeCustomRole  MongoRoleDefinitionType = "CustomRole"
)

func PossibleValuesForMongoRoleDefinitionType() []string {
	return []string{
		string(MongoRoleDefinitionTypeBuiltInRole),
		string(MongoRoleDefinitionTypeCustomRole),
	}
}
