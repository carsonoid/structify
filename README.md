# Structify!

Take in Kubernetes objects as YAML or JSON and print out golang structs for use in operators

## RoadMap

* [ ] Fix broken pointers to primitves in structs. Ex: `.spec.replicas` 
* [ ] Support Custom Resources when given a git repo and path to the types package
* [ ] Autodetect custom resources when they follow a known pattern
* [ ] Custom Domain + TLS
* [ ] Handle list objects and multiple yaml documents in the input

## Run Directly

```bash
go run main.go convert examples/job.yml
```

## Run Server

Start server:

```bash
go run main.go serve
```

Then open http://localhost:8080 to access the web interface
