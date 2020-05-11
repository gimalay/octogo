package com.example.todolist.ui.home

import androidx.compose.Composable
import androidx.compose.onActive
import androidx.compose.state
import androidx.compose.unaryPlus
import com.example.todolist.data.SourceData
import com.example.todolist.di.AppContainer

@Composable
fun HomeScreen(appContainer: AppContainer) {
    val loader = appContainer.loader
    val commander = appContainer.commander
    val navigator = appContainer.navigator

    val sourceData = +state { SourceData() }
    +onActive {
        sourceData.value.home = loader.getHome()
    }

    Home(
        projects = sourceData.value.home.projectsList,
        onAddProject = {
            commander.addHomeProject(it)
            sourceData.value.home = loader.getHome()
        },
        onRemoveProject = {
            commander.removeHomeProject(it)
            sourceData.value.home = loader.getHome()
        },
        onNavigateTo = { screen ->
            navigator.navigateTo(screen)
        }
    )
}