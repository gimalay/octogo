package com.example.todolist.di

import com.example.todolist.services.Commander
import com.example.todolist.services.GoBinding
import com.example.todolist.services.Loader
import com.example.todolist.services.Navigator

/**
 * Dependency Injection container at the application level.
 */
interface AppContainer {
    val loader: Loader
    val commander: Commander
    val navigator: Navigator
}

/**
 * Implementation for the Dependency Injection container at the application level.
 *
 * Variables are initialized lazily and the same instance is shared across the whole app.
 */
class AppContainerImpl : AppContainer {
    private val goBinding: GoBinding by lazy {
        GoBinding()
    }

    override val navigator: Navigator by lazy {
        Navigator()
    }

    override val loader: Loader by lazy {
        Loader(
            binding = this.goBinding
        )
    }

    override val commander: Commander by lazy {
        Commander(
           binding = this.goBinding
        )
    }
}