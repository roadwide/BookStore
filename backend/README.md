## Golang Standalone Web Service Template


**Generate ent files**

```bash
go generate ./ent
```

**Environment Variables**
```bash
export SQLITE_FILE=data.s3db
export JWT_SECRET_KEY=MxDNSc0AnSSoU7WUUWh9i
go run main.go
```

**Build and Serve**

```bash
docker build -t some-web-service:v1.0 .
docker run -d --name some-name some-web-service:v1.0
```
