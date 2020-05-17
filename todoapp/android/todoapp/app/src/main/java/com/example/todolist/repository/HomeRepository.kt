package com.example.todolist.repository

import com.example.todolist.model.HomeFilter
import com.example.todolist.model.HomeSorter
import com.example.todolist.model.UiModel
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.service.Loader

class HomeRepository(
    private val loader: Loader,
    private val ui: UiModel
) {
    private val homeProjectSortingComparator = { it: ViewModel.Home.Project ->
        val homeSorter = ui.homeSorter
        when (homeSorter.field) {
            ViewModel.Home.Project.ID_FIELD_NUMBER -> it.id.toString()
            ViewModel.Home.Project.NAME_FIELD_NUMBER -> it.name
            else -> null
        }
    }

    fun loadHome() {
        // TODO: Implement the loading with applied filter on GoLang side
        // TODO: Implement the loading with applied sorter on GoLang side
        val homeFilter = ui.homeFilter
        val homeSorter = ui.homeSorter

        val payload = Location.Home.newBuilder().build()
        loader.load(payload) { result ->
            val home = ViewModel.Home.parseFrom(result)
            val filteredProjects = home.projectsList.filter {
                it.id.toString().contains(homeFilter.projectId) &&
                it.name.contains(homeFilter.projectName)
            }

            val sortedProjects = if (homeSorter.isDesc) {
                filteredProjects.sortedByDescending(homeProjectSortingComparator)
            }
            else {
                filteredProjects.sortedBy(homeProjectSortingComparator)
            }

            ui.home = home
                .toBuilder()
                .clearProjects()
                .addAllProjects(sortedProjects)
                .build()
        }
    }

    fun applyHomeFilter(f: HomeFilter) {
        ui.homeFilter = f
        loadHome()
    }

    fun applyHomeSorter(s: HomeSorter) {
        ui.homeSorter = s
        loadHome()
    }
}
