import UIKit
import SwiftUI
import Binding
import Octogo
import SwiftProtobuf

class BindingAdapter: Octogo.BindingAdapterProtocol {
    private var binding: BindingBindingProtocol?

    func execute(commandType: Int, payload: Data, timestamp: String) throws {
        try binding!.execute(commandType, payload: payload, timestamp: timestamp)
    }

    func viewModel(commandType: Int, payload: Data, timestamp: String) throws -> Data {
        try binding!.viewModel(commandType, payload: payload, timestamp: timestamp)
    }

    func open(dbPath: String) throws {
        let error: ErrorPointer = nil
        let b = BindingNew(dbPath, error)
        self.binding = b
        if (error != nil) {
            print(error ?? "no error")
        }
    }

    func close() throws {
        try binding!.close()
    }
}

extension Commander {
    static let shared: Commander = {
        let binding = GoBinding(binding: BindingAdapter(), dbPath: GetDatabasePath())
        binding.open()
        let instance = Commander(binding: binding)
        return instance
    }()
}

extension ViewModel.Home: ViewModelMessage {
}
extension Location.Home: LocationMessage {
}
extension ViewModel.Project: ViewModelMessage {
}
extension Location.Project: LocationMessage {
}
extension ViewModel.Task: ViewModelMessage {
}
extension Location.Task: LocationMessage {
}

extension Command.NewProject: CommandMessage {}
extension Command.AddTask: CommandMessage {}
extension Command.NewTask: CommandMessage {}
extension Command.RemoveTask: CommandMessage {}
extension Command.RenameProject: CommandMessage {}

class HomeLoader: Loader<ViewModel.Home, Location.Home> {
}

class TaskLoader: Loader<ViewModel.Task, Location.Task> {
}

class ProjectLoader: Loader<ViewModel.Project, Location.Project> {
}

class SceneDelegate: UIResponder, UIWindowSceneDelegate {

    var window: UIWindow?


    func scene(_ scene: UIScene, willConnectTo session: UISceneSession, options connectionOptions: UIScene.ConnectionOptions) {
        // Use this method to optionally configure and attach the UIWindow `window` to the provided UIWindowScene `scene`.
        // If using a storyboard, the `window` property will automatically be initialized and attached to the scene.
        // This delegate does not imply the connecting scene or session are new (see `application:configurationForConnectingSceneSession` instead).

        // Use a UIHostingController as window root view controller
        if let windowScene = scene as? UIWindowScene {
            let window = UIWindow(windowScene: windowScene)
            let commander = Commander.shared

            window.rootViewController = UIHostingController(rootView: ContentView()
                    .environmentObject(Commander.shared)
                .environmentObject(HomeLoader(commandObserver: commander))
                .environmentObject(TaskLoader(commandObserver: commander))
                .environmentObject(ProjectLoader(commandObserver: commander))
            )
            self.window = window
            window.makeKeyAndVisible()
        }

//        HealthKitManager.sharedInstance.requestPermissions()
//        HealthKitManager.sharedInstance.readWorkoutsData()
    }

    func sceneDidDisconnect(_ scene: UIScene) {
        // Called as the scene is being released by the system.
        // This occurs shortly after the scene enters the background, or when its session is discarded.
        // Release any resources associated with this scene that can be re-created the next time the scene connects.
        // The scene may re-connect later, as its session was not neccessarily discarded (see `application:didDiscardSceneSessions` instead).
    }

    func sceneDidBecomeActive(_ scene: UIScene) {
        // Called when the scene has moved from an inactive state to an active state.
        // Use this method to restart any tasks that were paused (or not yet started) when the scene was inactive.
    }

    func sceneWillResignActive(_ scene: UIScene) {
        // Called when the scene will move from an active state to an inactive state.
        // This may occur due to temporary interruptions (ex. an incoming phone call).
    }

    func sceneWillEnterForeground(_ scene: UIScene) {
        // Called as the scene transitions from the background to the foreground.
        // Use this method to undo the changes made on entering the background.
    }

    func sceneDidEnterBackground(_ scene: UIScene) {
        // Called as the scene transitions from the foreground to the background.
        // Use this method to save data, release shared resources, and store enough scene-specific state information
        // to restore the scene back to its current state.
    }


}

func GetDatabasePath() -> String {
    let path = NSSearchPathForDirectoriesInDomains(
            .documentDirectory, .userDomainMask, true
    ).first!

    let fileName = "\(path)/boltdb";

    return fileName
}


enum MessageMappingError: Error {
    case cannotMapMessageType
}

extension SwiftProtobuf.Message {
    func type() throws -> Int {
        let mirror = Mirror(reflecting: self)

        for a in LocationType.allCases {
            if String(describing: a).lowercased() == String(describing: mirror.subjectType).lowercased() {
                return a.rawValue
            }
        }
        for a in CommandType.allCases {
            if String(describing: a).lowercased() == String(describing: mirror.subjectType).lowercased() {
                return a.rawValue
            }
        }

        throw MessageMappingError.cannotMapMessageType
    }

    init(data: Data) throws {
        try self.init(serializedData: data)
    }

    func encode() throws -> Data {
        try self.serializedData()
    }
}
