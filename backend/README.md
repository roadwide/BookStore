## Golang Standalone Web Service Template


**Generate ent files**

```bash
go generate ./ent
```

**Build and Serve**

```bash
docker build -t some-web-service:v1.0 .
docker run -d --name some-name some-web-service:v1.0
```
