import SwiftUI
import Octogo

struct ContentView: View {
    @EnvironmentObject var homeLoader: HomeLoader

    var body: some View {
        NavigationView {
            HomeView(model: homeLoader.model)
        }
    }
}

#if DEBUG
struct ContentView_Previews: PreviewProvider {
    static var previews: some View {


        return ContentView().environmentObject(HomeLoader(commandObserver: Commander.shared))
    }
}
#endif
