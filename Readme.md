2024/05/10
K:\videos\go\Exploring Go Modules by Mike Van Sickle - Pluralsight Jul 2019


Go 1.12



modules ecosystem to replace old workspaces

Standard workflow

Managing dependency versions
  

Advanced scenarios



mkdir exploring-go-modules
cd exploring-go-modules
go mod init github.com/yonyu/go-modules
code .
create main.go
go build .
./go-modules

-- import gorilla routing system
go get -u github/gorilla/mux v1.7.3

-- notice go.sum file