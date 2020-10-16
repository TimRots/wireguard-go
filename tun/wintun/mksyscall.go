/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2020 WireGuard LLC. All Rights Reserved.
 */

package wintun

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -systemdll=0 -output zsyscall_windows.go wintun_windows.go
