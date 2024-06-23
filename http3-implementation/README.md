# Basic HTTP/3 Implementation on top of QUIC (In Progress Learning Project)

* HTTP3 Spec: https://datatracker.ietf.org/doc/html/rfc9114
* QUIC Spec: https://www.rfc-editor.org/rfc/rfc9000.html
* QUIC implementation in go std lib: https://github.com/golang/go/issues/58547
* HTTP3 implemntation in go std lib: https://github.com/golang/go/issues/32204

## Features

* stream multiplexing
* per-stream flow control
* low-latency connection establishment.
* QUIC (Level 4/Transport Layer): Protocol for faster (than TCP) connection oriented communication

## QUIC

### Features/Overview

* uses TLS
* unit of data transfer: QUIC packets which contain frames with control information and application data
* reliable delivery and congestion control

### Streams

* can be unidirectional or bidirectional
* supports multiplexing
* can be created by either endpoint
* each stream has a stream id per connection
* frames can be ordered using a stream id and offset

## Implementation

- [ ] Implement a basic QUIC 1RTT Packet
- [ ] Implement Stream Connection
