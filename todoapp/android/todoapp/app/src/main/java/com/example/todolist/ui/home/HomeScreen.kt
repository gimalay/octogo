package com.example.todolist.ui.home

import androidx.compose.*
import com.example.todolist.di.AppContainer

@Composable
fun HomeScreen(appContainer: AppContainer) {
    val homeRepo = appContainer.homeRepository
    val homeCommander = appContainer.homeCommander
    val navigator = appContainer.navigator
    val ui = appContainer.ui

    +onActive { homeRepo.loadHome() }

    Home(
        projects = ui.home.projectsList,
        onAddProject = { projectName ->
            homeCommander.addProject(projectName)
        },
        onRemoveProject = { id ->
            homeCommander.removeProject(id)
        },
        onNavigateTo = { screen ->
            navigator.navigateTo(screen)
        },
        onApplyFilter = { filter ->
            homeRepo.applyHomeFilter(filter)
        }
    )
}