package com.example.todolist.service

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel

class Navigator {
    sealed class Screen {
        object Home : Screen()
        class Article(val project: ViewModel.Home.Project) : Screen()
    }

    private val navigation: Navigation = Navigation()

    fun getCurrentScreen(): Screen {
        return this.navigation.currentScreen
    }

    fun navigateTo(destination: Screen) {
        this.navigation.currentScreen = destination
    }
}

@Model
private class Navigation {
    var currentScreen: Navigator.Screen = Navigator.Screen.Home
}
