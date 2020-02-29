import Foundation
import SwiftProtobuf
import Binding

final class GoBinding {
    var binding: BindingBindingProtocol?

    private init() {
        open()
    }

    static let shared: GoBinding = {
        let instance = GoBinding()
        return instance
    }()

    func open() {
        let error: ErrorPointer = nil
        let b = BindingNew(GetDatabasePath(), error)
        self.binding = b
        if (error != nil) {
            print(error ?? "no error")
        }
    }

    func executeCommand(_ payload: SwiftProtobuf.Message, timestamp: Date) {
        do {
            let pd = try payload.serializedData()
            let type = self.getCommandType(payload)
            if type == CommandType.unknown {
                let s = try payload.jsonString()
                print("Unknown command type for: " + s)
            }

            let c = Command.with({
                $0.type = self.getCommandType(payload)
                $0.timestamp = Google_Protobuf_Timestamp(date: timestamp)
                $0.payload = pd
            })

            let cd = try c.serializedData()
            try self.binding?.execute(cd)
        } catch {
            print("Unexpected error: \(error).")
        }
    }

    func getCommandType(_ payload: SwiftProtobuf.Message) -> CommandType {
        switch payload {
        case is Command.NewTask:
            return CommandType.newTask
        case is Command.NewProject:
            return CommandType.newProject
        case is Command.RemoveTask:
            return CommandType.removeTask
        case is Command.AddTask:
            return CommandType.addTask
        case is Command.RenameProject:
            return CommandType.renameProject
        case is Command.RenameTask:
            return CommandType.renameTask
        default:
            return CommandType.unknown
        }
    }

    func getViewModel(location: Location) -> Data {
        do {
            let cd = try location.serializedData()
            let data = try binding!.viewModel(cd)
            return data
        } catch {

            print("Unexpected error: \(error).")
        }

        return Data()
    }

    func close() {
        do {
            try binding!.close()
        } catch {
            print("unexpected error: \(error).")
        }
    }
}

func GetDatabasePath() -> String {
    let path = NSSearchPathForDirectoriesInDomains(
            .documentDirectory, .userDomainMask, true
    ).first!

    let fileName = "\(path)/boltdb";

    return fileName
}

