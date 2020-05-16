package com.example.todolist.commanders

import com.example.todolist.model.CommandOuterClass
import com.example.todolist.model.UiModel
import com.example.todolist.repositories.HomeRepository
import com.example.todolist.services.Commander
import com.example.todolist.utils.newUUIDasByteString
import com.google.protobuf.ByteString

class HomeCommander(
    private val commander: Commander,
    private val repository: HomeRepository,
    private val ui: UiModel
) {
    fun addProject(projectName: String) {
        val newProject = CommandOuterClass.Command.NewProject
            .newBuilder()
            .setName(projectName)
            .setProjectID(newUUIDasByteString())
            .build()

        commander.execute(newProject) { repository.loadHome() }
//        commander.execute(newProject) {
//            val project = ViewModel.Home.Project.newBuilder().setName(projectName).build();
//            ui.home = ui.home.toBuilder().addProjects(project).build()
//        }
    }

    fun removeProject(projectId: ByteString) {
        val removableProject = CommandOuterClass.Command.DeleteProject
            .newBuilder()
            .setProjectID(projectId)
            .build()

        commander.execute(removableProject) { repository.loadHome() }
//        commander.execute(removableProject) {
//            val projects = ui.home.projectsList
//            val removableProjectIndex = projects.indexOf(projects.find { it.id == projectId })
//            ui.home = ui.home.toBuilder().removeProjects(removableProjectIndex).build()
//        }
    }
}