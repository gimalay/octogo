import Combine
import SwiftUI
import SwiftProtobuf

final class Commander: ObservableObject {
    private let serialQueue = DispatchQueue(label: "state.queue")

    static let shared: Commander = {
        let instance = Commander()
        return instance
    }()

    private init() {
    }


    func executeCommand(_ payload: SwiftProtobuf.Message, timestamp: Date = Date()) {
        serialQueue.async {
            GoBinding.shared.executeCommand(payload, timestamp: timestamp)
            self.serialQueue.async {
                UserData.shared.commandExecuted()
            }
        }
    }
}

