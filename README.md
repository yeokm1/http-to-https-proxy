# http-to-https-proxy
A proxy that upgrades HTTP connections to HTTPS for systems which cannot make HTTPS requests.

## Running the proxy

Download the latest binary corresponding to your platfrom from the [releases section](https://github.com/yeokm1/http-to-https-proxy/releases/).

### Default Configuration

```bash
./http-to-https-proxy 
2019/11/03 22:54:50 Starting HTTP to HTTPS proxy listening to 80, forward to 443 with listening buffer 5000
2019/11/03 22:54:50 You can supply the listening and forwarding port and buffer size as 3 command line arguments
2019/11/03 22:54:51 Received request to route to host ABC.com and url /api/endpoint
2019/11/03 22:54:52 End of handler
...
```

By default, proxy will listen to HTTP requests on port `80` and retransmit HTTPS via port `443`. Buffer size of `4096` is the buffer to receive destination server's response chunks before forwarding back original client.

### Modify ports and buffer size

```bash
./http-to-https-proxy 90 445 5000
2019/11/03 22:56:05 HTTP to HTTPS proxy v0.2 listening to 90, forward to 445 with listening buffer 5000
2019/11/03 22:56:05 You can supply the listening and forwarding port and buffer size as 3 command line arguments
...
```
All 3 arguments must be specified even if you only wish to change one of the values.

```bash
http-to-https-proxy.exe -i
2023/03/30 14:47:36 HTTP to HTTPS proxy v0.2 listening to 80, forward to 443 with listening buffer 4096
2023/03/30 14:47:36 Allow insecure TLS certificates
2023/03/30 14:47:36 You can supply the listening port, forward port, buffer size, insecure -i cert as command line args
```
If the server you are connecting to is using expired/insecure TLS certificates. You can add `-i` argument to allow those connections.

# Compiling

Just install the latest Go compiler for your platform. The latest at the time of writing is `1.20.5`. THe following was compiled on windows/amd64 platform using Powershell script `build.ps1`.
