// Upvest Confidential
//
// Copyright 2020 - 2023 Upvest GmbH. All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains the property
// of Upvest GmbH. The intellectual and technical concepts contained herein
// are proprietary to Upvest GmbH and are protected by trade secret or
// copyright law. Dissemination of this information or reproduction of this
// material is strictly forbidden unless prior written permission is
// obtained from Upvest GmbH.

package static

import (
	"embed"
)

//go:embed *
var FS embed.FS
