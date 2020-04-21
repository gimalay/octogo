package com.example.todolist.model

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.data.newUUIDasByteString
import com.google.protobuf.ByteString

@Model
class HomeModelMock : HomeModelBase {
    override var data: ViewModel.Home = ViewModel.Home.getDefaultInstance()

    override fun load() {
        val project1 = ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName("Test project 1")
            .build()

        val project2 = ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName("Test project 2")
            .build()

        this.data = ViewModel.Home
            .newBuilder()
            .addProjects(project1)
            .addProjects(project2)
            .build()
    }

    override fun addProject(projectName: String) {
        val project = ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName(projectName)
            .build()

        this.data = this.data
            .toBuilder()
            .addProjects(project)
            .build()
    }

    override fun removeProject(projectId: ByteString) {
        val removableProject = this.data.projectsList.find { it.id == projectId }
        val removableProjectIndex = this.data.projectsList.indexOf(removableProject)

        this.data = this.data
            .toBuilder()
            .removeProjects(removableProjectIndex)
            .build()
    }

}
