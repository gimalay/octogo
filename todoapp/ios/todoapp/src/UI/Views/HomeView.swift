import SwiftUI

struct HomeView: View {
    @EnvironmentObject var commander: Commander
    @EnvironmentObject var data: UserData

    var body: some View {
        List {
            ForEach(data.home.projects) { p in
                ProjectView(project: p).onTapGesture(perform: { self.deleteProject(id: p.id) })
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
                .onAppear(perform: { self.data.navigateTo(Location.Home()) })
    }

    func deleteProject(id: Data) {
        commander.executeCommand(Command.DeleteProject.with({
            $0.projectID = id
        }))
    }

    func newProject() {
        commander.executeCommand(Command.NewProject.with({
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
                HomeView()
                        .previewDevice(PreviewDevice(rawValue: deviceName))
                        .previewDisplayName(deviceName)
                        .environmentObject(Commander.shared)
                        .environmentObject(UserData.shared)
            }
        }
    }
}
#endif
