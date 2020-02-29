import SwiftUI

struct HomeView: View {
    @EnvironmentObject var commander: Commander
    @EnvironmentObject var data: UserData

    @State var showNewIdentityForm: Bool = false

    var body: some View {
        List {
            ForEach(data.home.projects) { p in
                ProjectView(project: p)
            }

            Section() {
                HStack {
                    Text("New Project")
                    Spacer()
                    Button(action: { self.showNewIdentityForm.toggle() }) {
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
