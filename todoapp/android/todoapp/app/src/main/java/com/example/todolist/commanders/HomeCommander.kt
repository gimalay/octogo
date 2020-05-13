package com.example.todolist.commanders

import com.example.todolist.model.CommandOuterClass
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.repositories.HomeRepository
import com.example.todolist.services.Commander
import com.example.todolist.utils.newUUIDasByteString
import com.google.protobuf.ByteString

class HomeCommander(
    private val commander: Commander,
    private val repository: HomeRepository
) {
    fun addProject(projectName: String) {
        val newProject = CommandOuterClass.Command.NewProject
            .newBuilder()
            .setName(projectName)
            .setProjectID(newUUIDasByteString())
            .build()

        commander.execute(newProject) { uiModel ->
            val project = ViewModel.Home.Project.newBuilder().setName(projectName).build();
            uiModel.home = uiModel.home.toBuilder().addProjects(project).build()
        }
    }

    fun removeProject(projectId: ByteString) {
        val removableProject = CommandOuterClass.Command.DeleteProject
            .newBuilder()
            .setProjectID(projectId)
            .build()

        commander.execute(removableProject) { uiModel ->
            val projects = uiModel.home.projectsList
            val removableProjectIndex = projects.indexOf(projects.find { it.id == projectId })
            uiModel.home = uiModel.home.toBuilder().removeProjects(removableProjectIndex).build()
        }
    }
}