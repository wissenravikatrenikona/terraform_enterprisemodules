package privateclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Encryption struct {
	KeyVaultProperties *EncryptionKeyVaultProperties `json:"keyVaultProperties,omitempty"`
	Status             *EncryptionState              `json:"status,omitempty"`
}
