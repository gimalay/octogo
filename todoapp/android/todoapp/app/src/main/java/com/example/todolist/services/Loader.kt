package com.example.todolist.services

import com.google.protobuf.Message

interface Loader {
    fun load(payload: Message, callback: (ByteArray) -> Unit)
}

class LoaderImpl(
    private val binding: GoBinding,
    private val executor: Executor
) : Loader {
    override fun load(payload: Message, callback: (ByteArray) -> Unit) {
        executor.run({ binding.read(payload) }, callback)
    }
}
