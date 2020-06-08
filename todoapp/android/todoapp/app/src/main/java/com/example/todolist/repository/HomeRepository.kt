package com.example.todolist.repository

import com.example.todolist.model.UiModel
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.service.Loader

class HomeRepository(
    private val loader: Loader,
    private val ui: UiModel
) {
    fun loadHome() {
        val payload = Location.Home
            .newBuilder()
            .setFilter(ui.homeFilter)
            .setSorter(ui.homeSorter)
            .build()

        loader.load(payload) { result ->
            ui.home = ViewModel.Home.parseFrom(result)
        }
    }

    fun applyHomeFilter(f: Location.Home.Filter) {
        ui.homeFilter = f
        loadHome()
    }

    fun applyHomeSorter(s: Location.Home.Sorter) {
        ui.homeSorter = s
        loadHome()
    }
}
