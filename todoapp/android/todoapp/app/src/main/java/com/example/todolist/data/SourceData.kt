package com.example.todolist.data

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel

@Model
class SourceData {
    var home: ViewModel.Home = ViewModel.Home.getDefaultInstance()
    var project: ViewModel.Project = ViewModel.Project.getDefaultInstance()
    // ...
}
