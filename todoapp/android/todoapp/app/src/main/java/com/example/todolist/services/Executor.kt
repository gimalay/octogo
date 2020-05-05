package com.example.todolist.services

import com.example.todolist.data.newUUIDasByteString
import com.example.todolist.di.ServiceLocator
import com.example.todolist.model.CommandOuterClass.Command
import com.google.protobuf.ByteString

class Executor {
    private val binding = ServiceLocator.goBinding

    fun addHomeProject(projectName: String) {
        val newProject = Command.NewProject
            .newBuilder()
            .setName(projectName)
            .setProjectID(newUUIDasByteString())
            .build()

        binding.execute(newProject)
    }

    fun removeHomeProject(projectId: ByteString) {
        val removableProject = Command.DeleteProject
            .newBuilder()
            .setProjectID(projectId)
            .build()

        binding.execute(removableProject)
    }

}
