// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main dplatformos-cli程序入口
package main

import (

	// 这一步是必需的，目的时让插件源码有机会进行匿名注册
	"github.com/33cn/dplatformos/cmd/cli/buildflags"
	_ "github.com/33cn/dplatformos/system"
	"github.com/33cn/dplatformos/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:28803"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "")
}
