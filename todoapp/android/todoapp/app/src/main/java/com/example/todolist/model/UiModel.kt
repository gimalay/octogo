package com.example.todolist.model

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel

@Model
class UiModel {
    var isLoading: Boolean = false

    var home: ViewModel.Home = ViewModel.Home.getDefaultInstance()
    var homeFilter: HomeFilter = HomeFilter.default
    var homeSorter: HomeSorter = HomeSorter.default

    var project: ViewModel.Project = ViewModel.Project.getDefaultInstance()

    // ...
}
