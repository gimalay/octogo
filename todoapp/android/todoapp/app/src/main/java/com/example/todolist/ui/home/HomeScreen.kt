package com.example.todolist.ui.home

import androidx.compose.*
import com.example.todolist.model.UiModel
import com.example.todolist.di.AppContainer

@Composable
fun HomeScreen(
    uiState: State<UiModel>,
    appContainer: AppContainer
) {
    val repo = appContainer.homeRepository
    val commander = appContainer.homeCommander
    val navigator = appContainer.navigator

    +onActive { repo.getHome() }

    Home(
        projects = uiState.value.home.projectsList,
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