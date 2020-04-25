package com.example.todolist.services

import com.example.todolist.data.newUUIDasByteString
import com.example.todolist.model.CommandOuterClass.Command
import com.google.protobuf.ByteString

object Executor {
    object Home {
        fun addProject(projectName: String) {
            val newProject = Command.NewProject
                .newBuilder()
                .setName(projectName)
                .setProjectID(newUUIDasByteString())
                .build()

            GoBinding.execute(newProject)
        }

        fun removeProject(projectId: ByteString) {
            val removableProject = Command.DeleteProject
                .newBuilder()
                .setProjectID(projectId)
                .build()

            GoBinding.execute(removableProject)
        }
    }

    object Project {
        // ...
    }
}
