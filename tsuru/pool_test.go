// Copyright 2015 tsuru-client authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"github.com/tsuru/tsuru/cmd"
	"github.com/tsuru/tsuru/cmd/cmdtest"
	"gopkg.in/check.v1"
	"net/http"
)

func (s *S) TestPoolListInfo(c *check.C) {
	c.Assert((&poolList{}).Info(), check.NotNil)
}

func (s *S) TestPoolListRun(c *check.C) {
	var stdout, stderr bytes.Buffer
	result := `[{"team": "test", "pools": ["pool"]}]`
	context := cmd.Context{
		Args:   []string{},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	expected := `+------+-------+
| Team | Pools |
+------+-------+
| test | pool  |
+------+-------+
`
	client := cmd.NewClient(&http.Client{Transport: &cmdtest.Transport{Message: result, Status: http.StatusOK}}, nil, manager)
	command := poolList{}
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(stdout.String(), check.Equals, expected)
}
