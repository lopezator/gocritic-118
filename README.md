# gocritic-118

This repro was tested with go1.18.1

If you use the provided binary (with the fix included), this panics:

```bash
./golangci-lint run -v ./...
```

But this doesn't:

```bash
./golangci-lint run -v --skip-files=slice/slice.go ./...
```

And this doesn't panic either:

```bash
./golangci-lint run -v --skip-dirs=slice ./...
```

You can test that with a regular `golanci-lint` binary, any of the three panics/errs.