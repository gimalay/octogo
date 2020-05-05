package com.example.todolist.ui

import androidx.compose.*
import com.example.todolist.data.SourceData
import com.example.todolist.services.Executor
import com.example.todolist.services.Loader

@Composable
fun App(loader: Loader, executor: Executor) {
    val sourceData = +state { SourceData() }
    +onActive {
        sourceData.value.home = loader.getHome()
    }

    HomeActivity(
        projects = sourceData.value.home.projectsList,
        onCopyProject = {
            executor.addHomeProject(it)
            sourceData.value.home = loader.getHome()
        },
        onRemoveProject = {
            executor.removeHomeProject(it)
            sourceData.value.home = loader.getHome()
        }
    )
}
