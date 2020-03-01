cd "$GOPATH"/src || exit
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/model/project.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/model/task.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/model/event.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/binding/command.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/binding/viewModel.proto

cd "$GOPATH"/src/github.com/gimalay/octogo/todoapp/binding || exit
protoc --swift_out=../ios/todoapp/src/Models command.proto
protoc --swift_out=../ios/todoapp/src/Models viewModel.proto

