# Generating Swagger Documentation

This tutorial shows how to install Swag and generate Swagger documentation for your Go APIs.

## Steps to Install and Use Swag

### 1. Install Swag

Install Swag using the `go install` command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2. Creating the example/example.go File

Access the `docs/` directory, create a folder for your package inside it, in this case, `example/`, and then create the `example.go` file inside it.

### 3. Example Code with Swagger Comments

Here is an example with Swagger comments for an `Example` function:

```go
package docs

import (
	_ "github.com/flambra/account/internal/domain"
	_ "github.com/flambra/helpers/hResp"
)

// Example godoc
//
//	@Summary		Example function
//	@Description	Example endpoint with provided details
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			example	body		hResp.DefaultResponse	true	"Example Request"
//	@Success		200		{object}	hResp.DefaultResponse
//	@Failure		400		{object}	hResp.DefaultResponse
//	@Failure		500		{object}	hResp.DefaultResponse
//	@Router			/example [post]
func Example() {}
```

### 4. Run `swag init`

After adding Swagger comments, run the `swag init` command again to generate updated documentation:

```bash
cd docs
```

```bash
swag fmt
```

```bash
swag init -g example/example.go -o ./ --parseDependency true
```