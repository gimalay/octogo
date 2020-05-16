package com.example.todolist.repositories

import com.example.todolist.model.UiModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.services.Loader
import com.google.protobuf.ByteString

class ProjectRepository(
    private val loader: Loader,
    private val ui: UiModel
) {

    fun loadProject(id: ByteString) {
        val payload = Location.Project
            .newBuilder()
            .setProjectID(id)
            .build()

        loader.load(payload) { result ->
            ui.project = ViewModel.Project.parseFrom(result)
        }
    }
}