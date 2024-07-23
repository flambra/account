#!/bin/bash

# Define the base directory
BASE_DIR="./"
SWAG_PATH="$HOME/go/bin/swag"

# Format the Go code
$SWAG_PATH fmt

# Iterate over each subdirectory in the base directory
for dir in "$BASE_DIR"/*/; do
    # Get the directory name without the trailing slash
    dir_name=$(basename "$dir")
    # Define the Go file to use as the entry point
    go_file="$dir_name/$dir_name.go"
    # Check if the Go file exists
    if [[ -f "$go_file" ]]; then
        echo "Generating Swagger documentation for $go_file"
        $SWAG_PATH init -g "$go_file" -o ./ --parseDependency true
    else
        echo "No Go file found for $dir_name"
    fi
done

echo "Swagger documentation generation completed." 