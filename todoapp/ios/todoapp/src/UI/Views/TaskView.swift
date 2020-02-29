import SwiftUI

struct TaskView: View {
    @EnvironmentObject var userData: Commander

    @GestureState private var didLongPress: Bool = false
    let notificationGenerator = UINotificationFeedbackGenerator()

    var task: ViewModel.Home.Project.Task

    var body: some View {
        HStack {
            Text(self.task.emoji).font(.largeTitle)
            Text(self.task.name).font(.body).truncationMode(.tail).lineLimit(1)

        }
    }
}

