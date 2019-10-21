# Structify!

Take in Kubernetes objects as YAML or JSON and print out golang structs for use in operators

# Run Directly

```bash
go run main.go convert examples/job.yml
```

# Run Server

Start server:

```bash
go run main.go serve
```

Then open http://localhost:8080 to access the web interface
