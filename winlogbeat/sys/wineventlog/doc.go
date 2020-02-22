// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

/*
Package wineventlog provides access to the Windows Event Log API used in
all versions of Windows since Vista (i.e. Windows 7+  and Windows Server 2008+).
This is distinct from the Event Logging API that was used in Windows XP,
Windows Server 2003, and Windows 2000.
*/
package wineventlog

// Add -trace to enable debug prints around syscalls.
//go:generate go get golang.org/x/sys/windows/mkwinsyscall
//go:generate $GOPATH/bin/mkwinsyscall.exe -systemdll -output zsyscall_windows.go syscall_windows.go
