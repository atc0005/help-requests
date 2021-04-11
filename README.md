<!-- omit in toc -->
# help-requests

Repo used to showcase content for help requests and bug reports

<!-- omit in toc -->
## Table of contents

- [Go](#go)
  - [Struct field alignment](#struct-field-alignment)

## Go

### Struct field alignment

This example [here](go/struct-alignment/main.go) is intended to show how
current versions of Go (1.15.11, 1.16.3) report `fieldalignment` linting
errors for structs which appear (perhaps due to my ignorance) to already be
aligned properly.

Structs:

```golang
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
```

Output:

```console
var "eCfgAscending.timeout" of type time.Duration has size 8 with Asignof 8 and Offsetof 0
var "eCfgAscending.notificationRateLimit" of type time.Duration has size 8 with Asignof 8 and Offsetof 8
var "eCfgAscending.template" of type *template.Template has size 8 with Asignof 8 and Offsetof 16
var "eCfgAscending.serverPort" of type int has size 8 with Asignof 8 and Offsetof 24
var "eCfgAscending.notificationRetries" of type int has size 8 with Asignof 8 and Offsetof 32
var "eCfgAscending.notificationRetryDelay" of type int has size 8 with Asignof 8 and Offsetof 40
var "eCfgAscending.server" of type string has size 16 with Asignof 8 and Offsetof 48
var "eCfgAscending.senderAddress" of type string has size 16 with Asignof 8 and Offsetof 64
var "eCfgAscending.clientIdentity" of type string has size 16 with Asignof 8 and Offsetof 80
var "eCfgAscending.recipientAddresses" of type []string has size 24 with Asignof 8 and Offsetof 96
var "eCfgAscending" of type main.emailConfigAscending has size 120 with Asignof 8


var "eCfgDescending.recipientAddresses" of type []string has size 24 with Asignof 8 and Offsetof 0
var "eCfgDescending.server" of type string has size 16 with Asignof 8 and Offsetof 24
var "eCfgDescending.senderAddress" of type string has size 16 with Asignof 8 and Offsetof 40
var "eCfgDescending.clientIdentity" of type string has size 16 with Asignof 8 and Offsetof 56
var "eCfgDescending.timeout" of type time.Duration has size 8 with Asignof 8 and Offsetof 72
var "eCfgDescending.notificationRateLimit" of type time.Duration has size 8 with Asignof 8 and Offsetof 80
var "eCfgDescending.template" of type *template.Template has size 8 with Asignof 8 and Offsetof 88
var "eCfgDescending.serverPort" of type int has size 8 with Asignof 8 and Offsetof 96
var "eCfgDescending.notificationRetries" of type int has size 8 with Asignof 8 and Offsetof 104
var "eCfgDescending.notificationRetryDelay" of type int has size 8 with Asignof 8 and Offsetof 112
var "eCfgDescending" of type main.emailConfigDescending has size 120 with Asignof 8
```
