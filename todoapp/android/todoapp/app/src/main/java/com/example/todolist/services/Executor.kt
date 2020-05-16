package com.example.todolist.services

import com.example.todolist.model.UiModel

class Executor(private val ui: UiModel) : ThreadsafeExecutor() {
    override fun <T>run(command: () -> T, callback: (T) -> Unit) {
        ui.isLoading = true
        run(
            command = command,
            onSuccess = callback,
            onFail = { throw it },
            onFinal = { ui.isLoading = false }
        )
    }
}

