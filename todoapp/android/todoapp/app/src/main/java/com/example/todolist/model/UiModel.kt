package com.example.todolist.model

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location

@Model
class UiModel {
    var isLoading: Boolean = false

    var home: ViewModel.Home = ViewModel.Home.getDefaultInstance()
    var homeFilter: Location.Home.Filter = Location.Home.Filter.getDefaultInstance()
    var homeSorter: Location.Home.Sorter = Location.Home.Sorter.getDefaultInstance()

    var project: ViewModel.Project = ViewModel.Project.getDefaultInstance()

    // ...
}
