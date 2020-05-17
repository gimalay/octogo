package com.example.todolist.commander

import com.example.todolist.model.CommandOuterClass
import com.example.todolist.model.UiModel
import com.example.todolist.repository.HomeRepository
import com.example.todolist.service.Commander
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
    }

    fun removeProject(projectId: ByteString) {
        val removableProject = CommandOuterClass.Command.DeleteProject
            .newBuilder()
            .setProjectID(projectId)
            .build()

        commander.execute(removableProject) { repository.loadHome() }
    }

}