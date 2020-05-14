import SwiftUI
import Octogo

struct HomeView: View {
    @EnvironmentObject var commander: Commander
    @EnvironmentObject var homeLoader: HomeLoader
    let model: ViewModel.Home

    var body: some View {
        List {
            ForEach(model.projects) { p in
                ProjectView(project: p).onTapGesture(perform: { })
            }

            Section() {
                HStack {
                    Text("New Project")
                    Spacer()
                    Button(action: newProject) {
                        Image(systemName: "plus.circle")
                    }
                }
            }
        }
                .navigationBarTitle(Text("Todo"), displayMode: .large)
                .listStyle(GroupedListStyle())
                .environment(\.horizontalSizeClass, .regular)
        .onAppear(perform: { self.homeLoader.load() })
    }

    func newProject() {
        commander.execute(Command.NewProject.with({
            $0.projectID = newId()
            $0.name = "New Project"
        }))

    }

}

func newId() -> Data {
    uuidToBytes(UUID())
}


func uuidToBytes(_ uuid: UUID) -> Data {
    withUnsafePointer(to: uuid) {
        Data(bytes: $0, count: MemoryLayout.size(ofValue: UUID().uuid))
    }
}

#if DEBUG
struct P986245_Previews: PreviewProvider {
    static var previews: some View {

        return ForEach(["iPhone SE", "iPhone XS Max"], id: \.self) { deviceName in
            NavigationView {
                HomeView(model: ViewModel.Home())
                        .previewDevice(PreviewDevice(rawValue: deviceName))
                        .previewDisplayName(deviceName)
                        .environmentObject(Commander.shared)
            }
        }
    }
}
#endif
