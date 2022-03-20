# TCP server with golang

## Understand server client server architect
HTTP Protocol: Request and Response

## HTTP request structure:

1. request line
Request-Line Format:
- **Method** SP **Request-URI** SP **HTTP-Version** CRLF
Ex:
- GET /path/to/file/index.html HTTP/1.1
2. headers
3. optional message body
---

## HTTP response Structure:

1. status line
Status-Line:
- **HTTP-Version** SP **Status-Code** SP **Reason-Phrase** CRLF
Ex:
- HTTP/1.1 200 OK
2. headers
3. optional message body

## small exercise
Open developer tool on Chrome. Use Option + ⌘ + J (on macOS)

---
What left to do?
1. Review TCP networking part