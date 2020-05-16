package com.example.todolist.ui.home

import androidx.compose.*
import com.example.todolist.di.AppContainer

@Composable
fun HomeScreen(appContainer: AppContainer) {
    val repo = appContainer.homeRepository
    val commander = appContainer.homeCommander
    val navigator = appContainer.navigator
    val ui = appContainer.ui

    +onActive { repo.loadHome() }

    Home(
        projects = ui.home.projectsList,
        onAddProject = { projectName ->
            commander.addProject(projectName)
        },
        onRemoveProject = { id ->
            commander.removeProject(id)
        },
        onNavigateTo = { screen ->
            navigator.navigateTo(screen)
        }
    )
}