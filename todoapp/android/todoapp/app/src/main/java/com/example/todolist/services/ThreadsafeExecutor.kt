package com.example.todolist.services

import android.os.Handler
import android.os.Looper
import java.util.concurrent.ExecutorService
import java.util.concurrent.Executors

abstract class ThreadsafeExecutor {
    private val executor: ExecutorService by lazy {
        Executors.newFixedThreadPool(4)
    }

    private val mainThreadHandler: Handler by lazy {
        Handler(Looper.getMainLooper())
    }

    abstract fun <T>run(command: () -> T, callback: (T) -> Unit = {})

    protected fun <T>run(
        command: () -> T,
        onSuccess: (T) -> Unit = {},
        onFail: (Exception) -> Unit = {},
        onFinal: () -> Unit = {}
    ) {
        executor.execute {
            try {
                val result = command()
                mainThreadHandler.post { onSuccess(result) }
            }
            catch (e: Exception) {
                mainThreadHandler.post { onFail(e) }
            }
            finally {
                mainThreadHandler.post { onFinal() }
            }
        }
    }
}
