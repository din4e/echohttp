# echohttp

A minimal HTTP request echo server. Outputs all HTTP request information to console.

一个极简的 HTTP 请求回显服务。将所有 HTTP 请求信息输出到控制台。

## Usage / 使用方法

```bash
# Windows
echohttp.exe -p 8080

# Linux / macOS
./echohttp -p 8080
```

### Options / 参数选项

| Flag | Default | Description |
|------|---------|-------------|
| `-p` | `8080` | Listen port / 监听端口 |

## Example / 示例

**Request / 请求:**
```bash
curl -X POST "http://localhost:8080/api/users?id=123" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer token123" \
  -d '{"name":"test"}'
```

**Output / 输出:**
```
====== HTTP Request ======
POST /api/users?id=123 HTTP/1.1
Host: localhost:8080
RemoteAddr: [::1]:54725
--- Headers ---
User-Agent: [curl/8.8.0]
Accept: [*/*]
Content-Type: [application/json]
Authorization: [Bearer token123]
Content-Length: [15]
--- Body ---
{"name":"test"}
==========================
```

## Build / 构建

```bash
# Windows
go build -ldflags="-s -w" -o echohttp.exe .

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o echohttp .
```

## License / 许可证

MIT
