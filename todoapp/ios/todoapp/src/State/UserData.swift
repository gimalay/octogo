import Combine
import SwiftUI
import SwiftProtobuf

final class UserData: ObservableObject {
    @Published var project: ViewModel.Project = ViewModel.Project()
    @Published var task: ViewModel.Task = ViewModel.Task()
    @Published var home: ViewModel.Home = ViewModel.Home()

    var location: Location = Location()

    private init() {
        navigateTo(Location.Home())
        update()
    }

    static let shared: UserData = {
        let instance = UserData()
        return instance
    }()

    private let serialQueue = DispatchQueue(label: "state.queue")

    func getLocationType(_ payload: SwiftProtobuf.Message) -> LocationType {
        switch payload {
        case is Location.Task:
            return LocationType.task
        case is Location.Home:
            return LocationType.home
        case is Location.Project:
            return LocationType.project
        default:
            return LocationType.unknown
        }
    }

    func navigateTo(_ payload: SwiftProtobuf.Message) {
        do {
            let pd = try payload.serializedData()
            location = Location.with({
                $0.type = getLocationType(payload)
                $0.payload = pd
            })

            update()

        } catch {
            print("Unexpected error")
        }
    }

    func commandExecuted() {
        update()
    }

    private func update() {
        serialQueue.async {
            let loc = self.location
            let data = GoBinding.shared.getViewModel(location: loc)

            DispatchQueue.main.async {
                self.publish(location: loc, data: data)
            }
        }
    }

    private func publish(location: Location, data: Data) {
        do {
            switch location.type {
            case LocationType.home:
                home = try ViewModel.Home(serializedData: data)
            case .project:
                project = try ViewModel.Project(serializedData: data)
            case .task:
                task = try ViewModel.Task(serializedData: data)
            case .unknown:
                return
            case .UNRECOGNIZED(_):
                return
            }
        } catch {
            print(error)
        }
    }

}
