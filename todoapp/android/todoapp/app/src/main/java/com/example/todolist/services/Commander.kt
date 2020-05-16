package com.example.todolist.services

import com.example.todolist.model.UiModel
import com.google.protobuf.Message

interface Commander {
    fun execute(payload: Message, callback: () -> Unit)
}

class CommanderImpl(
    private val binding: GoBinding,
    private val ui: UiModel,
    private val executor: ExecutorInBackground<Unit>
) : Commander {
    override fun execute(payload: Message, callback: () -> Unit) {
        ui.isLoading = true
        executor.executeInThreads(
            command = {
                binding.execute(payload)
            },
            onSuccess = {
                callback()
            },
            onFail = { e ->
                throw Error("Command cannot be executed. Message: " + e.message )
            },
            onFinal = {
                ui.isLoading = false
            }
        )
    }
}
