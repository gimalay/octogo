import SwiftUI

struct ContentView: View {
    @EnvironmentObject var userData: Commander

    var body: some View {
        NavigationView {
            HomeView()
        }
    }
}

#if DEBUG
struct ContentView_Previews: PreviewProvider {
    static var previews: some View {

        let data = Commander.shared

        return ContentView().environmentObject(data)
    }
}
#endif
