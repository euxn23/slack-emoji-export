.PHONY: x-build
x-build:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
