// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main dplatform-cli程序入口
package main

import (

	// 这一步是必需的，目的时让插件源码有机会进行匿名注册
	"github.com/33cn/dplatform/cmd/cli/buildflags"
	_ "github.com/33cn/dplatform/system"
	"github.com/33cn/dplatform/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:28803"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "")
}
