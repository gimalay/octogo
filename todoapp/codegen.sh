cd "$GOPATH"/src || exit
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/model/project.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/model/task.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/model/event.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/app/command.proto
protoc --go_out=plugins=grpc:. -I "$GOPATH"/src github.com/gimalay/octogo/todoapp/app/viewModel.proto

cd "$GOPATH"/src/github.com/gimalay/octogo/todoapp/app || exit
protoc --swift_out=../ios/todoapp/src/Models command.proto
protoc --swift_out=../ios/todoapp/src/Models viewModel.proto

protoc --java_out=../android/todoapp/app/src/main/java viewModel.proto
protoc --java_out=../android/todoapp/app/src/main/java command.proto