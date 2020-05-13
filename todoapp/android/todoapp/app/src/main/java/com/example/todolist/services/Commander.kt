package com.example.todolist.services

import com.example.todolist.model.UiModel
import com.google.protobuf.Message

interface Commander {
    fun execute(payload: Message, callback: (UiModel) -> Unit)
}

class CommanderImpl(
    private val binding: GoBinding,
    private val uiModel: UiModel,
    private val executor: ExecutorInBackground<Unit>
) : Commander {
    override fun execute(payload: Message, callback: (UiModel) -> Unit) {
        uiModel.isLoading = true
        executor.executeInThreads(
            command = {
                binding.execute(payload)
            },
            onSuccess = {
                callback(uiModel)
            },
            onFail = { e ->
                throw Error("Command cannot be executed. Message: " + e.message )
            },
            onFinal = {
                uiModel.isLoading = false
            }
        )
    }
}
