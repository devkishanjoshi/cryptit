
## Create a new module 
```
go mod init
go mod init <Module Path>
```



## Go mod file matches source code in module It adds missing modules as well as removes requirements on module dont provide any relevant packages
```
go mod tidy 
```

## Compiles & runs Program 
### Internally it creates a executable binary at temporary location & cleans when app exits. 
```
go run <File>
```
## compiles packages along with dependencies into an executable file
```
go build 
```
## compiles & move executable to $PATH/bin
```
go install
```

## Add dependency for package or upgrade/Downgrade it to its latest version  

```
go env GOPATH 

go get <packageName>
```


## Use a unpublished module to local Directory: 
```
go mod edit -replace github.com/devkishanjoshi/cryptit=../cryptit
```

## As Per comments it shows Documentation 

```
go doc 
```
