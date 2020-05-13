package com.example.todolist.services

import android.os.Handler
import java.util.concurrent.ExecutorService

class ExecutorInBackground<T> (
    private val executor: ExecutorService,
    private val mainThreadHandler: Handler
) {
    fun executeInThreads(
        command: () -> T,
        onSuccess: (T) -> Unit,
        onFail: (Exception) -> Unit = {},
        onFinal: () -> Unit = {}
    ) {
        executor.execute {
            try {
                val result = command()
                mainThreadHandler.post { onSuccess(result) }
            } catch (e: Exception) {
                mainThreadHandler.post { onFail(e) }
            } finally {
                mainThreadHandler.post { onFinal() }
            }
        }
    }
}
