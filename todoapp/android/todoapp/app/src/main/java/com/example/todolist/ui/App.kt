package com.example.todolist.ui

import androidx.compose.*
import com.example.todolist.data.SourceData
import com.example.todolist.services.Executor
import com.example.todolist.services.Loader

@Composable
fun App() {
    val sourceData = +state { SourceData() }
    +onActive {
        sourceData.value.home = Loader.Home.get()
    }

    HomeActivity(
        projects = sourceData.value.home.projectsList,
        onCopyProject = {
            Executor.Home.addProject(it)
            sourceData.value.home = Loader.Home.get()
        },
        onRemoveProject = {
            Executor.Home.removeProject(it)
            sourceData.value.home = Loader.Home.get()
        }
    )
}
