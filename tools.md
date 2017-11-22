

# Present

-base="": base path for slide template and static resources
-http="127.0.0.1:3999": HTTP service address (e.g., '127.0.0.1:3999')
-nacl=false: use Native Client environment playground (prevents non-Go code execution)
-notes=false: enable presenter notes (press 'N' from the browser to display them)
-orighost="": host component of web origin URL (e.g., 'localhost')
-play=true: enable playground (permit execution of arbitrary user code)

```
present -http="127.0.0.1:4000"
```

# Godoc

```
godoc -http="127.0.0.1:6060"
```