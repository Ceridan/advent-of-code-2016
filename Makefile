##############################################################################
# Set up
##############################################################################

.PHONY: direnv
direnv: # Set environment varaibales
	direnv allow .

.PHONY: gen
gen: # Generate next day
	@go run ./gen.go --day=$(day)

##############################################################################
# Development
##############################################################################
.PHONY: format
format: # Format code
	@go fmt ./...

.PHONY: test
test: # Test day XX
	@go test ./days/$(day)/day$(day).go ./days/$(day)/day$(day)_test.go

.PHONY: exec
exec: # Run day XX
	@go run ./days/$(day)/day$(day).go
