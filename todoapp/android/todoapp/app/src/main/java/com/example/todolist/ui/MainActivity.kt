package com.example.todolist.ui

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.ui.core.setContent
import com.example.todolist.TodoListApplication
import com.example.todolist.services.Navigator
import com.example.todolist.ui.article.ArticleScreen
import com.example.todolist.ui.home.HomeScreen

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        val appContainer = (this.application as TodoListApplication).container
        val navigator = appContainer.navigator
        setContent {
            App(navigator.getCurrentScreen()) { screen ->
                when (screen) {
                    is Navigator.Screen.Home -> HomeScreen(
                        appContainer = appContainer
                    )
                    is Navigator.Screen.Article -> ArticleScreen(
                        title = screen.project.name,
                        text = screen.project.id.toString(),
                        appContainer = appContainer
                    )
                }
            }
        }
    }

    override fun onBackPressed() {
        val appContainer = (this.application as TodoListApplication).container
        val navigator = appContainer.navigator
        when (navigator.getCurrentScreen()) {
            is Navigator.Screen.Article -> navigator.navigateTo(Navigator.Screen.Home)
            else -> super.onBackPressed()
        }
    }
}
