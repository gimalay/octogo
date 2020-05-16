package com.example.todolist.services

import com.google.protobuf.Message

interface Commander {
    fun execute(payload: Message, callback: (Unit) -> Unit)
}

class CommanderImpl(
    private val binding: GoBinding,
    private val executor: Executor
) : Commander {
    override fun execute(payload: Message, callback: (Unit) -> Unit) {
        executor.run({ binding.execute(payload) }, callback)
    }
}
