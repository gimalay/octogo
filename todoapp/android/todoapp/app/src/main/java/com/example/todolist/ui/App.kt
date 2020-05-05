package com.example.todolist.ui

import androidx.compose.*
import com.example.todolist.data.SourceData
import com.example.todolist.services.Executor
import com.example.todolist.services.Loader
import com.example.todolist.ui.home.Home

@Composable
fun App(loader: Loader, executor: Executor) {
    val sourceData = +state { SourceData() }
    +onActive {
        sourceData.value.home = loader.getHome()
    }

    Home(
        projects = sourceData.value.home.projectsList,
        onAddProject = {
            executor.addHomeProject(it)
            sourceData.value.home = loader.getHome()
        },
        onRemoveProject = {
            executor.removeHomeProject(it)
            sourceData.value.home = loader.getHome()
        }
    )
}
