Create a new module 
go mod init
go mod init <Module Path>


go mod tidy 
# Go mod file matches source code in module It adds missing modules as well as removes requirements on module dont provide any relevant packages


go run <File>

Compiles & runs Program 
Internally it creates a executable binary at temporary location & cleans when app exits. 

go build 
compiles packages along with dependencies into an executable file

go install
compiles & move executable to $PATH/bin

go env GOPATH 

go get <packageName>
Add dependency for package or upgrade/Downgrade it to its latest version  

Use a unpublished module to local Directory: 

go mod edit -replace github.com/devkishanjoshi/cryptit=../cryptit
