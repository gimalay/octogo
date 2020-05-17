package com.example.todolist.ui

import androidx.compose.*
import androidx.ui.animation.Crossfade
import androidx.ui.material.MaterialTheme
import com.example.todolist.di.AppContainer
import com.example.todolist.service.Navigator
import com.example.todolist.ui.article.ArticleScreen
import com.example.todolist.ui.home.HomeScreen

@Composable
fun App(appContainer: AppContainer) {
    val navigator = appContainer.navigator

    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {
        Crossfade(navigator.getCurrentScreen()) { screen ->
            when (screen) {
                is Navigator.Screen.Home -> HomeScreen(appContainer = appContainer)
                is Navigator.Screen.Article -> ArticleScreen(
                    appContainer = appContainer,
                    title = screen.project.name,
                    text = screen.project.id.toString()
                )
            }
        }
    }
}
