package com.example.todolist.di

import com.example.todolist.commander.HomeCommander
import com.example.todolist.model.UiModel
import com.example.todolist.repository.HomeRepository
import com.example.todolist.repository.ProjectRepository
import com.example.todolist.service.*

/**
 * Dependency Injection container at the application level.
 */
interface AppContainer {
    val navigator: Navigator
    val ui: UiModel

    val homeRepository: HomeRepository
    val homeCommander: HomeCommander

    val projectRepository: ProjectRepository

    fun destroy()
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

    private val executor: Executor by lazy {
        Executor(ui)
    }

    private val loader: Loader by lazy {
        LoaderImpl(
            binding = goBinding,
            executor = executor
        )
    }

    private val commander: Commander by lazy {
        CommanderImpl(
            binding = goBinding,
            executor = executor
        )
    }

    override val navigator: Navigator by lazy {
        Navigator()
    }

    override val ui: UiModel by lazy {
        UiModel()
    }

    override val homeCommander: HomeCommander by lazy {
        HomeCommander(
            commander = commander,
            repository = homeRepository,
            ui = ui
        )
    }

    override val homeRepository: HomeRepository by lazy {
        HomeRepository(
            loader = loader,
            ui = ui
        )
    }

    override val projectRepository: ProjectRepository by lazy {
        ProjectRepository(
            loader = loader,
            ui = ui
        )
    }

    override fun destroy() {
        goBinding.close()
    }

}