# gocritic-118

This repro was tested with go1.18.1

This should panic:

```bash
./golangci-lint run -v ./...
```

This shouldn't panic:

```bash
./golangci-lint run -v --skip-files=slice/slice.go ./...
```

This shouldn't panic either:

```bash
./golangci-lint run -v --skip-dirs=slice ./...
```