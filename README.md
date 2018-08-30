# Fasthttp simple reverse proxy
This is just for testing and trying to debug https://github.com/valyala/fasthttp/issues/348

## Run the proxy and an http test server
```bash
docker-compose up -d
```

By default, the port where the proxy listens is :8000.
If you need to change the ports, change them in the `docker-compose.yml`

```bash
curl -v http://127.0.0.1:8000/GET/200
```

The request will to through the proxy and hit the `http_test_server` and return
the response.