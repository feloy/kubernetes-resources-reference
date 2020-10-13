---
title: "HTTPGetAction"
description: "HTTPGetAction describes an action based on HTTP Get requests."
draft: false
collapsible: false
weight: 4
---
## HTTPGetAction
`import "k8s.io/api/core/v1"`
### HTTPGetAction
HTTPGetAction describes an action based on HTTP Get requests.
- **port** (IntOrString), required
  Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
*IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.*
- **host** (string)
  Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **httpHeaders** ([]HTTPHeader)
  Custom headers to set in the request. HTTP allows repeated headers.
*HTTPHeader describes a custom header to be used in HTTP probes*
  - **httpHeaders.name** (string), required
    The header field name
  - **httpHeaders.value** (string), required
    The header field value
- **path** (string)
  Path to access on the HTTP server.
- **scheme** (string)
  Scheme to use for connecting to the host. Defaults to HTTP.
