package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GremlinGraphGetProperties struct {
	Options  *OptionsResource                   `json:"options,omitempty"`
	Resource *GremlinGraphGetPropertiesResource `json:"resource,omitempty"`
}
