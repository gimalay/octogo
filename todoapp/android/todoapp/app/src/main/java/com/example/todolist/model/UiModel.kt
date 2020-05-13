package com.example.todolist.model

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel

@Model
class UiModel {
    var isLoading: Boolean = false
    var home: ViewModel.Home = ViewModel.Home.getDefaultInstance()
    var project: ViewModel.Project = ViewModel.Project.getDefaultInstance()
    // ...
}
