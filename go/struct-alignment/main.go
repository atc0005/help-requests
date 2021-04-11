// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/help-requests
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

// https://play.golang.org/p/DxWIJd_OgPM
//
// https://golang.org/pkg/unsafe/#Sizeof
// https://golang.org/pkg/unsafe/#Alignof
// https://golang.org/ref/spec#Size_and_alignment_guarantees
// https://groups.google.com/g/golang-nuts/c/XDfQUn4U_g8/m/sBbZyKnuk5AJ
// https://groups.google.com/g/golang-nuts/c/rmHMi1wrHYk
// https://github.com/golangci/golangci-lint/issues/541#issuecomment-790941025

// fieldalignment: struct with 104 pointer bytes could be 64 (govet)
// fieldalignment: struct with 96 pointer bytes could be 64 (govet)
//
// var "eCfgAscending" of type main.emailConfigAscending has size 120 with Asignof 8
// var "eCfgDescending" of type main.emailConfigDescending has size 120 with Asignof 8

import (
	"fmt"
	"html/template"
	"time"
	"unsafe"
)

func main() {

	type emailConfigAscending struct {
		timeout                time.Duration      // 8 bytes
		notificationRateLimit  time.Duration      // 8 bytes
		template               *template.Template // 8 bytes
		serverPort             int                // 8 bytes
		notificationRetries    int                // 8 bytes
		notificationRetryDelay int                // 8 bytes
		server                 string             // 16 bytes
		senderAddress          string             // 16 bytes
		clientIdentity         string             // 16 bytes
		recipientAddresses     []string           // 24 bytes
	}

	type emailConfigDescending struct {
		recipientAddresses     []string           // 24 bytes
		server                 string             // 16 bytes
		senderAddress          string             // 16 bytes
		clientIdentity         string             // 16 bytes
		timeout                time.Duration      // 8 bytes
		notificationRateLimit  time.Duration      // 8 bytes
		template               *template.Template // 8 bytes
		serverPort             int                // 8 bytes
		notificationRetries    int                // 8 bytes
		notificationRetryDelay int                // 8 bytes
	}

	var eCfgAscending emailConfigAscending
	var eCfgDescending emailConfigDescending

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.timeout",
		eCfgAscending.timeout,
		unsafe.Sizeof(eCfgAscending.timeout),
		unsafe.Alignof(eCfgAscending.timeout),
		unsafe.Offsetof(eCfgAscending.timeout),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.notificationRateLimit",
		eCfgAscending.notificationRateLimit,
		unsafe.Sizeof(eCfgAscending.notificationRateLimit),
		unsafe.Alignof(eCfgAscending.notificationRateLimit),
		unsafe.Offsetof(eCfgAscending.notificationRateLimit),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.template",
		eCfgAscending.template,
		unsafe.Sizeof(eCfgAscending.template),
		unsafe.Alignof(eCfgAscending.template),
		unsafe.Offsetof(eCfgAscending.template),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.serverPort",
		eCfgAscending.serverPort,
		unsafe.Sizeof(eCfgAscending.serverPort),
		unsafe.Alignof(eCfgAscending.serverPort),
		unsafe.Offsetof(eCfgAscending.serverPort),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.notificationRetries",
		eCfgAscending.notificationRetries,
		unsafe.Sizeof(eCfgAscending.notificationRetries),
		unsafe.Alignof(eCfgAscending.notificationRetries),
		unsafe.Offsetof(eCfgAscending.notificationRetries),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.notificationRetryDelay",
		eCfgAscending.notificationRetryDelay,
		unsafe.Sizeof(eCfgAscending.notificationRetryDelay),
		unsafe.Alignof(eCfgAscending.notificationRetryDelay),
		unsafe.Offsetof(eCfgAscending.notificationRetryDelay),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.server",
		eCfgAscending.server,
		unsafe.Sizeof(eCfgAscending.server),
		unsafe.Alignof(eCfgAscending.server),
		unsafe.Offsetof(eCfgAscending.server),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.senderAddress",
		eCfgAscending.senderAddress,
		unsafe.Sizeof(eCfgAscending.senderAddress),
		unsafe.Alignof(eCfgAscending.senderAddress),
		unsafe.Offsetof(eCfgAscending.senderAddress),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.clientIdentity",
		eCfgAscending.clientIdentity,
		unsafe.Sizeof(eCfgAscending.clientIdentity),
		unsafe.Alignof(eCfgAscending.clientIdentity),
		unsafe.Offsetof(eCfgAscending.clientIdentity),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgAscending.recipientAddresses",
		eCfgAscending.recipientAddresses,
		unsafe.Sizeof(eCfgAscending.recipientAddresses),
		unsafe.Alignof(eCfgAscending.recipientAddresses),
		unsafe.Offsetof(eCfgAscending.recipientAddresses),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d\n",
		"eCfgAscending",
		eCfgAscending,
		unsafe.Sizeof(eCfgAscending),
		unsafe.Alignof(eCfgAscending),
	)

	fmt.Println()
	fmt.Println()

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.recipientAddresses",
		eCfgDescending.recipientAddresses,
		unsafe.Sizeof(eCfgDescending.recipientAddresses),
		unsafe.Alignof(eCfgDescending.recipientAddresses),
		unsafe.Offsetof(eCfgDescending.recipientAddresses),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.server",
		eCfgDescending.server,
		unsafe.Sizeof(eCfgDescending.server),
		unsafe.Alignof(eCfgDescending.server),
		unsafe.Offsetof(eCfgDescending.server),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.senderAddress",
		eCfgDescending.senderAddress,
		unsafe.Sizeof(eCfgDescending.senderAddress),
		unsafe.Alignof(eCfgDescending.senderAddress),
		unsafe.Offsetof(eCfgDescending.senderAddress),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.clientIdentity",
		eCfgDescending.clientIdentity,
		unsafe.Sizeof(eCfgDescending.clientIdentity),
		unsafe.Alignof(eCfgDescending.clientIdentity),
		unsafe.Offsetof(eCfgDescending.clientIdentity),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.timeout",
		eCfgDescending.timeout,
		unsafe.Sizeof(eCfgDescending.timeout),
		unsafe.Alignof(eCfgDescending.timeout),
		unsafe.Offsetof(eCfgDescending.timeout),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.notificationRateLimit",
		eCfgDescending.notificationRateLimit,
		unsafe.Sizeof(eCfgDescending.notificationRateLimit),
		unsafe.Alignof(eCfgDescending.notificationRateLimit),
		unsafe.Offsetof(eCfgDescending.notificationRateLimit),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.template",
		eCfgDescending.template,
		unsafe.Sizeof(eCfgDescending.template),
		unsafe.Alignof(eCfgDescending.template),
		unsafe.Offsetof(eCfgDescending.template),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.serverPort",
		eCfgDescending.serverPort,
		unsafe.Sizeof(eCfgDescending.serverPort),
		unsafe.Alignof(eCfgDescending.serverPort),
		unsafe.Offsetof(eCfgDescending.serverPort),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.notificationRetries",
		eCfgDescending.notificationRetries,
		unsafe.Sizeof(eCfgDescending.notificationRetries),
		unsafe.Alignof(eCfgDescending.notificationRetries),
		unsafe.Offsetof(eCfgDescending.notificationRetries),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d and Offsetof %d\n",
		"eCfgDescending.notificationRetryDelay",
		eCfgDescending.notificationRetryDelay,
		unsafe.Sizeof(eCfgDescending.notificationRetryDelay),
		unsafe.Alignof(eCfgDescending.notificationRetryDelay),
		unsafe.Offsetof(eCfgDescending.notificationRetryDelay),
	)

	fmt.Printf(
		"var %q of type %T has size %d with Asignof %d\n",
		"eCfgDescending",
		eCfgDescending,
		unsafe.Sizeof(eCfgDescending),
		unsafe.Alignof(eCfgDescending),
	)

}
