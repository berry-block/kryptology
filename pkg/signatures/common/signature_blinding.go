//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/berry-block/kryptology/pkg/core/curves"
)

// SignatureBlinding is a value used for computing blind signatures
type SignatureBlinding = curves.PairingScalar
