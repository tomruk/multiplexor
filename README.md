# Multiplexor

## Modus operandi

Multiplexor is a protocol multiplexer and TLS (SSL) tunnel which helps to run multiple applications on 1 port. It let's you choose which protocol to use and whether to use TLS or not. You can choose any protocol according to your needs and your fantasy world. The server listens for incoming connections and distinguishes them by their protocols and forwards them. The client is used for transparently upgrading raw connections to SSL connections.

## Supported protocols

#### Transport protocols

You can use any transport protocol supported by Go. Except you cannot use UDP when listening to connections.

| Listen     | Connect    | Explanation   |
| ---------- | ---------- | ------------- |
| tcp        | tcp        | TCP           |
| tcp4       | tcp4       | TCP with IPv4 |
| tcp6       | tcp6       | TCP with IPv6 |
| x          | udp        | UDP           |
| x          | udp4       | UDP with IPv4 |
| x          | udp6       | UDP with IPv6 |
| unix       | unix       | Unix sockets  |

#### Match protocols

- HTTP
- SSH
- Socks 4
- Socks 5
- OpenVPN

## Build

Alternatively, you can simply use `go build` instead of `make`

```shell
make
sudo make install
sudo mkdir /etc/multiplexor
sudo cp server_example.yml /etc/multiplexor/server.yml
```

## systemd service

To start multiplexor-server on boot, if you have systemd on your system, you can install the systemd script on `scripts/multiplexor.service`

By default, it runs multiplexor-server as root which is not a good practice. But you can change it by altering and uncommenting `User=nobody` and `Group=nogroup` lines in the script.

Also, if you want to bind multiplexor-server to ports lower than 1024, you can grant that permission with setcap:

```shell
sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/local/bin/multiplexor-server
```

To install systemd service, copy `scripts/multiplexor.service` to the systemd service directory:

```shell
sudo cp scripts/multiplexor.service /lib/systemd/system
sudo systemctl start multiplexor
sudo systemctl enable multiplexor
```

## Server configuration

There are several locations multiplexor-server looks for `server.yml`:

- Current directory
- Path of the executable
- /etc/multiplexor

You can either put the `server.yml` in one of these locations or you can specify the path by running multiplexor-server with `-c` argument.

### Editing configuration

There are 2 main fields in the configuration file. First one is `listen` and the latter one is `rules`. For example configuration, see [server_example.yml](server_example.yml)

#### Listen

Here is an example of a listener which listens on 127.0.0.1:1337 with TCP and uses TLS. You may change the `transport` field according to the [supported protocols](https://github.com/tomruk/multiplexor#supported-protocols) table.

```yaml
transport: tcp
address: 127.0.0.1:1337
tls:
  enabled: true
  cert: cert.pem
  key: key.pem
```

You can disable TLS by removing the `tls` field or simpler by setting `enabled` in `tls` field to false. Here is the same example which doesn't have a TLS configuration.

```yaml
transport: tcp
address: 127.0.0.1:1337
```

#### Rules

Rules are used to specify which protocols to match and forward the incoming connection to.

Here is an example rule that matches the SSH protocol. You can change the `transport` field according to the supported protocols table.

```yaml
transport: tcp
address: 127.0.0.1:22
match: ssh
```

Another example that matches HTTP and forwards the connection to 127.0.0.1:6666 and uses UDP. Maybe i'm fantastic enough to use UDP for an HTTP server. Why not?

```yaml
transport: udp
address: 127.0.0.1:6666
match: http
```

You can also use TLS for outgoing connections.

```yaml
transport: tcp
address: 127.0.0.1:6666
match: http
tls:
  enabled: true
  ignoreBadCert: true
```

## Self signed SSL certificate generation

To do so, you don't need the `openssl` tool. Use the `-g` option and it will ask you the details to create the certificate.

```multiplexor-server -g```

## Client

The client is for proxying local connections to the remote server. There are times you need to use TLS but the application you use doesn't have TLS support.

For example, let's assume multiplexer-server is running on 10.10.10.100:443 with TLS and you want to connect it via SSH. You may use multiplexor-client.

```shell
# The line below listens on 127.0.0.1:1337 and
# transparently forwards incoming connections to 10.10.10.100:443
multiplexor-client --listen 127.0.0.1:1337 --connect 10.10.10.100:443

# Then we can use SSH like this
ssh root@127.0.0.1 -p 1337
```