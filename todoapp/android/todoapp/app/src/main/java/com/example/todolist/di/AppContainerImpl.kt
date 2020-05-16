package com.example.todolist.di

import android.os.Handler
import android.os.Looper
import com.example.todolist.commanders.HomeCommander
import com.example.todolist.model.UiModel
import com.example.todolist.repositories.HomeRepository
import com.example.todolist.repositories.ProjectRepository
import com.example.todolist.services.*
import java.util.concurrent.Executors

/**
 * Dependency Injection container at the application level.
 */
interface AppContainer {
    val navigator: Navigator
    val ui: UiModel

    val homeRepository: HomeRepository
    val homeCommander: HomeCommander

    val projectRepository: ProjectRepository
}

/**
 * Implementation for the Dependency Injection container at the application level.
 *
 * Variables are initialized lazily and the same instance is shared across the whole app.
 */
class AppContainerImpl : AppContainer {
    override val navigator: Navigator by lazy {
        Navigator()
    }

    override val ui: UiModel by lazy {
        UiModel()
    }

    private val goBinding: GoBinding by lazy {
        GoBinding()
    }

    private val loader: Loader by lazy {
        val executor = ExecutorInBackground<ByteArray>(
            executor = Executors.newFixedThreadPool(4),
            mainThreadHandler = Handler(Looper.getMainLooper())
        )

        LoaderImpl(
            binding = goBinding,
            executor = executor,
            ui = ui
        )
    }

    private val commander: Commander by lazy {
        val executor = ExecutorInBackground<Unit>(
            executor = Executors.newFixedThreadPool(4),
            mainThreadHandler = Handler(Looper.getMainLooper())
        )

        CommanderImpl(
            binding = goBinding,
            executor = executor,
            ui = ui
        )
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

}