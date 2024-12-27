##############################################################################
# Set up
##############################################################################

.PHONY: direnv
direnv: # Set environment varaibales
	direnv allow .

##############################################################################
# Development
##############################################################################
.PHONY: format
format: # Format code
	@go fmt ./...

.PHONY: test
test: # Test day XX
	@go test ./$(day)/day$(day).go

.PHONY: exec
exec: # Run day XX
	@go run ./$(day)/day$(day).go

.PHONY: gen
gen: # Generate next day
	@go run ./gen.go --day=$(day)
