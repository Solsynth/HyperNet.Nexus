# HyperNet.Nexus

Nexus is the core component of HyperNet,
which managed the resources allocation (such as the database allocate and more)
and the service gateway.

## Features

### Gateway

Gateway is one of the core features in Nexus, every HyperNet services will establish a connection to the Nexus
when they start. And it is required, if Nexus is died or it did not connected it successfully, the service will exit.

During the registration, the service will provide follow details:

1. An unique service ID - `id`
2. The human readable name of the service - `label`
3. The code for accessing the service - `type`
4. The grpc outbound address of the service - `grpc_addr`
5. The http outbound address of the service - `http_addr`

There is two kind of the gateway, the grpc one and the http one.

When the user requested the `/cgi/<type>`, the request was forwarded to the `http_addr`
of the service with load balancer included.

And when the internal service call `GetServiceGrpcConn(type)`, the Nexus will run a health check to ensure the
target service is alive and respond with a `grpc_addr`, the rest of the connection will be handled by the SDK.
It is also load balancer included.

### Allocation

The allocation is the additional service which make our developer experience better.
For simple, the allocation is basically a improved version of the shared configuration.
With some part of the configuration will be re-configured according to the service's demand.

Like the database allocation, service request with the database name they want,
and the Nexus will respond a configured connection string to the service, then the SDK will handle
the establishment of the connection.

At the same time, the allocated database will be added into the watchtower for auto maintenance
(auto remove the soft-deleted records, backup and more).

### Authorization

All the request forwarded by the Nexus will handle the authorization automatically.
The Authorization header field will be replaced by the internal one with full user data.

To implement the authorization, you must have at least one [Passport](https://github.com/Solsynth/HyperNet.Passport) instance alive.

For further usage, checkout the `hypernet/nexus/pkg/nex/sec` package.

## Installation

To run the Nexus, you need to have an etcd server and a nats server running.
The best way to run the Nexus is using docker compose. Here is an example with the required dependencies added
and health check configured:

```yaml
services:
  nexus:
    image: xsheep2010/nexus:nightly
    network_mode: "host"
    restart: unless-stopped
    depends_on:
      - nats
      - etcd
    volumes:
      - "/srv/hypernet/nexus.settings.toml:/settings.toml"
      - "/srv/hypernet/keys:/keys"
    healthcheck:
      test: ["CMD", "nc", "-z", "localhost", "7001"]
      interval: 30s
      timeout: 5s
      retries: 5
      start_period: 15s

  nats:
    image: nats:latest
    network_mode: "host"
    restart: unless-stopped

  etcd:
    image: bitnami/etcd:latest
    network_mode: "host"
    restart: unless-stopped
    environment:
      - ALLOW_NONE_AUTHENTICATION=yes
```

To add more services, you can according follow example:

```yaml
services:
  passport:
    image: xsheep2010/passport:nightly
    network_mode: "host"
    restart: unless-stopped
    depends_on:
      nexus:
        condition: service_healthy
    volumes:
      - "/srv/hypernet/passport.settings.toml:/settings.toml"
      - "/srv/hypernet/keys:/keys"
```
