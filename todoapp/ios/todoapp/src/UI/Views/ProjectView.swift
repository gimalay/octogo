import SwiftUI

struct ProjectView: View {
    var project: ViewModel.Home.Project

    var body: some View {
        Section(header: Text(project.name)) {
            ForEach(project.tasks, id: \.id) { t in
                TaskView(task: t)

            }
        }
    }
}

