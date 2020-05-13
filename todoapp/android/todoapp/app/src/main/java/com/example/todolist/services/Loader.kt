package com.example.todolist.services

import com.example.todolist.model.UiModel
import com.google.protobuf.Message

interface Loader {
    fun load(payload: Message, callback: (ByteArray) -> Unit)
}

class LoaderImpl(
    private val binding: GoBinding,
    private val uiModel: UiModel,
    private val executor: ExecutorInBackground<ByteArray>
) : Loader {
    override fun load(payload: Message, callback: (ByteArray) -> Unit) {
        uiModel.isLoading = true
        executor.executeInThreads(
            command = {
                binding.read(payload)
            },
            onSuccess = { result ->
                callback(result)
            },
            onFail = { e ->
                throw Error("Data cannot be read. Message: " + e.message )
            },
            onFinal = {
                uiModel.isLoading = false
            }
        )
    }
}
