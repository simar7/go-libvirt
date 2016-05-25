// Copyright 2016 The go-libvirt Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build integration

package libvirt

import (
	"net"
	"testing"
	"time"
)

func testConn(t *testing.T) net.Conn {
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		t.Fatal("failed to dial libvirt: %v", err)
	}

	return c
}

func TestConnectIntegration(t *testing.T) {
	conn := testConn(t)
	l := New(conn)
	defer l.Disconnect()

	if err := l.Connect(); err != nil {
		t.Error(err)
	}
}

func TestDisconnectIntegration(t *testing.T) {
	conn := testConn(t)
	l := New(conn)

	if err := l.Connect(); err != nil {
		t.Error(err)
	}

	if err := l.Disconnect(); err != nil {
		t.Error(err)
	}
}
