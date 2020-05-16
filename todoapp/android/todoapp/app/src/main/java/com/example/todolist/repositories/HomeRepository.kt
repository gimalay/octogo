package com.example.todolist.repositories

import com.example.todolist.model.UiModel
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.services.Loader

class HomeRepository(
    private val loader: Loader,
    private val ui: UiModel
) {
    fun loadHome() {
        val payload = Location.Home.newBuilder().build()
        loader.load(payload) { result ->
            ui.home = ViewModel.Home.parseFrom(result)
        }
    }

}
