# Image Signature Webhook

## Build
```
GOOS=linux go build .
docker build -t mkorejo/image-signature-webhook:0.1.0 .
docker push mkorejo/image-signature-webhook:0.1.0
```

## Usage

```
docker run mkorejo/image-signature-webhook:0.1.0 -h
Usage of /usr/local/bin/image-signature-webhook:
  -alsologtostderr
        log to standard error as well as files
  -grafeas string
        Grafeas URL (default "http://grafeas:8080")
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -tls-cert string
        TLS certificate (default "/etc/admission-controller/tls/cert.pem")
  -tls-key string
        TLS key (default "/etc/admission-controller/tls/key.pem")
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```
