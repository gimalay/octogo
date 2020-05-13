package com.example.todolist

import android.app.Application
import com.example.todolist.di.AppContainer
import com.example.todolist.di.AppContainerImpl

class TodoListApplication : Application() {

    // AppContainer instance used by the rest of classes to obtain dependencies
    lateinit var container: AppContainer

    override fun onCreate() {
        super.onCreate()
        this.container = AppContainerImpl()
    }
}