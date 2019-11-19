# cde-golang

Step 1: Create a desired folder structure
In this example, i have created two different folders customer and employee

Step 2: Run init mod
- Run the command `go mod init` it will return with an error should have a package name.
- Rerun the go mod init command with the package name at the end `github.com/haridhanakoti/golangPackage`
- Final command looks like `go mod init github.com/haridhanakoti/golangPackage`

Step 3: 
- New file named `go.mod` will be created on executing the above command

Step 4:
- Use the import statement to import the custom package like `emp "github.com/cdeGolang/golangPackage"`