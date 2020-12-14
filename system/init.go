// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package system 系统基础插件包
package system

import (
	_ "github.com/33cn/dplatformos/system/consensus/init" //register consensus init package
	_ "github.com/33cn/dplatformos/system/crypto/init"
	_ "github.com/33cn/dplatformos/system/dapp/init"
	_ "github.com/33cn/dplatformos/system/mempool/init"
	_ "github.com/33cn/dplatformos/system/p2p/init" // init p2p plugin
	_ "github.com/33cn/dplatformos/system/store/init"
)
