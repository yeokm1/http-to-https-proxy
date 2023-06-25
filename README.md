# http-to-https-proxy
A proxy that upgrades HTTP connections to HTTPS for systems which cannot make HTTPS requests.

## Running the proxy

Download the latest binary corresponding to your platfrom from the [releases section](https://github.com/yeokm1/http-to-https-proxy/releases/).

### Default Configuration

```bash
http-to-https-proxy.exe
2023/06/25 20:38:37 HTTP to HTTPS proxy v0.3 listening to 80, forward to 443 with listening buffer 4096
2023/06/25 20:38:37 You can supply the listening port, forward port, buffer size, insecure -i cert as command line args
2023/06/25 20:38:41 Request from 192.168.1.112:13519 to host api.openai.com and url /v1/chat/completions
2023/06/25 20:38:43 EOF reached
2023/06/25 20:38:43 End of handler
...
```

By default, proxy will listen to HTTP requests on port `80` and retransmit HTTPS via port `443`. Buffer size of `4096` is the buffer to receive destination server's response chunks before forwarding back original client.

### Modify ports and buffer size

```bash
http-to-https-proxy.exe 90 445 5000
2023/06/25 20:39:20 HTTP to HTTPS proxy v0.3 listening to 90, forward to 445 with listening buffer 5000
2023/06/25 20:39:20 You can supply the listening port, forward port, buffer size, insecure -i cert as command line args
...
```
All 3 arguments must be specified even if you only wish to change one of the values.

```bash
http-to-https-proxy.exe -i
2023/06/25 20:39:48 HTTP to HTTPS proxy v0.3 listening to 80, forward to 443 with listening buffer 4096
2023/06/25 20:39:48 Allow insecure TLS certificates
2023/06/25 20:39:48 You can supply the listening port, forward port, buffer size, insecure -i cert as command line args
```
If the server you are connecting to is using expired/insecure TLS certificates. You can add `-i` argument to allow those connections.

# Compiling

Just install the latest Go compiler for your platform. The latest at the time of writing is `1.20.5`. THe following was compiled on windows/amd64 platform using Powershell script `build.ps1`.
