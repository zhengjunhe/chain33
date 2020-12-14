// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

// Package main dplatform程序入口
package main

import (
	_ "github.com/33cn/dplatform/system"
	"github.com/33cn/dplatform/util/cli"
)

func main() {
	cli.RunDplatformOS("", "")
}
