package com.example.todolist.repository

import com.example.todolist.model.HomeFilter
import com.example.todolist.model.UiModel
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.service.Loader

class HomeRepository(
    private val loader: Loader,
    private val ui: UiModel
) {
    fun loadHome() {
        // TODO: Implement the loading with applied filter on GoLang side
        val payload = Location.Home.newBuilder().build()
        loader.load(payload) { result ->
            val home = ViewModel.Home.parseFrom(result)
            val filteredProjects = home.projectsList.filter {
                it.id.toString().contains(ui.homeFilter.projectId) &&
                it.name.contains(ui.homeFilter.projectName)
            }
            ui.home = home
                .toBuilder()
                .clearProjects()
                .addAllProjects(filteredProjects)
                .build()
        }
    }

    fun applyHomeFilter(f: HomeFilter) {
        ui.homeFilter = f
        loadHome()
    }
}
