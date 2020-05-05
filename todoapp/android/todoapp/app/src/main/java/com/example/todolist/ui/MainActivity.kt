package com.example.todolist.ui

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import androidx.ui.core.setContent
import com.example.todolist.di.ServiceLocator

class MainActivity : AppCompatActivity() {
    private val loader = ServiceLocator.loader
    private val executor = ServiceLocator.executor

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            App(loader, executor)
        }
    }
}
