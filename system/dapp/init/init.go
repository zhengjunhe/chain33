// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package init 初始化系统dapp包
package init

import (
	_ "github.com/33cn/dplatformos/system/dapp/coins"  // register coins package
	_ "github.com/33cn/dplatformos/system/dapp/manage" // register manage package
	_ "github.com/33cn/dplatformos/system/dapp/none"   // register none package
)
