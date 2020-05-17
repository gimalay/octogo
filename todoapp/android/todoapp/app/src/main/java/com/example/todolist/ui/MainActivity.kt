package com.example.todolist.ui

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.ui.core.setContent
import com.example.todolist.TodoListApplication
import com.example.todolist.service.Navigator

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        val appContainer = (this.application as TodoListApplication).container
        setContent { App(appContainer) }
    }

    override fun onBackPressed() {
        val appContainer = (this.application as TodoListApplication).container
        val navigator = appContainer.navigator
        when (navigator.getCurrentScreen()) {
            is Navigator.Screen.Article -> navigator.navigateTo(Navigator.Screen.Home)
            else -> super.onBackPressed()
        }
    }

    override fun onDestroy() {
        super.onDestroy()

        val appContainer = (this.application as TodoListApplication).container
        appContainer.destroy()
    }

}
